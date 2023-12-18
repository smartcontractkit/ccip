package v1_0_0

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/custom_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/logpollerutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	EXEC_EXECUTION_STATE_CHANGES = "Exec execution state changes"
	EXEC_TOKEN_POOL_ADDED        = "Token pool added"
	EXEC_TOKEN_POOL_REMOVED      = "Token pool removed"
)

var (
	abiOffRamp                                             = abihelpers.MustParseABI(evm_2_evm_offramp_1_0_0.EVM2EVMOffRampABI)
	abiCustomTokenPool                                     = abihelpers.MustParseABI(custom_token_pool.CustomTokenPoolABI)
	_                               ccipdata.OffRampReader = &OffRamp{}
	ExecutionStateChangedEvent                             = abihelpers.MustGetEventID("ExecutionStateChanged", abiOffRamp)
	PoolAddedEvent                                         = abihelpers.MustGetEventID("PoolAdded", abiOffRamp)
	PoolRemovedEvent                                       = abihelpers.MustGetEventID("PoolRemoved", abiOffRamp)
	ExecutionStateChangedSeqNrIndex                        = 1
)

var offRamp_poolAddedPoolRemovedEvents = []common.Hash{PoolAddedEvent, PoolRemovedEvent}

type ExecOnchainConfig evm_2_evm_offramp_1_0_0.EVM2EVMOffRampDynamicConfig

func (d ExecOnchainConfig) AbiString() string {
	return `
	[
		{
			"components": [
				{"name": "permissionLessExecutionThresholdSeconds", "type": "uint32"},
				{"name": "router", "type": "address"},
				{"name": "priceRegistry", "type": "address"},
				{"name": "maxTokensLength", "type": "uint16"},
				{"name": "maxDataSize", "type": "uint32"}
			],
			"type": "tuple"
		}
	]`
}

func (d ExecOnchainConfig) Validate() error {
	if d.PermissionLessExecutionThresholdSeconds == 0 {
		return errors.New("must set PermissionLessExecutionThresholdSeconds")
	}
	if d.Router == (common.Address{}) {
		return errors.New("must set Router address")
	}
	if d.PriceRegistry == (common.Address{}) {
		return errors.New("must set PriceRegistry address")
	}
	if d.MaxTokensLength == 0 {
		return errors.New("must set MaxTokensLength")
	}
	if d.MaxDataSize == 0 {
		return errors.New("must set MaxDataSize")
	}
	return nil
}

func (d ExecOnchainConfig) PermissionLessExecutionThresholdDuration() time.Duration {
	return time.Duration(d.PermissionLessExecutionThresholdSeconds) * time.Second
}

type OffRampTokens struct {
	destinationTokens []common.Address
	sourceTokens      []common.Address
	destinationPool   map[common.Address]common.Address
}

type OffRamp struct {
	offRamp                 *evm_2_evm_offramp_1_0_0.EVM2EVMOffRamp
	addr                    common.Address
	lp                      logpoller.LogPoller
	lggr                    logger.Logger
	ec                      client.Client
	evmBatchCaller          rpclib.EvmBatchCaller
	filters                 []logpoller.Filter
	estimator               gas.EvmFeeEstimator
	executionReportArgs     abi.Arguments
	eventIndex              int
	eventSig                common.Hash
	cachedOffRampTokens     cache.AutoSync[OffRampTokens]
	sourceToDestTokensCache sync.Map

	// Dynamic config
	configMu          sync.RWMutex
	gasPriceEstimator prices.GasPriceEstimatorExec
	offchainConfig    ccipdata.ExecOffchainConfig
	onchainConfig     ccipdata.ExecOnchainConfig
}

func (o *OffRamp) GetStaticConfig(ctx context.Context) (ccipdata.OffRampStaticConfig, error) {
	if o.offRamp == nil {
		return ccipdata.OffRampStaticConfig{}, fmt.Errorf("offramp not initialized")
	}
	c, err := o.offRamp.GetStaticConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return ccipdata.OffRampStaticConfig{}, fmt.Errorf("error while retrieving offramp config: %w", err)
	}
	return ccipdata.OffRampStaticConfig{
		CommitStore:         c.CommitStore,
		ChainSelector:       c.ChainSelector,
		SourceChainSelector: c.SourceChainSelector,
		OnRamp:              c.OnRamp,
		PrevOffRamp:         c.PrevOffRamp,
		ArmProxy:            c.ArmProxy,
	}, nil
}

func (o *OffRamp) GetExecutionState(ctx context.Context, sequenceNumber uint64) (uint8, error) {
	return o.offRamp.GetExecutionState(&bind.CallOpts{Context: ctx}, sequenceNumber)
}

func (o *OffRamp) GetSenderNonce(ctx context.Context, sender common.Address) (uint64, error) {
	return o.offRamp.GetSenderNonce(&bind.CallOpts{Context: ctx}, sender)
}

func (o *OffRamp) CurrentRateLimiterState(ctx context.Context) (evm_2_evm_offramp.RateLimiterTokenBucket, error) {
	state, err := o.offRamp.CurrentRateLimiterState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return *new(evm_2_evm_offramp.RateLimiterTokenBucket), err
	}
	return evm_2_evm_offramp.RateLimiterTokenBucket{
		Tokens:      state.Tokens,
		LastUpdated: state.LastUpdated,
		IsEnabled:   state.IsEnabled,
		Capacity:    state.Capacity,
		Rate:        state.Rate,
	}, nil
}

func (o *OffRamp) GetDestinationToken(ctx context.Context, address common.Address) (common.Address, error) {
	return o.offRamp.GetDestinationToken(&bind.CallOpts{Context: ctx}, address)
}

func (o *OffRamp) getDestinationTokensFromSourceTokens(ctx context.Context, tokenAddresses []common.Address) ([]common.Address, error) {
	destTokens := make([]common.Address, len(tokenAddresses))
	found := make(map[common.Address]bool)
	for i, tokenAddress := range tokenAddresses {
		if v, exists := o.sourceToDestTokensCache.Load(tokenAddress); exists {
			if destToken, isAddr := v.(common.Address); isAddr {
				destTokens[i] = destToken
				found[tokenAddress] = true
			} else {
				o.lggr.Errorf("source to dest cache contains invalid type %T", v)
			}
		}
	}

	if len(found) == len(tokenAddresses) {
		return destTokens, nil
	}

	evmCalls := make([]rpclib.EvmCall, 0, len(tokenAddresses))
	for _, sourceTk := range tokenAddresses {
		if !found[sourceTk] {
			evmCalls = append(evmCalls, rpclib.NewEvmCall(abiOffRamp, "getDestinationToken", o.addr, sourceTk))
		}
	}

	latestBlock, err := o.lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get latest block: %w", err)
	}

	results, err := o.evmBatchCaller.BatchCall(ctx, uint64(latestBlock.BlockNumber), evmCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call limit: %w", err)
	}

	destTokensFromRpc, err := rpclib.ParseOutputs[common.Address](results, func(d rpclib.DataAndErr) (common.Address, error) {
		return rpclib.ParseOutput[common.Address](d, 0)
	})
	if err != nil {
		return nil, fmt.Errorf("parse outputs: %w", err)
	}

	j := 0
	for i, sourceToken := range tokenAddresses {
		if !found[sourceToken] {
			destTokens[i] = destTokensFromRpc[j]
			o.sourceToDestTokensCache.Store(sourceToken, destTokens[i])
			j++
		}
	}

	seenDestTokens := mapset.NewSet[common.Address]()
	for _, destToken := range destTokens {
		if seenDestTokens.Contains(destToken) {
			return nil, fmt.Errorf("offRamp misconfig, destination token %s already exists", destToken)
		}
		seenDestTokens.Add(destToken)
	}

	return destTokens, nil
}

func (o *OffRamp) GetTokenPoolsRateLimits(ctx context.Context, poolAddresses []common.Address) ([]ccipdata.TokenBucketRateLimit, error) {
	if len(poolAddresses) == 0 {
		return nil, nil
	}

	evmCalls := make([]rpclib.EvmCall, 0, len(poolAddresses))
	for _, poolAddress := range poolAddresses {
		evmCalls = append(evmCalls, rpclib.NewEvmCall(
			abiCustomTokenPool,
			"currentOffRampRateLimiterState",
			poolAddress,
			o.addr,
		))
	}

	latestBlock, err := o.lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get latest block: %w", err)
	}

	results, err := o.evmBatchCaller.BatchCall(ctx, uint64(latestBlock.BlockNumber), evmCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call limit: %w", err)
	}

	rateLimits, err := rpclib.ParseOutputs[ccipdata.TokenBucketRateLimit](results, func(d rpclib.DataAndErr) (ccipdata.TokenBucketRateLimit, error) {
		return rpclib.ParseOutput[ccipdata.TokenBucketRateLimit](d, 0)
	})
	if err != nil {
		return nil, fmt.Errorf("parse outputs: %w", err)
	}

	return rateLimits, nil
}

func (o *OffRamp) GetSourceToDestTokensMapping(ctx context.Context) (map[common.Address]common.Address, error) {
	sourceTokens, err := o.getSourceTokens(ctx)
	if err != nil {
		return nil, err
	}

	destTokens, err := o.getDestinationTokensFromSourceTokens(ctx, sourceTokens)
	if err != nil {
		return nil, fmt.Errorf("get destination tokens from source tokens: %w", err)
	}

	srcToDstTokenMapping := make(map[common.Address]common.Address, len(sourceTokens))
	for i, sourceToken := range sourceTokens {
		srcToDstTokenMapping[sourceToken] = destTokens[i]
	}

	return srcToDstTokenMapping, nil
}

func (o *OffRamp) GetDestinationTokens(ctx context.Context) ([]common.Address, error) {
	cached, err := o.getCachedOffRampTokens(ctx)
	if err != nil {
		return nil, err
	}
	return cached.destinationTokens, nil
}

func (o *OffRamp) GetDestinationTokenPools(ctx context.Context) (map[common.Address]common.Address, error) {
	cached, err := o.getCachedOffRampTokens(ctx)
	if err != nil {
		return nil, err
	}
	return cached.destinationPool, nil
}

func (o *OffRamp) getSourceTokens(ctx context.Context) ([]common.Address, error) {
	cached, err := o.getCachedOffRampTokens(ctx)
	if err != nil {
		return nil, err
	}
	return cached.sourceTokens, nil
}

func (o *OffRamp) getCachedOffRampTokens(ctx context.Context) (OffRampTokens, error) {
	return o.cachedOffRampTokens.Get(ctx, func(ctx context.Context) (OffRampTokens, error) {
		destTokens, err := o.offRamp.GetDestinationTokens(&bind.CallOpts{Context: ctx})
		if err != nil {
			return OffRampTokens{}, fmt.Errorf("get destination tokens: %w", err)
		}
		sourceTokens, err := o.offRamp.GetSupportedTokens(&bind.CallOpts{Context: ctx})
		if err != nil {
			return OffRampTokens{}, err
		}
		destPools, err := o.getPoolsByDestTokens(ctx, destTokens)
		if err != nil {
			return OffRampTokens{}, fmt.Errorf("get pools by dest tokens: %w", err)
		}
		tokenToPool := make(map[common.Address]common.Address, len(destTokens))
		for i := range destTokens {
			tokenToPool[destTokens[i]] = destPools[i]
		}

		return OffRampTokens{
			destinationTokens: destTokens,
			sourceTokens:      sourceTokens,
			destinationPool:   tokenToPool,
		}, nil
	})
}

func (o *OffRamp) getPoolsByDestTokens(ctx context.Context, tokenAddrs []common.Address) ([]common.Address, error) {
	evmCalls := make([]rpclib.EvmCall, 0, len(tokenAddrs))
	for _, tk := range tokenAddrs {
		evmCalls = append(evmCalls, rpclib.NewEvmCall(
			abiOffRamp,
			"getPoolByDestToken",
			o.addr,
			tk,
		))
	}

	latestBlock, err := o.lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get latest block: %w", err)
	}

	results, err := o.evmBatchCaller.BatchCall(ctx, uint64(latestBlock.FinalizedBlockNumber), evmCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call limit: %w", err)
	}

	destPools, err := rpclib.ParseOutputs[common.Address](results, func(d rpclib.DataAndErr) (common.Address, error) {
		return rpclib.ParseOutput[common.Address](d, 0)
	})
	if err != nil {
		return nil, fmt.Errorf("parse outputs: %w", err)
	}

	return destPools, nil
}

func (o *OffRamp) OffchainConfig() ccipdata.ExecOffchainConfig {
	o.configMu.RLock()
	defer o.configMu.RUnlock()
	return o.offchainConfig
}

func (o *OffRamp) OnchainConfig() ccipdata.ExecOnchainConfig {
	o.configMu.RLock()
	defer o.configMu.RUnlock()
	return o.onchainConfig
}

func (o *OffRamp) GasPriceEstimator() prices.GasPriceEstimatorExec {
	o.configMu.RLock()
	defer o.configMu.RUnlock()
	return o.gasPriceEstimator
}

func (o *OffRamp) Address() common.Address {
	return o.addr
}

func (o *OffRamp) ChangeConfig(onchainConfig []byte, offchainConfig []byte) (common.Address, common.Address, error) {
	onchainConfigParsed, err := abihelpers.DecodeAbiStruct[ExecOnchainConfig](onchainConfig)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	offchainConfigParsed, err := ccipconfig.DecodeOffchainConfig[ccipdata.ExecOffchainConfig](offchainConfig)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}
	destRouter, err := router.NewRouter(onchainConfigParsed.Router, o.ec)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}
	destWrappedNative, err := destRouter.GetWrappedNative(nil)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}
	o.configMu.Lock()
	o.offchainConfig = ccipdata.ExecOffchainConfig{
		SourceFinalityDepth:         offchainConfigParsed.SourceFinalityDepth,
		DestFinalityDepth:           offchainConfigParsed.DestFinalityDepth,
		DestOptimisticConfirmations: offchainConfigParsed.DestOptimisticConfirmations,
		BatchGasLimit:               offchainConfigParsed.BatchGasLimit,
		RelativeBoostPerWaitHour:    offchainConfigParsed.RelativeBoostPerWaitHour,
		MaxGasPrice:                 offchainConfigParsed.MaxGasPrice,
		InflightCacheExpiry:         offchainConfigParsed.InflightCacheExpiry,
		RootSnoozeTime:              offchainConfigParsed.RootSnoozeTime,
	}
	o.onchainConfig = ccipdata.ExecOnchainConfig{PermissionLessExecutionThresholdSeconds: time.Second * time.Duration(onchainConfigParsed.PermissionLessExecutionThresholdSeconds)}
	o.gasPriceEstimator = prices.NewExecGasPriceEstimator(o.estimator, big.NewInt(int64(offchainConfigParsed.MaxGasPrice)), 0)
	o.configMu.Unlock()

	o.lggr.Infow("Starting exec plugin",
		"offchainConfig", onchainConfigParsed,
		"onchainConfig", offchainConfigParsed)
	return onchainConfigParsed.PriceRegistry, destWrappedNative, nil
}

func (o *OffRamp) Close(qopts ...pg.QOpt) error {
	return logpollerutil.UnregisterLpFilters(o.lp, o.filters, qopts...)
}

func (o *OffRamp) GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]ccipdata.Event[ccipdata.ExecutionStateChanged], error) {
	latestBlock, err := o.lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get lp latest block: %w", err)
	}

	logs, err := o.lp.IndexedLogsTopicRange(
		o.eventSig,
		o.addr,
		o.eventIndex,
		logpoller.EvmWord(seqNumMin),
		logpoller.EvmWord(seqNumMax),
		logpoller.Confirmations(confs),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return ccipdata.ParseLogs[ccipdata.ExecutionStateChanged](
		logs,
		o.lggr,
		func(log types.Log) (*ccipdata.ExecutionStateChanged, error) {
			sc, err := o.offRamp.ParseExecutionStateChanged(log)
			if err != nil {
				return nil, err
			}

			return &ccipdata.ExecutionStateChanged{
				SequenceNumber: sc.SequenceNumber,
				Finalized:      sc.Raw.BlockNumber >= uint64(latestBlock.FinalizedBlockNumber),
			}, nil
		},
	)
}

func encodeExecutionReport(args abi.Arguments, report ccipdata.ExecReport) ([]byte, error) {
	var msgs []evm_2_evm_offramp_1_0_0.InternalEVM2EVMMessage
	for _, msg := range report.Messages {
		var ta []evm_2_evm_offramp_1_0_0.ClientEVMTokenAmount
		for _, tokenAndAmount := range msg.TokenAmounts {
			ta = append(ta, evm_2_evm_offramp_1_0_0.ClientEVMTokenAmount{
				Token:  tokenAndAmount.Token,
				Amount: tokenAndAmount.Amount,
			})
		}
		msgs = append(msgs, evm_2_evm_offramp_1_0_0.InternalEVM2EVMMessage{
			SourceChainSelector: msg.SourceChainSelector,
			Sender:              msg.Sender,
			Receiver:            msg.Receiver,
			SequenceNumber:      msg.SequenceNumber,
			GasLimit:            msg.GasLimit,
			Strict:              msg.Strict,
			Nonce:               msg.Nonce,
			FeeToken:            msg.FeeToken,
			FeeTokenAmount:      msg.FeeTokenAmount,
			Data:                msg.Data,
			TokenAmounts:        ta,
			MessageId:           msg.MessageId,
		})
	}

	rep := evm_2_evm_offramp_1_0_0.InternalExecutionReport{
		Messages:          msgs,
		OffchainTokenData: report.OffchainTokenData,
		Proofs:            report.Proofs,
		ProofFlagBits:     report.ProofFlagBits,
	}
	return args.PackValues([]interface{}{&rep})
}

func (o *OffRamp) EncodeExecutionReport(report ccipdata.ExecReport) ([]byte, error) {
	return encodeExecutionReport(o.executionReportArgs, report)
}

func DecodeExecReport(args abi.Arguments, report []byte) (ccipdata.ExecReport, error) {
	unpacked, err := args.Unpack(report)
	if err != nil {
		return ccipdata.ExecReport{}, err
	}
	if len(unpacked) == 0 {
		return ccipdata.ExecReport{}, errors.New("assumptionViolation: expected at least one element")
	}

	erStruct, ok := unpacked[0].(struct {
		Messages []struct {
			SourceChainSelector uint64         `json:"sourceChainSelector"`
			SequenceNumber      uint64         `json:"sequenceNumber"`
			FeeTokenAmount      *big.Int       `json:"feeTokenAmount"`
			Sender              common.Address `json:"sender"`
			Nonce               uint64         `json:"nonce"`
			GasLimit            *big.Int       `json:"gasLimit"`
			Strict              bool           `json:"strict"`
			Receiver            common.Address `json:"receiver"`
			Data                []uint8        `json:"data"`
			TokenAmounts        []struct {
				Token  common.Address `json:"token"`
				Amount *big.Int       `json:"amount"`
			} `json:"tokenAmounts"`
			FeeToken  common.Address `json:"feeToken"`
			MessageId [32]uint8      `json:"messageId"`
		} `json:"messages"`
		OffchainTokenData [][][]uint8 `json:"offchainTokenData"`
		Proofs            [][32]uint8 `json:"proofs"`
		ProofFlagBits     *big.Int    `json:"proofFlagBits"`
	})

	if !ok {
		return ccipdata.ExecReport{}, fmt.Errorf("got %T", unpacked[0])
	}
	messages := []internal.EVM2EVMMessage{}
	for _, msg := range erStruct.Messages {
		var tokensAndAmounts []internal.TokenAmount
		for _, tokenAndAmount := range msg.TokenAmounts {
			tokensAndAmounts = append(tokensAndAmounts, internal.TokenAmount{
				Token:  tokenAndAmount.Token,
				Amount: tokenAndAmount.Amount,
			})
		}
		messages = append(messages, internal.EVM2EVMMessage{
			SequenceNumber:      msg.SequenceNumber,
			GasLimit:            msg.GasLimit,
			Nonce:               msg.Nonce,
			MessageId:           msg.MessageId,
			SourceChainSelector: msg.SourceChainSelector,
			Sender:              msg.Sender,
			Receiver:            msg.Receiver,
			Strict:              msg.Strict,
			FeeToken:            msg.FeeToken,
			FeeTokenAmount:      msg.FeeTokenAmount,
			Data:                msg.Data,
			TokenAmounts:        tokensAndAmounts,
			// TODO: Not needed for plugins, but should be recomputed for consistency.
			// Requires the offramp knowing about onramp version
			Hash: [32]byte{},
		})
	}

	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	return ccipdata.ExecReport{
		Messages:          messages,
		OffchainTokenData: erStruct.OffchainTokenData,
		Proofs:            erStruct.Proofs,
		ProofFlagBits:     new(big.Int).SetBytes(erStruct.ProofFlagBits.Bytes()),
	}, nil

}

func (o *OffRamp) DecodeExecutionReport(report []byte) (ccipdata.ExecReport, error) {
	return DecodeExecReport(o.executionReportArgs, report)
}

func (o *OffRamp) TokenEvents() []common.Hash {
	return offRamp_poolAddedPoolRemovedEvents
}

func (o *OffRamp) RegisterFilters(qopts ...pg.QOpt) error {
	return logpollerutil.RegisterLpFilters(o.lp, o.filters, qopts...)
}

func NewOffRamp(lggr logger.Logger, addr common.Address, ec client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator) (*OffRamp, error) {
	offRamp, err := evm_2_evm_offramp_1_0_0.NewEVM2EVMOffRamp(addr, ec)
	if err != nil {
		return nil, err
	}

	executionStateChangedSequenceNumberIndex := 1
	executionReportArgs := abihelpers.MustGetMethodInputs("manuallyExecute", abiOffRamp)[:1]
	filters := []logpoller.Filter{
		{
			Name:      logpoller.FilterName(EXEC_EXECUTION_STATE_CHANGES, addr.String()),
			EventSigs: []common.Hash{ExecutionStateChangedEvent},
			Addresses: []common.Address{addr},
		},
		{
			Name:      logpoller.FilterName(EXEC_TOKEN_POOL_ADDED, addr.String()),
			EventSigs: []common.Hash{PoolAddedEvent},
			Addresses: []common.Address{addr},
		},
		{
			Name:      logpoller.FilterName(EXEC_TOKEN_POOL_REMOVED, addr.String()),
			EventSigs: []common.Hash{PoolRemovedEvent},
			Addresses: []common.Address{addr},
		},
	}

	return &OffRamp{
		offRamp:             offRamp,
		ec:                  ec,
		addr:                addr,
		lggr:                lggr,
		lp:                  lp,
		filters:             filters,
		estimator:           estimator,
		executionReportArgs: executionReportArgs,
		eventSig:            ExecutionStateChangedEvent,
		eventIndex:          executionStateChangedSequenceNumberIndex,
		configMu:            sync.RWMutex{},
		evmBatchCaller: rpclib.NewDynamicLimitedBatchCaller(
			lggr,
			ec,
			rpclib.DefaultRpcBatchSizeLimit,
			rpclib.DefaultRpcBatchBackOffMultiplier,
		),
		cachedOffRampTokens: cache.NewLogpollerEventsBased[OffRampTokens](
			lp,
			offRamp_poolAddedPoolRemovedEvents,
			offRamp.Address(),
		),

		// values set on the fly after ChangeConfig is called
		gasPriceEstimator: prices.ExecGasPriceEstimator{},
		offchainConfig:    ccipdata.ExecOffchainConfig{},
		onchainConfig:     ccipdata.ExecOnchainConfig{},
	}, nil
}

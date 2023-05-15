package ccip

import (
	"context"
	"encoding/hex"
	"math/big"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/pkg/errors"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"

	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

// exec Report should make sure to cap returned payload to this limit
const MaxExecutionReportLength = 250_000

var (
	_ types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_ types.ReportingPlugin        = &ExecutionReportingPlugin{}
)

// ExecutionReportToEthTxMeta generates a txmgr.EthTxMeta from the given report.
// all the message ids will be added to the tx metadata.
func ExecutionReportToEthTxMeta(report []byte) (*txmgr.EthTxMeta, error) {
	execReport, err := abihelpers.DecodeExecutionReport(report)
	if err != nil {
		return nil, err
	}

	msgIDs := make([]string, len(execReport.EncodedMessages))
	for i, encMsg := range execReport.EncodedMessages {
		msg, err := abihelpers.DecodeMessage(encMsg)
		if err != nil {
			return nil, err
		}
		msgIDs[i] = hexutil.Encode(msg.MessageId[:])
	}

	return &txmgr.EthTxMeta{
		MessageIDs: msgIDs,
	}, nil
}

func MessagesFromExecutionReport(report types.Report) ([]uint64, [][]byte, error) {
	decodeExecutionReport, err := abihelpers.DecodeExecutionReport(report)
	if err != nil {
		return nil, nil, err
	}
	return decodeExecutionReport.SequenceNumbers, decodeExecutionReport.EncodedMessages, nil
}

type ExecutionPluginConfig struct {
	lggr                  logger.Logger
	sourceLP, destLP      logpoller.LogPoller
	onRamp                evm_2_evm_onramp.EVM2EVMOnRampInterface
	offRamp               evm_2_evm_offramp.EVM2EVMOffRampInterface
	commitStore           commit_store.CommitStoreInterface
	srcPriceRegistry      price_registry.PriceRegistryInterface
	srcWrappedNativeToken common.Address
	destClient            evmclient.Client
	destGasEstimator      txmgrtypes.FeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, common.Hash]
	leafHasher            hasher.LeafHasherInterface[[32]byte]
}

type ExecutionReportingPlugin struct {
	config                    ExecutionPluginConfig
	F                         int
	lggr                      logger.Logger
	inflightReports           *inflightReportsContainer
	snoozedRoots              map[[32]byte]time.Time
	destPriceRegistry         price_registry.PriceRegistryInterface
	destWrappedNative         common.Address
	onchainConfig             ccipconfig.ExecOnchainConfig
	offchainConfig            ccipconfig.ExecOffchainConfig
	srcToDstTokenMappingMu    sync.RWMutex
	srcToDstTokenMappingBlock int64
	srcToDstTokenMapping      map[common.Address]common.Address
}

type ExecutionReportingPluginFactory struct {
	config ExecutionPluginConfig
}

func NewExecutionReportingPluginFactory(config ExecutionPluginConfig) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{config: config}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	onchainConfig, err := abihelpers.DecodeAbiStruct[ccipconfig.ExecOnchainConfig](config.OnchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	offchainConfig, err := ccipconfig.DecodeOffchainConfig[ccipconfig.ExecOffchainConfig](config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	priceRegistry, err := price_registry.NewPriceRegistry(onchainConfig.PriceRegistry, rf.config.destClient)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	destRouter, err := router.NewRouter(onchainConfig.Router, rf.config.destClient)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	destWrappedNative, err := destRouter.GetWrappedNative(&bind.CallOpts{})
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	sourceToDestTokenMapping, err := generateSourceToDestTokenMapping(context.Background(), rf.config.offRamp)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	latestBlock, err := rf.config.destLP.LatestBlock()
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	rf.config.lggr.Infow("Starting exec plugin", "offchainConfig", offchainConfig, "onchainConfig", onchainConfig)

	return &ExecutionReportingPlugin{
			config:                    rf.config,
			F:                         config.F,
			lggr:                      rf.config.lggr.Named("ExecutionReportingPlugin"),
			snoozedRoots:              make(map[[32]byte]time.Time),
			inflightReports:           newInflightReportsContainer(offchainConfig.InflightCacheExpiry.Duration()),
			destPriceRegistry:         priceRegistry,
			destWrappedNative:         destWrappedNative,
			onchainConfig:             onchainConfig,
			offchainConfig:            offchainConfig,
			srcToDstTokenMappingBlock: latestBlock - int64(offchainConfig.DestIncomingConfirmations),
			srcToDstTokenMapping:      sourceToDestTokenMapping,
		}, types.ReportingPluginInfo{
			Name:          "CCIPExecution",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}, nil
}

func (r *ExecutionReportingPlugin) Query(context.Context, types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r *ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("ExecutionObservation")
	if isCommitStoreDownNow(ctx, lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	// Expire any inflight reports.
	r.inflightReports.expire(lggr)
	inFlight := r.inflightReports.getAll()

	observationBuildStart := time.Now()
	// IMPORTANT: We build executable set based on the leaders token prices, ensuring consistency across followers.
	executableObservations, err := r.getExecutableObservations(ctx, timestamp, inFlight)
	measureObservationBuildDuration(timestamp, time.Since(observationBuildStart))
	if err != nil {
		return nil, err
	}
	// cap observations which fits MaxObservationLength (after serialized)
	capped := sort.Search(len(executableObservations), func(i int) bool {
		var encoded []byte
		encoded, err = ExecutionObservation{Messages: executableObservations[:i+1]}.Marshal()
		if err != nil {
			// false makes Search keep looking to the right, always including any "erroring" ObservedMessage and allowing us to detect in the bottom
			return false
		}
		return len(encoded) > MaxObservationLength
	})
	if err != nil {
		return nil, err
	}
	executableObservations = executableObservations[:capped]
	lggr.Infof("executable observations %+v %v", executableObservations, abihelpers.EventSignatures.SendRequested)

	// Note can be empty
	return ExecutionObservation{Messages: executableObservations}.Marshal()
}

func (r *ExecutionReportingPlugin) getExecutableObservations(ctx context.Context, timestamp types.ReportTimestamp, inflight []InflightInternalExecutionReport) ([]ObservedMessage, error) {
	unexpiredReports, err := getUnexpiredCommitReports(r.config.destLP, r.config.commitStore, r.onchainConfig.PermissionLessExecutionThresholdDuration())
	if err != nil {
		return nil, err
	}
	r.lggr.Infow("unexpired roots", "n", len(unexpiredReports))
	if len(unexpiredReports) == 0 {
		return []ObservedMessage{}, nil
	}

	rpcPreprationStart := time.Now()
	// This could result in slightly different values on each call as
	// the function returns the allowed amount at the time of the last block.
	// Since this will only increase over time, the highest observed value will
	// always be the lower bound of what would be available on chain
	// since we already account for inflight txs.
	rateLimiterState, err := r.config.offRamp.CurrentRateLimiterState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	allowedTokenAmount := rateLimiterState.Tokens

	// Updates or sets the source to dest token mapping if needed
	if err = r.updateSourceToDestTokenMapping(ctx); err != nil {
		return nil, err
	}

	r.srcToDstTokenMappingMu.RLock()
	supportedDestTokens := make([]common.Address, 0, len(r.srcToDstTokenMapping))
	for _, destToken := range r.srcToDstTokenMapping {
		supportedDestTokens = append(supportedDestTokens, destToken)
	}
	r.srcToDstTokenMappingMu.RUnlock()

	srcTokensPrices, err := getTokensPrices(ctx, r.config.srcPriceRegistry, []common.Address{r.config.srcWrappedNativeToken})
	if err != nil {
		return nil, err
	}
	destTokensPrices, err := getTokensPrices(ctx, r.destPriceRegistry, append(supportedDestTokens, r.destWrappedNative))
	if err != nil {
		return nil, err
	}
	destGasPriceWei, _, err := r.config.destGasEstimator.GetFee(ctx, nil, 0, assets.NewWei(big.NewInt(int64(r.offchainConfig.MaxGasPrice))))
	if err != nil {
		return nil, errors.Wrap(err, "could not estimate destination gas price")
	}
	destGasPrice := destGasPriceWei.Legacy.ToInt()
	if destGasPriceWei.DynamicFeeCap != nil {
		destGasPrice = destGasPriceWei.DynamicFeeCap.ToInt()
	}
	measureBatchPrepareRPCDuration(timestamp, time.Since(rpcPreprationStart))

	r.lggr.Debugw("processing unexpired reports", "n", len(unexpiredReports))
	measureNumberOfReportsProcessed(timestamp, len(unexpiredReports))
	reportIterationStart := time.Now()
	defer func() {
		measureReportsIterationDuration(timestamp, time.Since(reportIterationStart))
	}()
	for _, unexpiredReport := range unexpiredReports {
		if ctx.Err() != nil {
			r.lggr.Warn("killed by context")
			break
		}
		snoozeUntil, haveSnoozed := r.snoozedRoots[unexpiredReport.MerkleRoot]
		if haveSnoozed && time.Now().Before(snoozeUntil) {
			continue
		}
		blessed, err := r.config.commitStore.IsBlessed(&bind.CallOpts{Context: ctx}, unexpiredReport.MerkleRoot)
		if err != nil {
			return nil, err
		}
		if !blessed {
			r.lggr.Infow("report is accepted but not blessed", "report", hexutil.Encode(unexpiredReport.MerkleRoot[:]))
			incSkippedRequests(reasonNotBlessed)
			continue
		}
		// Check this root for executable messages
		srcLogs, err := r.config.sourceLP.LogsDataWordRange(
			abihelpers.EventSignatures.SendRequested,
			r.config.onRamp.Address(),
			abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
			logpoller.EvmWord(unexpiredReport.Interval.Min),
			logpoller.EvmWord(unexpiredReport.Interval.Max),
			int(r.offchainConfig.SourceIncomingConfirmations),
		)
		if err != nil {
			return nil, err
		}
		if len(srcLogs) != int(unexpiredReport.Interval.Max-unexpiredReport.Interval.Min+1) {
			return nil, errors.Errorf("unexpected missing msgs in committed root %x have %d want %d", unexpiredReport.MerkleRoot, len(srcLogs), int(unexpiredReport.Interval.Max-unexpiredReport.Interval.Min+1))
		}
		// TODO: Reorg risk here? I.e. 1 message in a batch, we see its executed so we snooze forever,
		// then it gets reorged out and we'll never retry.
		executedMp, err := r.getExecutedSeqNrsInRange(unexpiredReport.Interval.Min, unexpiredReport.Interval.Max)
		if err != nil {
			return nil, err
		}

		r.lggr.Debugw("building next batch", "executedMp", len(executedMp))

		buildBatchDuration := time.Now()
		batch, allMessagesExecuted := r.buildBatch(srcLogs, executedMp, inflight, allowedTokenAmount,
			srcTokensPrices, destTokensPrices, destGasPrice)
		measureBatchBuildDuration(timestamp, time.Since(buildBatchDuration))

		// If all messages are already executed, snooze the root for the config.PermissionLessExecutionThresholdSeconds
		// so it will never be considered again.
		if allMessagesExecuted {
			r.lggr.Infof("Snoozing root %s forever since there are no executable txs anymore %v", hex.EncodeToString(unexpiredReport.MerkleRoot[:]), executedMp)
			r.snoozedRoots[unexpiredReport.MerkleRoot] = time.Now().Add(r.onchainConfig.PermissionLessExecutionThresholdDuration())
			incSkippedRequests(reasonAllExecuted)
			continue
		}
		if len(batch) != 0 {
			return batch, nil
		}
		r.snoozedRoots[unexpiredReport.MerkleRoot] = time.Now().Add(r.offchainConfig.RootSnoozeTime.Duration())
	}
	return []ObservedMessage{}, nil
}

// Updates the source to dest token mapping only if changes have happened.
// Uses the logPoller to look for PoolAdded and PoolRemoved logs and if
// it finds any, it will reload all tokens
func (r *ExecutionReportingPlugin) updateSourceToDestTokenMapping(ctx context.Context) error {
	r.srcToDstTokenMappingMu.RLock()
	lastPoolChangeBlock := r.srcToDstTokenMappingBlock
	r.srcToDstTokenMappingMu.RUnlock()

	poolEvents, err := r.config.destLP.LatestLogEventSigsAddrsWithConfs(
		lastPoolChangeBlock,
		[]common.Hash{abihelpers.EventSignatures.PoolAdded, abihelpers.EventSignatures.PoolRemoved},
		[]common.Address{r.config.offRamp.Address()},
		int(r.offchainConfig.DestIncomingConfirmations),
	)
	if err != nil {
		return err
	}

	// Update only if token pools have been added/removed
	if len(poolEvents) == 0 {
		return nil
	}
	highestBlockNumber := lastPoolChangeBlock
	for _, poolEvent := range poolEvents {
		if poolEvent.BlockNumber > highestBlockNumber {
			highestBlockNumber = poolEvent.BlockNumber
		}
	}

	newSrcToDestTokenMapping, err := generateSourceToDestTokenMapping(ctx, r.config.offRamp)
	if err != nil {
		return err
	}

	r.srcToDstTokenMappingMu.Lock()
	defer r.srcToDstTokenMappingMu.Unlock()
	r.srcToDstTokenMapping = newSrcToDestTokenMapping
	r.srcToDstTokenMappingBlock = highestBlockNumber + 1

	return nil
}

// Generates the source to dest token mapping based on the offRamp.
// NOTE: this queries the offRamp n+1 times, where n is the number of
// enabled tokens.
func generateSourceToDestTokenMapping(ctx context.Context, offRamp evm_2_evm_offramp.EVM2EVMOffRampInterface) (map[common.Address]common.Address, error) {
	srcToDstTokenMapping := make(map[common.Address]common.Address)
	sourceTokens, err := offRamp.GetSupportedTokens(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	for _, sourceToken := range sourceTokens {
		dst, err2 := offRamp.GetDestinationToken(&bind.CallOpts{Context: ctx}, sourceToken)
		if err2 != nil {
			return nil, err2
		}
		srcToDstTokenMapping[sourceToken] = dst
	}
	return srcToDstTokenMapping, nil
}

// Calculates a map that indicated whether a sequence number has already been executed
// before. It doesn't matter if the executed succeeded, since we don't retry previous
// attempts even if they failed.
func (r *ExecutionReportingPlugin) getExecutedSeqNrsInRange(min, max uint64) (map[uint64]struct{}, error) {
	executedLogs, err := r.config.destLP.IndexedLogsTopicRange(
		abihelpers.EventSignatures.ExecutionStateChanged,
		r.config.offRamp.Address(),
		abihelpers.EventSignatures.ExecutionStateChangedSequenceNumberIndex,
		logpoller.EvmWord(min),
		logpoller.EvmWord(max),
		int(r.offchainConfig.DestIncomingConfirmations),
	)
	if err != nil {
		return nil, err
	}
	executedMp := make(map[uint64]struct{})
	for _, executedLog := range executedLogs {
		exec, err := r.config.offRamp.ParseExecutionStateChanged(executedLog.GetGethLog())
		if err != nil {
			return nil, err
		}
		executedMp[exec.SequenceNumber] = struct{}{}
	}
	return executedMp, nil
}

// Builds a batch of transactions that can be executed, takes into account
// the available gas, rate limiting, execution state, nonce state, and
// profitability of execution.
func (r *ExecutionReportingPlugin) buildBatch(
	srcLogs []logpoller.Log,
	executedSeq map[uint64]struct{},
	inflight []InflightInternalExecutionReport,
	aggregateTokenLimit *big.Int,
	srcTokenPricesUSD map[common.Address]*big.Int,
	destTokenPricesUSD map[common.Address]*big.Int,
	execGasPriceEstimate *big.Int,
) (executableMessages []ObservedMessage, executedAllMessages bool) {
	r.srcToDstTokenMappingMu.RLock()
	defer r.srcToDstTokenMappingMu.RUnlock()

	inflightSeqNrs, inflightAggregateValue, maxInflightSenderNonces, err := r.inflight(inflight, destTokenPricesUSD, r.srcToDstTokenMapping)
	if err != nil {
		r.lggr.Errorw("Unexpected error computing inflight values", "err", err)
		return []ObservedMessage{}, false
	}
	availableGas := uint64(r.offchainConfig.BatchGasLimit)
	aggregateTokenLimit.Sub(aggregateTokenLimit, inflightAggregateValue)
	executedAllMessages = true
	expectedNonces := make(map[common.Address]uint64)
	for _, srcLog := range srcLogs {
		msg, err2 := r.config.onRamp.ParseCCIPSendRequested(gethtypes.Log{
			// Note this needs to change if we start indexing things.
			Topics: srcLog.GetTopics(),
			Data:   srcLog.Data,
		})
		if err2 != nil {
			r.lggr.Errorw("unable to parse message", "err", err2, "msg", msg)
			// Unable to parse so don't mark as executed
			executedAllMessages = false
			continue
		}
		lggr := r.lggr.With("messageID", hexutil.Encode(msg.Message.MessageId[:]))
		if _, executed := executedSeq[msg.Message.SequenceNumber]; executed {
			lggr.Infow("Skipping message already executed", "seqNr", msg.Message.SequenceNumber)
			continue
		}
		executedAllMessages = false
		if _, inflight := inflightSeqNrs[msg.Message.SequenceNumber]; inflight {
			lggr.Infow("Skipping message already inflight", "seqNr", msg.Message.SequenceNumber)
			continue
		}
		if _, ok := expectedNonces[msg.Message.Sender]; !ok {
			// First message in batch, need to populate expected nonce
			if maxInflight, ok := maxInflightSenderNonces[msg.Message.Sender]; ok {
				// Sender already has inflight nonce, populate from there
				expectedNonces[msg.Message.Sender] = maxInflight + 1
			} else {
				// Nothing inflight take from chain.
				// Chain holds existing nonce.
				nonce, err := r.config.offRamp.GetSenderNonce(nil, msg.Message.Sender)
				if err != nil {
					lggr.Errorw("unable to get sender nonce", "err", err)
					continue
				}
				expectedNonces[msg.Message.Sender] = nonce + 1
			}
		}
		// Check expected nonce is valid
		if msg.Message.Nonce != expectedNonces[msg.Message.Sender] {
			lggr.Warnw("Skipping message invalid nonce", "have", msg.Message.Nonce, "want", expectedNonces[msg.Message.Sender])
			continue
		}
		msgValue, err := aggregateTokenValue(destTokenPricesUSD, r.srcToDstTokenMapping, msg.Message.TokenAmounts)
		if err != nil {
			lggr.Errorw("Skipping message unable to compute aggregate value", "err", err)
			continue
		}
		// if token limit is smaller than message value skip message
		if aggregateTokenLimit.Cmp(msgValue) == -1 {
			lggr.Warnw("token limit is smaller than message value", "aggregateTokenLimit", aggregateTokenLimit.String(), "msgValue", msgValue.String())
			continue
		}
		// Fee boosting
		execCostUsd := computeExecCost(msg.Message.GasLimit, execGasPriceEstimate, destTokenPricesUSD[r.destWrappedNative])
		// calculating the source chain fee, dividing by 1e18 for denomination.
		// For example:
		// FeeToken=link; FeeTokenAmount=1e17 i.e. 0.1 link, price is 6e18 USD/link (1 USD = 1e18),
		// availableFee is 1e17*6e18/1e18 = 6e17 = 0.6 USD
		availableFee := big.NewInt(0).Mul(msg.Message.FeeTokenAmount, srcTokenPricesUSD[msg.Message.FeeToken])
		availableFee = availableFee.Div(availableFee, big.NewInt(1e18))
		availableFeeUsd := waitBoostedFee(time.Since(srcLog.BlockTimestamp), availableFee, r.offchainConfig.RelativeBoostPerWaitHour)
		if availableFeeUsd.Cmp(execCostUsd) < 0 {
			lggr.Infow("Insufficient remaining fee", "availableFeeUsd", availableFeeUsd, "execCostUsd", execCostUsd,
				"srcBlockTimestamp", srcLog.BlockTimestamp, "waitTime", time.Since(srcLog.BlockTimestamp), "boost", r.offchainConfig.RelativeBoostPerWaitHour)
			continue
		}

		messageMaxGas := msg.Message.GasLimit.Uint64() + maxGasOverHeadGas(len(srcLogs), len(msg.Message.Data), len(msg.Message.TokenAmounts))
		// Check sufficient gas in batch
		if availableGas < messageMaxGas {
			lggr.Infow("Insufficient remaining gas in batch limit", "availableGas", availableGas, "messageMaxGas", messageMaxGas)
			continue
		}
		availableGas -= messageMaxGas
		aggregateTokenLimit.Sub(aggregateTokenLimit, msgValue)

		var tokenData [][]byte

		// TODO add attestation data for USDC here
		for range msg.Message.TokenAmounts {
			tokenData = append(tokenData, []byte{})
		}

		lggr.Infow("Adding msg to batch", "seqNum", msg.Message.SequenceNumber, "nonce", msg.Message.Nonce)
		executableMessages = append(executableMessages, ObservedMessage{
			SeqNr:     msg.Message.SequenceNumber,
			TokenData: tokenData,
		})
		expectedNonces[msg.Message.Sender] = msg.Message.Nonce + 1
	}
	return executableMessages, executedAllMessages
}

func aggregateTokenValue(destTokenPricesUSD map[common.Address]*big.Int, srcToDst map[common.Address]common.Address, tokensAndAmount []evm_2_evm_onramp.ClientEVMTokenAmount) (*big.Int, error) {
	sum := big.NewInt(0)
	for i := 0; i < len(tokensAndAmount); i++ {
		price, ok := destTokenPricesUSD[srcToDst[tokensAndAmount[i].Token]]
		if !ok {
			return nil, errors.Errorf("do not have price for src token %v", tokensAndAmount[i].Token)
		}
		sum.Add(sum, new(big.Int).Quo(new(big.Int).Mul(price, tokensAndAmount[i].Amount), big.NewInt(1e18)))
	}
	return sum, nil
}

func (r *ExecutionReportingPlugin) parseSeqNr(log logpoller.Log) (uint64, error) {
	s, err := r.config.onRamp.ParseCCIPSendRequested(log.ToGethLog())
	if err != nil {
		return 0, err
	}
	return s.Message.SequenceNumber, nil
}

// Assumes non-empty report. Messages to execute can span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *ExecutionReportingPlugin) buildReport(ctx context.Context, lggr logger.Logger, observedMessages []ObservedMessage) ([]byte, error) {
	if err := validateSeqNumbers(ctx, r.config.commitStore, observedMessages); err != nil {
		return nil, err
	}
	commitReport, err := getCommitReportForSeqNum(r.config.destLP, r.config.commitStore, observedMessages[0].SeqNr)
	if err != nil {
		return nil, err
	}
	lggr.Infow("Building execution report", "observations", observedMessages, "merkleRoot", hexutil.Encode(commitReport.MerkleRoot[:]), "report", commitReport)

	msgsInRoot, leaves, tree, err := getProofData(ctx, lggr, r.config.leafHasher, r.parseSeqNr, r.config.onRamp.Address(), r.config.sourceLP, commitReport.Interval)
	if err != nil {
		return nil, err
	}

	// cap messages which fits MaxExecutionReportLength (after serialized)
	capped := sort.Search(len(observedMessages), func(i int) bool {
		report, _ := buildExecutionReportForMessages(msgsInRoot, leaves, tree, commitReport.Interval, observedMessages[:i+1])
		var encoded []byte
		encoded, err = abihelpers.EncodeExecutionReport(report)
		if err != nil {
			// false makes Search keep looking to the right, always including any "erroring" ObservedMessage and allowing us to detect in the bottom
			return false
		}
		return len(encoded) > MaxObservationLength
	})
	if err != nil {
		return nil, err
	}

	execReport, hashes := buildExecutionReportForMessages(msgsInRoot, leaves, tree, commitReport.Interval, observedMessages[:capped])
	encodedReport, err := abihelpers.EncodeExecutionReport(execReport)
	if err != nil {
		return nil, err
	}

	if capped < len(observedMessages) {
		lggr.Warnf(
			"Capping report to fit MaxExecutionReportLength: msgsCount %d -> %d, bytes %d, bytesLimit %d",
			len(observedMessages), capped, len(encodedReport), MaxExecutionReportLength,
		)
	}

	// Double check this verifies before sending.
	res, err := r.config.commitStore.Verify(&bind.CallOpts{Context: ctx}, hashes, execReport.Proofs, execReport.ProofFlagBits)
	if err != nil {
		lggr.Errorw("Unable to call verify", "observations", observedMessages[:capped], "root", commitReport.MerkleRoot[:], "seqRange", commitReport.Interval, "err", err)
		return nil, err
	}
	// No timestamp, means failed to verify root.
	if res.Cmp(big.NewInt(0)) == 0 {
		root := tree.Root()
		lggr.Errorf("Root does not verify for messages: %v, our inner root 0x%x", observedMessages[:capped], root)
		return nil, errors.New("root does not verify")
	}
	return encodedReport, nil
}

func (r *ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	parsableObservations := getParsableObservations[ExecutionObservation](lggr, observations)
	// Need at least F+1 observations
	if len(parsableObservations) <= r.F {
		lggr.Tracew("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}

	observedMessages := calculateObservedMessagesConsensus(lggr, parsableObservations, r.F)
	if len(observedMessages) == 0 {
		return false, nil, nil
	}

	report, err := r.buildReport(ctx, lggr, observedMessages)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report", "onRampAddr", r.config.onRamp.Address(), "observations", observedMessages)
	return true, report, nil
}

type seqNumTally struct {
	Tally     int
	TokenData [][]byte
}

func calculateObservedMessagesConsensus(lggr logger.Logger, observations []ExecutionObservation, f int) []ObservedMessage {
	tally := make(map[uint64]seqNumTally)
	for _, obs := range observations {
		for _, message := range obs.Messages {
			if val, ok := tally[message.SeqNr]; ok {
				// If we've already seen the seqNum we check if the token data is the same
				if !reflect.DeepEqual(message.TokenData, val.TokenData) {
					lggr.Warnf("Nodes reported different offchain token data [%v] [%v]", message.TokenData, val.TokenData)
				}
				val.Tally++
				tally[message.SeqNr] = val
				continue
			}
			// If we have not seen the seqNum we save a tally with the token data
			tally[message.SeqNr] = seqNumTally{
				Tally:     1,
				TokenData: message.TokenData,
			}
		}
	}
	var finalSequenceNumbers []ObservedMessage
	for seqNr, tallyInfo := range tally {
		// Note spec deviation - I think it's ok to rely on the batch builder for
		// capping the number of messages vs capping in two places/ways?
		if tallyInfo.Tally > f {
			finalSequenceNumbers = append(finalSequenceNumbers, ObservedMessage{
				SeqNr:     seqNr,
				TokenData: tallyInfo.TokenData,
			})
		}
	}
	// buildReport expects sorted sequence numbers (tally map is non-deterministic).
	sort.Slice(finalSequenceNumbers, func(i, j int) bool {
		return finalSequenceNumbers[i].SeqNr < finalSequenceNumbers[j].SeqNr
	})
	return finalSequenceNumbers
}

func (r *ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	seqNrs, encMsgs, err := MessagesFromExecutionReport(report)
	if err != nil {
		lggr.Errorw("unable to decode report", "err", err)
		return false, nil
	}
	lggr.Infof("Seq nums %v", seqNrs)
	// If the first message is executed already, this execution report is stale, and we do not accept it.
	stale, err := r.isStaleReport(seqNrs)
	if err != nil {
		return false, err
	}
	if stale {
		return false, nil
	}
	// Else just assume in flight
	if err = r.inflightReports.add(lggr, seqNrs, encMsgs); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	if isCommitStoreDownNow(ctx, r.config.lggr, r.config.commitStore) {
		return false, nil
	}
	seqNrs, _, err := MessagesFromExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for tx manager,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStaleReport(seqNrs)
	return !stale, err
}

func (r *ExecutionReportingPlugin) isStaleReport(seqNrs []uint64) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	msgState, err := r.config.offRamp.GetExecutionState(nil, seqNrs[0])
	if err != nil {
		// TODO: do we need to check for not present error?
		return true, err
	}
	if state := abihelpers.MessageExecutionState(msgState); state == abihelpers.ExecutionStateFailure || state == abihelpers.ExecutionStateSuccess {
		return true, nil
	}

	return false, nil
}

func (r *ExecutionReportingPlugin) Close() error {
	return nil
}

func (r *ExecutionReportingPlugin) inflight(
	inflight []InflightInternalExecutionReport,
	destTokenPrices map[common.Address]*big.Int,
	srcToDst map[common.Address]common.Address,
) (map[uint64]struct{}, *big.Int, map[common.Address]uint64, error) {
	inflightSeqNrs := make(map[uint64]struct{})
	inflightAggregateValue := big.NewInt(0)
	maxInflightSenderNonces := make(map[common.Address]uint64)
	for _, rep := range inflight {
		for _, seqNr := range rep.seqNrs {
			inflightSeqNrs[seqNr] = struct{}{}
		}
		for _, encMsg := range rep.encMessages {
			msg, err := r.config.onRamp.ParseCCIPSendRequested(gethtypes.Log{
				// Note this needs to change if we start indexing things.
				Topics: []common.Hash{abihelpers.EventSignatures.SendRequested},
				Data:   encMsg,
			})
			if err != nil {
				return nil, nil, nil, err
			}
			msgValue, err := aggregateTokenValue(destTokenPrices, srcToDst, msg.Message.TokenAmounts)
			if err != nil {
				return nil, nil, nil, err
			}
			inflightAggregateValue.Add(inflightAggregateValue, msgValue)
			maxInflightSenderNonce, ok := maxInflightSenderNonces[msg.Message.Sender]
			if !ok || msg.Message.Nonce > maxInflightSenderNonce {
				maxInflightSenderNonces[msg.Message.Sender] = msg.Message.Nonce
			}
		}
	}
	return inflightSeqNrs, inflightAggregateValue, maxInflightSenderNonces, nil
}

// getTokensPrices returns token prices of the given price registry,
// results include feeTokens and passed-in tokens
// price values are USD per 1e18 of smallest token denomination, in base units 1e18 (e.g. 5$ = 5e18 USD per 1e18 units).
// this function is used for price registry of both source and destination chains.
func getTokensPrices(ctx context.Context, priceRegistry price_registry.PriceRegistryInterface, tokens []common.Address) (map[common.Address]*big.Int, error) {
	prices := make(map[common.Address]*big.Int)

	// TODO cache and only check on changing config
	feeTokens, err := priceRegistry.GetFeeTokens(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "could not get source fee tokens")
	}

	wantedTokens := append(feeTokens, tokens...)
	wantedPrices, err := priceRegistry.GetTokenPrices(&bind.CallOpts{Context: ctx}, wantedTokens)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get token prices of %v", wantedTokens)
	}
	for i, token := range wantedTokens {
		prices[token] = wantedPrices[i].Value
	}

	return prices, nil
}

func getUnexpiredCommitReports(dstLogPoller logpoller.LogPoller, commitStore commit_store.CommitStoreInterface, permissionExecutionThreshold time.Duration) ([]commit_store.CommitStoreCommitReport, error) {
	logs, err := dstLogPoller.LogsCreatedAfter(abihelpers.EventSignatures.ReportAccepted, commitStore.Address(), time.Now().Add(-permissionExecutionThreshold))
	if err != nil {
		return nil, err
	}
	var reports []commit_store.CommitStoreCommitReport
	for _, log := range logs {
		reportAccepted, err := commitStore.ParseReportAccepted(gethtypes.Log{
			Topics: log.GetTopics(),
			Data:   log.Data,
		})
		if err != nil {
			return nil, err
		}
		reports = append(reports, reportAccepted.Report)
	}
	return reports, nil
}

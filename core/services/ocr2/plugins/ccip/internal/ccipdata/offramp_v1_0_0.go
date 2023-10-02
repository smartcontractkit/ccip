package ccipdata

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/logpollerutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	EXEC_EXECUTION_STATE_CHANGES = "Exec execution state changes"
	EXEC_TOKEN_POOL_ADDED        = "Token pool added"
	EXEC_TOKEN_POOL_REMOVED      = "Token pool removed"
)

var _ OffRampReader = &OffRampV1_0_0{}

type OffRampV1_0_0 struct {
	offRamp             *evm_2_evm_offramp.EVM2EVMOffRamp
	onRamp              *evm_2_evm_onramp_1_0_0.EVM2EVMOnRamp
	addr                common.Address
	lp                  logpoller.LogPoller
	lggr                logger.Logger
	ec                  client.Client
	filters             []logpoller.Filter
	estimator           gas.EvmFeeEstimator
	gasPriceEstimator   prices.GasPriceEstimatorExec
	executionReportArgs abi.Arguments
}

func (o *OffRampV1_0_0) DecodeExecutionReport(report []byte) (ExecReport, error) {
	unpacked, err := o.executionReportArgs.Unpack(report)
	if err != nil {
		return ExecReport{}, err
	}
	if len(unpacked) == 0 {
		return ExecReport{}, errors.New("assumptionViolation: expected at least one element")
	}
	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		Messages []struct {
			SourceChainSelector uint64         `json:"sourceChainSelector"`
			Sender              common.Address `json:"sender"`
			Receiver            common.Address `json:"receiver"`
			SequenceNumber      uint64         `json:"sequenceNumber"`
			GasLimit            *big.Int       `json:"gasLimit"`
			Strict              bool           `json:"strict"`
			Nonce               uint64         `json:"nonce"`
			FeeToken            common.Address `json:"feeToken"`
			FeeTokenAmount      *big.Int       `json:"feeTokenAmount"`
			Data                []uint8        `json:"data"`
			TokenAmounts        []struct {
				Token  common.Address `json:"token"`
				Amount *big.Int       `json:"amount"`
			} `json:"tokenAmounts"`
			SourceTokenData [][]byte `json:"sourceTokenData"`
			MessageId       [32]byte `json:"messageId"`
		} `json:"messages"`
		OffchainTokenData [][][]byte  `json:"offchainTokenData"`
		Proofs            [][32]uint8 `json:"proofs"`
		ProofFlagBits     *big.Int    `json:"proofFlagBits"`
	})
	if !ok {
		return ExecReport{}, fmt.Errorf("got %T", unpacked[0])
	}
	var messages []EVM2EVMMessage
	for _, msg := range erStruct.Messages {
		var tokensAndAmounts []evm_2_evm_offramp.ClientEVMTokenAmount
		for _, tokenAndAmount := range msg.TokenAmounts {
			tokensAndAmounts = append(tokensAndAmounts, evm_2_evm_offramp.ClientEVMTokenAmount{
				Token:  tokenAndAmount.Token,
				Amount: tokenAndAmount.Amount,
			})
		}
		// TODO
		messages = append(messages, EVM2EVMMessage{
			SequenceNumber:      msg.SequenceNumber,
			GasLimit:            nil,
			Nonce:               0,
			MessageId:           Hash{},
			SourceChainSelector: 0,
			Sender:              common.Address{},
			Receiver:            common.Address{},
			Strict:              false,
			FeeToken:            common.Address{},
			FeeTokenAmount:      nil,
			Data:                nil,
			TokenAmounts:        nil,
			SourceTokenData:     nil,
			Hash:                Hash{},
		})
	}

	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	return ExecReport{
		Messages:          messages,
		OffchainTokenData: erStruct.OffchainTokenData,
		Proofs:            erStruct.Proofs,
		ProofFlagBits:     new(big.Int).SetBytes(erStruct.ProofFlagBits.Bytes()),
	}, nil
}

func (o *OffRampV1_0_0) OffchainConfig() ExecOffchainConfig {
	//TODO implement me
	panic("implement me")
}

func (o *OffRampV1_0_0) OnchainConfig() ExecOnchainConfig {
	//TODO implement me
	panic("implement me")
}

func (o *OffRampV1_0_0) GasPriceEstimator() prices.GasPriceEstimatorExec {
	return o.gasPriceEstimator
}

func (o *OffRampV1_0_0) Address() common.Address {
	return o.addr
}

func (o *OffRampV1_0_0) ConfigChanged(onchainConfig []byte, offchainConfig []byte) (common.Address, common.Address, error) {
	onchainConfigParsed, err := abihelpers.DecodeAbiStruct[ExecOnchainConfigV1_0_0](onchainConfig)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	offchainConfigParsed, err := ccipconfig.DecodeOffchainConfig[ExecOffchainConfig](offchainConfig)
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
	o.gasPriceEstimator = prices.NewExecGasPriceEstimator(o.estimator, big.NewInt(int64(offchainConfigParsed.MaxGasPrice)), 0)
	o.lggr.Infow("Starting exec plugin",
		"offchainConfig", onchainConfigParsed,
		"onchainConfig", offchainConfigParsed)
	return onchainConfigParsed.PriceRegistry, destWrappedNative, nil
}

func (o *OffRampV1_0_0) GetDestinationTokens(ctx context.Context) ([]common.Address, error) {
	return o.offRamp.GetDestinationTokens(&bind.CallOpts{Context: ctx})
}

func (o *OffRampV1_0_0) Close(qopts ...pg.QOpt) error {
	return logpollerutil.UnregisterLpFilters(o.lp, o.filters, qopts...)
}

func (o *OffRampV1_0_0) GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[ExecutionStateChanged], error) {
	logs, err := o.lp.IndexedLogsTopicRange(
		abihelpers.EventSignatures.ExecutionStateChanged,
		o.addr,
		abihelpers.EventSignatures.ExecutionStateChangedSequenceNumberIndex,
		logpoller.EvmWord(seqNumMin),
		logpoller.EvmWord(seqNumMax),
		confs,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return parseLogs[ExecutionStateChanged](
		logs,
		o.lggr,
		func(log types.Log) (*ExecutionStateChanged, error) {
			sc, err := o.offRamp.ParseExecutionStateChanged(log)
			if err != nil {
				return nil, err
			}
			return &ExecutionStateChanged{SequenceNumber: sc.SequenceNumber}, nil
		},
	)
}

func (o *OffRampV1_0_0) EncodeExecutionReport(report ExecReport) ([]byte, error) {
	var msgs []evm_2_evm_offramp.InternalEVM2EVMMessage
	for _, msg := range report.Messages {
		var ta []evm_2_evm_offramp.ClientEVMTokenAmount
		for _, tokenAndAmount := range msg.TokenAmounts {
			ta = append(ta, evm_2_evm_offramp.ClientEVMTokenAmount{
				Token:  tokenAndAmount.Token,
				Amount: tokenAndAmount.Amount,
			})
		}
		msgs = append(msgs, evm_2_evm_offramp.InternalEVM2EVMMessage{
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
			SourceTokenData:     [][]byte{}, // Always empty in 1.0
			MessageId:           msg.MessageId,
		})
	}

	rep := evm_2_evm_offramp.InternalExecutionReport{
		Messages:          msgs,
		OffchainTokenData: report.OffchainTokenData,
		Proofs:            report.Proofs,
		ProofFlagBits:     report.ProofFlagBits,
	}
	return o.executionReportArgs.PackValues([]interface{}{&rep})
}

func NewOffRampV1_0_0(lggr logger.Logger, addr common.Address, ec client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator, srcClient client.Client) (*OffRampV1_0_0, error) {
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(addr, ec)
	if err != nil {
		return nil, err
	}
	offRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_offramp.EVM2EVMOffRampABI))
	if err != nil {
		panic(err)
	}
	manuallyExecuteMethod, ok := offRampABI.Methods["manuallyExecute"]
	if !ok {
		panic("missing event 'manuallyExecute'")
	}
	executionReportArgs := manuallyExecuteMethod.Inputs[:1]
	var filters = []logpoller.Filter{
		{
			Name:      logpoller.FilterName(EXEC_EXECUTION_STATE_CHANGES, addr.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.ExecutionStateChanged},
			Addresses: []common.Address{addr},
		},
		{
			Name:      logpoller.FilterName(EXEC_TOKEN_POOL_ADDED, addr.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.PoolAdded},
			Addresses: []common.Address{addr},
		},
		{
			Name:      logpoller.FilterName(EXEC_TOKEN_POOL_REMOVED, addr.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.PoolRemoved},
			Addresses: []common.Address{addr},
		},
	}
	if err := logpollerutil.RegisterLpFilters(lp, filters); err != nil {
		return nil, err
	}
	s, err := offRamp.GetStaticConfig(nil)
	if err != nil {
		return nil, err
	}
	// Must point to the correct version onramp
	contractType, version, err := ccipconfig.TypeAndVersion(s.OnRamp, srcClient)
	if err != nil {
		return nil, err
	}
	if contractType != ccipconfig.EVM2EVMOnRamp || version.String() != "1.0.0" {
		return nil, errors.Errorf("offramp points to invalid onramp %v expected 1.0.0", version.String())
	}
	onRamp, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(s.OnRamp, srcClient)
	if err != nil {
		return nil, err
	}
	return &OffRampV1_0_0{offRamp: offRamp, onRamp: onRamp, addr: addr, lggr: lggr, lp: lp, filters: filters, estimator: estimator, executionReportArgs: executionReportArgs}, nil
}

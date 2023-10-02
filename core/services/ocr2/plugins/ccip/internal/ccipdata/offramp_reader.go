package ccipdata

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

// Do not change the JSON format of this struct without consulting with
// the RDD people first.
type ExecOffchainConfig struct {
	SourceFinalityDepth         uint32
	DestOptimisticConfirmations uint32
	DestFinalityDepth           uint32
	BatchGasLimit               uint32
	RelativeBoostPerWaitHour    float64
	MaxGasPrice                 uint64
	InflightCacheExpiry         models.Duration
	RootSnoozeTime              models.Duration
}

func (c ExecOffchainConfig) Validate() error {
	if c.SourceFinalityDepth == 0 {
		return errors.New("must set SourceFinalityDepth")
	}
	if c.DestFinalityDepth == 0 {
		return errors.New("must set DestFinalityDepth")
	}
	if c.DestOptimisticConfirmations == 0 {
		return errors.New("must set DestOptimisticConfirmations")
	}
	if c.BatchGasLimit == 0 {
		return errors.New("must set BatchGasLimit")
	}
	if c.RelativeBoostPerWaitHour == 0 {
		return errors.New("must set RelativeBoostPerWaitHour")
	}
	if c.MaxGasPrice == 0 {
		return errors.New("must set MaxGasPrice")
	}
	if c.InflightCacheExpiry.Duration() == 0 {
		return errors.New("must set InflightCacheExpiry")
	}
	if c.RootSnoozeTime.Duration() == 0 {
		return errors.New("must set RootSnoozeTime")
	}

	return nil
}

type ExecOnchainConfig struct {
	PermissionLessExecutionThresholdSeconds time.Duration
}

type ExecOnchainConfigV1_0_0 evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig

func (d ExecOnchainConfigV1_0_0) AbiString() string {
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

func (d ExecOnchainConfigV1_0_0) Validate() error {
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

func (d ExecOnchainConfigV1_0_0) PermissionLessExecutionThresholdDuration() time.Duration {
	return time.Duration(d.PermissionLessExecutionThresholdSeconds) * time.Second
}

type ExecutionStateChanged struct {
	SequenceNumber uint64
}

type ExecReport struct {
	Messages          []internal.EVM2EVMMessage
	OffchainTokenData [][][]byte
	Proofs            [][32]byte
	ProofFlagBits     *big.Int
}

type OffRampReader interface {
	// Will error if messages are not a compatible verion
	EncodeExecutionReport(report ExecReport) ([]byte, error)

	DecodeExecutionReport(report []byte) (ExecReport, error)

	// GetExecutionStateChangesBetweenSeqNums returns all the execution state change events for the provided message sequence numbers (inclusive).
	GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[ExecutionStateChanged], error)

	GetDestinationTokens(ctx context.Context) ([]common.Address, error)

	Address() common.Address

	// Notifies the reader that the config has changed onchain
	ConfigChanged(onchainConfig []byte, offchainConfig []byte) (common.Address, common.Address, error)

	OffchainConfig() ExecOffchainConfig

	OnchainConfig() ExecOnchainConfig

	GasPriceEstimator() prices.GasPriceEstimatorExec

	Close(qopts ...pg.QOpt) error
}

// MessageExecutionState defines the execution states of CCIP messages.
type MessageExecutionState uint8

const (
	ExecutionStateUntouched MessageExecutionState = iota
	ExecutionStateInProgress
	ExecutionStateSuccess
	ExecutionStateFailure
)

func NewOffRampReader(lggr logger.Logger, addr common.Address, srcClient, destClient client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator) (OffRampReader, error) {
	_, version, err := ccipconfig.TypeAndVersion(addr, destClient)
	if err != nil {
		return nil, err
	}
	switch version.String() {
	case "1.0.0", "1.1.0":
		return NewOffRampV1_0_0(lggr, addr, destClient, lp, estimator, srcClient)
	case "1.2.0":
		return NewOffRampV1_2_0(lggr, addr, destClient, lp, estimator, srcClient)
	default:
		return nil, errors.Errorf("unsupported offramp verison %v", version.String())
	}
}

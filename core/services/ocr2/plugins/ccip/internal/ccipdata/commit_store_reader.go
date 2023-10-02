package ccipdata

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type CommitStoreInterval struct {
	Min, Max uint64
}

type CommitStoreReport struct {
	TokenPrices []TokenPrice
	GasPrices   []GasPrice
	Interval    CommitStoreInterval
	MerkleRoot  [32]byte
}

// Common to all versions
type CommitOnchainConfig commit_store.CommitStoreDynamicConfig

func (d CommitOnchainConfig) AbiString() string {
	return `
	[
		{
			"components": [
				{"name": "priceRegistry", "type": "address"}
			],
			"type": "tuple"
		}
	]`
}

func (d CommitOnchainConfig) Validate() error {
	if d.PriceRegistry == (common.Address{}) {
		return errors.New("must set Price Registry address")
	}
	return nil
}

type OffchainConfig struct {
	SourceFinalityDepth    uint32
	GasPriceDeviationPPB   uint32
	GasPriceHeartBeat      time.Duration
	TokenPriceDeviationPPB uint32
	TokenPriceHeartBeat    time.Duration
	InflightCacheExpiry    time.Duration
	DestFinalityDepth      uint32
}

type CommitStoreReader interface {
	GetExpectedNextSequenceNumber(context context.Context) (uint64, error)

	GetLatestPriceEpochAndRound(context context.Context) (uint64, error)

	// GetAcceptedCommitReportsGteSeqNum returns all the accepted commit reports that have sequence number greater than or equal to the provided.
	GetAcceptedCommitReportsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[CommitStoreReport], error)

	// GetAcceptedCommitReportsGteTimestamp returns all the commit reports with timestamp greater than or equal to the provided.
	GetAcceptedCommitReportsGteTimestamp(ctx context.Context, ts time.Time, confs int) ([]Event[CommitStoreReport], error)

	IsDown(ctx context.Context) bool

	IsBlessed(ctx context.Context, root [32]byte) (bool, error)

	// Notifies the reader that the config has changed onchain
	ConfigChanged(onchainConfig []byte, offchainConfig []byte) (common.Address, error)

	OffchainConfig() OffchainConfig

	GasPriceEstimator() prices.GasPriceEstimatorCommit

	EncodeCommitReport(report CommitStoreReport) ([]byte, error)

	DecodeCommitReport(report []byte) (CommitStoreReport, error)

	Verify(ctx context.Context, report ExecReport) bool

	Close(qopts ...pg.QOpt) error
}

func NewCommitStoreReader(lggr logger.Logger, address common.Address, ec client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator) (CommitStoreReader, error) {
	contractType, version, err := ccipconfig.TypeAndVersion(address, ec)
	if err != nil {
		return nil, errors.Errorf("expected %v got %v", ccipconfig.EVM2EVMOnRamp, contractType)
	}
	switch version.String() {
	case "1.0.0", "1.1.0":
		return NewCommitStoreV1_0_0(lggr, address, ec, lp, estimator)
	case "1.2.0":
		return NewCommitStoreV1_2_0(lggr, address, ec, lp, estimator)
	default:
		return nil, errors.Errorf("got unexpected version %v", version.String())
	}
}

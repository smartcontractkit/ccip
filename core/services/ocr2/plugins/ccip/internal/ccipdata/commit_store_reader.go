package ccipdata

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type Interval struct {
	Min, Max uint64
}

type CommitStoreReport struct {
	TokenPrices []TokenPrice
	GasPrices   []GasPrice
	Interval    Interval
	MerkleRoot  [32]byte
}

type CommitStoreReader interface {
	GetExpectedNextSequenceNumber(context context.Context) (uint64, error)

	GetLatestPriceEpochAndRound(context context.Context) (uint64, error)

	// GetAcceptedCommitReportsGteSeqNum returns all the accepted commit reports that have sequence number greater than or equal to the provided.
	GetAcceptedCommitReportsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[CommitStoreReport], error)

	// GetAcceptedCommitReportsGteTimestamp returns all the commit reports with timestamp greater than or equal to the provided.
	GetAcceptedCommitReportsGteTimestamp(ctx context.Context, ts time.Time, confs int) ([]Event[CommitStoreReport], error)

	IsDown(ctx context.Context) bool

	DestPriceRegistryFromOnC() common.Address

	Verify(ctx context.Context, report ExecReport) bool

	Close(qopts ...pg.QOpt) error
}

func NewCommitStoreReader(lggr logger.Logger, address common.Address, ec client.Client, lp logpoller.LogPoller) (CommitStoreReader, error) {
	contractType, version, err := ccipconfig.TypeAndVersion(address, ec)
	if err != nil {
		return nil, errors.Errorf("expected %v got %v", ccipconfig.EVM2EVMOnRamp, contractType)
	}
	switch version.String() {
	case "1.0.0":
		return NewCommitStoreV1_0_0(lggr, address, ec, lp)
	default:
		return nil, errors.Errorf("got unexpected version %v", version.String())
	}
}

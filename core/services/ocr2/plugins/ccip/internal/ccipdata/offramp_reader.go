package ccipdata

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type ExecutionStateChanged struct {
	SequenceNumber uint64
}

type ExecReport struct {
	Messages          []EVM2EVMMessage
	OffchainTokenData [][][]byte
	Proofs            [][32]byte
	ProofFlagBits     *big.Int
}

type OffRampReader interface {
	// Will error if messages are not a compatible verion
	EncodeExecutionReport(report ExecReport) ([]byte, error)

	// GetExecutionStateChangesBetweenSeqNums returns all the execution state change events for the provided message sequence numbers (inclusive).
	GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[ExecutionStateChanged], error)

	GetDestinationTokens(ctx context.Context) ([]common.Address, error)

	Address() common.Address
	Close(qopts ...pg.QOpt) error
}

func NewOffRampReader(lggr logger.Logger, addr common.Address, ec client.Client, lp logpoller.LogPoller) (OffRampReader, error) {
	_, version, err := ccipconfig.TypeAndVersion(addr, ec)
	if err != nil {
		return nil, err
	}
	switch version.String() {
	case "1.0.0", "1.1.0", "1.2.0":
		return NewOffRampV1_0_0(lggr, addr, ec, lp)
	default:
		return nil, errors.Errorf("unsupported offramp verison %v", version.String())
	}
}

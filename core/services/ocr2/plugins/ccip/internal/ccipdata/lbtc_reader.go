package ccipdata

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

// TODO: Implement lbtc token reader
type LBTCReader interface {
}

type LBTCReaderImpl struct {
}

func NewLBTCReader(lggr logger.Logger, jobID string, transmitter common.Address, lp logpoller.LogPoller, registerFilters bool) (*LBTCReaderImpl, error) {
	return &LBTCReaderImpl{}, nil
}

func CloseLBTCReader(lggr logger.Logger, jobID string, transmitter common.Address, lp logpoller.LogPoller) error {
	return nil
}

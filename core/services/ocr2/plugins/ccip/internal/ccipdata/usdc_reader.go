package ccipdata

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

const (
	MESSAGE_SENT_FILTER_NAME = "USDC message sent"
)

//go:generate mockery --quiet --name USDCReader --output . --filename usdc_reader_mock.go --inpackage --case=underscore
type USDCReader interface {
	// GetLastUSDCMessagePriorToLogIndexInTx returns the last USDC message that was sent before the provided log index in the given transaction.
	GetLastUSDCMessagePriorToLogIndexInTx(ctx context.Context, logIndex int64, txHash common.Hash) ([]byte, error)
	Close() error
}

type USDCReaderImpl struct {
	usdcMessageSent common.Hash
	lp              logpoller.LogPoller
	filterName      string
}

func (u *USDCReaderImpl) Close() error {
	return u.lp.UnregisterFilter(u.filterName)
}

func (u *USDCReaderImpl) GetLastUSDCMessagePriorToLogIndexInTx(ctx context.Context, logIndex int64, txHash common.Hash) ([]byte, error) {
	logs, err := u.lp.IndexedLogsByTxHash(
		u.usdcMessageSent,
		txHash,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	for i := range logs {
		current := logs[len(logs)-i-1]
		if current.LogIndex < logIndex {
			return current.Data, nil
		}
	}
	return nil, errors.Errorf("no USDC message found prior to log index %d in tx %s", logIndex, txHash.Hex())
}

func NewUSDCReader(transmitter common.Address, lp logpoller.LogPoller) (*USDCReaderImpl, error) {
	filterName := logpoller.FilterName(MESSAGE_SENT_FILTER_NAME, transmitter.Hex())
	eventSig := utils.Keccak256Fixed([]byte("MessageSent(bytes)"))
	if err := lp.RegisterFilter(logpoller.Filter{
		Name:      filterName,
		EventSigs: []common.Hash{eventSig},
		Addresses: []common.Address{transmitter},
	}); err != nil {
		return nil, err
	}
	return &USDCReaderImpl{
		lp:              lp,
		usdcMessageSent: eventSig,
		filterName:      filterName,
	}, nil
}

package tokendata

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
)

var (
	ErrNotReady = errors.New("token data not ready")
)

// Reader is an interface for fetching offchain token data
type Reader interface {
	// ReadTokenData returns the attestation bytes if ready, and throws an error if not ready.
	ReadTokenData(ctx context.Context, seqNum uint64, logIndex uint, txHash common.Hash) (tokenData []byte, err error)

	// GetSourceLogPollerFilters returns the filters that should be used for the source chain log poller
	GetSourceLogPollerFilters() []logpoller.Filter

	// GetSourceToken returns the token address on the source chain
	GetSourceToken() common.Address
}

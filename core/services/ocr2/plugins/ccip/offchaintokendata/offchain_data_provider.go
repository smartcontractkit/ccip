package offchaintokendata

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
)

// Provider is an interface for fetching offchain token data
type Provider interface {
	// IsTokenDataReady returns true if the attestation for the given sequence number is complete
	// and returns the attestation bytes if it is complete.
	// Note: this function can be called many times, the implementation should cache the result.
	IsTokenDataReady(ctx context.Context, seqNum uint64) (ready bool, tokenData []byte, err error)

	// GetSourceLogPollerFilters returns the filters that should be used for the source chain log poller
	GetSourceLogPollerFilters() []logpoller.Filter

	// GetSourceToken returns the token address on the source chain
	GetSourceToken() common.Address
}

package ccipdata

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
)

//go:generate mockery --quiet --name TokenPoolReader --filename token_pool_reader_mock.go --case=underscore
type TokenPoolReader interface {
	GetInboundTokenPoolRateLimitCall() (rpclib.EvmCall, error)
	GetInboundTokenPoolRateLimits(ctx context.Context, pools []TokenPoolReader) ([]TokenBucketRateLimit, error)
	Address() common.Address
}

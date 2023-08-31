package evmlogs

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
)

type RequestWithMeta[T any] struct {
	Request T
	BlockMeta
}

type BlockMeta struct {
	BlockTimestamp time.Time
	BlockNumber    int64
}

type Client interface {
	GetSendRequestsAfterNextMin(ctx context.Context, onRamp common.Address, nextMin uint64, confs int, checkFinalityTags bool) ([]RequestWithMeta[evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested], error)
	GetSendRequestsInSeqNumRange(ctx context.Context, onRamp common.Address, rangeMin, rangeMax uint64, confs int) ([]RequestWithMeta[evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested], error)
	GetTokenPriceUpdatesCreatedAfter(ctx context.Context, priceRegistry common.Address, ts time.Time, confs int) ([]RequestWithMeta[price_registry.PriceRegistryUsdPerTokenUpdated], error)
	GetGasPriceUpdatesCreatedAfter(ctx context.Context, priceRegistry common.Address, chainSelector uint64, ts time.Time, confs int) ([]RequestWithMeta[price_registry.PriceRegistryUsdPerUnitGasUpdated], error)
	GetExecutionStateChangesInRange(ctx context.Context, offRamp common.Address, rangeMin, rangeMax uint64, confs int) ([]RequestWithMeta[evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged], error)
	LatestBlock(ctx context.Context) (int64, error)
}

package observability

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
)

type ObservedPriceRegistry struct {
	price_registry.PriceRegistryInterface
	metric metricDetails
}

func NewObservedPriceRegistry(address common.Address, pluginName string, backend bind.ContractBackend) (price_registry.PriceRegistryInterface, error) {
	priceRegistry, err := price_registry.NewPriceRegistry(address, backend)
	if err != nil {
		return nil, err
	}

	return &ObservedPriceRegistry{
		PriceRegistryInterface: priceRegistry,
		metric: metricDetails{
			histogram:  priceRegistryHistogram,
			pluginName: pluginName,
		},
	}, nil
}

func (o *ObservedPriceRegistry) GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	return withObservedContract(o.metric, "GetFeeTokens", func() ([]common.Address, error) {
		return o.PriceRegistryInterface.GetFeeTokens(opts)
	})
}

func (o *ObservedPriceRegistry) GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]price_registry.InternalTimestampedUint192Value, error) {
	return withObservedContract(o.metric, "GetTokenPrices", func() ([]price_registry.InternalTimestampedUint192Value, error) {
		return o.PriceRegistryInterface.GetTokenPrices(opts, tokens)
	})
}

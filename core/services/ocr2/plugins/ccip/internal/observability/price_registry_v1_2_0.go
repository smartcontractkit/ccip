package observability

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
)

type ObservedPriceRegistryV1_2_0 struct {
	*ObservedPriceRegistryV1_0_0
	pr *price_registry.PriceRegistry
}

func NewObservedPriceRegistryV1_2_0(address common.Address, pluginName string, client client.Client) (*ObservedPriceRegistryV1_2_0, error) {
	v100, err := NewObservedPriceRegistryV1_0_0(address, pluginName, client)
	if err != nil {
		return nil, err
	}
	priceRegistry, err := price_registry.NewPriceRegistry(address, client)
	if err != nil {
		return nil, err
	}

	return &ObservedPriceRegistryV1_2_0{
		ObservedPriceRegistryV1_0_0: v100,
		pr:                          priceRegistry,
	}, nil
}

// Changed in 1.2.0
func (o *ObservedPriceRegistryV1_2_0) GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]price_registry.InternalTimestampedPackedUint224, error) {
	return withObservedContract(o.metric, "GetTokenPrices", func() ([]price_registry.InternalTimestampedPackedUint224, error) {
		return o.pr.GetTokenPrices(opts, tokens)
	})
}

package v1_6_0

import (
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
)

var (
	_ ccipdata.PriceRegistryReader = &PriceRegistry{}
)

type PriceRegistry struct {
	*v1_2_0.PriceRegistry
	pr *price_registry.PriceRegistry
}

func NewPriceRegistry(lggr logger.Logger, priceRegistryAddr common.Address, lp logpoller.LogPoller, ec client.Client, registerFilters bool) (*PriceRegistry, error) {
	v120, err := v1_2_0.NewPriceRegistry(lggr, priceRegistryAddr, lp, ec, registerFilters)
	if err != nil {
		return nil, err
	}
	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddr, ec)
	if err != nil {
		return nil, err
	}
	return &PriceRegistry{
		PriceRegistry: v120,
		pr:            priceRegistry,
	}, nil
}

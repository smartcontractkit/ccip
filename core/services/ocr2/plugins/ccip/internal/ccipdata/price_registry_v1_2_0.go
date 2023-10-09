package ccipdata

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/observability"
)

var (
	_ PriceRegistryReader = &PriceRegistryV1_2_0{}
)

type PriceRegistryV1_2_0 struct {
	*PriceRegistryV1_0_0
	obs *observability.ObservedPriceRegistryV1_2_0
}

func NewPriceRegistryV1_2_0(lggr logger.Logger, priceRegistryAddr common.Address, lp logpoller.LogPoller, ec client.Client) (*PriceRegistryV1_2_0, error) {
	v100, err := NewPriceRegistryV1_0_0(lggr, priceRegistryAddr, lp, ec)
	if err != nil {
		return nil, err
	}
	obs, err := observability.NewObservedPriceRegistryV1_2_0(priceRegistryAddr, ExecPluginLabel, ec)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryV1_2_0{
		PriceRegistryV1_0_0: v100,
		obs:                 obs,
	}, nil
}

// GetTokenPrices must be overridden to use the 1.2 ABI (return parameter changed from uint192 to uint224.
func (p *PriceRegistryV1_2_0) GetTokenPrices(ctx context.Context, wantedTokens []common.Address) ([]TokenPriceUpdate, error) {
	// Make call using 224 ABI.
	tps, err := p.obs.GetTokenPrices(&bind.CallOpts{Context: ctx}, wantedTokens)
	if err != nil {
		return nil, err
	}
	var tpu []TokenPriceUpdate
	for i, tp := range tps {
		tpu = append(tpu, TokenPriceUpdate{
			TokenPrice: TokenPrice{
				Token: wantedTokens[i],
				Value: tp.Value,
			},
			Timestamp: big.NewInt(int64(tp.Timestamp)), // TODO: valid conversion
		})
	}
	return tpu, nil
}

package v1_2_0

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_0_0"
)

var (
	_ ccipdata.PriceRegistryReader = &PriceRegistry{}
)

type PriceRegistry struct {
	*v1_0_0.PriceRegistry
	pr *price_registry.PriceRegistry
}

func NewPriceRegistry(lggr logger.Logger, priceRegistryAddr common.Address, lp logpoller.LogPoller, ec client.Client) (*PriceRegistry, error) {
	v100, err := v1_0_0.NewPriceRegistry(lggr, priceRegistryAddr, lp, ec)
	if err != nil {
		return nil, err
	}
	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddr, ec)
	if err != nil {
		return nil, err
	}
	return &PriceRegistry{
		PriceRegistry: v100,
		pr:            priceRegistry,
	}, nil
}

// GetTokenPrices must be overridden to use the 1.2 ABI (return parameter changed from uint192 to uint224)
// See https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/PriceRegistry.sol#L141
func (p *PriceRegistry) GetTokenPrices(ctx context.Context, wantedTokens []common.Address) ([]ccipdata.TokenPriceUpdate, error) {
	// Make call using 224 ABI.
	tps, err := p.pr.GetTokenPrices(&bind.CallOpts{Context: ctx}, wantedTokens)
	if err != nil {
		return nil, err
	}
	var tpu []ccipdata.TokenPriceUpdate
	for i, tp := range tps {
		tpu = append(tpu, ccipdata.TokenPriceUpdate{
			TokenPrice: ccipdata.TokenPrice{
				Token: wantedTokens[i],
				Value: tp.Value,
			},
			TimestampUnixSec: big.NewInt(int64(tp.Timestamp)),
		})
	}
	return tpu, nil
}

// ApplyPriceRegistryUpdate is a helper function used in tests only.
func ApplyPriceRegistryUpdate(t *testing.T, user *bind.TransactOpts, addr common.Address, ec client.Client, gasPrices []ccipdata.GasPrice, tokenPrices []ccipdata.TokenPrice) common.Hash {
	require.True(t, len(gasPrices) <= 1)
	pr, err := price_registry.NewPriceRegistry(addr, ec)
	require.NoError(t, err)
	o, err := pr.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, user.From, o)
	var tps []price_registry.InternalTokenPriceUpdate
	for _, tp := range tokenPrices {
		tps = append(tps, price_registry.InternalTokenPriceUpdate{
			SourceToken: tp.Token,
			UsdPerToken: tp.Value,
		})
	}
	var gps []price_registry.InternalGasPriceUpdate
	for _, gp := range gasPrices {
		gps = append(gps, price_registry.InternalGasPriceUpdate{
			DestChainSelector: gp.DestChainSelector,
			UsdPerUnitGas:     gp.Value,
		})
	}
	tx, err := pr.UpdatePrices(user, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: tps,
		GasPriceUpdates:   gps,
	})
	require.NoError(t, err)
	return tx.Hash()
}

package ccipdata

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/logpollerutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var (
	_ PriceRegistryReader = &PriceRegistryV1_2_0{}
)

type PriceRegistryV1_3_0 struct {
	*PriceRegistryV1_2_0
	pr            *price_registry.PriceRegistry
	feedIDAdded   common.Hash
	feedIDRemoved common.Hash
	v13Filters    []logpoller.Filter
}

func NewPriceRegistryV1_3_0(lggr logger.Logger, priceRegistryAddr common.Address, lp logpoller.LogPoller, ec client.Client) (*PriceRegistryV1_3_0, error) {
	v120, err := NewPriceRegistryV1_2_0(lggr, priceRegistryAddr, lp, ec)
	if err != nil {
		return nil, err
	}
	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddr, ec)
	if err != nil {
		return nil, err
	}
	prABI := abihelpers.MustParseABI(price_registry.PriceRegistryABI)
	feedIDAdded := abihelpers.MustGetEventID("FeedIdAdded", prABI)
	feedIDRemoved := abihelpers.MustGetEventID("FeedIdRemoved", prABI)
	var filters = []logpoller.Filter{
		{
			Name: logpoller.FilterName(FEED_ID_ADDED, priceRegistryAddr.String()),
			EventSigs: evmtypes.HashArray{
				feedIDAdded,
			},
			Addresses: evmtypes.AddressArray{
				priceRegistryAddr,
			},
		},
		{
			Name: logpoller.FilterName(FEED_ID_REMOVED, priceRegistryAddr.String()),
			EventSigs: evmtypes.HashArray{
				feedIDRemoved,
			},
			Addresses: evmtypes.AddressArray{
				priceRegistryAddr,
			},
		},
	}
	err = logpollerutil.RegisterLpFilters(lp, filters)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryV1_3_0{
		PriceRegistryV1_2_0: v120,
		pr:                  priceRegistry,
		feedIDAdded:         feedIDAdded,
		feedIDRemoved:       feedIDRemoved,
		v13Filters:          filters,
	}, nil
}

func (p *PriceRegistryV1_3_0) FeedIdEvents() []common.Hash {
	return []common.Hash{p.feedIDAdded, p.feedIDRemoved}
}

func (p *PriceRegistryV1_3_0) GetFeedIDsForTokens(ctx context.Context, tokenAddresses []common.Address) ([][32]byte, error) {
	return p.pr.GetFeedIds(&bind.CallOpts{Context: ctx}, tokenAddresses)
}

func (p *PriceRegistryV1_3_0) Close(opts ...pg.QOpt) error {
	return logpollerutil.UnregisterLpFilters(p.lp, append(p.filters, p.v13Filters...), opts...)
}

func ApplyPriceRegistryUpdateV1_3_0(
	t *testing.T,
	user *bind.TransactOpts,
	addr common.Address,
	ec client.Client,
	gasPrices []GasPrice,
	tokenPrices []TokenPrice,
) common.Hash {
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

package pricegetter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merclib"
)

// MercuryGetter is a PriceGetter that uses Mercury to get prices.
type MercuryGetter struct {
	mercClient merclib.MercuryClient
	prr        ccipdata.PriceRegistryReader
}

func NewMercuryGetter(mercClient merclib.MercuryClient, priceRegReader ccipdata.PriceRegistryReader) *MercuryGetter {
	return &MercuryGetter{mercClient: mercClient, prr: priceRegReader}
}

func (d *MercuryGetter) TokenPricesUSD(ctx context.Context, tokens []common.Address) (map[common.Address]*big.Int, error) {
	feedIDs, err := d.prr.GetFeedIDsForTokens(ctx, tokens)
	if err != nil {
		return nil, fmt.Errorf("error getting feed IDs: %w", err)
	}
	feedIDToAddrs := make(map[[32]byte]common.Address)
	for i, token := range tokens {
		feedIDToAddrs[feedIDs[i]] = token
	}
	rwcs, err := d.mercClient.BatchFetchPrices(ctx, feedIDs)
	if err != nil {
		return nil, fmt.Errorf("error fetching prices from mercury: %w", err)
	}
	if len(rwcs) != len(feedIDs) {
		return nil, fmt.Errorf("expected %d rwcs, got %d", len(feedIDs), len(rwcs))
	}
	ret := make(map[common.Address]*big.Int)
	var errs error
	for _, rwc := range rwcs {
		// TODO: do we have the decimals somewhere?
		tokenAddr, ok := feedIDToAddrs[rwc.FeedId]
		if !ok {
			// should be impossible but just in case
			errs = multierr.Append(errs, fmt.Errorf("feed ID %x not found in feedIDToAddrs", rwc.FeedId))
			continue
		}
		// TODO: benchmark price or something else?
		ret[tokenAddr] = rwc.V3Report.BenchmarkPrice
	}
	return ret, errs
}

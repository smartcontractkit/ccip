package db

import (
	"context"
	"fmt"

	cciporm "github.com/smartcontractkit/chainlink/v2/core/services/ccip"
)

func GetFilteredSortedLaneTokens(ctx context.Context, orm cciporm.ORM, ) (lerr error) {
	destFeeTokens, destBridgeableTokens, err := GetDestinationTokens(ctx, offRamp, priceRegistry)
	if err != nil {
		return nil, nil, fmt.Errorf("get tokens with batch limit: %w", err)
	}

	destTokensWithPrice, destTokensWithoutPrice, err := priceGetter.FilterConfiguredTokens(ctx, destBridgeableTokens)
	if err != nil {
		return nil, nil, fmt.Errorf("filter for priced tokens: %w", err)
	}

	return flattenedAndSortedTokens(destFeeTokens, destTokensWithPrice), destTokensWithoutPrice, nil
}

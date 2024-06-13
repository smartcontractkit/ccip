package reader

import (
	"context"
	"fmt"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type TokenPriceConfig struct {
	// This is mainly used for tokens on testnet to give them a price
	StaticPrices map[ocrtypes.Account]*cciptypes.BigInt `json:"staticPrice"`
}

type OnchainTokenPricesReader struct {
	TokenPriceConfig TokenPriceConfig
	// Reader for the chain that will have the token prices on-chain
	contractReader commontypes.ContractReader
}

// GetTokenPricesUSD TODO: Update interface to return cciptypes.BigInt
func (pr *OnchainTokenPricesReader) GetTokenPricesUSD(ctx context.Context, tokens []ocrtypes.Account) ([]*cciptypes.BigInt, error) {
	const (
		contractName = "PriceRegistry"
		functionName = "getTokenPrice"
	)

	var prices []*cciptypes.BigInt
	for _, token := range tokens {
		var price *cciptypes.BigInt
		if staticPrice, exists := pr.TokenPriceConfig.StaticPrices[token]; exists {
			price = staticPrice
		} else {
			err := pr.contractReader.GetLatestValue(ctx, contractName, functionName, token, &price)
			if err != nil {
				return nil, err
			}
		}

		prices = append(prices, price)
	}

	if len(prices) != len(tokens) {
		return nil, fmt.Errorf("failed to get all token prices")
	}

	return prices, nil
}

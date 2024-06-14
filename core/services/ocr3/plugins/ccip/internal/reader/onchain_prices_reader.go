package reader

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type TokenPriceConfig struct {
	// This is mainly used for tokens on testnet to give them a price
	StaticPrices map[ocrtypes.Account]*big.Int `json:"staticPrice"`
}

type OnchainTokenPricesReader struct {
	TokenPriceConfig TokenPriceConfig
	// Reader for the chain that will have the token prices on-chain
	contractReader commontypes.ContractReader
}

// GetTokenPricesUSD TODO: Update interface to return cciptypes.BigInt
func (pr *OnchainTokenPricesReader) GetTokenPricesUSD(ctx context.Context, tokens []ocrtypes.Account) ([]*big.Int, error) {
	const (
		contractName = "PriceAggregator"
		functionName = "getTokenPrice"
	)

	prices := make([]*big.Int, 0, len(tokens))
	errChan := make(chan error, 1)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(tokens))

	for idx, token := range tokens {
		go func(idx int, token ocrtypes.Account) {
			defer wg.Done()
			var price *big.Int
			if staticPrice, exists := pr.TokenPriceConfig.StaticPrices[token]; exists {
				price = staticPrice
			} else {
				if err := pr.contractReader.GetLatestValue(ctx, contractName, functionName, token, &price); err != nil {
					select {
					case errChan <- err:
					default:
					}
					cancel()
					return
				}
			}
			select {
			case <-ctx.Done():
				return
			default:
				prices[idx] = price
			}
		}(idx, token)
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(errChan)
	}()

	if err, ok := <-errChan; ok {
		return nil, fmt.Errorf("failed to get all token prices successfully: %v", err)
	}

	if len(prices) != len(tokens) {
		return nil, fmt.Errorf("failed to get all token prices successfully")
	}

	return prices, nil
}

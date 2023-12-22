package pricegetter

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type TokenPriceResult struct {
	TokenAddress common.Address
	Price        *big.Int
	Error        error
}

//go:generate mockery --quiet --name PriceGetter --output . --filename mock.go --inpackage --case=underscore
type PriceGetter interface {
	// TokenPricesUSD returns token prices in USD.
	//
	// All tokens should be returned in the same order as they were provided.
	// Top level error indicates problems not related to specific tokens. If a token is not found or can't be fetched,
	// the TokenPriceResult must contain an error for that token.
	TokenPricesUSD(ctx context.Context, tokens []common.Address) ([]TokenPriceResult, error)
}

// PriceResultsToMapSkipErrors transforms a slice of TokenPriceResult to a map of token address to price.
// It skips tokens with errors and only log this information.
//
// In Commit Plugin, we don't want to fail in case of spotting missing token, but rather log it.
// This should make CCIP more resilient. In case of missing tokens in the spec, Exec will skip messages containing those tokens.
func PriceResultsToMapSkipErrors(lggr logger.Logger, results []TokenPriceResult) map[common.Address]*big.Int {
	m := make(map[common.Address]*big.Int)
	for _, r := range results {
		if r.Error != nil {
			lggr.Errorw("error when fetching token price", "token", r.TokenAddress, "error", r.Error)
			continue
		}
		m[r.TokenAddress] = r.Price
	}
	return m
}

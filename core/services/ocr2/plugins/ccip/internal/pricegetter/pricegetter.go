package pricegetter

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenPriceResult struct {
	Price *big.Int
	Error error
}

//go:generate mockery --quiet --name PriceGetter --output . --filename mock.go --inpackage --case=underscore
type PriceGetter interface {
	// TokenPricesUSD returns token prices in USD.
	//
	// Top level error indicates problems not related to specific tokens. If a token is not found or can't be fetched,
	// the TokenPriceResult must contain an error for that token.
	TokenPricesUSD(ctx context.Context, tokens []common.Address) (map[common.Address]TokenPriceResult, error)
}

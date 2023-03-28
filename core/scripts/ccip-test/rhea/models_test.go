package rhea

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenPrice(t *testing.T) {
	var tt = []struct {
		token    Token
		expected *big.Int
	}{
		{
			LINK,
			new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
		{
			ANZ,
			new(big.Int).Mul(big.NewInt(1e12), big.NewInt(1e18)), // uses 6 decimals
		},
		{
			CACHEGOLD,
			new(big.Int).Mul(big.NewInt(60e10), big.NewInt(1e18)), // uses 8 decimals
		},
	}
	for _, tc := range tt {
		tc := tc
		a := tc.token.Price()
		assert.Equal(t, tc.expected, a)
	}
}

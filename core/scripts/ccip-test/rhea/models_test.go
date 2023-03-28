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
			big.NewInt(10),
		},
		{
			ANZ,
			big.NewInt(1e12), // uses 6 decimals
		},
		{
			CACHEGOLD,
			big.NewInt(60e10), // uses 8 decimals
		},
	}
	for _, tc := range tt {
		tc := tc
		a := tc.token.Price()
		assert.Equal(t, tc.expected, a)
	}
}

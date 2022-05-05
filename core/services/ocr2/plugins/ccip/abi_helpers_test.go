package ccip

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProofFlagToBits(t *testing.T) {
	var tt = []struct {
		flags    []bool
		expected *big.Int
	}{
		{
			[]bool{true, false, true},
			big.NewInt(5),
		},
		{
			[]bool{true, true, false}, // Note the bits are reversed, slightly easier to implement.
			big.NewInt(3),
		},
		{
			[]bool{false, true, true},
			big.NewInt(6),
		},
		{
			[]bool{false, false, false},
			big.NewInt(0),
		},
		{
			[]bool{true, true, true},
			big.NewInt(7),
		},
	}
	for _, tc := range tt {
		tc := tc
		a := ProofFlagsToBits(tc.flags)
		assert.Equal(t, tc.expected.String(), a.String())
	}
}

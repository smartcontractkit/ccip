package slicelib

import "math/big"

// BoolsToBitFlags transforms a list of boolean flags to a *big.Int encoded number.
func BoolsToBitFlags(bools []bool) *big.Int {
	encodedFlags := big.NewInt(0)
	for i := 0; i < len(bools); i++ {
		if bools[i] {
			encodedFlags.SetBit(encodedFlags, i, 1)
		}
	}
	return encodedFlags
}

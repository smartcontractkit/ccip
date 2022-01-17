package ccip

import (
	"math/big"
)

// Multi-chain tests using the sim have to be remapped to the default
// sim chainID because its a hardcoded constant in the geth code base and so
// and CHAINID op codes will ALWAYS be 1337.
func maybeRemapChainID(chainID *big.Int) *big.Int {
	testChainIDs := []*big.Int{big.NewInt(1000), big.NewInt(2000)}
	for _, testChainID := range testChainIDs {
		if chainID.Cmp(testChainID) == 0 {
			return big.NewInt(1337)
		}
	}
	return chainID
}

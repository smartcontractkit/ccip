package deployment

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/environment"
)

func TestDeployCCIPContracts(t *testing.T) {
	e := environment.NewMemoryEnvironment(t, environment.MemoryEnvironmentConfig{
		Chains: 1,
		Nodes:  1,
	})
	// Deploy all the CCIP contracts.
	require.NoError(t, DeployCCIPContracts(e))
	state, err := GenerateOnchainState(e)
	require.NoError(t, err)
	snap, err := state.Snapshot(e.AllChainSelectors())
	require.NoError(t, err)

	// Assert expect every deployed address to be in the address book.
	for name, chain := range snap.Chains {
		addrs, err := e.AddressBook.Addresses()
		require.NoError(t, err)
		evmChainID, _ := chainsel.ChainIdFromName(name)
		sel, _ := chainsel.SelectorFromChainId(evmChainID)
		assert.Contains(t, addrs[sel], chain.TokenAdminRegistry.String())
	}
}

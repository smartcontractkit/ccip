package deployment

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
)

func TestDeployCCIPContracts(t *testing.T) {
	e := deployment.NewMemoryEnvironment(t, deployment.MemoryEnvironmentConfig{
		Chains: 1,
		Nodes:  1,
	})
	// Deploy all the CCIP contracts.
	ab, err := DeployCCIPContracts(e, DeployCCIPContractConfig{})
	require.NoError(t, err)
	state, err := GenerateOnchainState(e, ab)
	require.NoError(t, err)
	snap, err := state.Snapshot(e.AllChainSelectors())
	require.NoError(t, err)

	// Assert expect every deployed address to be in the address book.
	for name, chain := range snap.Chains {
		addrs, err := ab.Addresses()
		require.NoError(t, err)
		evmChainID, _ := chainsel.ChainIdFromName(name)
		sel, _ := chainsel.SelectorFromChainId(evmChainID)
		assert.Contains(t, addrs[sel], chain.TokenAdminRegistry.String())
	}
}

package migrations

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	ccipdeployment "github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/deployment"
)

func Test0001_InitialDeploy(t *testing.T) {
	t.Skip() // WIP
	e := deployment.NewMemoryEnvironment(t, deployment.MemoryEnvironmentConfig{
		Chains: 1,
		Nodes:  1,
	})
	// Apply migration
	output, err := Apply0001(e, ccipdeployment.DeployCCIPContractConfig{})
	require.NoError(t, err)

	state, err := ccipdeployment.GenerateOnchainState(e, output.AddressBook)
	require.NoError(t, err)

	// TODO: Validate jobs
	// Apply jobs
	for nodeIDs, jobs := range output.JobSpecs {
		for _, job := range jobs {
			_, err := e.Offchain.ProposeJob(context.Background(), nodeIDs, job)
			require.NoError(t, err)
		}
	}
	// TODO: Inspect proposal
	// Apply proposal
	require.NoError(t, ccipdeployment.ApplyProposal(e, output.Proposals[0], state))

	// TODO: Inspect onchain state
	// TODO: Send traffic

	//snap, err := state.Snapshot(e.AllChainSelectors())
	//require.NoError(t, err)
	//
	//// Assert expect every deployed address to be in the address book.
	//for name, chain := range snap.Chains {
	//	addrs, err := ab.Addresses()
	//	require.NoError(t, err)
	//	evmChainID, _ := chainsel.ChainIdFromName(name)
	//	sel, _ := chainsel.SelectorFromChainId(evmChainID)
	//	assert.Contains(t, addrs[sel], chain.TokenAdminRegistry.String())
	//}
}

package migrations

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	jobv1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/job/v1"

	ccipdeployment "github.com/smartcontractkit/chainlink/integration-tests/deployment/ccip"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func Test0001_InitialDeploy(t *testing.T) {
	lggr := logger.TestLogger(t)
	chains := memory.NewMemoryChains(t, 3)
	t.Log(chains, lggr)

	homeChainSel := uint64(0)
	homeChainEVM := uint64(0)
	// First chain is home chain.
	for chainSel := range chains {
		homeChainEVM, _ = chainsel.ChainIdFromSelector(chainSel)
		homeChainSel = chainSel
		break
	}
	ab, err := ccipdeployment.DeployCapReg(lggr, chains, homeChainSel)
	require.NoError(t, err)
	fmt.Println(homeChainEVM, ab)

	addrs, err := ab.AddressesForChain(homeChainSel)
	require.NoError(t, err)
	require.Len(t, addrs, 1)
	capReg := common.Address{}
	for addr := range addrs {
		capReg = common.HexToAddress(addr)
		break
	}

	e := memory.NewMemoryEnvironmentExistingChains(t, lggr, chains, memory.MemoryEnvironmentConfig{
		Chains:     3,
		Nodes:      4,
		Bootstraps: 1,
		RegistryConfig: memory.RegistryConfig{
			EVMChainID: homeChainEVM,
			Contract:   capReg,
		},
	})
	// Apply migration
	output, err := Apply0001(e, ccipdeployment.DeployCCIPContractConfig{})
	require.NoError(t, err)

	// Before we can add the jobs, we need to add the config to the cap registry
	/*
		ccipCapabilityID, err := homeChainUni.capabilityRegistry.GetHashedCapabilityId(
			callCtx, CapabilityLabelledName, CapabilityVersion)
		require.NoError(t, err, "failed to get hashed capability id for ccip")
		require.NotEqual(t, [32]byte{}, ccipCapabilityID, "ccip capability id is empty")

		// Need to Add nodes and assign capabilities to them before creating DONS
		homeChainUni.AddNodes(t, p2pIDs, [][32]byte{ccipCapabilityID})

		for _, uni := range universes {
			t.Logf("Adding chainconfig for chain %d", uni.chainID)
			AddChainConfig(t, homeChainUni, getSelector(uni.chainID), p2pIDs, fChain)
		}

		cfgs, err := homeChainUni.ccipConfig.GetAllChainConfigs(callCtx)
		require.NoError(t, err)
		require.Len(t, cfgs, numChains)

		// Create a DON for each chain
		for _, uni := range universes {
			// Add nodes and give them the capability
			t.Log("Adding DON for universe: ", uni.chainID)
			chainSelector := getSelector(uni.chainID)
			homeChainUni.AddDON(
				t,
				ccipCapabilityID,
				chainSelector,
				uni,
				fChain,
				bootstrapP2PID,
				p2pIDs,
				oracles[uni.chainID],
			)
		}
	*/

	// Apply the jobs.
	for nodeID, jobs := range output.JobSpecs {
		for _, job := range jobs {
			_, err := e.Offchain.ProposeJob(context.Background(),
				&jobv1.ProposeJobRequest{
					NodeId: nodeID,
					Spec:   job,
				})
			require.NoError(t, err)
		}
	}
	// TODO: With the jobs, we should be able to send traffic through the new deployment.

	// TODO: Apply the proposal.
}

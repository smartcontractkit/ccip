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
	require.Len(t, addrs, 2)
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
	state, err := ccipdeployment.GenerateOnchainState(e, ab)
	require.NoError(t, err)
	// Apply migration
	output, err := Apply0001(e, ccipdeployment.DeployCCIPContractConfig{
		HomeChainSel: homeChainSel,
		// Capreg/config already exist.
		CCIPOnChainState: state,
	})
	require.NoError(t, err)

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

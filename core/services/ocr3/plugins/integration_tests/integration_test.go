package integration_tests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/consul/sdk/freeport"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	deployments "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment"
	"github.com/stretchr/testify/require"
	"testing"
)

func confirm(tx common.Hash, chain deployments.Chain) error {
	chain.Client.(*client.SimulatedBackendClient).Commit()
	return nil
}

func TestCCIP(t *testing.T) {
	numNodes := 4
	ports := freeport.GetN(t, numNodes)
	capabilitiesPorts := freeport.GetN(t, numNodes)
	var nodes []*ocr3Node
	var nodeIds []string
	chains := createChains(t, 3)
	homeChainSelector := chainsel.TEST_90000001
	// say first one is home chain
	addressBook := deployments.NewMemoryAddressBook()
	require.NoError(t, deployments.DeployCapabilityRegistry(addressBook, chains[homeChainSelector.Selector], confirm))
	onChainState, err := deployments.GenerateOnchainState(chains, addressBook)
	require.NoError(t, err)
	t.Log(onChainState.CapabilityRegistry.Address())

	for i := 0; i < numNodes; i++ {
		ocrNode := setupNodeOCR3(t, ports[i], capabilitiesPorts[i], nil, chains, onChainState.CapabilityRegistry.Address(), homeChainSelector.Selector)
		nodeIds = append(nodeIds, ocrNode.peerID)
	}
	jobServiceClient := MockJobDistributor{nodes: nodes}
	require.NoError(t, deployments.DeployNewCCIPToExistingDON(addressBook, nodeIds, chains, jobServiceClient, confirm))
	// Do all the in memory testing here
}

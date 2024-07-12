package ccip_integration_tests

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/libocr/commontypes"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"
)

func TestIntegration_OCR3Nodes(t *testing.T) {
	numChains := 3
	homeChainUni, universes := createUniverses(t, numChains)
	//homeChainUni, universes := createUniverses(t, 3)
	numNodes := 4
	t.Log("creating ocr3 nodes")
	var (
		oracles      = make(map[uint64][]confighelper2.OracleIdentityExtra)
		transmitters = make(map[uint64][]common.Address)
		apps         []chainlink.Application
		nodes        []*ocr3Node
		p2pIDs       [][32]byte

		// The bootstrap node will be the first node (index 0)
		bootstrapPort  int
		bootstrapP2PID p2pkey.PeerID
		bootStrappers  []commontypes.BootstrapperLocator
	)

	ports := freeport.GetN(t, numNodes)
	for i := 0; i < numNodes; i++ {
		node := setupNodeOCR3(t, ports[i], bootStrappers, universes)

		apps = append(apps, node.app)
		for chainID, transmitter := range node.transmitters {
			transmitters[chainID] = append(transmitters[chainID], transmitter)
		}
		for chainID, transmitter := range node.transmitters {
			identity := confighelper2.OracleIdentityExtra{
				OracleIdentity: confighelper2.OracleIdentity{
					OnchainPublicKey:  node.keybundle.PublicKey(),
					TransmitAccount:   ocrtypes.Account(transmitter.Hex()),
					OffchainPublicKey: node.keybundle.OffchainPublicKey(),
					PeerID:            node.peerID,
				},
				ConfigEncryptionPublicKey: node.keybundle.ConfigEncryptionPublicKey(),
			}
			oracles[chainID] = append(oracles[chainID], identity)
		}
		nodes = append(nodes, node)
		peerID, err := p2pkey.MakePeerID(node.peerID)
		require.NoError(t, err)
		p2pIDs = append(p2pIDs, peerID)

		// First Node is the bootstrap node
		if i == 0 {
			bootstrapPort = ports[i]
			bootstrapP2PID = peerID
			bootStrappers = []commontypes.BootstrapperLocator{
				{PeerID: node.peerID, Addrs: []string{
					fmt.Sprintf("127.0.0.1:%d", bootstrapPort),
				}},
			}
		}
	}

	t.Log("starting ticker to commit blocks")
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	tickCtx, tickCancel := context.WithCancel(testutils.Context(t))
	defer tickCancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-tick.C:
				for _, uni := range universes {
					uni.backend.Commit()
				}
			case <-tickCtx.Done():
				return
			}
		}
	}()
	t.Cleanup(func() {
		tickCancel()
		wg.Wait()
	})

	bootstrapNode := nodes[0]

	t.Log("adding bootstrap node job")
	err := bootstrapNode.app.Start(testutils.Context(t))
	require.NoError(t, err, "failed to start bootstrap node")
	t.Cleanup(func() {
		require.NoError(t, bootstrapNode.app.Stop())
	})

	evmChains := bootstrapNode.app.GetRelayers().LegacyEVMChains()
	require.NotNil(t, evmChains)
	require.Len(t, evmChains.Slice(), numChains)

	t.Log("creating ocr3 jobs")
	for i := 0; i < numNodes; i++ {
		//err := apps[i].Start(testutils.Context(t))
		//require.NoError(t, err)
		//tApp := apps[i]
		//t.Cleanup(func() {
		//	require.NoError(t, tApp.Stop())
		//})
		//
		//ccipSpecToml := createCCIPSpecToml(nodes[i].peerID, bootstrapP2PID.String(), bootstrapPort, nodes[i].keybundle.ID())
		//t.Log("Creating ccip job with spec:\n", ccipSpecToml)

		//ccipJob, err2 := ccipcapability(ccipSpecToml)
		//require.NoError(t, err2, "failed to validate ccip job")
		//err2 = apps[i].AddJobV2(testutils.Context(t), &ccipJob)
		//require.NoError(t, err2, "failed to add ccip job")
	}

	// add the ccip dons to the capability registry.
	ccipCapabilityID, err := homeChainUni.capabilityRegistry.GetHashedCapabilityId(nil, CapabilityLabelledName, CapabilityVersion)
	require.NoError(t, err, "failed to get hashed capability id for ccip")
	require.NotEqual(t, [32]byte{}, ccipCapabilityID, "ccip capability id is empty")

	homeChainUni.AddNodes(t, p2pIDs, [][32]byte{ccipCapabilityID})
	// create a DON for each chain
	for _, uni := range universes {
		// Add nodes and give them the capability
		t.Log("AddingDON for universe: ", uni.chainID)
		homeChainUni.AddDON(t,
			ccipCapabilityID,
			uni.chainID,
			uni.offramp.Address().Bytes(),
			1, // f
			bootstrapP2PID,
			p2pIDs,
			oracles[uni.chainID],
		)
	}
}

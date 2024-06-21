package internal

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/keystone_capability_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"
)

func TestIntegration_CCIPCapability(t *testing.T) {
	t.Skip("not ready yet")

	numChains := 3
	owner, chains := createChains(t, numChains)
	homeChainUni, universes := deployContracts(t, owner, chains)
	fullyConnectCCIPContracts(t, owner, universes)

	// add the ccip capability to the capability registry.
	_, err := homeChainUni.capabilityRegistry.AddCapabilities(owner, []keystone_capability_registry.CapabilityRegistryCapability{
		{
			LabelledName:          "ccip",
			Version:               "v1.0.0",
			CapabilityType:        2, // consensus. not used (?)
			ResponseType:          0, // report. not used (?)
			ConfigurationContract: homeChainUni.ccipConfigContract,
		},
	})
	require.NoError(t, err, "failed to add capabilities to the capability registry")
	homeChainUni.backend.Commit()

	numNodes := 4

	t.Log("creating ocr3 nodes")
	var (
		oracles      = make(map[int64][]confighelper2.OracleIdentityExtra)
		transmitters = make(map[int64][]common.Address)
		apps         []chainlink.Application
		nodes        []*ocr3Node
		p2pIDs       [][32]byte

		// The bootstrap node will be the first node (index 0)
		bootstrapPort  int
		bootstrapP2PID p2pkey.PeerID
	)
	ports := freeport.GetN(t, numNodes)
	for i := 0; i < numNodes; i++ {
		node := setupNodeOCR3(t, owner, ports[i], chains)

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

		if i == 0 {
			bootstrapPort = ports[i]
			bootstrapP2PID = peerID
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
				for _, backend := range chains {
					backend.Commit()
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

	t.Log("creating ocr3 jobs")
	for i := 0; i < numNodes; i++ {
		err = apps[i].Start(testutils.Context(t))
		require.NoError(t, err)
		tapp := apps[i]
		t.Cleanup(func() {
			require.NoError(t, tapp.Stop())
		})

		ccipSpecToml := createCCIPSpecToml(nodes[i].peerID, bootstrapP2PID.String(), bootstrapPort, nodes[i].keybundle.ID())
		t.Log("Creating ccip job with spec:\n", ccipSpecToml)
		ccipJob, err2 := ccipcapability.ValidatedCCIPSpec(ccipSpecToml)
		require.NoError(t, err2, "failed to validate ccip job")
		err2 = apps[i].AddJobV2(testutils.Context(t), &ccipJob)
		require.NoError(t, err2, "failed to add ccip job")
	}

	// add the ccip dons to the capability registry.
	ccipCapabilityID, err := homeChainUni.capabilityRegistry.GetHashedCapabilityId(nil, "ccip", "v1.0.0")
	require.NoError(t, err, "failed to get hashed capability id for ccip")
	// create a DON for each chain
	for _, uni := range universes {
		homeChainUni.capabilityRegistry.AddDON(owner, p2pIDs, []keystone_capability_registry.CapabilityRegistryCapabilityConfiguration{
			{
				CapabilityId: ccipCapabilityID,
				Config:       donOCRConfig(t, uni, oracles[int64(uni.chainID)]),
			},
		}, false, false, 1 /* f value: unused for ccip */)
	}
}

package internal

import (
	"context"
	"fmt"
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
	"github.com/smartcontractkit/libocr/commontypes"
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

	// add the bootstrap and ccip capabilities to the capability registry.
	_, err := homeChainUni.capabilityRegistry.AddCapabilities(owner, []keystone_capability_registry.CapabilityRegistryCapability{
		{
			LabelledName:          "ccipbootstrap",
			Version:               "v1.0.0",
			CapabilityType:        2,                // consensus. not used (?)
			ResponseType:          0,                // report. not used (?)
			ConfigurationContract: common.Address{}, // no config contract for bootstrap DON.
		},
		{
			LabelledName:          "ccip",
			Version:               "v1.0.0",
			CapabilityType:        2,                // consensus. not used (?)
			ResponseType:          0,                // report. not used (?)
			ConfigurationContract: common.Address{}, // TODO: deploy
		},
	})
	require.NoError(t, err, "failed to add capabilities to the capability registry")
	homeChainUni.backend.Commit()

	t.Log("Creating bootstrap node")
	bootstrapNodePort := freeport.GetOne(t)
	bootstrapNode := setupNodeOCR3(t, owner, bootstrapNodePort, chains, nil)
	bootstrapP2PID, err := p2pkey.MakePeerID(bootstrapNode.peerID)
	require.NoError(t, err)
	numNodes := 4

	t.Log("creating ocr3 nodes")
	var (
		oracles      = make(map[int64][]confighelper2.OracleIdentityExtra)
		transmitters = make(map[int64][]common.Address)
		apps         []chainlink.Application
		nodes        []*ocr3Node
		p2pIDs       [][32]byte
	)
	ports := freeport.GetN(t, numNodes)
	for i := 0; i < numNodes; i++ {
		// Supply the bootstrap IP and port as a V2 peer address
		bootstrappers := []commontypes.BootstrapperLocator{
			{PeerID: bootstrapNode.peerID, Addrs: []string{
				fmt.Sprintf("127.0.0.1:%d", bootstrapNodePort),
			}},
		}
		node := setupNodeOCR3(t, owner, ports[i], chains, bootstrappers)

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

	t.Log("adding bootstrap node job")
	err = bootstrapNode.app.Start(testutils.Context(t))
	require.NoError(t, err, "failed to start bootstrap node")
	t.Cleanup(func() {
		require.NoError(t, bootstrapNode.app.Stop())
	})

	specToml := fmt.Sprintf(`
type = "ccipbootstrap"
capabilityVersion = "v1.0.0"
capabilityLabelledName = "ccipbootstrap"
[relayConfig]
evmChainID = %d
`, homeChainID)
	t.Log("Creating bootstrap job with spec:\n", specToml)
	bootstrapJob, err := ccipcapability.ValidatedCCIPBootstrapSpec(specToml)
	require.NoError(t, err, "failed to validate ccipbootstrap job")
	err = bootstrapNode.app.AddJobV2(testutils.Context(t), &bootstrapJob)
	require.NoError(t, err, "failed to add bootstrap job")

	t.Log("creating ocr3 jobs")
	for i := 0; i < numNodes; i++ {
		err = apps[i].Start(testutils.Context(t))
		require.NoError(t, err)
		tapp := apps[i]
		t.Cleanup(func() {
			require.NoError(t, tapp.Stop())
		})

		ccipSpecToml := fmt.Sprintf(`
type = "ccip"
capabilityVersion = "v1.0.0"
capabilityLabelledName = "ccip"
p2pKeyID = "%s"
[ocrKeyBundleIDs]
evm = "%s"
[relayConfigs.evm.chainReaderConfig.contracts.Offramp]
contractABI = "the abi"

[relayConfigs.evm.chainReaderConfig.contracts.Offramp.configs.getStuff]
chainSpecificName = "getStuffEVM"

[pluginConfig]
tokenPricesPipeline = "the pipeline"`, nodes[i].peerID, nodes[i].keybundle.ID())
		t.Log("Creating ccip job with spec:\n", ccipSpecToml)
		ccipJob, err2 := ccipcapability.ValidatedCCIPSpec(ccipSpecToml)
		require.NoError(t, err2, "failed to validate ccip job")
		err2 = apps[i].AddJobV2(testutils.Context(t), &ccipJob)
		require.NoError(t, err2, "failed to add ccip job")
	}

	// add the bootstrap and ccip dons to the capability registry.
	bootstrapCapabilityID, err := homeChainUni.capabilityRegistry.GetHashedCapabilityId(nil, "ccipbootstrap", "v1.0.0")
	require.NoError(t, err, "failed to get hashed capability id for ccipbootstrap")
	homeChainUni.capabilityRegistry.AddDON(owner, [][32]byte{bootstrapP2PID}, []keystone_capability_registry.CapabilityRegistryCapabilityConfiguration{
		{
			CapabilityId: bootstrapCapabilityID,
			Config:       donBootstrapConfig(t),
		},
	}, false, false, 1 /* f value: unused for bootstrap */)

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

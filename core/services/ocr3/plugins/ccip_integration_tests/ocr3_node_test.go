package ccip_integration_tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/consul/sdk/freeport"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
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
	numNodes := 4
	t.Log("creating ocr3 nodes")
	var (
		oracles = make(map[uint64][]confighelper2.OracleIdentityExtra)
		apps    []chainlink.Application
		nodes   []*ocr3Node
		p2pIDs  [][32]byte

		// The bootstrap node will be the first node (index 0)
		bootstrapPort  int
		bootstrapP2PID p2pkey.PeerID
		bootstrappers  []commontypes.BootstrapperLocator
	)

	ports := freeport.GetN(t, numNodes)
	capabilitiesPorts := freeport.GetN(t, numNodes)
	for i := 0; i < numNodes; i++ {
		node := setupNodeOCR3(t, ports[i], capabilitiesPorts[i], bootstrappers, universes, homeChainUni)

		apps = append(apps, node.app)
		for chainID, transmitter := range node.transmitters {
			//transmitters[chainID] = append(transmitters[chainID], transmitter)
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
			bootstrappers = []commontypes.BootstrapperLocator{
				{PeerID: node.peerID, Addrs: []string{
					fmt.Sprintf("127.0.0.1:%d", bootstrapPort),
				}},
			}
		}
	}

	// Start committing periodically in the background for all the chains
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()
	commitBlocksBackground(t, universes, tick)

	ctx := testutils.Context(t)

	ccipCapabilityID, err := homeChainUni.capabilityRegistry.GetHashedCapabilityId(&bind.CallOpts{
		Context: ctx,
	}, CapabilityLabelledName, CapabilityVersion)
	require.NoError(t, err, "failed to get hashed capability id for ccip")
	require.NotEqual(t, [32]byte{}, ccipCapabilityID, "ccip capability id is empty")

	// Need to Add nodes and assign capabilities to them before creating DONS

	homeChainUni.AddNodes(t, p2pIDs, [][32]byte{ccipCapabilityID})

	// Add homechain configs
	for _, uni := range universes {
		AddChainConfig(t, homeChainUni, getSelector(uni.chainID), p2pIDs, 1)
	}

	cfgs, err3 := homeChainUni.ccipConfig.GetAllChainConfigs(&bind.CallOpts{})
	require.NoError(t, err3)
	t.Logf("homechain_configs %+v", cfgs)
	require.Len(t, cfgs, numChains)

	t.Log("creating ocr3 jobs")
	for i := 0; i < len(nodes); i++ {
		err := nodes[i].app.Start(ctx)
		require.NoError(t, err)
		tApp := apps[i]
		t.Cleanup(func() {
			require.NoError(t, tApp.Stop())
		})

		jb := mustGetJobSpec(t, bootstrapP2PID, bootstrapPort, nodes[i].peerID, nodes[i].keybundle.ID())
		require.NoErrorf(t, tApp.AddJobV2(ctx, &jb), "Wasn't able to create ccip job for node %d", i)
	}

	// Create a DON for each chain
	for _, uni := range universes {
		// Add nodes and give them the capability
		t.Log("AddingDON for universe: ", uni.chainID)
		chainSelector := getSelector(uni.chainID)
		homeChainUni.AddDON(t,
			ccipCapabilityID,
			chainSelector,
			uni.offramp.Address().Bytes(),
			1, // f
			bootstrapP2PID,
			p2pIDs,
			oracles[uni.chainID],
		)
	}

	pingPongs := initializePingPongContracts(t, universes)
	for chainID, universe := range universes {
		for otherChain, pingPong := range pingPongs[chainID] {
			t.Log("PingPong From: ", chainID, " To: ", otherChain)
			_, err2 := pingPong.StartPingPong(universe.owner)
			require.NoError(t, err2)
			universe.backend.Commit()

			logIter, err3 := universe.onramp.FilterCCIPSendRequested(&bind.FilterOpts{Start: 0}, nil)
			require.NoError(t, err3)
			// Iterate until latest event
			for logIter.Next() {
			}
			log := logIter.Event
			require.Equal(t, getSelector(otherChain), log.DestChainSelector)
			require.Equal(t, pingPong.Address(), log.Message.Sender)
			chainPingPongAddr := pingPongs[otherChain][chainID].Address().Bytes()
			// With chain agnostic addresses we need to pad the address to the correct length if the receiver is zero prefixed
			paddedAddr := common.LeftPadBytes(chainPingPongAddr, len(log.Message.Receiver))
			require.Equal(t, paddedAddr, log.Message.Receiver)
			sink := make(chan *evm_2_evm_multi_offramp.EVM2EVMMultiOffRampCommitReportAccepted)
			subscipriton, err := universe.offramp.WatchCommitReportAccepted(&bind.WatchOpts{}, sink)
			require.NoError(t, err)

			for {
				select {
				case <-time.After(5 * time.Second):
					t.Log("Timed out waiting for commit report")
				case <-subscipriton.Err():
					t.Log("Error waiting for commit report")
				case report := <-sink:
					t.Log("Received commit report: ", report)
					break
				}
			}
		}
	}

}

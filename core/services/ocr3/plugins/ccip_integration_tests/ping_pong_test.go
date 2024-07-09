package ccip_integration_tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	chainsel "github.com/smartcontractkit/chain-selectors"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	pp "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ping_pong_demo"
)

var (
	chainAID = chainsel.TEST_90000001.EvmChainID
	chainBID = chainsel.TEST_90000002.EvmChainID
)

func TestPingPong(t *testing.T) {
	owner, chains := createChains(t, 4)

	homeChainUni, universes := deployContracts(t, owner, chains)
	fullyConnectCCIPContracts(t, owner, universes)
	_, err := homeChainUni.capabilityRegistry.AddCapabilities(owner, []kcr.CapabilitiesRegistryCapability{
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

	pingPongs := initializePingPongContracts(t, owner, universes)
	for chainID, universe := range universes {
		for otherChain, pingPong := range pingPongs[chainID] {
			println(otherChain)
			_, err = pingPong.StartPingPong(owner)
			require.NoError(t, err)
			universe.backend.Commit()
		}
	}
}

// InitializeContracts initializes ping pong contracts on all chains and
// connects them all to each other.
func initializePingPongContracts(
	t *testing.T,
	owner *bind.TransactOpts,
	chainUniverses map[uint64]onchainUniverse,
) map[uint64]map[uint64]*pp.PingPongDemo {
	pingPongs := make(map[uint64]map[uint64]*pp.PingPongDemo)
	chainIDs := maps.Keys(chainUniverses)
	// For each chain initialize N ping pong contracts, where N is the (number of chains - 1)
	for chainID, universe := range chainUniverses {
		chainsToConnectTo := filter(chainIDs, func(chainIDArg uint64) bool {
			return chainIDArg != chainID
		})
		pingPongs[chainID] = make(map[uint64]*pp.PingPongDemo)
		for _, chainToConnect := range chainsToConnectTo {
			backend := universe.backend
			pingPongAddr, _, _, err := pp.DeployPingPongDemo(owner, backend, universe.router.Address(), universe.linkToken.Address())
			require.NoError(t, err)
			backend.Commit()
			pingPong, err := pp.NewPingPongDemo(pingPongAddr, backend)
			require.NoError(t, err)
			universe.backend.Commit()
			// Fund the ping pong contract with LINK
			_, err = universe.linkToken.Transfer(owner, pingPong.Address(), Link(10))
			universe.backend.Commit()
			require.NoError(t, err)
			pingPongs[chainID][chainToConnect] = pingPong
		}
	}

	// Connect each ping pong contract to its counterpart on the other chain
	for chainID, universe := range chainUniverses {
		for chainToConnect, pingPong := range pingPongs[chainID] {
			println("Setting counterpart ping pong contract on chain", chainID, "to chain", chainToConnect)
			_, err := pingPong.SetCounterpart(
				owner,
				chainUniverses[chainToConnect].chainID,
				// This is the address of the ping pong contract on the other chain
				pingPongs[chainToConnect][chainID].Address(),
			)
			require.NoError(t, err)
			universe.backend.Commit()
		}
	}
	return pingPongs
}

package ccip_integration_tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	pp "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ping_pong_demo"
)

func TestPingPong(t *testing.T) {
	owner, chains := createChains(t, 4)

	homeChainUni, universes := setupUniverses(t, owner, chains)
	setupInitialConfigs(t, owner, universes, homeChainUni)

	pingPongs := initializePingPongContracts(t, owner, universes)
	for chainID, universe := range universes {
		for otherChain, pingPong := range pingPongs[chainID] {
			println("PingPong From: ", chainID, " To: ", otherChain)
			_, err := pingPong.StartPingPong(owner)
			require.NoError(t, err)
			nCommits := 100
			// Give time for the logPoller to catch up
			for i := 0; i < nCommits; i++ {
				universe.backend.Commit()
			}
			block, err := universe.logPoller.LatestBlock(testutils.Context(t))
			require.NoError(t, err)
			logs, err := universe.logPoller.Logs(testutils.Context(t), block.BlockNumber-int64(nCommits), block.BlockNumber,
				evm_2_evm_multi_onramp.EVM2EVMMultiOnRampCCIPSendRequested{}.Topic(), universe.onramp.Address())
			require.NoError(t, err)
			require.Len(t, logs, 1)
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
			_, err = universe.linkToken.Transfer(owner, pingPong.Address(), e18Mult(10))
			universe.backend.Commit()
			require.NoError(t, err)
			pingPongs[chainID][chainToConnect] = pingPong
		}
	}

	// Set up each ping pong contract to its counterpart on the other chain
	for chainID, universe := range chainUniverses {
		for chainToConnect, pingPong := range pingPongs[chainID] {
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

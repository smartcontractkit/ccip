package ccip_integration_tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/stretchr/testify/require"

	pp "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/link_token"
)

var (
	chainAID = chainsel.TEST_90000001.EvmChainID
	chainBID = chainsel.TEST_90000002.EvmChainID
)

func TestPingPong(t *testing.T) {
	owner, chainA := createChain(t)
	_, chainB := createChain(t)

	//====================================================InitializeContracts========================================
	pingPongA := initializeChainContracts(t, chainAID, owner, chainA)
	pingPongB := initializeChainContracts(t, chainBID, owner, chainB)

	//====================================================Prepare PingPongs========================================
	_, err := pingPongA.SetCounterpart(owner, chainBID, pingPongB.Address())
	require.NoError(t, err)
	_, err = pingPongB.SetCounterpart(owner, chainAID, pingPongA.Address())
	require.NoError(t, err)

	//====================================================Start PingPong========================================
	_, err = pingPongA.StartPingPong(owner)
	require.NoError(t, err)
	_, err = pingPongB.StartPingPong(owner)
	require.NoError(t, err)
}

func createChain(t *testing.T) (*bind.TransactOpts, *backends.SimulatedBackend) {
	owner := testutils.MustNewSimTransactor(t)

	chain := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.From: core.GenesisAccount{
			Balance: assets.Ether(10_000).ToInt(),
		},
	}, 30e6)

	chain.Commit()
	return owner, chain
}

func initializeChainContracts(t *testing.T,
	chainID uint64,
	owner *bind.TransactOpts,
	backend *backends.SimulatedBackend) *pp.PingPongDemo {
	linkAddr, _, _, err := link_token.DeployLinkToken(owner, backend)
	require.NoErrorf(t, err, "failed to deploy link token on chain id %d", chainID)
	backend.Commit()

	linkToken, err := link_token.NewLinkToken(linkAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, linkToken, nil)

	rmnAddr, _, _, err := mock_arm_contract.DeployMockARMContract(owner, backend)
	require.NoErrorf(t, err, "failed to deploy mock arm on chain id %d", chainID)
	backend.Commit()

	rmn, err := mock_arm_contract.NewMockARMContract(rmnAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, rmn, nil)

	rmnProxyAddr, _, _, err := arm_proxy_contract.DeployARMProxyContract(owner, backend, rmnAddr)
	require.NoErrorf(t, err, "failed to deploy arm proxy on chain id %d", chainID)
	backend.Commit()

	rmnProxy, err := arm_proxy_contract.NewARMProxyContract(rmnProxyAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, rmnProxy, nil)

	wethAddr, _, _, err := weth9.DeployWETH9(owner, backend)
	require.NoErrorf(t, err, "failed to deploy weth contract on chain id %d", chainID)
	backend.Commit()

	weth, err := weth9.NewWETH9(wethAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, weth, nil)

	routerAddr, _, _, err := router.DeployRouter(owner, backend, wethAddr, rmnProxyAddr)
	require.NoErrorf(t, err, "failed to deploy router on chain id %d", chainID)
	backend.Commit()

	rout, err := router.NewRouter(routerAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, rout, nil)

	pingPongAddr, _, _, err := pp.DeployPingPongDemo(owner, backend, routerAddr, linkAddr)
	require.NoError(t, err)
	pingPong, err := pp.NewPingPongDemo(pingPongAddr, backend)
	require.NotEqual(t, pingPongAddr, nil)

	return pingPong
}

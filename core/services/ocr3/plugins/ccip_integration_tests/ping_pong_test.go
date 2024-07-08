package ccip_integration_tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
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
	//ownerA, chainA := createChain(t)
	//ownerB, chainB := createChain(t)
	owner, chains := createChains(t, 4)

	//====================================================InitializeContracts========================================
	//pingPongA := initializeChainContracts(t, chainAID, ownerA, chainA)
	//pingPongB := initializeChainContracts(t, chainBID, ownerB, chainB)
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

	//====================================================Prepare PingPongs========================================
	//_, err := pingPongA.SetCounterpart(ownerA, chainBID, pingPongB.Address())
	//require.NoError(t, err)
	//_, err = pingPongB.SetCounterpart(ownerB, chainAID, pingPongA.Address())
	//require.NoError(t, err)
	//
	////====================================================Start PingPong========================================
	//_, err = pingPongA.StartPingPong(ownerA)
	//require.NoError(t, err)
	//_, err = pingPongB.StartPingPong(ownerB)
	//require.NoError(t, err)
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
	priceRegistryAddr, _, _, err := price_registry.DeployPriceRegistry(owner, backend, []common.Address{}, []common.Address{
		linkToken.Address(),
	}, 24*60*60, []price_registry.PriceRegistryTokenPriceFeedUpdate{})
	require.NoError(t, err, "failed to deploy price registry on chain id %d", chainID)
	backend.Commit()

	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, priceRegistry, nil)

	tarAddr, _, _, err := token_admin_registry.DeployTokenAdminRegistry(owner, backend)
	require.NoErrorf(t, err, "failed to deploy token admin registry on chain id %d", chainID)
	backend.Commit()

	tokenAdminRegistry, err := token_admin_registry.NewTokenAdminRegistry(tarAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, tokenAdminRegistry, nil)

	chainSelector, ok := chainsel.EvmChainIdToChainSelector()[uint64(chainID)]
	require.Truef(t, ok, "chain selector for chain id %d not found", chainID)

	onrampAddr, _, _, err := evm_2_evm_multi_onramp.DeployEVM2EVMMultiOnRamp(
		owner,
		backend,
		evm_2_evm_multi_onramp.EVM2EVMMultiOnRampStaticConfig{
			LinkToken:     linkAddr,
			ChainSelector: chainSelector,
			RmnProxy:      rmnProxyAddr,
		},
		evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDynamicConfig{
			Router:        routerAddr,
			PriceRegistry: priceRegistryAddr,
		},
		// can set this later once all chains are deployed
		[]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainConfigArgs{},
		// disabled for simplicity
		[]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs{},
		[]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampTokenTransferFeeConfigArgs{},
	)
	require.NoErrorf(t, err, "failed to deploy onramp on chain id %d", chainID)
	backend.Commit()

	onramp, err := evm_2_evm_multi_onramp.NewEVM2EVMMultiOnRamp(onrampAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, onramp, nil)

	offrampAddr, _, _, err := evm_2_evm_multi_offramp.DeployEVM2EVMMultiOffRamp(
		owner,
		backend,
		evm_2_evm_multi_offramp.EVM2EVMMultiOffRampStaticConfig{
			ChainSelector:      chainSelector,
			RmnProxy:           rmnProxyAddr,
			TokenAdminRegistry: tarAddr,
		},
		// can fill this in later once all chains are deployed
		[]evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{},
	)
	require.NoErrorf(t, err, "failed to deploy offramp on chain id %d", chainID)
	backend.Commit()

	offramp, err := evm_2_evm_multi_offramp.NewEVM2EVMMultiOffRamp(offrampAddr, backend)
	require.NoError(t, err)
	require.NotEqual(t, offramp, nil)

	pingPongAddr, _, _, err := pp.DeployPingPongDemo(owner, backend, routerAddr, linkAddr)
	require.NoError(t, err)
	pingPong, err := pp.NewPingPongDemo(pingPongAddr, backend)
	require.NotEqual(t, pingPongAddr, nil)

	return pingPong
}

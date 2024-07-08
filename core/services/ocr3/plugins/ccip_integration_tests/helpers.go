package ccip_integration_tests

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/nonce_manager"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/link_token"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

var (
	homeChainID = int64(chainsel.GETH_TESTNET.EvmChainID)
)

type ocr3Node struct {
	app          chainlink.Application
	peerID       string
	transmitters map[int64]common.Address
	keybundle    ocr2key.KeyBundle
	db           *sqlx.DB
}

type homeChain struct {
	backend            *backends.SimulatedBackend
	chainID            uint64
	capabilityRegistry *kcr.CapabilitiesRegistry
	ccipConfigContract common.Address // TODO: deploy
}

type onchainUniverse struct {
	backend            *backends.SimulatedBackend
	chainID            uint64
	linkToken          *link_token.LinkToken
	weth               *weth9.WETH9
	router             *router.Router
	rmnProxy           *arm_proxy_contract.ARMProxyContract
	rmn                *mock_arm_contract.MockARMContract
	onramp             *evm_2_evm_multi_onramp.EVM2EVMMultiOnRamp
	offramp            *evm_2_evm_multi_offramp.EVM2EVMMultiOffRamp
	priceRegistry      *price_registry.PriceRegistry
	tokenAdminRegistry *token_admin_registry.TokenAdminRegistry
}

func deployContracts(
	t *testing.T,
	owner *bind.TransactOpts,
	chains map[int64]*backends.SimulatedBackend,
) (homeChainUni homeChain, universes map[int64]onchainUniverse) {
	require.Len(t, chains, 4, "must have 4 chains total, 1 home chain and 3 non-home-chains")

	// deploy the capability registry on the home chain
	homeChainBackend, ok := chains[homeChainID]
	require.True(t, ok, "home chain backend not available")

	addr, _, _, err := kcr.DeployCapabilitiesRegistry(owner, homeChainBackend)
	require.NoError(t, err, "failed to deploy capability registry on home chain")
	homeChainBackend.Commit()

	capabilityRegistry, err := kcr.NewCapabilitiesRegistry(addr, homeChainBackend)
	require.NoError(t, err)

	// deploy the ccip contracts on the non-home-chain chains (total of 3).
	universes = make(map[int64]onchainUniverse)
	for chainID, backend := range chains {
		if chainID == homeChainID {
			continue
		}

		// contracts to deploy:
		// 0. link token
		// 1. onramp
		// 2. offramp
		// 3. price registry
		// 4. router
		// 5. rmn
		linkAddr, _, _, err := link_token.DeployLinkToken(owner, backend)
		require.NoErrorf(t, err, "failed to deploy link token on chain id %d", chainID)
		backend.Commit()

		linkToken, err := link_token.NewLinkToken(linkAddr, backend)
		require.NoError(t, err)

		rmnAddr, _, _, err := mock_arm_contract.DeployMockARMContract(owner, backend)
		require.NoErrorf(t, err, "failed to deploy mock arm on chain id %d", chainID)
		backend.Commit()

		rmn, err := mock_arm_contract.NewMockARMContract(rmnAddr, backend)
		require.NoError(t, err)

		rmnProxyAddr, _, _, err := arm_proxy_contract.DeployARMProxyContract(owner, backend, rmnAddr)
		require.NoErrorf(t, err, "failed to deploy arm proxy on chain id %d", chainID)
		backend.Commit()

		rmnProxy, err := arm_proxy_contract.NewARMProxyContract(rmnProxyAddr, backend)
		require.NoError(t, err)

		wethAddr, _, _, err := weth9.DeployWETH9(owner, backend)
		require.NoErrorf(t, err, "failed to deploy weth contract on chain id %d", chainID)
		backend.Commit()

		weth, err := weth9.NewWETH9(wethAddr, backend)
		require.NoError(t, err)

		routerAddr, _, _, err := router.DeployRouter(owner, backend, wethAddr, rmnProxyAddr)
		require.NoErrorf(t, err, "failed to deploy router on chain id %d", chainID)
		backend.Commit()

		rout, err := router.NewRouter(routerAddr, backend)
		require.NoError(t, err)

		priceRegistryAddr, _, _, err := price_registry.DeployPriceRegistry(owner, backend, []common.Address{}, []common.Address{
			linkToken.Address(),
		}, 24*60*60, []price_registry.PriceRegistryTokenPriceFeedUpdate{})
		require.NoError(t, err, "failed to deploy price registry on chain id %d", chainID)
		backend.Commit()

		priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddr, backend)
		require.NoError(t, err)

		tarAddr, _, _, err := token_admin_registry.DeployTokenAdminRegistry(owner, backend)
		require.NoErrorf(t, err, "failed to deploy token admin registry on chain id %d", chainID)
		backend.Commit()

		tokenAdminRegistry, err := token_admin_registry.NewTokenAdminRegistry(tarAddr, backend)
		require.NoError(t, err)

		chainSelector, ok := chainsel.EvmChainIdToChainSelector()[uint64(chainID)]
		require.Truef(t, ok, "chain selector for chain id %d not found", chainID)

		nonceManagerAddr, _, _, err := nonce_manager.DeployNonceManager(owner, backend, []common.Address{owner.From})
		nonceManager, err := nonce_manager.NewNonceManager(nonceManagerAddr, backend)
		require.NoErrorf(t, err, "failed to deploy nonce_manager on chain id %d", chainID)
		require.NotEqual(t, nonceManager, nil)
		backend.Commit()

		//`withdrawFeeTokens` onRamp function is not part of the message flow
		// so we can set this to any address
		feeAggregator := testutils.NewAddress()

		onrampAddr, _, _, err := evm_2_evm_multi_onramp.DeployEVM2EVMMultiOnRamp(
			owner,
			backend,
			evm_2_evm_multi_onramp.EVM2EVMMultiOnRampStaticConfig{
				LinkToken:          linkAddr,
				ChainSelector:      chainSelector,
				RmnProxy:           rmnProxyAddr,
				MaxFeeJuelsPerMsg:  big.NewInt(1e18),
				NonceManager:       nonceManagerAddr,
				TokenAdminRegistry: tarAddr,
			},
			evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDynamicConfig{
				Router:        routerAddr,
				PriceRegistry: priceRegistryAddr,
				FeeAggregator: feeAggregator,
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

		offrampAddr, _, _, err := evm_2_evm_multi_offramp.DeployEVM2EVMMultiOffRamp(
			owner,
			backend,
			evm_2_evm_multi_offramp.EVM2EVMMultiOffRampStaticConfig{
				ChainSelector:      chainSelector,
				RmnProxy:           rmnProxyAddr,
				TokenAdminRegistry: tarAddr,
				NonceManager:       nonceManagerAddr,
			},
			// can fill this in later once all chains are deployed
			[]evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{},
		)
		require.NoErrorf(t, err, "failed to deploy offramp on chain id %d", chainID)
		backend.Commit()

		offramp, err := evm_2_evm_multi_offramp.NewEVM2EVMMultiOffRamp(offrampAddr, backend)
		require.NoError(t, err)

		universes[chainID] = onchainUniverse{
			backend:            backend,
			chainID:            uint64(chainID),
			linkToken:          linkToken,
			weth:               weth,
			router:             rout,
			rmnProxy:           rmnProxy,
			rmn:                rmn,
			onramp:             onramp,
			offramp:            offramp,
			priceRegistry:      priceRegistry,
			tokenAdminRegistry: tokenAdminRegistry,
		}
	}

	return homeChain{
		backend:            homeChainBackend,
		chainID:            uint64(homeChainID),
		capabilityRegistry: capabilityRegistry,
	}, universes
}

func fullyConnectCCIPContracts(
	t *testing.T,
	owner *bind.TransactOpts,
	universes map[int64]onchainUniverse,
) {
	chainIDs := maps.Keys(universes)
	for chainID, uni := range universes {
		chainsToConnectTo := filter(chainIDs, func(chainIDArg int64) bool {
			return chainIDArg != chainID
		})

		// we are forming a fully-connected graph, so in each iteration we connect
		// the current chain (referenced by chainID) to all other chains that are not
		// ourselves.
		var (
			onrampDestChainConfigArgs    []evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainConfigArgs
			routerOnrampUpdates          []router.RouterOnRamp
			routerOfframpUpdates         []router.RouterOffRamp
			offrampSourceChainConfigArgs []evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs
		)
		for _, chainToConnect := range chainsToConnectTo {
			chainSelector, ok := chainsel.EvmChainIdToChainSelector()[uint64(chainToConnect)]
			require.Truef(t, ok, "chain selector not found for chain id %d", chainToConnect)
			onrampDestChainConfigArgs = append(onrampDestChainConfigArgs, evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainConfigArgs{
				DestChainSelector: chainSelector,
				DynamicConfig: evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainDynamicConfig{
					IsEnabled:                         true,
					MaxNumberOfTokensPerMsg:           10,
					MaxDataBytes:                      256,
					MaxPerMsgGasLimit:                 3_000_000,
					DestGasOverhead:                   50_000,
					DefaultTokenFeeUSDCents:           1,
					DestGasPerPayloadByte:             10,
					DestDataAvailabilityOverheadGas:   0,
					DestGasPerDataAvailabilityByte:    100,
					DestDataAvailabilityMultiplierBps: 1,
					DefaultTokenDestGasOverhead:       50_000,
					DefaultTokenDestBytesOverhead:     32,
					DefaultTxGasLimit:                 200_000,
					GasMultiplierWeiPerEth:            1,
					NetworkFeeUSDCents:                1,
				},
			})

			remoteUni, ok := universes[chainID]
			require.Truef(t, ok, "could not find universe for chain id %d", chainID)

			offrampSourceChainConfigArgs = append(offrampSourceChainConfigArgs, evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{
				SourceChainSelector: chainSelector,
				IsEnabled:           true,
				OnRamp:              remoteUni.onramp.Address(),
			})

			// onramps are multi-dest and offramps are multi-source.
			// so set the same ramp for all the chain selectors.
			routerOnrampUpdates = append(routerOnrampUpdates, router.RouterOnRamp{
				DestChainSelector: chainSelector,
				OnRamp:            uni.onramp.Address(),
			})
			routerOfframpUpdates = append(routerOfframpUpdates, router.RouterOffRamp{
				SourceChainSelector: chainSelector,
				OffRamp:             uni.offramp.Address(),
			})
		}

		_, err := uni.onramp.ApplyDestChainConfigUpdates(owner, onrampDestChainConfigArgs)
		require.NoErrorf(t, err, "failed to apply dest chain config updates on onramp on chain id %d", chainID)
		uni.backend.Commit()

		_, err = uni.offramp.ApplySourceChainConfigUpdates(owner, offrampSourceChainConfigArgs)
		require.NoErrorf(t, err, "failed to apply source chain config updates on offramp on chain id %d", chainID)
		uni.backend.Commit()

		_, err = uni.router.ApplyRampUpdates(owner, routerOnrampUpdates, []router.RouterOffRamp{}, routerOfframpUpdates)
		require.NoErrorf(t, err, "failed to apply ramp updates on router on chain id %d", chainID)
		uni.backend.Commit()
	}
}

func filter[T any](s []T, cond func(arg T) bool) (r []T) {
	for _, v := range s {
		if cond(v) {
			r = append(r, v)
		}
	}
	return
}

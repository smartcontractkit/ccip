package internal

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jmoiron/sqlx"
	chainsel "github.com/smartcontractkit/chain-selectors"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3confighelper"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	v2toml "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml"
	encodeutils "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	evmutils "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/keystone_capability_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/link_token"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/logger/audit"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/plugins"
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
	capabilityRegistry *keystone_capability_registry.CapabilityRegistry
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

	addr, _, _, err := keystone_capability_registry.DeployCapabilityRegistry(owner, homeChainBackend)
	require.NoError(t, err, "failed to deploy capability registry on home chain")
	homeChainBackend.Commit()

	capabilityRegistry, err := keystone_capability_registry.NewCapabilityRegistry(addr, homeChainBackend)
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
			evm_2_evm_multi_onramp.RateLimiterConfig{
				IsEnabled: false,
			},
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

// setupNodeOCR3 creates a chainlink node and any associated keys in order to run
// ccip.
func setupNodeOCR3(
	t *testing.T,
	owner *bind.TransactOpts,
	port int,
	chainIDToBackend map[int64]*backends.SimulatedBackend,
) *ocr3Node {
	// Do not want to load fixtures as they contain a dummy chainID.
	config, db := heavyweight.FullTestDBNoFixturesV2(t, func(c *chainlink.Config, s *chainlink.Secrets) {
		c.Insecure.OCRDevelopmentMode = ptr(true) // Disables ocr spec validation so we can have fast polling for the test.

		c.Feature.LogPoller = ptr(true)

		// P2P V2 configs.
		c.P2P.V2.Enabled = ptr(true)
		c.P2P.V2.DeltaDial = config.MustNewDuration(500 * time.Millisecond)
		c.P2P.V2.DeltaReconcile = config.MustNewDuration(5 * time.Second)
		c.P2P.V2.ListenAddresses = &[]string{fmt.Sprintf("127.0.0.1:%d", port)}

		// OCR configs
		c.OCR.Enabled = ptr(false)
		c.OCR.DefaultTransactionQueueDepth = ptr(uint32(200))
		c.OCR2.Enabled = ptr(true)

		c.EVM[0].LogPollInterval = config.MustNewDuration(500 * time.Millisecond)
		c.EVM[0].GasEstimator.LimitDefault = ptr[uint64](3_500_000)
		c.EVM[0].Transactions.ForwardersEnabled = ptr(false)
		c.OCR2.ContractPollInterval = config.MustNewDuration(5 * time.Second)

		var chains v2toml.EVMConfigs
		for chainID := range chainIDToBackend {
			chains = append(chains, createConfigV2Chain(big.NewInt(chainID)))
		}
		c.EVM = chains
		c.OCR2.ContractPollInterval = config.MustNewDuration(5 * time.Second)
	})

	lggr := logger.TestLogger(t)
	lggr.SetLogLevel(zapcore.InfoLevel)
	ctx := testutils.Context(t)
	clients := make(map[int64]client.Client)

	for chainID, backend := range chainIDToBackend {
		clients[chainID] = client.NewSimulatedBackendClient(t, backend, big.NewInt(chainID))
	}

	master := keystore.New(db, utils.FastScryptParams, lggr)

	keystore := KeystoreSim{
		eks: &EthKeystoreSim{
			Eth: master.Eth(),
			t:   t,
		},
		csa: master.CSA(),
	}
	mailMon := mailbox.NewMonitor("ccip", lggr.Named("mailbox"))
	evmOpts := chainlink.EVMFactoryConfig{
		ChainOpts: legacyevm.ChainOpts{
			AppConfig: config,
			GenEthClient: func(i *big.Int) client.Client {
				t.Log("genning eth client for chain id:", i.String())
				client, ok := clients[i.Int64()]
				if !ok {
					t.Fatal("no backend for chainID", i)
				}
				return client
			},
			MailMon: mailMon,
			DS:      db,
		},
		CSAETHKeystore: keystore,
	}
	relayerFactory := chainlink.RelayerFactory{
		Logger:       lggr,
		LoopRegistry: plugins.NewLoopRegistry(lggr.Named("LoopRegistry"), config.Tracing()),
		GRPCOpts:     loop.GRPCOpts{},
	}
	initOps := []chainlink.CoreRelayerChainInitFunc{chainlink.InitEVM(testutils.Context(t), relayerFactory, evmOpts)}
	rci, err := chainlink.NewCoreRelayerChainInteroperators(initOps...)
	require.NoError(t, err)

	app, err := chainlink.NewApplication(chainlink.ApplicationOpts{
		Config:                     config,
		DS:                         db,
		KeyStore:                   master,
		RelayerChainInteroperators: rci,
		Logger:                     lggr,
		ExternalInitiatorManager:   nil,
		CloseLogger:                lggr.Sync,
		UnrestrictedHTTPClient:     &http.Client{},
		RestrictedHTTPClient:       &http.Client{},
		AuditLogger:                audit.NoopLogger,
		MailMon:                    mailMon,
		LoopRegistry:               plugins.NewLoopRegistry(lggr, config.Tracing()),
	})
	require.NoError(t, err)
	require.NoError(t, app.GetKeyStore().Unlock(ctx, "password"))
	_, err = app.GetKeyStore().P2P().Create(ctx)
	require.NoError(t, err)

	p2pIDs, err := app.GetKeyStore().P2P().GetAll()
	require.NoError(t, err)
	require.Len(t, p2pIDs, 1)
	peerID := p2pIDs[0].PeerID()

	// create a transmitter for each chain
	transmitters := make(map[int64]common.Address)
	for chainID, backend := range chainIDToBackend {
		addrs, err2 := app.GetKeyStore().Eth().EnabledAddressesForChain(testutils.Context(t), big.NewInt(chainID))
		require.NoError(t, err2)
		if len(addrs) == 1 {
			// just fund the address
			fundAddress(t, owner, addrs[0], assets.Ether(10).ToInt(), backend)
			transmitters[chainID] = addrs[0]
		} else {
			// create key and fund it
			_, err3 := app.GetKeyStore().Eth().Create(testutils.Context(t), big.NewInt(chainID))
			require.NoError(t, err3, "failed to create key for chain", chainID)
			sendingKeys, err3 := app.GetKeyStore().Eth().EnabledAddressesForChain(testutils.Context(t), big.NewInt(chainID))
			require.NoError(t, err3)
			require.Len(t, sendingKeys, 1)
			fundAddress(t, owner, sendingKeys[0], assets.Ether(10).ToInt(), backend)
			transmitters[chainID] = sendingKeys[0]
		}
	}
	require.Len(t, transmitters, len(chainIDToBackend))

	keybundle, err := app.GetKeyStore().OCR2().Create(ctx, chaintype.EVM)
	require.NoError(t, err)

	return &ocr3Node{
		// can't use this app because it doesn't have the right toml config
		// missing bootstrapp
		// app:          app,
		peerID:       peerID.Raw(),
		transmitters: transmitters,
		keybundle:    keybundle,
		db:           db,
	}
}

func createChains(t *testing.T, numChains int) (owner *bind.TransactOpts, chains map[int64]*backends.SimulatedBackend) {
	owner = testutils.MustNewSimTransactor(t)
	chains = make(map[int64]*backends.SimulatedBackend)

	chains[homeChainID] = backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.From: core.GenesisAccount{
			Balance: assets.Ether(10_000).ToInt(),
		},
	}, 30e6)

	for chainID := int64(chainsel.TEST_90000001.EvmChainID); chainID < int64(chainsel.TEST_90000020.EvmChainID); chainID++ {
		chains[chainID] = backends.NewSimulatedBackend(core.GenesisAlloc{
			owner.From: core.GenesisAccount{
				Balance: assets.Ether(10000).ToInt(),
			},
		}, 30e6)

		if len(chains) == numChains {
			break
		}
	}
	return
}

func ptr[T any](v T) *T { return &v }

var _ keystore.Eth = &EthKeystoreSim{}

type EthKeystoreSim struct {
	keystore.Eth
	t *testing.T
}

// override
func (e *EthKeystoreSim) SignTx(ctx context.Context, address common.Address, tx *gethtypes.Transaction, chainID *big.Int) (*gethtypes.Transaction, error) {
	// always sign with chain id 1337 for the simulated backend
	e.t.Log("always signing tx for chain id:", chainID.String(), "with chain id 1337, tx hash:", tx.Hash())
	return e.Eth.SignTx(ctx, address, tx, big.NewInt(1337))
}

type KeystoreSim struct {
	eks keystore.Eth
	csa keystore.CSA
}

func (e KeystoreSim) Eth() keystore.Eth {
	return e.eks
}

func (e KeystoreSim) CSA() keystore.CSA {
	return e.csa
}

func fundAddress(t *testing.T, from *bind.TransactOpts, to common.Address, amount *big.Int, backend *backends.SimulatedBackend) {
	nonce, err := backend.PendingNonceAt(testutils.Context(t), from.From)
	require.NoError(t, err)
	gp, err := backend.SuggestGasPrice(testutils.Context(t))
	require.NoError(t, err)
	rawTx := gethtypes.NewTx(&gethtypes.LegacyTx{
		Nonce:    nonce,
		GasPrice: gp,
		Gas:      21000,
		To:       &to,
		Value:    amount,
	})
	signedTx, err := from.Signer(from.From, rawTx)
	require.NoError(t, err)
	err = backend.SendTransaction(testutils.Context(t), signedTx)
	require.NoError(t, err)
	backend.Commit()
}

func createConfigV2Chain(chainID *big.Int) *v2toml.EVMConfig {
	chain := v2toml.Defaults((*evmutils.Big)(chainID))
	chain.GasEstimator.LimitDefault = ptr(uint64(4e6))
	chain.LogPollInterval = config.MustNewDuration(500 * time.Millisecond)
	chain.Transactions.ForwardersEnabled = ptr(false)
	chain.FinalityDepth = ptr(uint32(2))
	return &v2toml.EVMConfig{
		ChainID: (*evmutils.Big)(chainID),
		Enabled: ptr(true),
		Chain:   chain,
		Nodes:   v2toml.EVMNodes{&v2toml.Node{}},
	}
}

func donOCRConfig(t *testing.T, uni onchainUniverse, oracles []confighelper2.OracleIdentityExtra) []byte {
	var schedule []int
	for range oracles {
		schedule = append(schedule, 1)
	}
	offchainConfig, onchainConfig := []byte{}, []byte{}
	f := uint8(1)
	_, _, f, _, offchainConfigVersion, offchainConfig, err := ocr3confighelper.ContractSetConfigArgsForTests(
		30*time.Second, // deltaProgress
		10*time.Second, // deltaResend
		20*time.Second, // deltaInitial
		2*time.Second,  // deltaRound
		20*time.Second, // deltaGrace
		10*time.Second, // deltaCertifiedCommitRequest
		10*time.Second, // deltaStage
		3,              // rmax
		schedule,
		oracles,
		offchainConfig,
		50*time.Millisecond, // maxDurationQuery
		5*time.Second,       // maxDurationObservation
		10*time.Second,      // maxDurationShouldAcceptAttestedReport
		10*time.Second,      // maxDurationShouldTransmitAcceptedReport
		int(f),
		onchainConfig)
	require.NoError(t, err, "failed to create contract config")
	/*
		struct OCR3Config {
			PluginType pluginType; // ────────╮ The plugin that the configuration is for.
			uint64 chainSelector; //          | The (remote) chain that the configuration is for.
			uint8 F; //                       | The "big F" parameter for the role DON.
			uint64 offchainConfigVersion; // ─╯ The version of the offchain configuration.
			bytes32 offrampAddress; // The remote chain combined (offramp|commit store) address.
			bytes32[2][] signers; // An associative array that contains (p2p id, onchain signer public key) pairs.
			bytes32[2][] transmitters; // An associative array that contains (p2p id, transmitter) pairs.
			bytes offchainConfig; // The offchain configuration for the OCR3 protocol. Protobuf encoded.
		}
	*/
	ocrConfigABI := `
{
	[
		{
			"type": "uint8"
		},
		{
			"type": "uint64"
		},
		{
			"type": "uint8"
		},
		{
			"type": "uint64"
		},
		{
			"type": "bytes32"
		},
		{
			"type": "bytes32[2][]"
		},
		{
			"type": "bytes32[2][]"
		},
		{
			"type": "bytes"
		}
	]
}`
	chainSelector, ok := chainsel.EvmChainIdToChainSelector()[uni.chainID]
	require.True(t, ok, "chain selector not found for chain id", uni.chainID)
	var offrampAddressBytes32 [32]byte
	copy(offrampAddressBytes32[:], uni.offramp.Address().Bytes())
	commitConfig, err := encodeutils.ABIEncode(ocrConfigABI,
		uint8(0),              // pluginType
		chainSelector,         // chainSelector
		f,                     // F
		offchainConfigVersion, // offchainConfigVersion
		offrampAddressBytes32, // offrampAddress
		nil,                   // TODO signers
		nil,                   // TODO transmitters
		offchainConfig,        // offchainConfig
	)
	require.NoError(t, err, "failed to encode commit OCR3 config")
	// TODO: implement
	return commitConfig
}

func createCCIPSpecToml(nodeP2PID, bootstrapP2PID string, bootstrapPort int, ocrKeyBundleID string) string {
	return fmt.Sprintf(`
type = "ccip"
capabilityVersion = "v1.0.0"
capabilityLabelledName = "ccip"
p2pKeyID = "%s"
p2pV2Bootstrappers = ["%s"]
[ocrKeyBundleIDs]
evm = "%s"
[relayConfigs.evm.chainReaderConfig.contracts.Offramp]
contractABI = "the abi"

[relayConfigs.evm.chainReaderConfig.contracts.Offramp.configs.getStuff]
chainSpecificName = "getStuffEVM"

[pluginConfig]
tokenPricesPipeline = "the pipeline"`,
		nodeP2PID,
		fmt.Sprintf("%s@127.0.0.1:%d", bootstrapP2PID,
			bootstrapPort,
		),
		ocrKeyBundleID,
	)
}

package ccip_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/smartcontractkit/libocr/commontypes"
	ocrnetworking "github.com/smartcontractkit/libocr/networking"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	eth "github.com/smartcontractkit/chainlink/core/chains/evm/client"
	evmclient "github.com/smartcontractkit/chainlink/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/core/chains/evm/headtracker"
	httypes "github.com/smartcontractkit/chainlink/core/chains/evm/headtracker/types"
	"github.com/smartcontractkit/chainlink/core/chains/evm/log"
	"github.com/smartcontractkit/chainlink/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/chains/terra"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/mock_v3_aggregator_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/sender_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/configtest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/validate"
	"github.com/smartcontractkit/chainlink/core/services/ocrbootstrap"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/services/pg"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func setupChain(t *testing.T) (*backends.SimulatedBackend, *bind.TransactOpts) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	chain := backends.NewSimulatedBackend(core.GenesisAlloc{
		user.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1000000000000000000))}},
		ethconfig.Defaults.Miner.GasCeil)
	return chain, user
}

type CCIPContracts struct {
	sourceUser, destUser           *bind.TransactOpts
	sourceChain, destChain         *backends.SimulatedBackend
	sourcePool, destPool           *native_token_pool.NativeTokenPool
	onRamp                         *onramp.OnRamp
	sourceLinkToken, destLinkToken *link_token_interface.LinkToken
	offRamp                        *offramp.OffRamp
	messageReceiver                *simple_message_receiver.SimpleMessageReceiver
	senderDapp                     *sender_dapp.SenderDapp
	receiverDapp                   *receiver_dapp.ReceiverDapp
	executor                       *message_executor.MessageExecutor
}

func setupCCIPContracts(t *testing.T) CCIPContracts {
	sourceChain, sourceUser := setupChain(t)
	destChain, destUser := setupChain(t)

	// Deploy link token and pool on source chain
	sourceLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()
	sourceLinkToken, err := link_token_interface.NewLinkToken(sourceLinkTokenAddress, sourceChain)
	require.NoError(t, err)
	sourcePoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(sourceUser,
		sourceChain,
		sourceLinkTokenAddress,
		native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		}, native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		})
	require.NoError(t, err)
	sourceChain.Commit()
	sourcePool, err := native_token_pool.NewNativeTokenPool(sourcePoolAddress, sourceChain)
	require.NoError(t, err)

	// Deploy link token and pool on destination chain
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	destLinkToken, err := link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
	destPoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(destUser, destChain, destLinkTokenAddress,
		native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		}, native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		})
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := native_token_pool.NewNativeTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Float the offramp pool with 1M juels
	// Dest user is the owner of the dest pool, so he can store
	o, err := destPool.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, destUser.From.String(), o.String())
	b, err := destLinkToken.BalanceOf(nil, destUser.From)
	require.NoError(t, err)
	t.Log("balance", b)
	_, err = destLinkToken.Approve(destUser, destPoolAddress, big.NewInt(1000000))
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.LockOrBurn(destUser, destUser.From, big.NewInt(1000000))
	require.NoError(t, err)

	afnSourceAddress, _, _, err := afn_contract.DeployAFNContract(
		sourceUser,
		sourceChain,
		[]common.Address{sourceUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	sourceChain.Commit()

	// LINK/ETH price
	feedAddress, _, _, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(sourceUser, sourceChain, 18, big.NewInt(6000000000000000))
	require.NoError(t, err)

	// Deploy onramp source chain
	onRampAddress, _, _, err := onramp.DeployOnRamp(
		sourceUser,                               // user
		sourceChain,                              // client
		sourceChainID,                            // source chain id
		[]*big.Int{destChainID},                  // destinationChainIds
		[]common.Address{sourceLinkTokenAddress}, // tokens
		[]common.Address{sourcePoolAddress},      // pools
		[]common.Address{feedAddress},            // Feeds
		[]common.Address{},                       // allow list
		afnSourceAddress,                         // AFN
		big.NewInt(86400),                        //maxTimeWithoutAFNSignal 86400 seconds = one day
		onramp.OnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e12,
			MaxTokensLength:  5,
		},
	)
	require.NoError(t, err)
	// We do this so onRamp.Address() works
	onRamp, err := onramp.NewOnRamp(onRampAddress, sourceChain)
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, onRampAddress, true)
	require.NoError(t, err)

	afnDestAddress, _, _, err := afn_contract.DeployAFNContract(
		destUser,
		destChain,
		[]common.Address{destUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	destChain.Commit()

	feedDestAddress, _, _, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(destUser, destChain, 18, big.NewInt(6000000000000000))
	require.NoError(t, err)

	// Deploy offramp dest chain
	offRampAddress, _, _, err := offramp.DeployOffRamp(
		destUser,                                 // user
		destChain,                                // client
		sourceChainID,                            // source chain id
		destChainID,                              // dest chain id
		[]common.Address{sourceLinkTokenAddress}, // source tokens
		[]common.Address{destPoolAddress},        // dest pool addresses
		[]common.Address{feedDestAddress},        // feeds
		afnDestAddress,                           // AFN address
		big.NewInt(86400),                        // max timeout without AFN signal  86400 seconds = one day
		offramp.OffRampInterfaceOffRampConfig{
			ExecutionFeeJuels:     0,
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
	)
	require.NoError(t, err)
	offRamp, err := offramp.NewOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	// Set the pool to be the offramp
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	destChain.Commit()

	// Deploy offramp contract token receiver
	messageReceiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	messageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(messageReceiverAddress, destChain)
	require.NoError(t, err)
	// Deploy offramp token receiver dapp
	receiverDappAddress, _, _, err := receiver_dapp.DeployReceiverDapp(destUser, destChain, offRampAddress, destLinkTokenAddress)
	require.NoError(t, err)
	eoaTokenReceiver, err := receiver_dapp.NewReceiverDapp(receiverDappAddress, destChain)
	require.NoError(t, err)
	// Deploy onramp token sender dapp
	senderDappAddress, _, _, err := sender_dapp.DeploySenderDapp(sourceUser, sourceChain, onRampAddress, destChainID, receiverDappAddress)
	require.NoError(t, err)
	eoaTokenSender, err := sender_dapp.NewSenderDapp(senderDappAddress, sourceChain)
	require.NoError(t, err)

	// Need to commit here, or we will hit the block gas limit when deploying the executor
	sourceChain.Commit()
	destChain.Commit()

	// Deploy the message executor ocr2 contract
	executorAddress, _, _, err := message_executor.DeployMessageExecutor(destUser, destChain, offRampAddress, false)
	require.NoError(t, err)
	executor, err := message_executor.NewMessageExecutor(executorAddress, destChain)
	require.NoError(t, err)

	sourceChain.Commit()
	destChain.Commit()

	return CCIPContracts{
		sourceUser:      sourceUser,
		destUser:        destUser,
		sourceChain:     sourceChain,
		destChain:       destChain,
		sourcePool:      sourcePool,
		destPool:        destPool,
		onRamp:          onRamp,
		sourceLinkToken: sourceLinkToken,
		destLinkToken:   destLinkToken,
		offRamp:         offRamp,
		messageReceiver: messageReceiver,
		receiverDapp:    eoaTokenReceiver,
		senderDapp:      eoaTokenSender,
		executor:        executor,
	}
}

var (
	sourceChainID = big.NewInt(1000)
	destChainID   = big.NewInt(1337)
)

type EthKeyStoreSim struct {
	keystore.Eth
}

func (ks EthKeyStoreSim) SignTx(address common.Address, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if chainID.String() == "1000" {
		// A terrible hack, just for the multichain test. All simulation clients run on chainID 1337.
		// We let the destChain actually use 1337 to make sure the config digests are properly generated.
		return ks.Eth.SignTx(address, tx, big.NewInt(1337))
	}
	return ks.Eth.SignTx(address, tx, chainID)
}

var _ keystore.Eth = EthKeyStoreSim{}

type testCheckerFactory struct {
	err error
}

func (t *testCheckerFactory) BuildChecker(spec txmgr.TransmitCheckerSpec) (txmgr.TransmitChecker, error) {
	return &testChecker{t.err}, nil
}

type testChecker struct {
	err error
}

func (t *testChecker) Check(
	_ context.Context,
	_ logger.Logger,
	_ txmgr.EthTx,
	_ txmgr.EthTxAttempt,
) error {
	return t.err
}

func setupNodeCCIP(t *testing.T, owner *bind.TransactOpts, port int64, dbName string, sourceChain *backends.SimulatedBackend, destChain *backends.SimulatedBackend) (chainlink.Application, string, common.Address, ocr2key.KeyBundle, *configtest.TestGeneralConfig, func()) {
	p2paddresses := []string{
		fmt.Sprintf("127.0.0.1:%d", port),
	}
	// Do not want to load fixtures as they contain a dummy chainID.
	config, db := heavyweight.FullTestDB(t, fmt.Sprintf("%s%d", dbName, port), true, false)
	config.Overrides.FeatureOffchainReporting = null.BoolFrom(false)
	config.Overrides.FeatureOffchainReporting2 = null.BoolFrom(true)
	config.Overrides.FeatureCCIP = null.BoolFrom(true)
	config.Overrides.GlobalGasEstimatorMode = null.NewString("FixedPrice", true)
	config.Overrides.SetP2PV2DeltaDial(500 * time.Millisecond)
	config.Overrides.SetP2PV2DeltaReconcile(5 * time.Second)
	config.Overrides.DefaultChainID = nil
	config.Overrides.P2PListenPort = null.NewInt(0, true)
	config.Overrides.P2PV2ListenAddresses = p2paddresses
	config.Overrides.P2PV2AnnounceAddresses = p2paddresses
	config.Overrides.P2PNetworkingStack = ocrnetworking.NetworkingStackV2
	// Disables ocr spec validation so we can have fast polling for the test.
	config.Overrides.Dev = null.BoolFrom(true)

	var lggr = logger.TestLogger(t)
	eventBroadcaster := pg.NewEventBroadcaster(config.DatabaseURL(), 0, 0, lggr, uuid.NewV1())

	// We fake different chainIDs using the wrapped sim cltest.SimulatedBackend
	chainORM := evm.NewORM(db, lggr, config)
	_, err := chainORM.CreateChain(*utils.NewBig(sourceChainID), evmtypes.ChainCfg{})
	require.NoError(t, err)
	_, err = chainORM.CreateChain(*utils.NewBig(destChainID), evmtypes.ChainCfg{})
	require.NoError(t, err)
	sourceClient := evmclient.NewSimulatedBackendClient(t, sourceChain, sourceChainID)
	destClient := evmclient.NewSimulatedBackendClient(t, destChain, destChainID)

	keyStore := keystore.New(db, utils.FastScryptParams, lggr, config)
	simEthKeyStore := EthKeyStoreSim{Eth: keyStore.Eth()}
	cfg := cltest.NewTestGeneralConfig(t)
	evmCfg := evmtest.NewChainScopedConfig(t, cfg)
	checkerFactory := &testCheckerFactory{}

	// Create our chainset manually so we can have custom eth clients
	// (the wrapped sims faking different chainIDs)
	evmChain, err := evm.LoadChainSet(evm.ChainSetOpts{
		ORM:              chainORM,
		Config:           config,
		Logger:           lggr,
		DB:               db,
		KeyStore:         simEthKeyStore,
		EventBroadcaster: eventBroadcaster,
		GenEthClient: func(c evmtypes.Chain) eth.Client {
			if c.ID.String() == sourceChainID.String() {
				return sourceClient
			} else if c.ID.String() == destChainID.String() {
				return destClient
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenHeadTracker: func(c evmtypes.Chain, hb httypes.HeadBroadcaster) httypes.HeadTracker {
			if c.ID.String() == sourceChainID.String() {
				return headtracker.NewHeadTracker(
					lggr, sourceClient,
					evmtest.NewChainScopedConfig(t, config),
					hb,
					headtracker.NewHeadSaver(lggr, headtracker.NewORM(db, lggr, pgtest.NewPGCfg(false), *sourceClient.ChainID()), evmCfg),
				)
			} else if c.ID.String() == destChainID.String() {
				return headtracker.NewHeadTracker(
					lggr,
					destClient,
					evmtest.NewChainScopedConfig(t, config),
					hb,
					headtracker.NewHeadSaver(lggr, headtracker.NewORM(db, lggr, pgtest.NewPGCfg(false), *destClient.ChainID()), evmCfg),
				)
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenLogBroadcaster: func(c evmtypes.Chain) log.Broadcaster {
			if c.ID.String() == sourceChainID.String() {
				t.Log("Generating log broadcaster source")
				return log.NewBroadcaster(log.NewORM(db, lggr, config, *sourceChainID), sourceClient,
					evmtest.NewChainScopedConfig(t, config), lggr, nil)
			} else if c.ID.String() == destChainID.String() {
				return log.NewBroadcaster(log.NewORM(db, lggr, config, *destChainID), destClient,
					evmtest.NewChainScopedConfig(t, config), lggr, nil)
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenTxManager: func(c evmtypes.Chain) txmgr.TxManager {
			if c.ID.String() == sourceChainID.String() {
				return txmgr.NewTxm(db, sourceClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr, checkerFactory)
			} else if c.ID.String() == destChainID.String() {
				return txmgr.NewTxm(db, destClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr, checkerFactory)
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
	})
	if err != nil {
		lggr.Fatal(err)
	}
	terraChain, err := terra.NewChainSet(terra.ChainSetOpts{
		Config:           cfg,
		Logger:           lggr,
		DB:               db,
		KeyStore:         keyStore.Terra(),
		EventBroadcaster: eventBroadcaster,
		ORM:              terra.NewORM(db, lggr, cfg),
	})

	app, err := chainlink.NewApplication(chainlink.ApplicationOpts{
		Config:           config,
		EventBroadcaster: eventBroadcaster,
		SqlxDB:           db,
		KeyStore:         keyStore,
		Chains: chainlink.Chains{
			EVM:   evmChain,
			Terra: terraChain,
		},
		Logger:                   lggr,
		ExternalInitiatorManager: nil,
		CloseLogger: func() error {
			return nil
		},
	})
	require.NoError(t, err)
	require.NoError(t, app.GetKeyStore().Unlock("password"))
	_, err = app.GetKeyStore().P2P().Create()
	require.NoError(t, err)

	p2pIDs, err := app.GetKeyStore().P2P().GetAll()
	require.NoError(t, err)
	require.Len(t, p2pIDs, 1)
	peerID := p2pIDs[0].PeerID()
	config.Overrides.P2PPeerID = peerID

	_, err = app.GetKeyStore().Eth().Create(destChainID)
	require.NoError(t, err)
	sendingKeys, err := app.GetKeyStore().Eth().SendingKeys()
	require.NoError(t, err)
	require.Len(t, sendingKeys, 1)
	transmitter := sendingKeys[0].Address.Address()
	s, err := app.GetKeyStore().Eth().GetState(sendingKeys[0].ID())
	require.NoError(t, err)
	lggr.Debug(fmt.Sprintf("Transmitter address %s chainID %s", transmitter, s.EVMChainID.String()))

	// Fund the relayTransmitter address with some ETH
	n, err := destChain.NonceAt(context.Background(), owner.From, nil)
	require.NoError(t, err)

	tx := types.NewTransaction(n, transmitter, big.NewInt(1000000000000000000), 21000, big.NewInt(1000000000), nil)
	signedTx, err := owner.Signer(owner.From, tx)
	require.NoError(t, err)
	err = destChain.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)
	destChain.Commit()

	kb, err := app.GetKeyStore().OCR2().Create(chaintype.EVM)
	require.NoError(t, err)
	return app, peerID.Raw(), transmitter, kb, config, func() {
		app.Stop()
	}
}

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := setupCCIPContracts(t)
	lggr := logger.TestLogger(t)
	// Oracles need ETH on the destination chain
	bootstrapNodePort := int64(19599)
	appBootstrap, bootstrapPeerID, _, _, _, _ := setupNodeCCIP(t, ccipContracts.destUser, bootstrapNodePort, "bootstrap_ccip", ccipContracts.sourceChain, ccipContracts.destChain)
	var (
		oracles      []confighelper2.OracleIdentityExtra
		transmitters []common.Address
		kbs          []ocr2key.KeyBundle
		apps         []chainlink.Application
	)
	// Set up the minimum 4 oracles all funded with destination ETH
	for i := int64(0); i < 4; i++ {
		app, peerID, transmitter, kb, cfg, _ := setupNodeCCIP(t, ccipContracts.destUser, bootstrapNodePort+1+i, fmt.Sprintf("oracle_ccip%d", i), ccipContracts.sourceChain, ccipContracts.destChain)
		// Supply the bootstrap IP and port as a V2 peer address
		cfg.Overrides.P2PV2Bootstrappers = []commontypes.BootstrapperLocator{
			{PeerID: bootstrapPeerID, Addrs: []string{
				fmt.Sprintf("127.0.0.1:%d", bootstrapNodePort),
			}},
		}
		kbs = append(kbs, kb)
		apps = append(apps, app)
		transmitters = append(transmitters, transmitter)
		offchainPublicKey, _ := hex.DecodeString(strings.TrimPrefix(kb.OnChainPublicKey(), "0x"))
		oracles = append(oracles, confighelper2.OracleIdentityExtra{
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  offchainPublicKey,
				TransmitAccount:   ocrtypes2.Account(transmitter.String()),
				OffchainPublicKey: kb.OffchainPublicKey(),
				PeerID:            peerID,
			},
			ConfigEncryptionPublicKey: kb.ConfigEncryptionPublicKey(),
		})
	}

	reportingPluginConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   1,
	}.Encode()
	require.NoError(t, err)

	setupOnchainConfig(t, ccipContracts, oracles, reportingPluginConfig)
	ctx := context.Background()
	err = appBootstrap.Start(ctx)
	require.NoError(t, err)
	defer appBootstrap.Stop()

	// Add the bootstrap job
	chainSet := appBootstrap.GetChains()
	require.NotNil(t, chainSet)
	ocrJob, err := ocrbootstrap.ValidatedBootstrapSpecToml(fmt.Sprintf(`
type               	= "bootstrap"
relay 				= "evm"
schemaVersion      	= 1
name               	= "boot"
contractID    	    = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"
[relayConfig]
chainID = %s
`, ccipContracts.offRamp.Address(), destChainID))
	require.NoError(t, err)
	err = appBootstrap.AddJobV2(context.Background(), &ocrJob)
	require.NoError(t, err)

	// For each oracle add a relayer and job
	for i := 0; i < 4; i++ {
		err = apps[i].Start(ctx)
		require.NoError(t, err)
		defer apps[i].Stop()
		// Wait for peer wrapper to start
		time.Sleep(1 * time.Second)
		ccipJob, err := validate.ValidatedOracleSpecToml(apps[0].GetConfig(), fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-relay"
relay 				= "evm"
schemaVersion      	= 1
name               	= "ccip-job-%d"
contractID 			= "%s"
ocrKeyBundleID      = "%s"
transmitterID 		= "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampID            = "%s"
sourceChainID       = %s
destChainID         = %s

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.offRamp.Address(), kbs[i].ID(), transmitters[i], ccipContracts.onRamp.Address(), sourceChainID, destChainID, destChainID))
		require.NoError(t, err)
		err = apps[i].AddJobV2(context.Background(), &ccipJob)
		require.NoError(t, err)
		// Add executor job
		ccipExecutionJob, err := validate.ValidatedOracleSpecToml(apps[0].GetConfig(), fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-execution"
relay 				= "evm"
schemaVersion      	= 1
name               	= "ccip-executor-job-%d"
contractID 			= "%s"
ocrKeyBundleID      = "%s"
transmitterID 		= "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampID            = "%s"
offRampID           = "%s"
sourceChainID       = %s
destChainID         = %s

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.executor.Address(), kbs[i].ID(), transmitters[i], ccipContracts.onRamp.Address(), ccipContracts.offRamp.Address(), sourceChainID, destChainID, destChainID))
		require.NoError(t, err)
		err = apps[i].AddJobV2(context.Background(), &ccipExecutionJob)
		require.NoError(t, err)
	}
	// Send a request.
	// Jobs are booting but that is ok, the log broadcaster
	// will backfill this request log.
	ccipContracts.sourceUser.GasLimit = 500000
	_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.onRamp.Address(), big.NewInt(100))
	ccipContracts.sourceChain.Commit()
	msg := onramp.CCIPMessagePayload{
		Receiver:           ccipContracts.messageReceiver.Address(),
		DestinationChainId: destChainID,
		Data:               []byte("hello xchain world"),
		Tokens:             []common.Address{ccipContracts.sourceLinkToken.Address()},
		Amounts:            []*big.Int{big.NewInt(100)},
		Options:            nil,
	}
	tx, err := ccipContracts.onRamp.RequestCrossChainSend(ccipContracts.sourceUser, msg)
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()
	rec, err := ccipContracts.sourceChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), rec.Status)

	reportingPluginConfig, err = ccip.OffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   0,
	}.Encode()
	require.NoError(t, err)

	setupOnchainConfig(t, ccipContracts, oracles, reportingPluginConfig)

	// Request should appear on all nodes eventually
	for i := 0; i < 4; i++ {
		var reqs []*ccip.Request
		ccipReqORM := ccip.NewORM(apps[i].GetSqlxDB(), lggr, pgtest.NewPGCfg(false))
		gomega.NewGomegaWithT(t).Eventually(func() bool {
			ccipContracts.sourceChain.Commit()
			reqs, err = ccipReqORM.Requests(sourceChainID, destChainID, big.NewInt(0), nil, ccip.RequestStatusUnstarted, nil, nil)
			return len(reqs) == 1
		}, 5*time.Second, 1*time.Second).Should(gomega.BeTrue())
	}

	// Once all nodes have the request, the reporting plugin should run to generate and submit a report onchain.
	// So we should eventually see a successful offramp submission.
	// Note that since we only send blocks here, it's likely that all the nodes will enter the transmission
	// phase before someone has submitted, so 1 report will succeed and 3 will revert.
	var report offramp.CCIPRelayReport
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		report, err = ccipContracts.offRamp.GetLastReport(nil)
		require.NoError(t, err)
		ccipContracts.destChain.Commit()
		t.Log("last report", report.MinSequenceNumber.String(), report.MaxSequenceNumber.String())
		return report.MinSequenceNumber.String() == "1" && report.MaxSequenceNumber.String() == "1"
	}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())

	// We should see the request in a fulfilled state on all nodes
	// after the offramp submission. There should be no
	// remaining valid requests.
	for i := 0; i < 4; i++ {
		gomega.NewGomegaWithT(t).Eventually(func() bool {
			ccipReqORM := ccip.NewORM(apps[i].GetSqlxDB(), lggr, pgtest.NewPGCfg(false))
			ccipContracts.destChain.Commit()
			reqs, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, report.MaxSequenceNumber, ccip.RequestStatusRelayConfirmed, nil, nil)
			require.NoError(t, err)
			valid, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, nil, ccip.RequestStatusUnstarted, nil, nil)
			require.NoError(t, err)
			return len(reqs) == 1 && len(valid) == 0
		}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())
	}

	// Now the merkle root is across.
	// Let's try to execute a request as an external party.
	// The raw log in the merkle root should be the abi-encoded version of the CCIPMessage
	ccipReqORM := ccip.NewORM(apps[0].GetSqlxDB(), lggr, pgtest.NewPGCfg(false))
	reqs, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, report.MaxSequenceNumber, "", nil, nil)
	require.NoError(t, err)
	root, proof := ccip.GenerateMerkleProof(32, [][]byte{reqs[0].Raw}, 0)
	// Root should match the report root
	require.True(t, bytes.Equal(root[:], report.MerkleRoot[:]))

	// Path should verify.
	genRoot := ccip.GenerateMerkleRoot(reqs[0].Raw, proof)
	require.True(t, bytes.Equal(root[:], genRoot[:]))
	exists, err := ccipContracts.offRamp.GetMerkleRoot(nil, report.MerkleRoot)
	require.NoError(t, err)
	require.True(t, exists.Int64() > 0)

	h, err := utils.Keccak256(append([]byte{0x00}, reqs[0].Raw...))
	var leaf [32]byte
	copy(leaf[:], h)
	decodedMsg, err := ccip.DecodeCCIPMessage(reqs[0].Raw)
	require.NoError(t, err)
	offRampProof := offramp.CCIPMerkleProof{
		Path:  proof.PathForExecute(),
		Index: proof.Index(),
	}
	onchainRoot, err := ccipContracts.offRamp.MerkleRoot(nil, *decodedMsg, offRampProof)
	require.NoError(t, err)
	require.Equal(t, genRoot, onchainRoot)

	// Execute the Message
	tx, err = ccipContracts.offRamp.ExecuteTransaction(ccipContracts.destUser, *decodedMsg, offRampProof, false)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()

	// We should now have the Message in the offchain receiver
	receivedMsg, err := ccipContracts.messageReceiver.SMessage(nil)
	require.NoError(t, err)
	assert.Equal(t, "hello xchain world", string(receivedMsg.Payload.Data))

	// Now let's send an EOA to EOA request
	// We can just use the sourceUser and destUser
	startBalanceSource, err := ccipContracts.sourceLinkToken.BalanceOf(nil, ccipContracts.sourceUser.From)
	require.NoError(t, err)
	startBalanceDest, err := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destUser.From)
	require.NoError(t, err)
	t.Log(startBalanceSource, startBalanceDest)

	ccipContracts.sourceUser.GasLimit = 500000
	// Approve the sender contract to take the tokens
	_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.senderDapp.Address(), big.NewInt(100))
	ccipContracts.sourceChain.Commit()
	// Send the tokens. Should invoke the onramp.
	// Only the destUser can execute.
	tx, err = ccipContracts.senderDapp.SendTokens(ccipContracts.sourceUser, ccipContracts.destUser.From, []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, ccipContracts.destUser.From)
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()

	// DON should eventually send another report
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		report, err = ccipContracts.offRamp.GetLastReport(nil)
		require.NoError(t, err)
		ccipContracts.destChain.Commit()
		return report.MinSequenceNumber.String() == "2" && report.MaxSequenceNumber.String() == "2"
	}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())

	eoaReq, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, report.MaxSequenceNumber, "", nil, nil)
	require.NoError(t, err)
	root, proof = ccip.GenerateMerkleProof(32, [][]byte{eoaReq[0].Raw}, 0)
	// Root should match the report root
	require.True(t, bytes.Equal(root[:], report.MerkleRoot[:]))

	// Execute the Message
	decodedMsg, err = ccip.DecodeCCIPMessage(eoaReq[0].Raw)
	require.NoError(t, err)
	ccip.MakeCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
	tx, err = ccipContracts.offRamp.ExecuteTransaction(ccipContracts.destUser, *decodedMsg, offramp.CCIPMerkleProof{
		Path:  proof.PathForExecute(),
		Index: proof.Index(),
	}, false)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()

	// The destination user's balance should increase
	endBalanceSource, err := ccipContracts.sourceLinkToken.BalanceOf(nil, ccipContracts.sourceUser.From)
	require.NoError(t, err)
	endBalanceDest, err := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destUser.From)
	require.NoError(t, err)
	t.Log("Start balances", startBalanceSource, startBalanceDest)
	t.Log("End balances", endBalanceSource, endBalanceDest)
	assert.Equal(t, "100", big.NewInt(0).Sub(startBalanceSource, endBalanceSource).String())
	assert.Equal(t, "100", big.NewInt(0).Sub(endBalanceDest, startBalanceDest).String())

	// Now let's send a request flagged for oracle execution
	_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.onRamp.Address(), big.NewInt(100))
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()
	require.NoError(t, err)
	msg = onramp.CCIPMessagePayload{
		Receiver:           ccipContracts.messageReceiver.Address(),
		Data:               []byte("hey DON, execute for me"),
		Tokens:             []common.Address{ccipContracts.sourceLinkToken.Address()},
		Amounts:            []*big.Int{big.NewInt(100)},
		DestinationChainId: destChainID,
		Executor:           ccipContracts.executor.Address(),
		Options:            []byte{},
	}
	_, err = ccipContracts.onRamp.RequestCrossChainSend(ccipContracts.sourceUser, msg)
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()

	// Should first be relayed, seq number 3
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		report, err = ccipContracts.offRamp.GetLastReport(nil)
		require.NoError(t, err)
		ccipContracts.destChain.Commit()
		return report.MinSequenceNumber.String() == "3" && report.MaxSequenceNumber.String() == "3"
	}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())

	// Should see the 3rd message be executed
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		it, err := ccipContracts.offRamp.FilterCrossChainMessageExecuted(nil, nil)
		require.NoError(t, err)
		ecount := 0
		for it.Next() {
			t.Log("executed", it.Event.SequenceNumber)
			ecount++
		}
		ccipContracts.destChain.Commit()
		return ecount == 3
	}, 20*time.Second, 1*time.Second).Should(gomega.BeTrue())
	// In total, we should see 3 relay reports containing seq 1,2,3
	// and 3 execution_confirmed messages
	reqs, err = ccipReqORM.Requests(sourceChainID, destChainID, big.NewInt(1), big.NewInt(3), ccip.RequestStatusExecutionConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqs, 3)
	_, err = ccipReqORM.RelayReport(big.NewInt(1))
	require.NoError(t, err)
	_, err = ccipReqORM.RelayReport(big.NewInt(2))
	require.NoError(t, err)
	_, err = ccipReqORM.RelayReport(big.NewInt(3))
	require.NoError(t, err)
}

func setupOnchainConfig(t *testing.T, ccipContracts CCIPContracts, oracles []confighelper2.OracleIdentityExtra, reportingPluginConfig []byte) {
	// Note We do NOT set the payees, payment is done in the OCR2Base implementation
	// Set the offramp config.
	signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		2*time.Second,        // deltaProgress
		1*time.Second,        // deltaResend
		1*time.Second,        // deltaRound
		500*time.Millisecond, // deltaGrace
		2*time.Second,        // deltaStage
		3,
		[]int{1, 1, 1, 1},
		oracles,
		reportingPluginConfig,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		1, // faults
		nil,
	)

	require.NoError(t, err)
	lggr := logger.TestLogger(t)
	lggr.Infow("Setting Config on Oracle Contract",
		"signers", signers,
		"transmitters", transmitters,
		"threshold", threshold,
		"onchainConfig", onchainConfig,
		"encodedConfigVersion", offchainConfigVersion,
	)

	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	require.NoError(t, err)
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	require.NoError(t, err)

	// Set the DON on the offramp
	_, err = ccipContracts.offRamp.SetConfig(
		ccipContracts.destUser,
		signerAddresses,
		transmitterAddresses,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()

	// Same DON on the message executor
	_, err = ccipContracts.executor.SetConfig(
		ccipContracts.destUser,
		signerAddresses,
		transmitterAddresses,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()
}

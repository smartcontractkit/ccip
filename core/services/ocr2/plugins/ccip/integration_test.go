package ccip_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
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
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/mock_v3_aggregator_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_router"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp_router"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/sender_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/configtest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
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
	onRampRouter                   *onramp_router.OnRampRouter
	sourceLinkToken, destLinkToken *link_token_interface.LinkToken
	offRamp                        *offramp.OffRamp
	offRampRouter                  *offramp_router.OffRampRouter
	messageReceiver                *simple_message_receiver.SimpleMessageReceiver
	senderDapp                     *sender_dapp.SenderDapp
	receiverDapp                   *receiver_dapp.ReceiverDapp
	executor                       *offramp_executor.OffRampExecutor
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
			Rate:     big.NewInt(1e9),
			Capacity: big.NewInt(1e18),
		}, native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1e9),
			Capacity: big.NewInt(1e18),
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
			Capacity: big.NewInt(1e18),
		}, native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e18),
		})
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := native_token_pool.NewNativeTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	releaseBucket, err := destPool.GetReleaseOrMintBucket(nil)
	require.NoError(t, err)
	t.Logf("release bucket %v %v\n", releaseBucket.Tokens, releaseBucket.Capacity)

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
	// Create onramp router
	onRampRouterAddress, _, _, err := onramp_router.DeployOnRampRouter(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()

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
			Router:           onRampRouterAddress,
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
	sourceChain.Commit()
	onRampRouter, err := onramp_router.NewOnRampRouter(onRampRouterAddress, sourceChain)
	require.NoError(t, err)
	_, err = onRampRouter.SetOnRamp(sourceUser, destChainID, onRampAddress)
	require.NoError(t, err)
	sourceChain.Commit()

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
		// We set this above the current unix timestamp
		// so we do not even have to send a heartbeat for it to be healthy.
		big.NewInt(time.Now().Unix()*2),
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
	// Create offRamp router
	offRampRouterAddress, _, offRampRouter, err := offramp_router.DeployOffRampRouter(destUser, destChain, []common.Address{offRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	_, err = offRamp.SetRouter(destUser, offRampRouterAddress)
	require.NoError(t, err)
	offRampRouter, err = offramp_router.NewOffRampRouter(offRampRouterAddress, destChain)
	require.NoError(t, err)

	// Deploy offramp contract token receiver
	messageReceiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	messageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(messageReceiverAddress, destChain)
	require.NoError(t, err)
	// Deploy offramp token receiver dapp
	receiverDappAddress, _, _, err := receiver_dapp.DeployReceiverDapp(destUser, destChain, offRampRouterAddress, destLinkTokenAddress)
	require.NoError(t, err)
	eoaTokenReceiver, err := receiver_dapp.NewReceiverDapp(receiverDappAddress, destChain)
	require.NoError(t, err)
	// Deploy onramp token sender dapp
	senderDappAddress, _, _, err := sender_dapp.DeploySenderDapp(sourceUser, sourceChain, onRampRouterAddress, destChainID, receiverDappAddress)
	require.NoError(t, err)
	eoaTokenSender, err := sender_dapp.NewSenderDapp(senderDappAddress, sourceChain)
	require.NoError(t, err)

	// Need to commit here, or we will hit the block gas limit when deploying the executor
	sourceChain.Commit()
	destChain.Commit()

	// Deploy the message executor ocr2 contract
	executorAddress, _, _, err := offramp_executor.DeployOffRampExecutor(destUser, destChain, offRampAddress, false)
	require.NoError(t, err)
	executor, err := offramp_executor.NewOffRampExecutor(executorAddress, destChain)
	require.NoError(t, err)

	sourceChain.Commit()
	destChain.Commit()

	// Ensure we have at least finality blocks.
	for i := 0; i < 50; i++ {
		sourceChain.Commit()
		destChain.Commit()
	}
	// (practically) Infinite approval source user (owner of all the link) to onramp router.
	sourceUser.GasLimit = 500000
	_, err = sourceLinkToken.Approve(sourceUser, onRampRouter.Address(), big.NewInt(math.MaxInt64))
	sourceChain.Commit()

	return CCIPContracts{
		sourceUser:      sourceUser,
		destUser:        destUser,
		sourceChain:     sourceChain,
		destChain:       destChain,
		sourcePool:      sourcePool,
		destPool:        destPool,
		onRamp:          onRamp,
		onRampRouter:    onRampRouter,
		sourceLinkToken: sourceLinkToken,
		destLinkToken:   destLinkToken,
		offRamp:         offRamp,
		offRampRouter:   offRampRouter,
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
		// We let the destChain actually use 1337 to make sure the offchainConfig digests are properly generated.
		return ks.Eth.SignTx(address, tx, big.NewInt(1337))
	}
	return ks.Eth.SignTx(address, tx, chainID)
}

var _ keystore.Eth = EthKeyStoreSim{}

func setupNodeCCIP(t *testing.T, owner *bind.TransactOpts, port int64, dbName string, sourceChain *backends.SimulatedBackend, destChain *backends.SimulatedBackend) (chainlink.Application, string, common.Address, ocr2key.KeyBundle, *configtest.TestGeneralConfig, func()) {
	p2paddresses := []string{
		fmt.Sprintf("127.0.0.1:%d", port),
	}
	// Do not want to load fixtures as they contain a dummy chainID.
	config, db := heavyweight.FullTestDBNoFixtures(t, fmt.Sprintf("%s%d", dbName, port))
	config.Overrides.FeatureOffchainReporting = null.BoolFrom(false)
	config.Overrides.FeatureOffchainReporting2 = null.BoolFrom(true)
	config.Overrides.FeatureCCIP = null.BoolFrom(true)
	config.Overrides.FeatureLogPoller = null.BoolFrom(true)
	config.Overrides.GlobalGasEstimatorMode = null.NewString("FixedPrice", true)
	config.Overrides.SetP2PV2DeltaDial(500 * time.Millisecond)
	config.Overrides.SetP2PV2DeltaReconcile(5 * time.Second)
	config.Overrides.DefaultChainID = nil
	config.Overrides.P2PListenPort = null.NewInt(0, true)
	config.Overrides.P2PV2ListenAddresses = p2paddresses
	config.Overrides.P2PV2AnnounceAddresses = p2paddresses
	config.Overrides.P2PNetworkingStack = ocrnetworking.NetworkingStackV2
	// NOTE: For the executor jobs, the default of 500k is insufficient for a 3 message batch
	config.Overrides.GlobalEvmGasLimitDefault = null.NewInt(600000, true)
	// Disables ocr spec validation so we can have fast polling for the test.
	config.Overrides.Dev = null.BoolFrom(true)

	var lggr = logger.TestLogger(t)
	eventBroadcaster := pg.NewEventBroadcaster(config.DatabaseURL(), 0, 0, lggr, uuid.NewV1())

	// We fake different chainIDs using the wrapped sim cltest.SimulatedBackend
	chainORM := evm.NewORM(db, lggr, config)
	_, err := chainORM.CreateChain(*utils.NewBig(sourceChainID), &evmtypes.ChainCfg{})
	require.NoError(t, err)
	_, err = chainORM.CreateChain(*utils.NewBig(destChainID), &evmtypes.ChainCfg{})
	require.NoError(t, err)
	sourceClient := evmclient.NewSimulatedBackendClient(t, sourceChain, sourceChainID)
	destClient := evmclient.NewSimulatedBackendClient(t, destChain, destChainID)

	keyStore := keystore.New(db, utils.FastScryptParams, lggr, config)
	simEthKeyStore := EthKeyStoreSim{Eth: keyStore.Eth()}
	cfg := cltest.NewTestGeneralConfig(t)
	evmCfg := evmtest.NewChainScopedConfig(t, cfg)

	// Create our chainset manually so we can have custom eth clients
	// (the wrapped sims faking different chainIDs)
	var (
		sourceLp logpoller.LogPoller = logpoller.NewLogPoller(logpoller.NewORM(sourceChainID, db, lggr, config), sourceClient,
			lggr, 500*time.Millisecond, 10, 2)
		destLp logpoller.LogPoller = logpoller.NewLogPoller(logpoller.NewORM(destChainID, db, lggr, config), destClient,
			lggr, 500*time.Millisecond, 10, 2)
	)
	evmChain, err := evm.LoadChainSet(evm.ChainSetOpts{
		ORM:              chainORM,
		Config:           config,
		Logger:           lggr,
		DB:               db,
		KeyStore:         simEthKeyStore,
		EventBroadcaster: eventBroadcaster,
		GenEthClient: func(c evmtypes.DBChain) eth.Client {
			if c.ID.String() == sourceChainID.String() {
				return sourceClient
			} else if c.ID.String() == destChainID.String() {
				return destClient
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenHeadTracker: func(c evmtypes.DBChain, hb httypes.HeadBroadcaster) httypes.HeadTracker {
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
		GenLogPoller: func(c evmtypes.DBChain) logpoller.LogPoller {
			if c.ID.String() == sourceChainID.String() {
				t.Log("Generating log broadcaster source")
				return sourceLp
			} else if c.ID.String() == destChainID.String() {
				return destLp
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenLogBroadcaster: func(c evmtypes.DBChain) log.Broadcaster {
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
		GenTxManager: func(c evmtypes.DBChain) txmgr.TxManager {
			if c.ID.String() == sourceChainID.String() {
				return txmgr.NewTxm(db, sourceClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr, &txmgr.CheckerFactory{sourceClient}, sourceLp)
			} else if c.ID.String() == destChainID.String() {
				return txmgr.NewTxm(db, destClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr, &txmgr.CheckerFactory{destClient}, destLp)
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
	})
	if err != nil {
		lggr.Fatal(err)
	}
	app, err := chainlink.NewApplication(chainlink.ApplicationOpts{
		Config:           config,
		EventBroadcaster: eventBroadcaster,
		SqlxDB:           db,
		KeyStore:         keyStore,
		Chains: chainlink.Chains{
			EVM: evmChain,
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
	sendingKeys, err := app.GetKeyStore().Eth().SendingKeys(destChainID)
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

type Node struct {
	app         chainlink.Application
	transmitter common.Address
	kb          ocr2key.KeyBundle
}

func (node *Node) eventuallyHasReqSeqNum(t *testing.T, ccipContracts CCIPContracts, seqNum int) logpoller.Log {
	c, err := node.app.GetChains().EVM.Get(sourceChainID)
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.sourceChain.Commit()
		ccipContracts.destChain.Commit()
		lgs, err := c.LogPoller().LogsDataWordRange(ccip.CrossChainSendRequested, ccipContracts.onRamp.Address(), 2, ccip.EvmWord(uint64(seqNum)), ccip.EvmWord(uint64(seqNum)), 1)
		require.NoError(t, err)
		if len(lgs) == 1 {
			log = lgs[0]
			return true
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue())
	return log
}

func (node *Node) eventuallyHasExecutedSeqNum(t *testing.T, ccipContracts CCIPContracts, seqNum int) logpoller.Log {
	c, err := node.app.GetChains().EVM.Get(destChainID)
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.sourceChain.Commit()
		ccipContracts.destChain.Commit()
		lgs, err := c.LogPoller().IndexedLogsTopicRange(ccip.CrossChainMessageExecuted, ccipContracts.offRamp.Address(), 1, ccip.EvmWord(uint64(seqNum)), ccip.EvmWord(uint64(seqNum)), 1)
		require.NoError(t, err)
		if len(lgs) == 1 {
			log = lgs[0]
			return true
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue())
	return log
}

func (node *Node) addJob(t *testing.T, spec string) {
	ccipJob, err := validate.ValidatedOracleSpecToml(node.app.GetConfig(), spec)
	require.NoError(t, err)
	err = node.app.AddJobV2(context.Background(), &ccipJob)
	require.NoError(t, err)
}

func (node *Node) addBootstrapJob(t *testing.T, spec string) {
	ccipJob, err := ocrbootstrap.ValidatedBootstrapSpecToml(spec)
	require.NoError(t, err)
	err = node.app.AddJobV2(context.Background(), &ccipJob)
	require.NoError(t, err)
}

func allNodesHaveReqSeqNum(t *testing.T, ccipContracts CCIPContracts, nodes []Node, seqNum int) logpoller.Log {
	var log logpoller.Log
	for _, node := range nodes {
		log = node.eventuallyHasReqSeqNum(t, ccipContracts, seqNum)
	}
	return log
}

func allNodesHaveExecutedSeqNum(t *testing.T, ccipContracts CCIPContracts, nodes []Node, seqNum int) logpoller.Log {
	var log logpoller.Log
	for _, node := range nodes {
		log = node.eventuallyHasExecutedSeqNum(t, ccipContracts, seqNum)
	}
	return log
}

func queueRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, executor common.Address) *gethtypes.Transaction {
	msg := onramp_router.CCIPMessagePayload{
		Receiver:           ccipContracts.messageReceiver.Address(),
		DestinationChainId: destChainID,
		Data:               []byte(msgPayload),
		Tokens:             tokens,
		Amounts:            amounts,
		Executor:           executor,
	}
	tx, err := ccipContracts.onRampRouter.RequestCrossChainSend(ccipContracts.sourceUser, msg)
	require.NoError(t, err)
	return tx
}

func confirmTxs(t *testing.T, txs []*gethtypes.Transaction, chain *backends.SimulatedBackend) {
	chain.Commit()
	for _, tx := range txs {
		rec, err := chain.TransactionReceipt(context.Background(), tx.Hash())
		require.NoError(t, err)
		require.Equal(t, uint64(1), rec.Status)
	}
}

func sendRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, executor common.Address) {
	tx := queueRequest(t, ccipContracts, msgPayload, tokens, amounts, executor)
	confirmTxs(t, []*gethtypes.Transaction{tx}, ccipContracts.sourceChain)
}

func eventuallyReportRelayed(t *testing.T, ccipContracts CCIPContracts, min, max int) offramp.CCIPRelayReport {
	var report offramp.CCIPRelayReport
	var err error
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		report, err = ccipContracts.offRamp.GetLastReport(nil)
		require.NoError(t, err)
		ccipContracts.sourceChain.Commit()
		ccipContracts.destChain.Commit()
		t.Log("last report", report.MinSequenceNumber, report.MaxSequenceNumber)
		return report.MinSequenceNumber == uint64(min) && report.MaxSequenceNumber == uint64(max)
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue())
	return report
}

func executeMessage(t *testing.T, ccipContracts CCIPContracts, req logpoller.Log, allReqs []logpoller.Log, report offramp.CCIPRelayReport) {
	// Double check root exists onchain.
	exists, err := ccipContracts.offRamp.GetMerkleRoot(nil, report.MerkleRoot)
	require.NoError(t, err)
	require.True(t, exists.Int64() > 0)

	// Build full tree for report
	mctx := merklemulti.NewKeccakCtx()
	var leafHashes [][32]byte
	for _, otherReq := range allReqs {
		leafHashes = append(leafHashes, mctx.HashLeaf(otherReq.Data))
	}
	tree := merklemulti.NewTree(mctx, leafHashes)

	// Generate our proof in the execution ramp report format.
	decodedMsg, err := ccip.DecodeCCIPMessage(req.Data)
	require.NoError(t, err)
	index := decodedMsg.SequenceNumber - report.MinSequenceNumber
	proof := tree.Prove([]int{int(index)})
	require.Equal(t, tree.Root(), report.MerkleRoot)
	require.NoError(t, err, "hashes %v index %d", proof.Hashes, index)
	offRampProof := offramp.CCIPExecutionReport{
		Messages:       []offramp.CCIPMessage{*decodedMsg},
		Proofs:         proof.Hashes,
		ProofFlagsBits: ccip.ProofFlagsToBits(proof.SourceFlags),
	}
	onchainRoot, err := ccipContracts.offRamp.MerkleRoot(nil, offRampProof)
	require.NoError(t, err)
	require.Equal(t, tree.Root(), onchainRoot)

	// Execute.
	tx, err := ccipContracts.offRamp.ExecuteTransaction(ccipContracts.destUser, offRampProof, false)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()
	rec, err := ccipContracts.destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), rec.Status)
}

func setupAndStartNodes(t *testing.T, ctx context.Context, ccipContracts CCIPContracts, bootstrapNodePort int64) (Node, []Node) {
	appBootstrap, bootstrapPeerID, bootstrapTransmitter, bootstrapKb, _, _ := setupNodeCCIP(t, ccipContracts.destUser, bootstrapNodePort, "bootstrap_ccip", ccipContracts.sourceChain, ccipContracts.destChain)
	var (
		oracles []confighelper2.OracleIdentityExtra
		nodes   []Node
	)
	err := appBootstrap.Start(ctx)
	require.NoError(t, err)
	t.Cleanup(func() {
		appBootstrap.Stop()
	})
	bootstrapNode := Node{
		appBootstrap, bootstrapTransmitter, bootstrapKb,
	}
	// Set up the minimum 4 oracles all funded with destination ETH
	for i := int64(0); i < 4; i++ {
		app, peerID, transmitter, kb, cfg, _ := setupNodeCCIP(t, ccipContracts.destUser, bootstrapNodePort+1+i, fmt.Sprintf("oracle_ccip%d", i), ccipContracts.sourceChain, ccipContracts.destChain)
		// Supply the bootstrap IP and port as a V2 peer address
		cfg.Overrides.P2PV2Bootstrappers = []commontypes.BootstrapperLocator{
			{PeerID: bootstrapPeerID, Addrs: []string{
				fmt.Sprintf("127.0.0.1:%d", bootstrapNodePort),
			}},
		}
		nodes = append(nodes, Node{
			app, transmitter, kb,
		})
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
		err := app.Start(ctx)
		require.NoError(t, err)
		t.Cleanup(func() {
			app.Stop()
		})
	}

	reportingPluginConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   1,
	}.Encode()
	require.NoError(t, err)
	setupOnchainConfig(t, ccipContracts, oracles, reportingPluginConfig)
	return bootstrapNode, nodes
}

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := setupCCIPContracts(t)
	bootstrapNodePort := int64(19599)
	ctx := context.Background()
	// Starts nodes and configures them in the OCR contracts.
	bootstrapNode, nodes := setupAndStartNodes(t, ctx, ccipContracts, bootstrapNodePort)

	// Add the bootstrap job
	bootstrapNode.addBootstrapJob(t, fmt.Sprintf(`
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

	// For each node add a relayer and executor job.
	for i, node := range nodes {
		node.addJob(t, fmt.Sprintf(`
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
pollPeriod          = "1s"

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.offRamp.Address(), node.kb.ID(), node.transmitter, ccipContracts.onRamp.Address(), sourceChainID, destChainID, destChainID))
		node.addJob(t, fmt.Sprintf(`
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
pollPeriod          = "1s"

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.executor.Address(), node.kb.ID(), node.transmitter, ccipContracts.onRamp.Address(), ccipContracts.offRamp.Address(), sourceChainID, destChainID, destChainID))
	}
	// With jobs present, replay the config log.
	b, err := ccipContracts.destChain.BlockByNumber(context.Background(), nil)
	require.NoError(t, err)
	for _, node := range nodes {
		c, err := node.app.GetChains().EVM.Get(destChainID)
		require.NoError(t, err)
		require.NoError(t, c.LogPoller().Replay(context.Background(), int64(b.NumberU64()-1)))
	}

	currentSeqNum := 1
	t.Run("single self-execute", func(t *testing.T) {
		sendRequest(t, ccipContracts, "single req", []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, common.Address{})
		req := allNodesHaveReqSeqNum(t, ccipContracts, nodes, currentSeqNum)
		report := eventuallyReportRelayed(t, ccipContracts, currentSeqNum, currentSeqNum)
		executeMessage(t, ccipContracts, req, []logpoller.Log{req}, report)
		allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, currentSeqNum)
		receivedMsg, err := ccipContracts.messageReceiver.SMessage(nil)
		require.NoError(t, err)
		assert.Equal(t, "single req", string(receivedMsg.Payload.Data))
		currentSeqNum++
	})

	t.Run("batch self-execute", func(t *testing.T) {
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			txs = append(txs, queueRequest(t, ccipContracts, fmt.Sprintf("batch request %d", currentSeqNum+i), []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, common.Address{}))
		}
		// Send a batch of requests in a single block
		confirmTxs(t, txs, ccipContracts.sourceChain)
		// All nodes should have all 3.
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, allNodesHaveReqSeqNum(t, ccipContracts, nodes, currentSeqNum+i))
		}
		// Should see a report with the full range
		report := eventuallyReportRelayed(t, ccipContracts, currentSeqNum, currentSeqNum+n-1)
		// Execute them all
		for _, req := range reqs {
			executeMessage(t, ccipContracts, req, reqs, report)
		}
		for i := range reqs {
			allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, currentSeqNum+i)
		}
		receivedMsg, err := ccipContracts.messageReceiver.SMessage(nil)
		require.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("batch request %d", currentSeqNum+n-1), string(receivedMsg.Payload.Data))
		currentSeqNum += n
	})

	t.Run("single auto-execute", func(t *testing.T) {
		sendRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, ccipContracts.executor.Address())
		allNodesHaveReqSeqNum(t, ccipContracts, nodes, currentSeqNum)
		eventuallyReportRelayed(t, ccipContracts, currentSeqNum, currentSeqNum)
		allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, currentSeqNum)
		currentSeqNum++
	})

	t.Run("batch auto-execute", func(t *testing.T) {
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			txs = append(txs, queueRequest(t, ccipContracts, fmt.Sprintf("batch request %d", currentSeqNum+i), []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, ccipContracts.executor.Address()))
		}
		// Send a batch of requests in a single block
		confirmTxs(t, txs, ccipContracts.sourceChain)
		// All nodes should have all 3.
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, allNodesHaveReqSeqNum(t, ccipContracts, nodes, currentSeqNum+i))
		}
		// Should see a report with the full range
		eventuallyReportRelayed(t, ccipContracts, currentSeqNum, currentSeqNum+n-1)
		// Should all be executed
		for i := range reqs {
			allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, currentSeqNum+i)
		}
		currentSeqNum += n
	})

	t.Run("eoa2eoa", func(t *testing.T) {
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
		_, err = ccipContracts.senderDapp.SendTokens(ccipContracts.sourceUser, ccipContracts.destUser.From, []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, ccipContracts.destUser.From)
		require.NoError(t, err)
		ccipContracts.sourceChain.Commit()

		req2 := allNodesHaveReqSeqNum(t, ccipContracts, nodes, currentSeqNum)
		report2 := eventuallyReportRelayed(t, ccipContracts, currentSeqNum, currentSeqNum)
		executeMessage(t, ccipContracts, req2, []logpoller.Log{req2}, report2)
		allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, currentSeqNum)

		// The destination user's balance should increase
		endBalanceSource, err := ccipContracts.sourceLinkToken.BalanceOf(nil, ccipContracts.sourceUser.From)
		require.NoError(t, err)
		endBalanceDest, err := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destUser.From)
		require.NoError(t, err)
		assert.Equal(t, "100", big.NewInt(0).Sub(startBalanceSource, endBalanceSource).String())
		assert.Equal(t, "100", big.NewInt(0).Sub(endBalanceDest, startBalanceDest).String())
	})
}

func setupOnchainConfig(t *testing.T, ccipContracts CCIPContracts, oracles []confighelper2.OracleIdentityExtra, reportingPluginConfig []byte) {
	// Note We do NOT set the payees, payment is done in the OCR2Base implementation
	// Set the offramp offchainConfig.
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

package ccip_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
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
	"github.com/smartcontractkit/chainlink/core/chains/evm/client"
	eth "github.com/smartcontractkit/chainlink/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/core/chains/evm/headtracker"
	httypes "github.com/smartcontractkit/chainlink/core/chains/evm/headtracker/types"
	"github.com/smartcontractkit/chainlink/core/chains/evm/log"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
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
		user.From: {Balance: big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(1e18))}},
		ethconfig.Defaults.Miner.GasCeil)
	return chain, user
}

type MaybeRevertReceiver struct {
	Receiver *maybe_revert_message_receiver.MaybeRevertMessageReceiver
	Strict   bool
}

type CCIPContracts struct {
	sourceUser, destUser           *bind.TransactOpts
	sourceChain, destChain         *backends.SimulatedBackend
	sourcePool, destPool           *native_token_pool.NativeTokenPool
	sourceLinkToken, destLinkToken *link_token_interface.LinkToken
	blobVerifier                   *blob_verifier.BlobVerifier
	receivers                      []MaybeRevertReceiver

	// Toll contracts
	tollOnRampFees    map[common.Address]*big.Int
	tollOnRampRouter  *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter
	tollOnRamp        *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp
	tollOffRampRouter *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter
	tollOffRamp       *any_2_evm_toll_offramp.EVM2EVMTollOffRamp

	// Sub contracts
	subOnRampFee     *big.Int
	subOnRampRouter  *evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouter
	subOnRamp        *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRamp
	subOffRampRouter *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter
	subOffRamp       *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp
}

func setupCCIPContracts(t *testing.T) CCIPContracts {
	sourceChain, sourceUser := setupChain(t)
	destChain, destUser := setupChain(t)

	var hundredLink = big.NewInt(0)
	// 100 LINK
	hundredLink.SetString("100000000000000000000", 10)

	// Deploy link token and pool on source chain
	sourceLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()
	sourceLinkToken, err := link_token_interface.NewLinkToken(sourceLinkTokenAddress, sourceChain)
	require.NoError(t, err)
	sourcePoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(sourceUser,
		sourceChain,
		sourceLinkTokenAddress)
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
	destPoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := native_token_pool.NewNativeTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Float the offramp pool
	o, err := destPool.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, destUser.From.String(), o.String())
	_, err = destLinkToken.Transfer(destUser, destPoolAddress, hundredLink)
	require.NoError(t, err)
	destChain.Commit()

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

	// Create onramp router
	onRampRouterAddress, _, _, err := evm_2_any_toll_onramp_router.DeployEVM2AnyTollOnRampRouter(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()

	// Deploy onramp source chain
	onRampAddress, _, _, err := evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
		sourceUser,                               // user
		sourceChain,                              // client
		sourceChainID,                            // source chain id
		destChainID,                              // destinationChainIds
		[]common.Address{sourceLinkTokenAddress}, // tokens
		[]common.Address{sourcePoolAddress},      // pools
		[]common.Address{},                       // allow list
		afnSourceAddress,                         // AFN
		evm_2_evm_toll_onramp.BaseOnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e12,
			MaxTokensLength:  5,
		},
		evm_2_evm_toll_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
		onRampRouterAddress,
	)
	require.NoError(t, err)
	tollOnRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(onRampAddress, sourceChain)
	require.NoError(t, err)
	tollOnRampFees := map[common.Address]*big.Int{
		sourceLinkTokenAddress: big.NewInt(1), // 1 juel relay fee.
	}
	_, err = tollOnRamp.SetFeeConfig(sourceUser, evm_2_evm_toll_onramp.Any2EVMTollOnRampInterfaceFeeConfig{
		Fees:      []*big.Int{tollOnRampFees[sourceLinkTokenAddress]},
		FeeTokens: []common.Address{sourceLinkTokenAddress},
	})
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, onRampAddress, true)
	require.NoError(t, err)
	sourceChain.Commit()
	_, err = tollOnRamp.SetPrices(sourceUser, []common.Address{sourceLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)
	tollOnRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(onRampRouterAddress, sourceChain)
	require.NoError(t, err)
	_, err = tollOnRampRouter.SetOnRamp(sourceUser, destChainID, onRampAddress)
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

	// Deploy offramp dest chain
	blobVerifierAddress, _, _, err := blob_verifier.DeployBlobVerifier(
		destUser,    // user
		destChain,   // client
		destChainID, // dest chain id
		sourceChainID,
		afnDestAddress, // AFN address
		blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{},
	)
	require.NoError(t, err)
	blobVerifier, err := blob_verifier.NewBlobVerifier(blobVerifierAddress, destChain)
	require.NoError(t, err)
	// Set the pool to be the offramp
	destChain.Commit()
	offRampAddress, _, _, err := any_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
		destUser,
		destChain,
		sourceChainID,
		destChainID,
		any_2_evm_toll_offramp.BaseOffRampInterfaceOffRampConfig{
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
		blobVerifier.Address(),
		onRampAddress,
		afnDestAddress,
		[]common.Address{sourceLinkTokenAddress},
		[]common.Address{destPoolAddress},
		any_2_evm_toll_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
	)
	require.NoError(t, err)
	tollOffRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	// Create offRampAddr router
	offRampRouterAddress, _, tollOffRampRouter, err := any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(destUser, destChain, []common.Address{offRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	tollOffRampRouter, err = any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(offRampRouterAddress, destChain)
	require.NoError(t, err)
	_, err = tollOffRamp.SetRouter(destUser, offRampRouterAddress)
	require.NoError(t, err)
	_, err = tollOffRamp.SetPrices(destUser, []common.Address{destLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	// Deploy 2 revertable (one SS one non-SS)
	revertingMessageReceiver1Address, _, _, err := maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(destUser, destChain, false)
	require.NoError(t, err)
	revertingMessageReceiver1, _ := maybe_revert_message_receiver.NewMaybeRevertMessageReceiver(revertingMessageReceiver1Address, destChain)
	revertingMessageReceiver2Address, _, _, err := maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(destUser, destChain, false)
	require.NoError(t, err)
	revertingMessageReceiver2, _ := maybe_revert_message_receiver.NewMaybeRevertMessageReceiver(revertingMessageReceiver2Address, destChain)
	// Need to commit here, or we will hit the block gas limit when deploying the executor
	sourceChain.Commit()
	destChain.Commit()

	// Setup subscription contracts.
	subOnRampRouterAddress, _, _, err := evm_2_any_subscription_onramp_router.DeployEVM2AnySubscriptionOnRampRouter(
		sourceUser, sourceChain, evm_2_any_subscription_onramp_router.Any2EVMSubscriptionOnRampRouterInterfaceRouterConfig{
			Fee:      big.NewInt(0),
			FeeToken: sourceLinkTokenAddress,
			FeeAdmin: sourceUser.From,
		})
	require.NoError(t, err)
	subOnRampRouter, _ := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(subOnRampRouterAddress, sourceChain)
	subOnRampAddress, _, _, err := evm_2_evm_subscription_onramp.DeployEVM2EVMSubscriptionOnRamp(sourceUser, sourceChain, sourceChainID, destChainID,
		[]common.Address{sourceLinkTokenAddress}, // tokens
		[]common.Address{sourcePoolAddress},      // pools
		[]common.Address{},                       // allow list
		afnSourceAddress,                         // AFN
		evm_2_evm_subscription_onramp.BaseOnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e12,
			MaxTokensLength:  5,
		},
		evm_2_evm_subscription_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
		onRampRouterAddress,
	)
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, subOnRampAddress, true)
	require.NoError(t, err)
	subOnRamp, _ := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(subOnRampAddress, sourceChain)
	_, err = subOnRamp.SetRouter(sourceUser, subOnRampRouterAddress)
	require.NoError(t, err)
	_, err = subOnRampRouter.SetOnRamp(sourceUser, destChainID, subOnRampAddress)
	require.NoError(t, err)
	subOnRampFee := big.NewInt(1)
	_, err = subOnRampRouter.SetFee(sourceUser, subOnRampFee)
	require.NoError(t, err)
	sourceChain.Commit()
	_, err = subOnRamp.SetPrices(sourceUser, []common.Address{sourceLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	subOffRampRouterAddress, _, _, err := any_2_evm_subscription_offramp_router.DeployAny2EVMSubscriptionOffRampRouter(
		destUser, destChain, []common.Address{}, any_2_evm_subscription_offramp_router.SubscriptionInterfaceSubscriptionConfig{
			SetSubscriptionSenderDelay: 0,
			WithdrawalDelay:            0,
			FeeToken:                   destLinkTokenAddress,
		})
	require.NoError(t, err)
	destChain.Commit()
	subOffRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(subOffRampRouterAddress, destChain)
	require.NoError(t, err)
	subOffRampAddress, _, _, err := any_2_evm_subscription_offramp.DeployEVM2EVMSubscriptionOffRamp(destUser, destChain, sourceChainID, destChainID,
		any_2_evm_subscription_offramp.BaseOffRampInterfaceOffRampConfig{
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
		blobVerifier.Address(),
		subOnRampAddress,
		afnDestAddress,
		[]common.Address{sourceLinkTokenAddress},
		[]common.Address{destPoolAddress},
		any_2_evm_subscription_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: hundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
	)
	require.NoError(t, err)
	subOffRamp, _ := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(subOffRampAddress, destChain)
	_, err = destPool.SetOffRamp(destUser, subOffRampAddress, true)
	require.NoError(t, err)
	_, err = subOffRamp.SetRouter(destUser, subOffRampRouterAddress)
	require.NoError(t, err)
	_, err = subOffRampRouter.AddOffRamp(destUser, subOffRampAddress)
	require.NoError(t, err)
	destChain.Commit()
	_, err = subOffRamp.SetPrices(destUser, []common.Address{destLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	_, err = subOffRampRouter.GetSupportedTokensForExecutionFee(nil)
	require.NoError(t, err)
	_, err = subOffRampRouter.GetFeeToken(nil)
	require.NoError(t, err)
	// Enable onramps on blob verifier.
	blobVerifier.SetConfig(destUser, blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{
		OnRamps:          []common.Address{onRampAddress, subOnRampAddress},
		MinSeqNrByOnRamp: []uint64{1, 1},
	})
	// Ensure we have at least finality blocks.
	for i := 0; i < 50; i++ {
		sourceChain.Commit()
		destChain.Commit()
	}
	return CCIPContracts{
		sourceUser:      sourceUser,
		destUser:        destUser,
		sourceChain:     sourceChain,
		destChain:       destChain,
		sourcePool:      sourcePool,
		destPool:        destPool,
		sourceLinkToken: sourceLinkToken,
		destLinkToken:   destLinkToken,
		blobVerifier:    blobVerifier,
		receivers:       []MaybeRevertReceiver{{Receiver: revertingMessageReceiver1, Strict: false}, {Receiver: revertingMessageReceiver2, Strict: true}},

		// Toll
		tollOnRampFees:    tollOnRampFees,
		tollOnRamp:        tollOnRamp,
		tollOnRampRouter:  tollOnRampRouter,
		tollOffRampRouter: tollOffRampRouter,
		tollOffRamp:       tollOffRamp,

		// Sub
		subOnRampFee:     subOnRampFee,
		subOnRamp:        subOnRamp,
		subOnRampRouter:  subOnRampRouter,
		subOffRampRouter: subOffRampRouter,
		subOffRamp:       subOffRamp,
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
	config.Overrides.GlobalEvmGasLimitDefault = null.NewInt(1500000, true)
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
	sourceClient := client.NewSimulatedBackendClient(t, sourceChain, sourceChainID)
	destClient := client.NewSimulatedBackendClient(t, destChain, destChainID)

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
	evmChain, err := evm.LoadChainSet(nil, evm.ChainSetOpts{
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
		UnrestrictedHTTPClient: &http.Client{},
		RestrictedHTTPClient:   &http.Client{},
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
	sendingKeys, err := app.GetKeyStore().Eth().EnabledKeysForChain(destChainID)
	require.NoError(t, err)
	require.Len(t, sendingKeys, 1)
	transmitter := sendingKeys[0].Address
	s, err := app.GetKeyStore().Eth().GetState(sendingKeys[0].ID(), destChainID)
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

func (node *Node) eventuallyHasReqSeqNum(t *testing.T, ccipContracts CCIPContracts, eventSig common.Hash, onRamp common.Address, seqNum int) logpoller.Log {
	c, err := node.app.GetChains().EVM.Get(sourceChainID)
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.sourceChain.Commit()
		ccipContracts.destChain.Commit()
		lgs, err := c.LogPoller().LogsDataWordRange(eventSig, onRamp, ccip.SendRequestedSequenceNumberIndex, ccip.EvmWord(uint64(seqNum)), ccip.EvmWord(uint64(seqNum)), 1)
		require.NoError(t, err)
		t.Log("Send requested", len(lgs))
		if len(lgs) == 1 {
			log = lgs[0]
			return true
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "eventually has seq num")
	return log
}

func (node *Node) eventuallyHasExecutedSeqNum(t *testing.T, ccipContracts CCIPContracts, offRamp common.Address, seqNum int) logpoller.Log {
	c, err := node.app.GetChains().EVM.Get(destChainID)
	require.NoError(t, err)
	var log logpoller.Log
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		ccipContracts.sourceChain.Commit()
		ccipContracts.destChain.Commit()
		lgs, err := c.LogPoller().IndexedLogsTopicRange(
			ccip.ExecutionStateChanged,
			offRamp,
			ccip.CrossChainMessageExecutedSequenceNumberIndex,
			ccip.EvmWord(uint64(seqNum)),
			ccip.EvmWord(uint64(seqNum)),
			1)
		require.NoError(t, err)
		t.Log("Executed logs", lgs)
		if len(lgs) == 1 {
			log = lgs[0]
			return true
		}
		return false
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "eventually has executed seq num")
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

func allNodesHaveReqSeqNum(t *testing.T, ccipContracts CCIPContracts, eventSig common.Hash, onRamp common.Address, nodes []Node, seqNum int) logpoller.Log {
	var log logpoller.Log
	for _, node := range nodes {
		log = node.eventuallyHasReqSeqNum(t, ccipContracts, eventSig, onRamp, seqNum)
	}
	return log
}

func allNodesHaveExecutedSeqNum(t *testing.T, ccipContracts CCIPContracts, offRamp common.Address, nodes []Node, seqNum int) logpoller.Log {
	var log logpoller.Log
	for _, node := range nodes {
		log = node.eventuallyHasExecutedSeqNum(t, ccipContracts, offRamp, seqNum)
	}
	return log
}

func queueSubRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, gasLimit *big.Int, receiver common.Address) *gethtypes.Transaction {
	msg := evm_2_any_subscription_onramp_router.CCIPEVM2AnySubscriptionMessage{
		Receiver: receiver,
		Data:     []byte(msgPayload),
		Tokens:   tokens,
		Amounts:  amounts,
		GasLimit: gasLimit,
	}
	tx, err := ccipContracts.subOnRampRouter.CcipSend(ccipContracts.sourceUser, destChainID, msg)
	require.NoError(t, err)
	return tx
}

func queueRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, feeToken common.Address, feeTokenAmount *big.Int, gasLimit *big.Int) *gethtypes.Transaction {
	msg := evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage{
		Receiver:       ccipContracts.receivers[0].Receiver.Address(),
		Data:           []byte(msgPayload),
		Tokens:         tokens,
		Amounts:        amounts,
		FeeToken:       feeToken,
		FeeTokenAmount: feeTokenAmount,
		GasLimit:       gasLimit,
	}
	tx, err := ccipContracts.tollOnRampRouter.CcipSend(ccipContracts.sourceUser, destChainID, msg)
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

func sendSubRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, gasLimit *big.Int, receiver common.Address) {
	tx := queueSubRequest(t, ccipContracts, msgPayload, tokens, amounts, gasLimit, receiver)
	confirmTxs(t, []*gethtypes.Transaction{tx}, ccipContracts.sourceChain)
}

func sendRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []common.Address, amounts []*big.Int, feeToken common.Address, feeTokenAmount *big.Int, gasLimit *big.Int) {
	tx := queueRequest(t, ccipContracts, msgPayload, tokens, amounts, feeToken, feeTokenAmount, gasLimit)
	confirmTxs(t, []*gethtypes.Transaction{tx}, ccipContracts.sourceChain)
}

func eventuallyReportRelayed(t *testing.T, ccipContracts CCIPContracts, onRamp common.Address, min, max int) {
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		minSeqNum, err := ccipContracts.blobVerifier.GetExpectedNextSequenceNumber(nil, onRamp)
		require.NoError(t, err)
		ccipContracts.sourceChain.Commit()
		ccipContracts.destChain.Commit()
		t.Log("min seq num reported", minSeqNum)
		return minSeqNum > uint64(max)
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "eventually report relayed")
}

/*
func executeMessage(t *testing.T, ccipContracts CCIPContracts, req logpoller.Log, allReqs []logpoller.Log, report blob_verifier.CCIPRelayReport) {
	// Double check root exists onchain.
	exists, err := ccipContracts.blobVerifier.GetMerkleRoot(nil, report.MerkleRoot)
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
	offRampProof := blob_verifier.CCIPExecutionReport{
		Messages:       []blob_verifier.CCIPAny2EVMTollMessage{*decodedMsg},
		Proofs:         proof.Hashes,
		ProofFlagsBits: ccip.ProofFlagsToBits(proof.SourceFlags),
	}
	onchainRoot, err := ccipContracts.blobVerifier.MerkleRoot(nil, offRampProof)
	require.NoError(t, err)
	require.Equal(t, tree.Root(), onchainRoot)

	// Execute.
	tx, err := ccipContracts.blobVerifier.ExecuteTransaction(ccipContracts.destUser, offRampProof, false)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()
	rec, err := ccipContracts.destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), rec.Status)
}
*/

func setupAndStartNodes(t *testing.T, ctx context.Context, ccipContracts CCIPContracts, bootstrapNodePort int64) (Node, []Node, int64) {
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
	configBlock := setupOnchainConfig(t, ccipContracts, oracles, reportingPluginConfig)
	return bootstrapNode, nodes, configBlock
}

func AssertTxSuccess(t *testing.T, ccipContracts CCIPContracts, log logpoller.Log) {
	executionStateChanged, err := ccipContracts.tollOffRamp.ParseExecutionStateChanged(log.GetGethLog())
	require.NoError(t, err)
	if ccip.MessageExecutionState(executionStateChanged.State) != ccip.Success {
		t.Log("Execution failed")
		t.Fail()
	}
}

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := setupCCIPContracts(t)
	bootstrapNodePort := int64(19599)
	ctx := context.Background()
	// Starts nodes and configures them in the OCR contracts.
	bootstrapNode, nodes, configBlock := setupAndStartNodes(t, ctx, ccipContracts, bootstrapNodePort)

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
`, ccipContracts.blobVerifier.Address(), destChainID))

	linkEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"JuelsPerETH": "200000000000000000000"}`))
		require.NoError(t, err)
	}))
	defer linkEth.Close()
	// For each node add a relayer and executor job.
	for i, node := range nodes {
		node.addJob(t, fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-relay"
relay               = "evm"
schemaVersion      	= 1
name               	= "ccip-relay-%d"
contractID 			= "%s"
ocrKeyBundleID      = "%s"
transmitterID 		= "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampIDs            = ["%s", "%s"]
sourceChainID       = %s
destChainID         = %s
pollPeriod          = "1s"
destStartBlock      = %d

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.blobVerifier.Address(), node.kb.ID(), node.transmitter, ccipContracts.tollOnRamp.Address(), ccipContracts.subOnRamp.Address(), sourceChainID, destChainID, configBlock, destChainID))
		node.addJob(t, fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-execution"
relay               = "evm"
schemaVersion       = 1
name                = "ccip-executor-toll-%d"
contractID          = "%s"
ocrKeyBundleID      = "%s"
transmitterID       = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampID            = "%s"
blobVerifierID      = "%s"
sourceChainID       = %s
destChainID         = %s
pollPeriod          = "1s"
destStartBlock      = %d
tokensPerFeeCoinPipeline = """
// Price 1 
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="JuelsPerETH"];
link->link_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse)}"];
"""

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.tollOffRamp.Address(), node.kb.ID(), node.transmitter, ccipContracts.tollOnRamp.Address(), ccipContracts.blobVerifier.Address(), sourceChainID, destChainID, configBlock, linkEth.URL, ccipContracts.destLinkToken.Address(), destChainID))
		node.addJob(t, fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-execution"
relay               = "evm"
schemaVersion       = 1
name                = "ccip-executor-subscription-%d"
contractID 			= "%s"
ocrKeyBundleID      = "%s"
transmitterID       = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampID            = "%s"
blobVerifierID      = "%s"
sourceChainID       = %s
destChainID         = %s
pollPeriod          = "1s"
destStartBlock      = %d
tokensPerFeeCoinPipeline = """
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="JuelsPerETH"];
link->link_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse)}"];
"""

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.subOffRamp.Address(), node.kb.ID(), node.transmitter, ccipContracts.subOnRamp.Address(), ccipContracts.blobVerifier.Address(), sourceChainID, destChainID, configBlock, linkEth.URL, ccipContracts.destLinkToken.Address(), destChainID))
	}
	// Replay for bootstrap.
	bc, err := bootstrapNode.app.GetChains().EVM.Get(destChainID)
	require.NoError(t, err)
	require.NoError(t, bc.LogPoller().Replay(context.Background(), configBlock))
	/*
		t.Run("single self-execute", func(t *testing.T) {
			sendRequest(t, ccipContracts, "single req", []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, big.NewInt(0), big.NewInt(0))
			req := allNodesHaveReqSeqNum(t, ccipContracts, nodes, tollCurrentSeqNum)
			report := eventuallyReportRelayed(t, ccipContracts, tollCurrentSeqNum, tollCurrentSeqNum)
			executeMessage(t, ccipContracts, req, []logpoller.Log{req}, report)
			allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, tollCurrentSeqNum)
			receivedMsg, err := ccipContracts.successReceivers[0].SMessage(nil)
			require.NoError(t, err)
			assert.Equal(t, "single req", string(receivedMsg.Data))
			tollCurrentSeqNum++
		})

		t.Run("batch self-execute", func(t *testing.T) {
			var txs []*gethtypes.Transaction
			n := 3
			for i := 0; i < n; i++ {
				txs = append(txs, queueRequest(t, ccipContracts, fmt.Sprintf("batch request %d", tollCurrentSeqNum+i), []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{big.NewInt(100)}, big.NewInt(0), big.NewInt(0)))
			}
			// Send a batch of requests in a single block
			confirmTxs(t, txs, ccipContracts.sourceChain)
			// All nodes should have all 3.
			var reqs []logpoller.Log
			for i := 0; i < n; i++ {
				reqs = append(reqs, allNodesHaveReqSeqNum(t, ccipContracts, nodes, tollCurrentSeqNum+i))
			}
			// Should see a report with the full range
			report := eventuallyReportRelayed(t, ccipContracts, tollCurrentSeqNum, tollCurrentSeqNum+n-1)
			// Execute them all
			for _, req := range reqs {
				executeMessage(t, ccipContracts, req, reqs, report)
			}
			for i := range reqs {
				allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, tollCurrentSeqNum+i)
			}
			receivedMsg, err := ccipContracts.successReceivers[0].SMessage(nil)
			require.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("batch request %d", tollCurrentSeqNum+n-1), string(receivedMsg.Data))
			tollCurrentSeqNum += n
		})
	*/

	tollCurrentSeqNum := 1
	subCurrentSeqNum := 1
	// Create src sub and dst sub, funded for 10 msgs with 100k callback and 1 token.
	// Needs to be sufficient to cover default gas price of 200gwei.
	// Costs ~7 link for 100k callback @ 200gwei.
	for _, receiver := range ccipContracts.receivers {
		relayFee := big.NewInt(0).Mul(ccipContracts.subOnRampFee, big.NewInt(10))
		_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.subOnRampRouter.Address(), relayFee)
		require.NoError(t, err)
		_, err = ccipContracts.subOnRampRouter.FundSubscription(ccipContracts.sourceUser, relayFee)
		require.NoError(t, err)
		ccipContracts.sourceChain.Commit()
		subscriptionBalance := big.NewInt(0).Mul(big.NewInt(80), big.NewInt(1e18))
		_, err = ccipContracts.destLinkToken.Approve(ccipContracts.destUser, ccipContracts.subOffRampRouter.Address(), subscriptionBalance)
		require.NoError(t, err)
		_, err = ccipContracts.subOffRampRouter.CreateSubscription(ccipContracts.destUser, any_2_evm_subscription_offramp_router.SubscriptionInterfaceOffRampSubscription{
			Senders:          []common.Address{ccipContracts.sourceUser.From},
			Receiver:         receiver.Receiver.Address(),
			StrictSequencing: receiver.Strict,
			Balance:          subscriptionBalance,
		})
		require.NoError(t, err)
	}

	ccipContracts.destChain.Commit()

	t.Run("single auto-execute toll", func(t *testing.T) {
		// Approve router to take source token.
		tokenAmount := big.NewInt(100)
		// Example sending a msg with 100k callback execution on eth mainnet.
		// Gas price 200e9 wei/gas * 1e5 gas * (2e20 juels/eth / 1e18wei/eth)
		// 4e18 juels = 4 link, which does not include gas used outside the callback.
		// Gas outside the callback for 1 token is ~100k in the worst case.
		feeTokenAmount := big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18))
		_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.tollOnRampRouter.Address(), big.NewInt(0).Add(tokenAmount, feeTokenAmount))
		require.NoError(t, err)
		ccipContracts.sourceChain.Commit()

		startReceiver, _ := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.receivers[0].Receiver.Address())
		startPool, _ := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destPool.Address())
		sendRequest(t, ccipContracts, "hey DON, execute for me",
			[]common.Address{ccipContracts.sourceLinkToken.Address()},
			[]*big.Int{tokenAmount}, ccipContracts.sourceLinkToken.Address(),
			feeTokenAmount,
			big.NewInt(100_000))
		allNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPTollSendRequested, ccipContracts.tollOnRamp.Address(), nodes, tollCurrentSeqNum)
		eventuallyReportRelayed(t, ccipContracts, ccipContracts.tollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum)
		executionLog := allNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.tollOffRamp.Address(), nodes, tollCurrentSeqNum)
		AssertTxSuccess(t, ccipContracts, executionLog)
		endReceiver, _ := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.receivers[0].Receiver.Address())
		endPool, _ := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destPool.Address())
		// We expect that the receiver should have their tokens.
		assert.Equal(t, tokenAmount.String(), big.NewInt(0).Sub(endReceiver, startReceiver).String())
		// TODO: Assert change forwarding. Change forwarding not implemented yet??
		t.Log(startPool, endPool)
		tollCurrentSeqNum++
	})

	t.Run("single auto-execute subscription", func(t *testing.T) {
		tokenAmount := big.NewInt(100)
		subBefore, err := ccipContracts.subOffRampRouter.GetSubscription(nil, ccipContracts.receivers[0].Receiver.Address())
		require.NoError(t, err)
		_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.subOnRampRouter.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.sourceChain.Commit()

		startReceiver, _ := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.receivers[0].Receiver.Address())
		sendSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.sourceLinkToken.Address()},
			[]*big.Int{tokenAmount}, big.NewInt(100_000), ccipContracts.receivers[0].Receiver.Address())
		allNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.subOnRamp.Address(), nodes, subCurrentSeqNum)
		eventuallyReportRelayed(t, ccipContracts, ccipContracts.subOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		executionLog := allNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.subOffRamp.Address(), nodes, subCurrentSeqNum)
		AssertTxSuccess(t, ccipContracts, executionLog)

		endReceiver, _ := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.receivers[0].Receiver.Address())
		subAfter, err := ccipContracts.subOffRampRouter.GetSubscription(nil, ccipContracts.receivers[0].Receiver.Address())
		require.NoError(t, err)
		// Subscription should decrease. Tricky to measure exactly since we measure gas consumption directly in sub offramp.
		assert.Equal(t, 1, subBefore.Balance.Cmp(subAfter.Balance), "before %v after %v", subBefore.Balance, subAfter.Balance)
		// We should see the tokenAmount transferred.
		assert.Equal(t, tokenAmount.String(), big.NewInt(0).Sub(endReceiver, startReceiver).String(), "before %v after %v", startReceiver, endReceiver)
		subCurrentSeqNum++
	})

	t.Run("batch auto-execute toll", func(t *testing.T) {
		tokenAmount := big.NewInt(100)
		feeTokenAmount := big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18))
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.tollOnRampRouter.Address(), big.NewInt(0).Add(tokenAmount, feeTokenAmount))
			require.NoError(t, err)
			txs = append(txs, queueRequest(t, ccipContracts, fmt.Sprintf("batch request %d", tollCurrentSeqNum+i), []common.Address{ccipContracts.sourceLinkToken.Address()}, []*big.Int{tokenAmount},
				ccipContracts.sourceLinkToken.Address(), feeTokenAmount, big.NewInt(100_000)))
		}
		// Send a batch of requests in a single block
		confirmTxs(t, txs, ccipContracts.sourceChain)
		// All nodes should have all 3.
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, allNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPTollSendRequested, ccipContracts.tollOnRamp.Address(), nodes, tollCurrentSeqNum+i))
		}
		// Should see a report with the full range
		eventuallyReportRelayed(t, ccipContracts, ccipContracts.tollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum+n-1)
		// Should all be executed
		for i := range reqs {
			executionLog := allNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.tollOffRamp.Address(), nodes, tollCurrentSeqNum+i)
			AssertTxSuccess(t, ccipContracts, executionLog)
		}
		tollCurrentSeqNum += n
	})

	t.Run("batch auto-execute subscription", func(t *testing.T) {
		subBefore, _ := ccipContracts.subOffRampRouter.GetSubscription(nil, ccipContracts.receivers[0].Receiver.Address())
		tokenAmount := big.NewInt(100)
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.subOnRampRouter.Address(), tokenAmount)
			require.NoError(t, err)
			txs = append(txs, queueSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.sourceLinkToken.Address()},
				[]*big.Int{tokenAmount}, big.NewInt(100_000), ccipContracts.receivers[0].Receiver.Address()))
		}
		ccipContracts.sourceChain.Commit()
		// Send a batch of requests in a single block
		confirmTxs(t, txs, ccipContracts.sourceChain)
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, allNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.subOnRamp.Address(), nodes, subCurrentSeqNum+i))
		}
		// Should see a report with the full range
		eventuallyReportRelayed(t, ccipContracts, ccipContracts.subOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum+n-1)
		// Should all be executed
		for i := range reqs {
			executionLog := allNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.subOffRamp.Address(), nodes, subCurrentSeqNum+i)
			AssertTxSuccess(t, ccipContracts, executionLog)
		}
		// Check subscription balance after
		subAfter, _ := ccipContracts.subOffRampRouter.GetSubscription(nil, ccipContracts.receivers[0].Receiver.Address())
		require.Equal(t, 1, subBefore.Balance.Cmp(subAfter.Balance), "before %v after %v", subBefore.Balance, subAfter.Balance)
		b := big.NewInt(0).Sub(subBefore.Balance, subAfter.Balance).String()
		t.Log("Balance change", b)
		subCurrentSeqNum += n
	})

	t.Run("single strict sequencing auto-execute subscription", func(t *testing.T) {
		tokenAmount := big.NewInt(100)
		_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.subOnRampRouter.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.sourceChain.Commit()
		sendSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.sourceLinkToken.Address()},
			[]*big.Int{tokenAmount}, big.NewInt(100_000), ccipContracts.receivers[1].Receiver.Address())
		allNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.subOnRamp.Address(), nodes, subCurrentSeqNum)
		eventuallyReportRelayed(t, ccipContracts, ccipContracts.subOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		executionLog := allNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.subOffRamp.Address(), nodes, subCurrentSeqNum)
		AssertTxSuccess(t, ccipContracts, executionLog)
		subCurrentSeqNum++
	})

	// TODO: One reverted SS should block subsequent SS
	// then we should be able to manually execute.

	/*
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
			msg := evm_2_evm_toll_onramp_router.CCIPEVM2AnyTollMessage{
				Receiver:       ccipContracts.destUser.From,
				Data:           nil,
				Tokens:         []common.Address{ccipContracts.sourceLinkToken.Address()},
				Amounts:        []*big.Int{big.NewInt(100)},
				FeeToken:       ccipContracts.sourceLinkToken.Address(),
				FeeTokenAmount: feeTokenAmount,
				GasLimit:       big.NewInt(1e6),
			}
			_, err = ccipContracts.tollOnRampRouter.CcipSend(ccipContracts.sourceUser, destChainID, msg)
			require.NoError(t, err)
			ccipContracts.sourceChain.Commit()

			allNodesHaveReqSeqNum(t, ccipContracts, nodes, tollCurrentSeqNum)
			eventuallyReportRelayed(t, ccipContracts, tollCurrentSeqNum, tollCurrentSeqNum)
			//executeMessage(t, ccipContracts, req2, []logpoller.Log{req2}, report2)
			allNodesHaveExecutedSeqNum(t, ccipContracts, nodes, tollCurrentSeqNum)

			// The destination user's balance should increase
			endBalanceSource, err := ccipContracts.sourceLinkToken.BalanceOf(nil, ccipContracts.sourceUser.From)
			require.NoError(t, err)
			endBalanceDest, err := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destUser.From)
			require.NoError(t, err)
			assert.Equal(t, "100", big.NewInt(0).Sub(startBalanceSource, endBalanceSource).String())
			assert.Equal(t, "100", big.NewInt(0).Sub(endBalanceDest, startBalanceDest).String())
		})
	*/
}

func setupOnchainConfig(t *testing.T, ccipContracts CCIPContracts, oracles []confighelper2.OracleIdentityExtra, reportingPluginConfig []byte) int64 {
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
		50*time.Millisecond,  // Max duration query
		200*time.Millisecond, // Max duration observation
		100*time.Millisecond,
		100*time.Millisecond,
		100*time.Millisecond,
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

	blockBeforeConfig, err := ccipContracts.destChain.BlockByNumber(context.Background(), nil)
	require.NoError(t, err)
	// Set the DON on the offramp
	_, err = ccipContracts.blobVerifier.SetConfig0(
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

	// Same DON on the toll offramp
	_, err = ccipContracts.tollOffRamp.SetConfig(
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

	// Same DON on the sub offramp
	_, err = ccipContracts.subOffRamp.SetConfig(
		ccipContracts.destUser,
		signerAddresses,
		transmitterAddresses,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	return blockBeforeConfig.Number().Int64()
}

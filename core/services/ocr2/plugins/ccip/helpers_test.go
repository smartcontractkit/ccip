package ccip

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/stretchr/testify/require"

	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store_helper"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type OCR2TestContract interface {
	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)
}

func generateAndSetTestOCR2Config(contract OCR2TestContract, owner *bind.TransactOpts, onChainConfig []byte) (*types.Transaction, error) {
	var signers []common.Address
	var transmitters []common.Address

	for i := 0; i < 4; i++ {
		signers = append(signers, utils.RandomAddress())
		transmitters = append(transmitters, utils.RandomAddress())
	}

	return contract.SetOCR2Config(owner, signers, transmitters, 1, onChainConfig, 2, nil)
}

type ccipPluginTestHarness struct {
	sourceChainID, destChainID uint64
	lggr                       logger.Logger
	owner                      *bind.TransactOpts // Has all the link and 100ETH

	sourceClient *backends.SimulatedBackend
	sourceLP     logpoller.LogPollerTest
	destClient   *backends.SimulatedBackend
	destLP       logpoller.LogPollerTest

	sourceRouter *router.Router
	onRamp       *evm_2_evm_onramp.EVM2EVMOnRamp

	destRouter          *router.Router
	offRamp             *evm_2_evm_offramp.EVM2EVMOffRamp
	commitStoreHelper   *commit_store_helper.CommitStoreHelper
	commitStore         *commit_store.CommitStore
	priceRegistry       *price_registry.PriceRegistry
	receiver            *simple_message_receiver.SimpleMessageReceiver
	commitOnchainConfig CommitOnchainConfig
	execOnchainConfig   ExecOnchainConfig

	sourceFeeTokenAddress common.Address
	destFeeTokenAddress   common.Address
	sourceNativeAddress   common.Address
}

func (th *ccipPluginTestHarness) flushLogs(t *testing.T) {
	th.sourceClient.Commit()
	th.sourceLP.PollAndSaveLogs(testutils.Context(t), th.sourceClient.Blockchain().CurrentBlock().Number.Int64())

	th.destClient.Commit()
	th.destLP.PollAndSaveLogs(testutils.Context(t), th.destClient.Blockchain().CurrentBlock().Number.Int64())
}

func deployTokenAndPool(
	t *testing.T,
	owner *bind.TransactOpts,
	client *backends.SimulatedBackend,
	liquidityAmount *big.Int,
) (token *link_token_interface.LinkToken, pool *lock_release_token_pool.LockReleaseTokenPool) {
	// Deploy source link token
	sourceFeeTokenAddress, _, _, err := link_token_interface.DeployLinkToken(owner, client)
	require.NoError(t, err)
	token, err = link_token_interface.NewLinkToken(sourceFeeTokenAddress, client)
	require.NoError(t, err)

	// Deploy source pool
	sourcePoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		owner,
		client,
		sourceFeeTokenAddress,
		lock_release_token_pool.RateLimiterConfig{
			Capacity:  big.NewInt(8e18),
			Rate:      big.NewInt(1e18),
			IsEnabled: true,
		},
	)
	require.NoError(t, err)
	client.Commit()
	pool, err = lock_release_token_pool.NewLockReleaseTokenPool(sourcePoolAddress, client)
	require.NoError(t, err)

	_, err = token.Approve(owner, sourcePoolAddress, liquidityAmount)
	require.NoError(t, err)
	client.Commit()
	_, err = pool.AddLiquidity(owner, liquidityAmount)
	require.NoError(t, err)
	client.Commit()

	return token, pool
}

func deploySourceCcipContracts(
	t *testing.T,
	owner *bind.TransactOpts,
	sourceClient *backends.SimulatedBackend,
	destChainID uint64,
	sourceFeeTokenAddress, sourceNativeAddress common.Address,
	sourcePool *lock_release_token_pool.LockReleaseTokenPool,
) (sourceRouter *router.Router, onRamp *evm_2_evm_onramp.EVM2EVMOnRamp) {
	// deploy source afn
	sourceAfnAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		owner,
		sourceClient,
	)
	require.NoError(t, err)

	// deploy source router
	sourceRouterAddress, _, _, err := router.DeployRouter(owner, sourceClient, sourceNativeAddress)
	require.NoError(t, err)
	sourceClient.Commit()

	sourceRouter, err = router.NewRouter(sourceRouterAddress, sourceClient)
	require.NoError(t, err)

	// deploy priceRegistry
	sourcePriceRegistryAddress, _, _, err := price_registry.DeployPriceRegistry(owner, sourceClient, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: sourceFeeTokenAddress,
				UsdPerToken: big.NewInt(6e18), // $6
			},
		},
		DestChainId:   destChainID,
		UsdPerUnitGas: big.NewInt(2000e9), // $2000
	}, []common.Address{}, []common.Address{sourceFeeTokenAddress}, uint32(time.Hour.Seconds())) // 1h
	require.NoError(t, err)
	sourceClient.Commit()

	// deploy source onRamp
	onRampAddress, _, _, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		owner,
		sourceClient,
		evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
			LinkToken:         sourceFeeTokenAddress,
			ChainId:           sourceClient.Blockchain().Config().ChainID.Uint64(),
			DestChainId:       destChainID,
			DefaultTxGasLimit: 200_000,
		},
		evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
			Router:          sourceRouterAddress,
			PriceRegistry:   sourcePriceRegistryAddress,
			MaxDataSize:     1e5,
			MaxTokensLength: 5,
			MaxGasLimit:     4e6,
			Afn:             sourceAfnAddress,
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool{
			{Token: sourceFeeTokenAddress, Pool: sourcePool.Address()},
		},
		[]common.Address{},
		evm_2_evm_onramp.RateLimiterConfig{
			Capacity:  big.NewInt(8e18),
			Rate:      big.NewInt(1e18),
			IsEnabled: true,
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{{Token: sourceFeeTokenAddress, Multiplier: 1e18, FeeAmount: big.NewInt(0), DestGasOverhead: 0}},
		[]evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
	)
	require.NoError(t, err)
	onRamp, err = evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, sourceClient)
	require.NoError(t, err)
	sourceClient.Commit()

	// register onramp in router and pool
	_, err = sourcePool.ApplyRampUpdates(owner, []lock_release_token_pool.TokenPoolRampUpdate{{Ramp: onRampAddress, Allowed: true}}, []lock_release_token_pool.TokenPoolRampUpdate{})
	require.NoError(t, err)
	_, err = sourceRouter.ApplyRampUpdates(owner, []router.RouterOnRampUpdate{{OnRamp: onRampAddress, DestChainId: destChainID}}, []router.RouterOffRampUpdate{})
	require.NoError(t, err)
	sourceClient.Commit()

	return sourceRouter, onRamp
}

func deployDestCcipContracts(
	t *testing.T,
	owner *bind.TransactOpts,
	destClient *backends.SimulatedBackend,
	sourceChainID uint64,
	onRampAddress, sourceFeeTokenAddress, destFeeTokenAddress common.Address,
	destPool *lock_release_token_pool.LockReleaseTokenPool,
) (
	afnAddress common.Address,
	priceRegistry *price_registry.PriceRegistry,
	destRouter *router.Router,
	offRamp *evm_2_evm_offramp.EVM2EVMOffRamp,
	commitStoreHelper *commit_store_helper.CommitStoreHelper,
) {
	afnAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		owner,
		destClient,
	)
	require.NoError(t, err)
	destClient.Commit()

	// deploy priceRegistry
	priceRegistryAddress, _, _, err := price_registry.DeployPriceRegistry(owner, destClient, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: destFeeTokenAddress,
				UsdPerToken: big.NewInt(8e18), // $8
			},
		},
		DestChainId:   sourceChainID,
		UsdPerUnitGas: big.NewInt(2000e9), // $2000
	}, []common.Address{}, []common.Address{destFeeTokenAddress}, uint32(time.Hour.Seconds())) // 1h
	require.NoError(t, err)
	destClient.Commit()

	priceRegistry, err = price_registry.NewPriceRegistry(priceRegistryAddress, destClient)
	require.NoError(t, err)

	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		owner,      // user
		destClient, // client
		commit_store_helper.CommitStoreStaticConfig{
			ChainId:       destClient.Blockchain().Config().ChainID.Uint64(),
			SourceChainId: sourceChainID,
			OnRamp:        onRampAddress,
		},
	)
	require.NoError(t, err)
	commitStoreHelper, err = commit_store_helper.NewCommitStoreHelper(commitStoreAddress, destClient)
	require.NoError(t, err)
	destClient.Commit()

	_, err = priceRegistry.ApplyPriceUpdatersUpdates(owner, []common.Address{commitStoreAddress}, []common.Address{})
	require.NoError(t, err)

	// deploy dest router
	destRouterAddress, _, _, err := router.DeployRouter(owner, destClient, common.Address{})
	require.NoError(t, err)
	destClient.Commit()
	destRouter, err = router.NewRouter(destRouterAddress, destClient)
	require.NoError(t, err)

	// deploy offramp
	offRampAddress, _, _, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(
		owner,
		destClient,
		evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
			CommitStore:   commitStoreHelper.Address(),
			ChainId:       destClient.Blockchain().Config().ChainID.Uint64(),
			SourceChainId: sourceChainID,
			OnRamp:        onRampAddress,
		},
		[]common.Address{sourceFeeTokenAddress},
		[]common.Address{destPool.Address()},
		evm_2_evm_offramp.RateLimiterConfig{
			Capacity:  big.NewInt(1e18),
			Rate:      big.NewInt(1e18),
			IsEnabled: true,
		},
	)
	require.NoError(t, err)
	offRamp, err = evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddress, destClient)
	require.NoError(t, err)

	// register offramp in router and pool
	_, err = destPool.ApplyRampUpdates(owner, []lock_release_token_pool.TokenPoolRampUpdate{}, []lock_release_token_pool.TokenPoolRampUpdate{{Ramp: offRampAddress, Allowed: true}})
	require.NoError(t, err)
	_, err = destRouter.ApplyRampUpdates(owner, nil, []router.RouterOffRampUpdate{
		{SourceChainId: sourceChainID, OffRamps: []common.Address{offRampAddress}},
	})
	require.NoError(t, err)

	return afnAddress, priceRegistry, destRouter, offRamp, commitStoreHelper
}

func setupCcipTestHarness(t *testing.T) ccipPluginTestHarness {
	destChainID := testutils.SimulatedChainID.Uint64()
	sourceChainID := destChainID - 1

	lggr := logger.TestLogger(t)
	owner := testutils.MustNewSimTransactor(t)

	// db, clients and logpollers
	db := pgtest.NewSqlxDB(t)
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_log_poller_blocks_evm_chain_id_fkey DEFERRED`)))
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_log_poller_filters_evm_chain_id_fkey DEFERRED`)))
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_logs_evm_chain_id_fkey DEFERRED`)))

	sourceSimulatedChainID := big.NewInt(int64(sourceChainID))
	sourceORM := logpoller.NewORM(sourceSimulatedChainID, db, lggr, pgtest.NewQConfig(true))
	sourceClient := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		owner.From: {
			Balance: new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10*ethconfig.Defaults.Miner.GasCeil) // 80M gas
	var sourceLP logpoller.LogPollerTest = logpoller.NewLogPoller(sourceORM, evmclient.NewSimulatedBackendClient(t, sourceClient, sourceSimulatedChainID), lggr.Named("sourceLP"), 1*time.Hour, 2, 3, 2, 1000)

	destORM := logpoller.NewORM(testutils.SimulatedChainID, db, lggr, pgtest.NewQConfig(true))
	destClient := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		owner.From: {
			Balance: new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10*ethconfig.Defaults.Miner.GasCeil) // 80M gas
	var destLP logpoller.LogPollerTest = logpoller.NewLogPoller(destORM, evmclient.NewSimulatedBackendClient(t, destClient, testutils.SimulatedChainID), lggr.Named("destLP"), 1*time.Hour, 2, 3, 2, 1000)

	liquidityAmount := new(big.Int).Mul(big.NewInt(20), big.NewInt(1e18))

	// deploy tokens and pools
	sourceFeeToken, sourcePool := deployTokenAndPool(t, owner, sourceClient, liquidityAmount)
	destFeeToken, destPool := deployTokenAndPool(t, owner, destClient, liquidityAmount)
	sourceFeeTokenAddress, destFeeTokenAddress := sourceFeeToken.Address(), destFeeToken.Address()

	// deploy source ccip contracts
	sourceNativeAddress := testutils.NewAddress() // dummy for now, we don't need the native token contract
	sourceRouter, onRamp := deploySourceCcipContracts(t, owner, sourceClient, destChainID, sourceFeeTokenAddress, sourceNativeAddress, sourcePool)

	// deploy dest ccip contracts
	afnAddress, priceRegistry, destRouter, offRamp, commitStoreHelper := deployDestCcipContracts(
		t,
		owner,
		destClient,
		sourceChainID,
		onRamp.Address(),
		sourceFeeTokenAddress,
		destFeeTokenAddress,
		destPool,
	)

	// since CommitStoreHelper inherits CommitStore, it's safe for testing purposes to initialize on the same address
	commitStore, err := commit_store.NewCommitStore(commitStoreHelper.Address(), destClient)
	require.NoError(t, err)

	// deploy receiver
	receiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(owner, destClient)
	require.NoError(t, err)
	receiver, err := simple_message_receiver.NewSimpleMessageReceiver(receiverAddress, destClient)
	require.NoError(t, err)
	destClient.Commit()

	// onChain configs
	commitOnchainConfig := CommitOnchainConfig{Afn: afnAddress, PriceRegistry: priceRegistry.Address()}
	encodedCommitOnchainConfig, err := EncodeAbiStruct(commitOnchainConfig)
	require.NoError(t, err)

	_, err = generateAndSetTestOCR2Config(commitStoreHelper, owner, encodedCommitOnchainConfig)
	require.NoError(t, err)

	execOnchainConfig := ExecOnchainConfig{
		PermissionLessExecutionThresholdSeconds: 600,
		Router:                                  destRouter.Address(),
		PriceRegistry:                           priceRegistry.Address(),
		Afn:                                     afnAddress,
		MaxTokensLength:                         5,
		MaxDataSize:                             200_000,
	}
	encodedExecOnchainConfig, err := EncodeAbiStruct(execOnchainConfig)
	require.NoError(t, err)

	_, err = generateAndSetTestOCR2Config(offRamp, owner, encodedExecOnchainConfig)
	require.NoError(t, err)
	destClient.Commit()

	// register filters in logPoller
	eventsSignatures := GetEventSignatures()
	require.NoError(t, sourceLP.RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(COMMIT_CCIP_SENDS, onRamp.Address().String()),
		EventSigs: []common.Hash{eventsSignatures.SendRequested}, Addresses: []common.Address{onRamp.Address()},
	}))
	require.NoError(t, destLP.RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(COMMIT_PRICE_UPDATES, priceRegistry.Address()),
		EventSigs: []common.Hash{UsdPerUnitGasUpdated, UsdPerTokenUpdated}, Addresses: []common.Address{priceRegistry.Address()},
	}))

	// approve router
	_, err = sourceFeeToken.Approve(owner, sourceRouter.Address(), liquidityAmount)
	require.NoError(t, err)
	sourceClient.Commit()

	th := ccipPluginTestHarness{
		destChainID:   destChainID,
		sourceChainID: sourceChainID,
		lggr:          lggr,
		owner:         owner,

		sourceClient: sourceClient,
		sourceLP:     sourceLP,
		destClient:   destClient,
		destLP:       destLP,

		sourceRouter: sourceRouter,
		onRamp:       onRamp,

		destRouter:        destRouter,
		offRamp:           offRamp,
		commitStoreHelper: commitStoreHelper,
		commitStore:       commitStore,
		receiver:          receiver,
		priceRegistry:     priceRegistry,

		sourceFeeTokenAddress: sourceFeeTokenAddress,
		destFeeTokenAddress:   destFeeTokenAddress,
		sourceNativeAddress:   sourceNativeAddress,

		commitOnchainConfig: commitOnchainConfig,
		execOnchainConfig:   execOnchainConfig,
	}

	th.flushLogs(t)
	return th
}

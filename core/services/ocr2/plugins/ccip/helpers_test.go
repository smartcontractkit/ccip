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

type ccipPluginTestHarness struct {
	// Has all the link and 100ETH
	sourceChainID, destChainID uint64
	lggr                       logger.Logger
	owner                      *bind.TransactOpts
	client                     *backends.SimulatedBackend
	logPoller                  logpoller.LogPollerTest
	flushLogs                  func()

	offRamp           *evm_2_evm_offramp.EVM2EVMOffRamp
	commitStoreHelper *commit_store_helper.CommitStoreHelper
	commitStore       *commit_store.CommitStore
	receiver          *simple_message_receiver.SimpleMessageReceiver
	priceRegistry     *price_registry.PriceRegistry

	feeTokenAddress, sourceNativeAddress common.Address
	commitOnchainConfig                  CommitOnchainConfig
	execOnchainConfig                    ExecOnchainConfig
}

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

func setupCcipTestHarness(t *testing.T) ccipPluginTestHarness {
	destChainID := testutils.SimulatedChainID.Uint64()
	sourceChainID := destChainID - 1

	lggr := logger.TestLogger(t)
	owner := testutils.MustNewSimTransactor(t)

	db := pgtest.NewSqlxDB(t)
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_log_poller_blocks_evm_chain_id_fkey DEFERRED`)))
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_log_poller_filters_evm_chain_id_fkey DEFERRED`)))
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_logs_evm_chain_id_fkey DEFERRED`)))
	orm := logpoller.NewORM(testutils.SimulatedChainID, db, lggr, pgtest.NewQConfig(true))
	client := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		owner.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10*ethconfig.Defaults.Miner.GasCeil) // 80M gas
	var logPoller logpoller.LogPollerTest = logpoller.NewLogPoller(orm, evmclient.NewSimulatedBackendClient(t, client, testutils.SimulatedChainID), lggr, 1*time.Hour, 2, 3, 2, 1000)

	prevStart := int64(1)
	flushLogs := func() {
		client.Commit()
		logPoller.PollAndSaveLogs(testutils.Context(t), prevStart)
		latestBlock, err := orm.SelectLatestBlock()
		require.NoError(t, err)
		prevStart = latestBlock.BlockNumber + 1
	}

	// Deploy link token
	feeTokenAddress, _, _, err := link_token_interface.DeployLinkToken(owner, client)
	require.NoError(t, err)
	flushLogs()
	feeToken, err := link_token_interface.NewLinkToken(feeTokenAddress, client)
	require.NoError(t, err)

	// Deploy destination pool
	destPoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		owner,
		client,
		feeTokenAddress,
		lock_release_token_pool.RateLimiterConfig{
			Capacity:  big.NewInt(1e18),
			Rate:      big.NewInt(1e18),
			IsEnabled: true,
		},
	)
	require.NoError(t, err)
	flushLogs()
	destPool, err := lock_release_token_pool.NewLockReleaseTokenPool(destPoolAddress, client)
	require.NoError(t, err)

	// Fund dest pool
	liquidityAmount := big.NewInt(1000000)
	_, err = feeToken.Approve(owner, destPoolAddress, liquidityAmount)
	require.NoError(t, err)
	flushLogs()
	_, err = destPool.AddLiquidity(owner, liquidityAmount)
	require.NoError(t, err)
	flushLogs()

	afnAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		owner,
		client,
	)
	require.NoError(t, err)
	flushLogs()

	// deploy priceRegistry
	priceRegistryAddress, _, _, err := price_registry.DeployPriceRegistry(owner, client, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: feeTokenAddress,
				UsdPerToken: big.NewInt(8e18), // $8
			},
		},
		DestChainId:   sourceChainID,
		UsdPerUnitGas: big.NewInt(2000e9), // $2000
	}, []common.Address{}, []common.Address{feeTokenAddress}, uint32(time.Hour.Seconds())) // 1h
	require.NoError(t, err)
	flushLogs()

	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddress, client)
	require.NoError(t, err)

	onRampAddress := testutils.NewAddress()
	sourceFeeTokenAddress := testutils.NewAddress()
	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		owner,  // user
		client, // client
		commit_store_helper.CommitStoreStaticConfig{
			ChainId:       destChainID,
			SourceChainId: sourceChainID,
			OnRamp:        onRampAddress,
		},
	)
	require.NoError(t, err)
	commitStoreHelper, err := commit_store_helper.NewCommitStoreHelper(commitStoreAddress, client)
	require.NoError(t, err)
	// since CommitStoreHelper inherits CommitStore, it's safe for testing purposes to initialize on the same address
	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, client)
	require.NoError(t, err)
	flushLogs()

	_, err = priceRegistry.ApplyPriceUpdatersUpdates(owner, []common.Address{commitStoreAddress}, []common.Address{})
	require.NoError(t, err)

	routerAddress, _, routerContract, err := router.DeployRouter(owner, client, common.Address{})
	require.NoError(t, err)
	flushLogs()
	offRampAddress, _, _, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(
		owner,
		client,
		evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
			CommitStore:   commitStore.Address(),
			ChainId:       destChainID,
			SourceChainId: sourceChainID,
			OnRamp:        onRampAddress,
		},
		[]common.Address{sourceFeeTokenAddress},
		[]common.Address{destPoolAddress},
		evm_2_evm_offramp.RateLimiterConfig{
			Capacity:  big.NewInt(1e18),
			Rate:      big.NewInt(1e18),
			IsEnabled: true,
		},
	)
	require.NoError(t, err)
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddress, client)
	require.NoError(t, err)
	_, err = destPool.ApplyRampUpdates(owner, []lock_release_token_pool.IPoolRampUpdate{}, []lock_release_token_pool.IPoolRampUpdate{{Ramp: offRampAddress, Allowed: true}})
	require.NoError(t, err)
	receiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(owner, client)
	require.NoError(t, err)
	receiver, err := simple_message_receiver.NewSimpleMessageReceiver(receiverAddress, client)
	require.NoError(t, err)
	flushLogs()
	_, err = routerContract.ApplyRampUpdates(owner, nil, []router.RouterOffRampUpdate{
		{SourceChainId: sourceChainID, OffRamps: []common.Address{offRampAddress}},
	})
	require.NoError(t, err)
	flushLogs()

	commitOnchainConfig := CommitOnchainConfig{Afn: afnAddress, PriceRegistry: priceRegistryAddress}
	encodedCommitOnchainConfig, err := EncodeAbiStruct(commitOnchainConfig)
	require.NoError(t, err)

	_, err = generateAndSetTestOCR2Config(commitStoreHelper, owner, encodedCommitOnchainConfig)
	require.NoError(t, err)

	execOnchainConfig := ExecOnchainConfig{PermissionLessExecutionThresholdSeconds: 600, Router: routerAddress, PriceRegistry: priceRegistryAddress, Afn: afnAddress, MaxTokensLength: 5, MaxDataSize: 200_000}
	encodedExecOnchainConfig, err := EncodeAbiStruct(execOnchainConfig)
	require.NoError(t, err)

	_, err = generateAndSetTestOCR2Config(offRamp, owner, encodedExecOnchainConfig)
	require.NoError(t, err)
	flushLogs()

	// register price update filters in logPoller
	err = logPoller.RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(COMMIT_PRICE_UPDATES, priceRegistryAddress),
		EventSigs: []common.Hash{UsdPerUnitGasUpdated, UsdPerTokenUpdated}, Addresses: []common.Address{priceRegistryAddress},
	})
	require.NoError(t, err)

	return ccipPluginTestHarness{
		destChainID:   destChainID,
		sourceChainID: sourceChainID,
		lggr:          lggr,
		owner:         owner,
		client:        client,
		logPoller:     logPoller,
		flushLogs:     flushLogs,

		offRamp:           offRamp,
		commitStoreHelper: commitStoreHelper,
		commitStore:       commitStore,
		receiver:          receiver,
		priceRegistry:     priceRegistry,

		feeTokenAddress:     feeTokenAddress,
		sourceNativeAddress: testutils.NewAddress(),

		commitOnchainConfig: commitOnchainConfig,
		execOnchainConfig:   execOnchainConfig,
	}
}

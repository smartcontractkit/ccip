package ccipdata_test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store_helper"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestOffRampFilters(t *testing.T) {
	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) ccipdata.Closer {
		c, err := ccipdata.NewOffRampV1_0_0(logger.TestLogger(t), addr, new(mocks.Client), lp, nil)
		require.NoError(t, err)
		return c
	}, 3)
	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) ccipdata.Closer {
		c, err := ccipdata.NewOffRampV1_2_0(logger.TestLogger(t), addr, new(mocks.Client), lp, nil)
		require.NoError(t, err)
		return c
	}, 3)
}

func TestExecOffchainConfig_Encoding(t *testing.T) {
	tests := map[string]struct {
		want      ccipdata.ExecOffchainConfig
		expectErr bool
	}{
		"encodes and decodes config with all fields set": {
			want: ccipdata.ExecOffchainConfig{
				SourceFinalityDepth:         3,
				DestOptimisticConfirmations: 6,
				DestFinalityDepth:           3,
				BatchGasLimit:               5_000_000,
				RelativeBoostPerWaitHour:    0.07,
				MaxGasPrice:                 200e9,
				InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
				RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
			},
		},
		"fails decoding when all fields present but with 0 values": {
			want: ccipdata.ExecOffchainConfig{
				SourceFinalityDepth:         0,
				DestFinalityDepth:           0,
				DestOptimisticConfirmations: 0,
				BatchGasLimit:               0,
				RelativeBoostPerWaitHour:    0,
				MaxGasPrice:                 0,
				InflightCacheExpiry:         models.MustMakeDuration(0),
				RootSnoozeTime:              models.MustMakeDuration(0),
			},
			expectErr: true,
		},
		"fails decoding when all fields are missing": {
			want:      ccipdata.ExecOffchainConfig{},
			expectErr: true,
		},
		"fails decoding when some fields are missing": {
			want: ccipdata.ExecOffchainConfig{
				SourceFinalityDepth: 99999999,
				InflightCacheExpiry: models.MustMakeDuration(64 * time.Second),
			},
			expectErr: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			exp := tc.want
			encode, err := ccipconfig.EncodeOffchainConfig(&exp)
			require.NoError(t, err)
			got, err := ccipconfig.DecodeOffchainConfig[ccipdata.ExecOffchainConfig](encode)

			if tc.expectErr {
				require.ErrorContains(t, err, "must set")
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

func TestExecOnchainConfig100(t *testing.T) {
	tests := []struct {
		name      string
		want      ccipdata.ExecOnchainConfigV1_0_0
		expectErr bool
	}{
		{
			name: "encodes and decodes config with all fields set",
			want: ccipdata.ExecOnchainConfigV1_0_0{
				PermissionLessExecutionThresholdSeconds: rand.Uint32(),
				Router:                                  utils.RandomAddress(),
				PriceRegistry:                           utils.RandomAddress(),
				MaxTokensLength:                         uint16(rand.Uint32()),
				MaxDataSize:                             rand.Uint32(),
			},
		},
		{
			name: "encodes and fails decoding config with missing fields",
			want: ccipdata.ExecOnchainConfigV1_0_0{
				PermissionLessExecutionThresholdSeconds: rand.Uint32(),
				MaxDataSize:                             rand.Uint32(),
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := abihelpers.EncodeAbiStruct(tt.want)
			require.NoError(t, err)

			decoded, err := abihelpers.DecodeAbiStruct[ccipdata.ExecOnchainConfigV1_0_0](encoded)
			if tt.expectErr {
				require.ErrorContains(t, err, "must set")
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, decoded)
			}
		})
	}
}

func TestExecOnchainConfig120(t *testing.T) {
	tests := []struct {
		name      string
		want      ccipdata.ExecOnchainConfigV1_2_0
		expectErr bool
	}{
		{
			name: "encodes and decodes config with all fields set",
			want: ccipdata.ExecOnchainConfigV1_2_0{
				PermissionLessExecutionThresholdSeconds: rand.Uint32(),
				Router:                                  utils.RandomAddress(),
				PriceRegistry:                           utils.RandomAddress(),
				MaxNumberOfTokensPerMsg:                 uint16(rand.Uint32()),
				MaxDataBytes:                            rand.Uint32(),
				MaxPoolReleaseOrMintGas:                 rand.Uint32(),
			},
		},
		{
			name: "encodes and fails decoding config with missing fields",
			want: ccipdata.ExecOnchainConfigV1_2_0{
				PermissionLessExecutionThresholdSeconds: rand.Uint32(),
				MaxDataBytes:                            rand.Uint32(),
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := abihelpers.EncodeAbiStruct(tt.want)
			require.NoError(t, err)

			decoded, err := abihelpers.DecodeAbiStruct[ccipdata.ExecOnchainConfigV1_2_0](encoded)
			if tt.expectErr {
				require.ErrorContains(t, err, "must set")
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, decoded)
			}
		})
	}
}

// The versions to test.
func getVersions() []string {
	return []string{ccipdata.V1_0_0, ccipdata.V1_1_0, ccipdata.V1_2_0}
}

func TestOffRampReaderInit(t *testing.T) {
	for _, version := range getVersions() {
		t.Run("OffRampReader_"+version, func(t *testing.T) {
			setupAndTestOffRampReader(t, version)
		})
	}
}

func newSimulation(t *testing.T) (*bind.TransactOpts, *client.SimulatedBackendClient) {
	user := testutils.MustNewSimTransactor(t)
	sim := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		user.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(50), big.NewInt(1e18)),
		},
	}, 10e8)
	balance, err := sim.BalanceAt(user.Context, user.From, nil)
	require.NoError(t, err)
	logger.TestLogger(t).Debug("User balance: ", balance)
	ec := client.NewSimulatedBackendClient(t, sim, testutils.SimulatedChainID)
	return user, ec
}

// Should be moved to a common test utils package.
func assertNonRevert(t *testing.T, tx *types.Transaction, bc *client.SimulatedBackendClient, user *bind.TransactOpts) {
	receipt, err := bc.TransactionReceipt(user.Context, tx.Hash())
	require.NoError(t, err)
	require.NotEqual(t, uint64(0), receipt.Status, "Transaction should not have reverted")
	logger.TestLogger(t).Debug("Transaction status: ", receipt.Status)
}

func setupAndTestOffRampReader(t *testing.T, version string) {

	user, bc := newSimulation(t)
	log := logger.TestLogger(t)
	orm := logpoller.NewORM(testutils.SimulatedChainID, pgtest.NewSqlxDB(t), log, pgtest.NewQConfig(true))
	lp := logpoller.NewLogPoller(
		orm,
		bc,
		log,
		100*time.Millisecond, 2, 3, 2, 1000)
	//log.Debug("Created user ", user.From)

	// Set gas limit to avoid issue with gas estimator.
	user.GasPrice = big.NewInt(1000000000)
	//user.GasLimit = 0
	user.GasLimit = 10_000_000
	//user.GasLimit = 1_000_000

	// Create the version-specific offRamp:
	// For version 1.0 and 1.1:
	// evm_2_evm_offramp_1_0_0.DeployEVM2EVMOffRamp(..)
	// For version 1.2 only:
	// evm_2_evm_offramp.DeployEVM2EVMOffRamp(..)

	// Deploy a mock ARM.
	armAddr, tx, _, err := mock_arm_contract.DeployMockARMContract(user, bc)
	bc.Commit()
	assertNonRevert(t, tx, bc, user)
	log.Debug("ARM deployed at ", armAddr)

	// Deploy a CommitStore.
	onRampAddr := utils.RandomAddress()
	csAddr, tx, cs, err := commit_store_helper.DeployCommitStoreHelper(user, bc, commit_store_helper.CommitStoreStaticConfig{
		ChainSelector:       testutils.SimulatedChainID.Uint64(),
		SourceChainSelector: testutils.SimulatedChainID.Uint64(),
		OnRamp:              onRampAddr,
		ArmProxy:            armAddr,
	})
	bc.Commit()
	assertNonRevert(t, tx, bc, user)
	log.Debug("Transaction: ", tx)
	log.Debug("CommitStore deployed at ", csAddr)
	log.Debug("CommitStore deployed at ", cs.Address())
	callOpts := &bind.CallOpts{
		From:    user.From,
		Context: context.Background(),
	}

	// Test the deployed CommitStore.
	number, err := cs.GetExpectedNextSequenceNumber(callOpts)
	require.NoError(t, err)
	log.Debug("Expected next sequence number: ", number)
	tav, err := cs.TypeAndVersion(callOpts)
	require.NoError(t, err)
	log.Debug("Tav: ", tav)

	// Deploy the OffRamp.
	staticConfig := evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
		CommitStore:         csAddr,
		ChainSelector:       testutils.SimulatedChainID.Uint64(),
		SourceChainSelector: testutils.SimulatedChainID.Uint64(),
		OnRamp:              onRampAddr,
		PrevOffRamp:         common.Address{},
		ArmProxy:            armAddr,
	}
	sourceTokens := []common.Address{
		//utils.RandomAddress(), // Need to be IERC20.
	}
	pools := []common.Address{
		//utils.RandomAddress(), // Need to be IPool.
	}
	rateLimiterConfig := evm_2_evm_offramp.RateLimiterConfig{
		IsEnabled: false,
		Capacity:  big.NewInt(0),
		Rate:      big.NewInt(0),
	}

	log.Debug("Will deploy offRamp with static config: ", staticConfig)

	offRampAddr, tx, offRamp, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(user, bc, staticConfig, sourceTokens, pools, rateLimiterConfig)
	bc.Commit()
	require.NoError(t, err)
	assertNonRevert(t, tx, bc, user)

	log.Debug("\tTransaction: ", tx)
	log.Debug("\tTx hash: ", tx.Hash())
	log.Debug("\tTx chain ID: ", tx.ChainId())
	log.Debug("\tTx gas price: ", tx.GasPrice())
	log.Debug("\tTx gas : ", tx.Gas())
	log.Debug("\tTx gas fee cap: ", tx.GasFeeCap())
	log.Debug("\tTx cost: ", tx.Cost())
	log.Debug("\tOffRamp at: ", offRampAddr)
	log.Debug("\tOffRamp: ", offRamp.Address())

	//
	receipt, err := bc.TransactionReceipt(user.Context, tx.Hash())
	require.NoError(t, err)
	log.Debug("\tReceipt: block number: ", receipt.BlockNumber)
	log.Debug("\tReceipt: status: ", receipt.Status)
	log.Debug("\tReceipt: gas used: ", receipt.GasUsed)
	log.Debug("\tReceipt: contract address: ", receipt.ContractAddress)
	log.Debug("\tReceipt: type: ", receipt.Type)
	log.Debug("\tReceipt: logs: ", receipt.Logs)

	block, err := bc.BlockByHash(user.Context, receipt.BlockHash)
	require.NoError(t, err)
	log.Debug("Block: number: ", block.Number())
	log.Debug("Block: txs: ", block.Transactions())
	log.Debug("Block: tx: ", block.Transaction(tx.Hash()))

	code, err := bc.Backend().CodeAt(user.Context, offRampAddr, receipt.BlockNumber)
	require.NoError(t, err)
	log.Debug("Code: ", code)
	//

	//or, err = evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddr, bc)
	//require.NoError(t, err)
	//log.Debug("OffRamp: ", or)

	// tmp: exploring blocks for contract presence.
	for block := 0; block < 4; block++ {
		code, err := bc.CodeAt(user.Context, offRampAddr, big.NewInt(int64(block)))
		require.NoError(t, err)
		log.Debug("Code at block ", block, ": ", code)
	}
	bc.Commit()

	log.Debug("OffRamp address: ", offRamp.Address())

	scCs, err := offRamp.GetStaticConfig(callOpts)
	require.NoError(t, err)
	log.Debug("Static config - cs:", scCs)

	owner, err := offRamp.Owner(callOpts)
	require.NoError(t, err)
	require.Equal(t, user.From, owner)

	tav, err = offRamp.TypeAndVersion(callOpts)
	require.NoError(t, err)
	log.Debug("Tav: ", tav)

	// Deploy the version-specific reader.
	reader, err := ccipdata.NewOffRampReader(log, offRampAddr, bc, lp, nil)
	require.NoError(t, err)
	assertNonRevert(t, tx, bc, user)
	require.Equal(t, offRampAddr, reader.Address())

	res, err := reader.GetDestinationTokens(user.Context)
	require.NoError(t, err)
	require.Equal(t, []common.Address{}, res)
}

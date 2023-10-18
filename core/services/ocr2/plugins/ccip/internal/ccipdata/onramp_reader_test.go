package ccipdata

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_1_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type readerTH struct {
	lp     logpoller.LogPollerTest
	ec     client.Client
	log    logger.Logger
	user   *bind.TransactOpts
	reader OnRampReader

	// Expected state
	//blockTs              []uint64
	//expectedFeeTokens    []common.Address
	//expectedGasUpdates   map[uint64][]ccipdata.GasPrice
	//expectedTokenUpdates map[uint64][]ccipdata.TokenPrice
	//dest                 uint64
}

// The versions to test.
func getVersions() []string {
	return []string{V1_0_0, V1_1_0, V1_2_0}
}

func Test_OnRampReader(t *testing.T) {
	// Assert all readers produce the same expected results.
	for _, version := range getVersions() {
		t.Run("OnRampReader_"+version, func(t *testing.T) {
			setupAndTestOnRampReader(t, version)
		})
	}
}

func setupAndTestOnRampReader(t *testing.T, version string) {
	th := setupOnRampReaderTH(t, version)
	testOnRampReader(t, th, version)
}

func setupOnRampReaderTH(t *testing.T, version string) readerTH {
	user, bc := newSim(t)
	log := logger.TestLogger(t)
	lp := logpoller.NewLogPoller(logpoller.NewORM(testutils.SimulatedChainID, pgtest.NewSqlxDB(t), log, pgtest.NewQConfig(true)),
		bc,
		log,
		100*time.Millisecond, 2, 3, 2, 1000)

	// Setup onRamp.
	//var onRampAddress common.Address
	//switch version {
	//case V1_0_0:
	//	onRampAddress = setupOnRampV1_0_0(t, user, bc, log)
	//case V1_1_0:
	//	onRampAddress = setupOnRampV1_1_0(t, user, bc, log)
	//default:
	//	require.Fail(t, "Unknown version: ", version)
	//}
	onRampAddress := setupOnRampV1_1_0(t, user, bc, log)

	// Create the reader.
	log.Debug("Creating the onRampReader...")
	reader, err := NewOnRampReader(log, testutils.SimulatedChainID.Uint64(), testutils.SimulatedChainID.Uint64(), onRampAddress, lp, bc, true)
	require.NoError(t, err)

	return readerTH{
		lp:     lp,
		ec:     bc,
		log:    log,
		user:   user,
		reader: reader,
	}
}

func setupOnRampV1_0_0(t *testing.T, user *bind.TransactOpts, bc *client.SimulatedBackendClient, log logger.SugaredLogger) common.Address {

	linkTokenAddress := common.HexToAddress("0x000011")

	staticConfig := evm_2_evm_onramp_1_0_0.EVM2EVMOnRampStaticConfig{
		LinkToken:         linkTokenAddress,
		ChainSelector:     testutils.SimulatedChainID.Uint64(),
		DestChainSelector: testutils.SimulatedChainID.Uint64(),
		DefaultTxGasLimit: 30000,
		MaxNopFeesJuels:   big.NewInt(1000000),
		PrevOnRamp:        common.HexToAddress("0x000009"),
		ArmProxy:          common.HexToAddress("0x000008"),
	}

	dynamicConfig := evm_2_evm_onramp_1_0_0.EVM2EVMOnRampDynamicConfig{
		Router:          common.HexToAddress("0x000100"),
		MaxTokensLength: 4,
		PriceRegistry:   common.HexToAddress("0x000066"),
		MaxDataSize:     100000,
		MaxGasLimit:     100000,
	}

	rateLimiterConfig := evm_2_evm_onramp_1_0_0.RateLimiterConfig{
		IsEnabled: false,
		Capacity:  big.NewInt(5),
		Rate:      big.NewInt(5),
	}

	allowList := []common.Address{user.From}
	feeTokenConfigs := []evm_2_evm_onramp_1_0_0.EVM2EVMOnRampFeeTokenConfigArgs{
		{
			Token:                 linkTokenAddress,
			GasMultiplier:         1,
			NetworkFeeAmountUSD:   big.NewInt(0),
			DestGasOverhead:       50,
			DestGasPerPayloadByte: 60,
			Enabled:               false,
		},
	}
	tokenTransferConfigArgs := []evm_2_evm_onramp_1_0_0.EVM2EVMOnRampTokenTransferFeeConfigArgs{
		{
			Token:  common.HexToAddress("0x111111"),
			MinFee: 10,
			MaxFee: 1000,
			Ratio:  1,
		},
	}
	nopsAndWeights := []evm_2_evm_onramp_1_0_0.EVM2EVMOnRampNopAndWeight{
		{
			Nop:    common.HexToAddress("0x222222222"),
			Weight: 1,
		},
	}
	//tokenAndPool := []evm_2_evm_onramp_1_0_0.InternalPoolUpdate{
	//	{
	//		Token: linkTokenAddress,
	//		Pool:  common.HexToAddress("0x333111111"),
	//	},
	//}
	tokenAndPool := []evm_2_evm_onramp_1_0_0.InternalPoolUpdate{}

	user.GasPrice = big.NewInt(10000000000)
	user.GasLimit = 0

	onRampAddress, transaction, onRamp, err := evm_2_evm_onramp_1_0_0.DeployEVM2EVMOnRamp(
		user,
		bc,
		staticConfig,
		dynamicConfig,
		tokenAndPool,
		allowList,
		rateLimiterConfig,
		feeTokenConfigs,
		tokenTransferConfigArgs,
		nopsAndWeights,
	)
	bc.Commit()

	require.NoError(t, err)
	require.NotNil(t, onRampAddress)
	require.NotNil(t, transaction)
	require.NotNil(t, onRamp)

	// Test calls to onRamp (sanity check).
	callOpts := bind.CallOpts{
		From:    user.From,
		Context: context.Background(),
	}
	bc.Commit()

	res, err := onRamp.GetDynamicConfig(&callOpts)
	require.NoError(t, err)
	require.NotNil(t, res)
	log.Debug("DynamicConfig: ", res)

	tav, err := onRamp.TypeAndVersion(&callOpts)
	require.NoError(t, err)
	require.NotNil(t, tav)
	log.Debug("TypeAndVersion: ", tav)
	return onRampAddress
}

func setupOnRampV1_1_0(t *testing.T, user *bind.TransactOpts, bc *client.SimulatedBackendClient, log logger.SugaredLogger) common.Address {

	linkTokenAddress := common.HexToAddress("0x000011")

	staticConfig := evm_2_evm_onramp_1_1_0.EVM2EVMOnRampStaticConfig{
		LinkToken:         linkTokenAddress,
		ChainSelector:     testutils.SimulatedChainID.Uint64(),
		DestChainSelector: testutils.SimulatedChainID.Uint64(),
		DefaultTxGasLimit: 30000,
		MaxNopFeesJuels:   big.NewInt(1000000),
		PrevOnRamp:        common.HexToAddress("0x000009"),
		ArmProxy:          common.HexToAddress("0x000008"),
	}

	dynamicConfig := evm_2_evm_onramp_1_1_0.EVM2EVMOnRampDynamicConfig{
		Router:          common.HexToAddress("0x000100"),
		MaxTokensLength: 4,
		PriceRegistry:   common.HexToAddress("0x000066"),
		MaxDataSize:     100000,
		MaxGasLimit:     100000,
	}

	rateLimiterConfig := evm_2_evm_onramp_1_1_0.RateLimiterConfig{
		IsEnabled: false,
		Capacity:  big.NewInt(5),
		Rate:      big.NewInt(5),
	}

	allowList := []common.Address{user.From}
	feeTokenConfigs := []evm_2_evm_onramp_1_1_0.EVM2EVMOnRampFeeTokenConfigArgs{
		{
			Token:                  linkTokenAddress,
			NetworkFeeUSD:          0,
			MinTokenTransferFeeUSD: 0,
			MaxTokenTransferFeeUSD: 0,
			GasMultiplier:          0,
			PremiumMultiplier:      0,
			Enabled:                false,
		},
		//GasMultiplier:         1,
		//NetworkFeeAmountUSD:   big.NewInt(0),
		//DestGasOverhead:       50,
		//DestGasPerPayloadByte: 60,
		//Enabled:               false,
	}
	tokenTransferConfigArgs := []evm_2_evm_onramp_1_1_0.EVM2EVMOnRampTokenTransferFeeConfigArgs{
		{
			Token:           linkTokenAddress,
			Ratio:           0,
			DestGasOverhead: 0,
		},
	}
	nopsAndWeights := []evm_2_evm_onramp_1_1_0.EVM2EVMOnRampNopAndWeight{
		{
			Nop:    common.HexToAddress("0x222222222"),
			Weight: 1,
		},
	}
	tokenAndPool := []evm_2_evm_onramp_1_1_0.InternalPoolUpdate{}

	user.GasPrice = big.NewInt(10000000000)
	user.GasLimit = 0

	onRampAddress, transaction, onRamp, err := evm_2_evm_onramp_1_1_0.DeployEVM2EVMOnRamp(
		user,
		bc,
		staticConfig,
		dynamicConfig,
		tokenAndPool,
		allowList,
		rateLimiterConfig,
		feeTokenConfigs,
		tokenTransferConfigArgs,
		nopsAndWeights,
	)
	bc.Commit()

	require.NoError(t, err)
	require.NotNil(t, onRampAddress)
	require.NotNil(t, transaction)
	require.NotNil(t, onRamp)

	// Test calls to onRamp (sanity check).
	callOpts := bind.CallOpts{
		From:    user.From,
		Context: context.Background(),
	}
	bc.Commit()

	res, err := onRamp.GetDynamicConfig(&callOpts)
	require.NoError(t, err)
	require.NotNil(t, res)
	log.Debug("DynamicConfig: ", res)

	tav, err := onRamp.TypeAndVersion(&callOpts)
	require.NoError(t, err)
	require.NotNil(t, tav)
	log.Debug("TypeAndVersion: ", tav)
	return onRampAddress
}

func newSim(t *testing.T) (*bind.TransactOpts, *client.SimulatedBackendClient) {
	user := testutils.MustNewSimTransactor(t)
	sim := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		user.From: {
			Balance: big.NewInt(3).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10e6)
	backendClient := client.NewSimulatedBackendClient(t, sim, testutils.SimulatedChainID)
	return user, backendClient
}

func testOnRampReader(t *testing.T, th readerTH, version string) {
	switch version {
	case V1_0_0:
		testV1_0_0(t, th)
	case V1_1_0:
		testV1_1_0(t, th)
	case V1_2_0:
		testV1_2_0(t, th)
	default:
		require.Fail(t, "Unknown version: ", version)
	}
}

func testV1_0_0(t *testing.T, th readerTH) {
	res, err := th.reader.RouterAddress()
	require.NoError(t, err)
	require.Equal(t, "0x0000000000000000000000000000000000000100", res.Hex())

	// TODO
}

func testV1_1_0(t *testing.T, th readerTH) {
	res, err := th.reader.RouterAddress()
	require.NoError(t, err)
	require.Equal(t, "0x0000000000000000000000000000000000000100", res.Hex()) // TODO fix actual value.

	// TODO
}

func testV1_2_0(t *testing.T, th readerTH) {
	res, err := th.reader.RouterAddress()
	require.NoError(t, err)
	require.Equal(t, "0x0000000000000000000000000000000000000100", res.Hex()) // TODO fix actual value.

	// TODO
}

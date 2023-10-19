package ccipdata

import (
	"context"
	"errors"
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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_1_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type readerTH struct {
	lp     logpoller.LogPollerTest
	ec     client.Client
	log    logger.Logger
	user   *bind.TransactOpts
	reader OnRampReader
}

// The versions to test.
func getVersions() []string {
	return []string{V1_0_0, V1_1_0, V1_2_0}
}

func TestOnRampReaderInit(t *testing.T) {
	for _, version := range getVersions() {
		t.Run("OnRampReader_"+version, func(t *testing.T) {
			setupAndTestOnRampReader(t, version)
		})
	}
}

func setupAndTestOnRampReader(t *testing.T, version string) {
	th := setupOnRampReaderTH(t, version)
	testVersionSpecificOnRampReader(t, th, version)
}

func setupOnRampReaderTH(t *testing.T, version string) readerTH {
	user, bc := newSim(t)
	log := logger.TestLogger(t)
	orm := logpoller.NewORM(testutils.SimulatedChainID, pgtest.NewSqlxDB(t), log, pgtest.NewQConfig(true))
	lp := logpoller.NewLogPoller(
		orm,
		bc,
		log,
		100*time.Millisecond, 2, 3, 2, 1000)

	// Setup onRamp.
	var onRampAddress common.Address
	switch version {
	case V1_0_0:
		onRampAddress = setupOnRampV1_0_0(t, user, bc, log)
	case V1_1_0:
		onRampAddress = setupOnRampV1_1_0(t, user, bc, log)
	case V1_2_0:
		onRampAddress = setupOnRampV1_2_0(t, user, bc, log)
	default:
		require.Fail(t, "Unknown version: ", version)
	}

	// Insert log messages.
	topic := []byte("Commit ccip sends - " + onRampAddress.String())
	topics := [][]byte{topic}
	err := orm.InsertLogs([]logpoller.Log{
		{
			EvmChainId:  utils.NewBigI(testutils.SimulatedChainID.Int64()),
			LogIndex:    0,
			BlockHash:   common.HexToHash("0x3333"),
			BlockNumber: 1,
			EventSig:    common.BytesToHash([]byte(onRampAddress.String())),
			Topics:      topics,
			Address:     onRampAddress,
			TxHash:      common.HexToHash("0x000011"),
			Data:        append([]byte("hello 0"), byte(0)),
		},
		{
			EvmChainId:  utils.NewBigI(testutils.SimulatedChainID.Int64()),
			LogIndex:    0,
			BlockHash:   common.HexToHash("0x3333"),
			BlockNumber: 2,
			EventSig:    common.BytesToHash([]byte(onRampAddress.String())),
			Topics:      topics,
			Address:     onRampAddress,
			TxHash:      common.HexToHash("0x11111"),
			Data:        append([]byte("hello 1"), byte(1)),
		},
		{
			EvmChainId:  utils.NewBigI(testutils.SimulatedChainID.Int64()),
			LogIndex:    0,
			BlockHash:   common.HexToHash("0x3333"),
			BlockNumber: 3,
			EventSig:    common.BytesToHash([]byte(onRampAddress.String())),
			Topics:      topics,
			Address:     onRampAddress,
			TxHash:      common.HexToHash("0x22222"),
			Data:        append([]byte("hello 2"), byte(2)),
		},
	})
	require.NoError(t, err)
	bc.Commit()

	// Create the version-specific reader.
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
		Router:          common.HexToAddress("0x000110"),
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

func setupOnRampV1_2_0(t *testing.T, user *bind.TransactOpts, bc *client.SimulatedBackendClient, log logger.SugaredLogger) common.Address {

	linkTokenAddress := common.HexToAddress("0x000011")
	staticConfig := evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
		LinkToken:         linkTokenAddress,
		ChainSelector:     testutils.SimulatedChainID.Uint64(),
		DestChainSelector: testutils.SimulatedChainID.Uint64(),
		DefaultTxGasLimit: 30000,
		MaxNopFeesJuels:   big.NewInt(1000000),
		PrevOnRamp:        common.HexToAddress("0x000009"),
		ArmProxy:          common.HexToAddress("0x000008"),
	}
	dynamicConfig := evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
		Router:                            common.HexToAddress("0x000120"),
		MaxNumberOfTokensPerMsg:           0,
		DestGasOverhead:                   0,
		DestGasPerPayloadByte:             0,
		DestDataAvailabilityOverheadGas:   0,
		DestGasPerDataAvailabilityByte:    0,
		DestDataAvailabilityMultiplierBps: 0,
		PriceRegistry:                     common.HexToAddress("0x000777"),
		MaxDataBytes:                      0,
		MaxPerMsgGasLimit:                 0,
	}
	rateLimiterConfig := evm_2_evm_onramp.RateLimiterConfig{
		IsEnabled: false,
		Capacity:  big.NewInt(5),
		Rate:      big.NewInt(5),
	}
	feeTokenConfigs := []evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
		{
			Token:                      linkTokenAddress,
			NetworkFeeUSDCents:         0,
			GasMultiplierWeiPerEth:     0,
			PremiumMultiplierWeiPerEth: 0,
			Enabled:                    false,
		},
	}
	tokenTransferConfigArgs := []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
		{
			Token:             linkTokenAddress,
			MinFeeUSDCents:    0,
			MaxFeeUSDCents:    0,
			DeciBps:           0,
			DestGasOverhead:   0,
			DestBytesOverhead: 0,
		},
	}
	nopsAndWeights := []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{
		{
			Nop:    common.HexToAddress("0x222222222"),
			Weight: 1,
		},
	}
	tokenAndPool := []evm_2_evm_onramp.InternalPoolUpdate{}
	user.GasPrice = big.NewInt(10000000000)
	user.GasLimit = 0
	onRampAddress, transaction, onRamp, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		user,
		bc,
		staticConfig,
		dynamicConfig,
		tokenAndPool,
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

func testVersionSpecificOnRampReader(t *testing.T, th readerTH, version string) {
	switch version {
	case V1_0_0:
		testOnRampReader(t, th, common.HexToAddress("0x0000000000000000000000000000000000000100"))
	case V1_1_0:
		testOnRampReader(t, th, common.HexToAddress("0x0000000000000000000000000000000000000110"))
	case V1_2_0:
		testOnRampReader(t, th, common.HexToAddress("0x0000000000000000000000000000000000000120"))
	default:
		require.Fail(t, "Unknown version: ", version)
	}
}

func testOnRampReader(t *testing.T, th readerTH, expectedRouterAddress common.Address) {

	res, err := th.reader.RouterAddress()
	require.NoError(t, err)
	require.Equal(t, expectedRouterAddress, res)

	//th.lp.PollAndSaveLogs(th.user.Context, 3)

	_, err = th.reader.GetSendRequestsGteSeqNum(th.user.Context, 0, 0)
	require.Error(t, err, errors.New("latest finalized header is nil")) // requires logs to be polled.

	msg, err := th.reader.GetSendRequestsBetweenSeqNums(th.user.Context, 0, 10, 0)
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, []Event[internal.EVM2EVMMessage]{}, msg)
}

package ccipdata

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
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

func TestNewOnRampReader__noContractAtAddress(t *testing.T) {
	_, bc := NewSimulation(t)
	log := logger.TestLogger(t)
	orm := logpoller.NewORM(testutils.SimulatedChainID, pgtest.NewSqlxDB(t), log, pgtest.NewQConfig(true))
	lp := logpoller.NewLogPoller(
		orm,
		bc,
		log,
		100*time.Millisecond, 2, 3, 2, 1000)

	_, err := NewOnRampReader(
		logger.TestLogger(t),
		testutils.SimulatedChainID.Uint64(), testutils.SimulatedChainID.Uint64(),
		common.Address{},
		lp,
		bc,
		true)
	assert.EqualError(t, err, "expected 'EVM2EVMOnRamp' got '' (no contract code at given address)")
}

func TestOnRampReaderInit(t *testing.T) {

	tests := []struct {
		name    string
		version string
	}{
		{
			name:    "OnRampReader_V1_0_0",
			version: V1_0_0,
		},
		{
			name:    "OnRampReader_V1_1_0",
			version: V1_1_0,
		},
		{
			name:    "OnRampReader_V1_2_0",
			version: V1_2_0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := setupOnRampReaderTH(t, test.version)
			testVersionSpecificOnRampReader(t, th, test.version)
		})
	}
}

func setupOnRampReaderTH(t *testing.T, version string) readerTH {
	user, bc := NewSimulation(t)
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
		PrevOnRamp:        common.Address{},
		ArmProxy:          utils.RandomAddress(),
	}
	dynamicConfig := evm_2_evm_onramp_1_0_0.EVM2EVMOnRampDynamicConfig{
		Router:          common.HexToAddress("0x000100"),
		MaxTokensLength: 4,
		PriceRegistry:   utils.RandomAddress(),
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
			Token:  utils.RandomAddress(),
			MinFee: 10,
			MaxFee: 1000,
			Ratio:  1,
		},
	}
	nopsAndWeights := []evm_2_evm_onramp_1_0_0.EVM2EVMOnRampNopAndWeight{
		{
			Nop:    utils.RandomAddress(),
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
	AssertNonRevert(t, transaction, bc, user)
	require.NotNil(t, onRampAddress)
	require.NotNil(t, transaction)
	require.NotNil(t, onRamp)
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
	AssertNonRevert(t, transaction, bc, user)
	require.NotNil(t, onRampAddress)
	require.NotNil(t, transaction)
	require.NotNil(t, onRamp)
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
	AssertNonRevert(t, transaction, bc, user)
	require.NotNil(t, onRampAddress)
	require.NotNil(t, transaction)
	require.NotNil(t, onRamp)

	return onRampAddress
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

	_, err = th.reader.GetSendRequestsGteSeqNum(th.user.Context, 0, 0)
	require.Error(t, err, errors.New("latest finalized header is nil")) // requires logs to be polled.

	msg, err := th.reader.GetSendRequestsBetweenSeqNums(th.user.Context, 0, 10, 0)
	require.NoError(t, err)
	require.NotNil(t, msg)
	require.Equal(t, []Event[internal.EVM2EVMMessage]{}, msg)
}

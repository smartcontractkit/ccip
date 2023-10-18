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

func Test_OnRampReader(t *testing.T) {

	// Test all versions.
	versions := []string{V1_0_0, V1_1_0, V1_2_0}

	// Assert all readers produce the same expected results.
	for _, version := range versions {
		t.Run("OnRampReader_"+version, func(t *testing.T) {
			testOnRampReader(t, version)
		})
	}
}

func testOnRampReader(t *testing.T, version string) {
	th := setupOnRampReaderTH(t, version)
	res, err := th.reader.RouterAddress()
	require.NoError(t, err)
	require.Equal(t, res.Hex(), "0x5550000000000000000000000000000000000001") // TODO fix actual value.
}

func setupOnRampReaderTH(t *testing.T, version string) readerTH {
	user, bc := newSim(t)
	log := logger.TestLogger(t)
	lp := logpoller.NewLogPoller(logpoller.NewORM(testutils.SimulatedChainID, pgtest.NewSqlxDB(t), log, pgtest.NewQConfig(true)),
		bc,
		log,
		100*time.Millisecond, 2, 3, 2, 1000)

	//addr, _, _, err := price_registry_1_0_0.DeployPriceRegistry(user, bc, nil, feeTokens, 1000)
	//onramp_v1_0_0, _, _, err := onramp_v1price_registry_1_0_0.DeployPriceRegistry(user, bc, nil, []common.Address{}, 1000)
	//onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")
	//onRampAddress := common.HexToAddress("0x1110000000001")

	//ramp, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampAddress, bc)
	//require.NoError(t, err)
	//require.Equal(t, onRampAddress, ramp.Address())
	//log.Debug("Ramp: ", ramp.Address().Hex())

	// Create the onRamp.
	// TODO init based on version.

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
		Router:          common.HexToAddress("0x000055"),
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
	tokenAndPool := []evm_2_evm_onramp_1_0_0.InternalPoolUpdate{
		{
			Token: common.HexToAddress("0x111111"),
			Pool:  common.HexToAddress("0x333111111"),
		},
	}

	user.GasPrice = big.NewInt(10000000000)
	user.GasLimit = 600000

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
	log.Debug("Deployed onRamp at: ", onRampAddress.Hex())
	log.Debug("OnRamp: ", onRamp.Address())
	log.Debug("Transaction: ", transaction.Hash().Hex())

	require.NoError(t, err)
	require.NotNil(t, onRampAddress)
	require.NotNil(t, transaction)
	require.NotNil(t, onRamp)
	//require.Equal(t, onRampAddress, t2)

	require.NotNil(t, onRamp.Address())
	log.Debug("OnRamp: ", onRamp.Address())

	callOpts := bind.CallOpts{
		From:    user.From,
		Context: context.Background(),
	}
	bc.Commit()

	res, err := onRamp.GetDynamicConfig(&callOpts)
	require.NoError(t, err)
	require.NotNil(t, res)

	tav, err := onRamp.TypeAndVersion(&callOpts)
	require.NoError(t, err)
	require.NotNil(t, tav)
	//or, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampAddress, bc)
	//require.NoError(t, err)

	//func NewOnRampReader(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source bc.Client, finalityTags bool, qopts ...pg.QOpt) (OnRampReader, error)

	// TODO

	log.Debug("Creating the onRampReader...")
	reader, err := NewOnRampReader(log, testutils.SimulatedChainID.Uint64(), testutils.SimulatedChainID.Uint64(), onRampAddress, lp, bc, true)
	require.NoError(t, err)

	//var reader OnRampReader
	//var err error
	//switch version {
	//case V1_0_0:
	//	reader, err = NewOnRampV1_0_0(log, 1, 4, onRampAddress, lp, bc, false)
	//case V1_1_0:
	//	reader, err = NewOnRampV1_1_0(log, 1, 4, onRampAddress, lp, bc, false)
	//case V1_2_0:
	//	reader, err = NewOnRampV1_2_0(log, 1, 4, onRampAddress, lp, bc, false)
	//default:
	//	require.Fail(t, "Unexpected version: "+version)
	//}
	require.NoError(t, err)

	return readerTH{
		lp:     lp,
		ec:     bc,
		log:    log,
		user:   user,
		reader: reader,
	}
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

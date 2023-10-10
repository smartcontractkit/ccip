package ccipdata_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

func TestPriceRegistryFilters(t *testing.T) {
	cl := mocks.NewClient(t)
	cl.On("ConfiguredChainID").Return(big.NewInt(1))

	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) ccipdata.Closer {
		c, err := ccipdata.NewPriceRegistryV1_0_0(logger.TestLogger(t), addr, lp, cl)
		require.NoError(t, err)
		return c
	}, 3)

	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) ccipdata.Closer {
		c, err := ccipdata.NewPriceRegistryV1_2_0(logger.TestLogger(t), addr, lp, cl)
		require.NoError(t, err)
		return c
	}, 3)
}

type priceRegReaderTH struct {
	lp      logpoller.LogPollerTest
	ec      client.Client
	lggr    logger.Logger
	user    *bind.TransactOpts
	sim     *backends.SimulatedBackend
	readers []ccipdata.PriceRegistryReader

	// Expected state
	expectedFeeTokens  []common.Address
	expectedGasUpdates []ccipdata.GasPrice
}

// setupPriceRegistryReaderTH instatiates all versions of the price registry reader
// with a snapshot of data so reader tests can do multi-version assertions.
func setupPriceRegistryReaderTH(t *testing.T) priceRegReaderTH {
	user := testutils.MustNewSimTransactor(t)
	sim := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		user.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10e6)
	sim.Commit()
	ec := client.NewSimulatedBackendClient(t, sim, testutils.SimulatedChainID)
	lggr := logger.TestLogger(t)
	// TODO: We should be able to use an in memory log poller ORM here to speed up the tests.
	lp := logpoller.NewLogPoller(logpoller.NewORM(testutils.SimulatedChainID, pgtest.NewSqlxDB(t), lggr, pgtest.NewQConfig(true)), ec, lggr, 100*time.Millisecond, 2, 3, 2, 1000)

	feeTokens := []common.Address{randomAddress(), randomAddress()}
	gasPriceUpdates := []ccipdata.GasPrice{
		{
			DestChainSelector: uint64(10),
			Value:             big.NewInt(11),
		},
	}
	addr, _, pr, err := price_registry_1_0_0.DeployPriceRegistry(user, sim, nil, feeTokens, 1000)
	require.NoError(t, err)
	_, err = pr.UpdatePrices(user, price_registry_1_0_0.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry_1_0_0.InternalTokenPriceUpdate{},
		DestChainSelector: gasPriceUpdates[0].DestChainSelector,
		UsdPerUnitGas:     gasPriceUpdates[0].Value,
	})
	require.NoError(t, err)
	sim.Commit()
	pr10, err := ccipdata.NewPriceRegistryReader(lggr, addr, lp, ec)
	require.NoError(t, err)

	addr, _, pr2, err := price_registry.DeployPriceRegistry(user, sim, nil, feeTokens, 1000)
	require.NoError(t, err)
	_, err = pr2.UpdatePrices(user, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{},
		GasPriceUpdates: []price_registry.InternalGasPriceUpdate{
			{
				DestChainSelector: gasPriceUpdates[0].DestChainSelector,
				UsdPerUnitGas:     gasPriceUpdates[0].Value,
			},
		},
	})
	require.NoError(t, err)
	sim.Commit()
	pr12r, err := ccipdata.NewPriceRegistryReader(lggr, addr, lp, ec)
	require.NoError(t, err)

	// Capture all lp data.
	lp.PollAndSaveLogs(context.Background(), 1)

	return priceRegReaderTH{
		lp:                 lp,
		ec:                 ec,
		lggr:               lggr,
		user:               user,
		sim:                sim,
		readers:            []ccipdata.PriceRegistryReader{pr10, pr12r},
		expectedFeeTokens:  feeTokens,
		expectedGasUpdates: gasPriceUpdates,
	}
}

func TestPriceRegistryReader(t *testing.T) {
	th := setupPriceRegistryReaderTH(t)
	for _, pr := range th.readers {
		// Assert fee token read.
		gotFeeTokens, err := pr.GetFeeTokens(context.Background())
		require.NoError(t, err)
		assert.Equal(t, th.expectedFeeTokens, gotFeeTokens)

		// Assert latest gas price read.
		updates, err := pr.GetGasPriceUpdatesCreatedAfter(context.Background(), th.expectedGasUpdates[0].DestChainSelector, time.Unix(0, 0), 0)
		require.NoError(t, err)
		require.Equal(t, len(updates), 1)
		assert.Equal(t, updates[0].Data.GasPrice, th.expectedGasUpdates[0])

		// TODO test other price reader methods.
	}
}

package commitplugin

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestCommitInflight(t *testing.T) {
	lggr := logger.TestLogger(t)
	c := newInflightCommitReportsContainer(time.Hour)

	c.inFlightPriceUpdates = append(c.inFlightPriceUpdates, inflightPriceUpdate{
		gasPrices:     []ccipdata.GasPrice{},
		createdAt:     time.Now(),
		epochAndRound: ccipcalc.MergeEpochAndRound(2, 4),
	})

	// Initially should be empty
	inflightUpdate, hasUpdate := c.getLatestInflightGasPriceUpdate()
	assert.Equal(t, inflightUpdate, update{})
	assert.False(t, hasUpdate)
	assert.Equal(t, uint64(0), c.maxInflightSeqNr())

	epochAndRound := uint64(1)

	// Add a single report inflight
	root1 := utils.Keccak256Fixed(hexutil.MustDecode("0xaa"))
	require.NoError(t, c.add(lggr, ccipdata.CommitStoreReport{
		Interval:   ccipdata.CommitStoreInterval{Min: 1, Max: 2},
		MerkleRoot: root1,
		GasPrices: []ccipdata.GasPrice{
			{DestChainSelector: 123, Value: big.NewInt(999)},
		},
	}, epochAndRound))
	inflightUpdate, hasUpdate = c.getLatestInflightGasPriceUpdate()
	assert.Equal(t, big.NewInt(999), inflightUpdate.value)
	assert.True(t, hasUpdate)
	assert.Equal(t, uint64(2), c.maxInflightSeqNr())
	epochAndRound++

	// Add another price report
	root2 := utils.Keccak256Fixed(hexutil.MustDecode("0xab"))
	require.NoError(t, c.add(lggr, ccipdata.CommitStoreReport{
		Interval:   ccipdata.CommitStoreInterval{Min: 3, Max: 4},
		MerkleRoot: root2,
		GasPrices: []ccipdata.GasPrice{
			{DestChainSelector: 321, Value: big.NewInt(888)},
		},
	}, epochAndRound))
	inflightUpdate, hasUpdate = c.getLatestInflightGasPriceUpdate()
	assert.Equal(t, big.NewInt(888), inflightUpdate.value)
	assert.True(t, hasUpdate)
	assert.Equal(t, uint64(4), c.maxInflightSeqNr())
	epochAndRound++

	// Add gas price updates
	require.NoError(t, c.add(lggr, ccipdata.CommitStoreReport{
		GasPrices: []ccipdata.GasPrice{
			{
				DestChainSelector: uint64(1),
				Value:             big.NewInt(1),
			},
		}}, epochAndRound))

	inflightUpdate, hasUpdate = c.getLatestInflightGasPriceUpdate()
	assert.Equal(t, big.NewInt(1), inflightUpdate.value)
	assert.True(t, hasUpdate)
	assert.Equal(t, uint64(4), c.maxInflightSeqNr())
	epochAndRound++

	// Add a token price update
	token := common.HexToAddress("0xa")
	require.NoError(t, c.add(lggr, ccipdata.CommitStoreReport{
		TokenPrices: []ccipdata.TokenPrice{
			{
				Token: token,
				Value: big.NewInt(10),
			},
		},
		GasPrices: []ccipdata.GasPrice{{}},
	}, epochAndRound))
	// Apply cache price to existing
	latestInflightTokenPriceUpdates := c.latestInflightTokenPriceUpdates()
	require.Equal(t, len(latestInflightTokenPriceUpdates), 1)
	assert.Equal(t, big.NewInt(10), latestInflightTokenPriceUpdates[token].value)

	// larger epoch and round overrides existing price update
	c.inFlightPriceUpdates = append(c.inFlightPriceUpdates, inflightPriceUpdate{
		tokenPrices: []ccipdata.TokenPrice{
			{Token: token, Value: big.NewInt(9999)},
		},
		gasPrices: []ccipdata.GasPrice{
			{
				DestChainSelector: uint64(1),
				Value:             big.NewInt(999),
			},
		},
		createdAt:     time.Now(),
		epochAndRound: ccipcalc.MergeEpochAndRound(999, 99),
	})
	latestInflightTokenPriceUpdates = c.latestInflightTokenPriceUpdates()
	require.Equal(t, len(latestInflightTokenPriceUpdates), 1)
	assert.Equal(t, big.NewInt(9999), latestInflightTokenPriceUpdates[token].value)
}

func Test_inflightCommitReportsContainer_expire(t *testing.T) {
	c := &inflightCommitReportsContainer{
		cacheExpiry: time.Minute,
		inFlight: map[[32]byte]inflightReport{
			common.HexToHash("1"): {
				report:    ccipdata.CommitStoreReport{},
				createdAt: time.Now().Add(-5 * time.Minute),
			},
			common.HexToHash("2"): {
				report:    ccipdata.CommitStoreReport{},
				createdAt: time.Now().Add(-10 * time.Second),
			},
		},
		inFlightPriceUpdates: []inflightPriceUpdate{
			{
				gasPrices:     []ccipdata.GasPrice{{DestChainSelector: 100, Value: big.NewInt(0)}},
				createdAt:     time.Now().Add(-priceExpiryMultiplier * time.Minute),
				epochAndRound: ccipcalc.MergeEpochAndRound(10, 5),
			},
			{
				gasPrices:     []ccipdata.GasPrice{{DestChainSelector: 200, Value: big.NewInt(0)}},
				createdAt:     time.Now().Add(-priceExpiryMultiplier * time.Second),
				epochAndRound: ccipcalc.MergeEpochAndRound(20, 5),
			},
		},
	}
	c.expire(logger.NullLogger)

	assert.Len(t, c.inFlight, 1)
	_, exists := c.inFlight[common.HexToHash("2")]
	assert.True(t, exists)

	assert.Len(t, c.inFlightPriceUpdates, 1)
	assert.Equal(t, ccipcalc.MergeEpochAndRound(20, 5), c.inFlightPriceUpdates[0].epochAndRound)
}

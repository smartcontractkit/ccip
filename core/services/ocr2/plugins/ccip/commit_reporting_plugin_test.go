package ccip

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/v2/common/txmgr/types/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type commitTestHarness = struct {
	ccipPluginTestHarness
	plugin           *CommitReportingPlugin
	mockFeeEstimator *mocks.FeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, evmtypes.TxHash]
}

func setupCommitTestHarness(t *testing.T) commitTestHarness {
	th := setupCcipTestHarness(t)

	sourceFeeEstimator := mocks.NewFeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, evmtypes.TxHash](t)

	plugin := CommitReportingPlugin{
		config: CommitPluginConfig{
			lggr:               th.lggr,
			dest:               th.logPoller,
			commitStore:        th.commitStore,
			priceRegistry:      th.priceRegistry,
			priceGetter:        fakePriceGetter{},
			sourceNative:       utils.RandomAddress(),
			sourceFeeEstimator: sourceFeeEstimator,
			sourceChainID:      th.sourceChainID,
		},
		inFlight: map[[32]byte]InflightReport{},
		offchainConfig: CommitOffchainConfig{
			FeeUpdateHeartBeat:  models.MustMakeDuration(12 * time.Hour),
			InflightCacheExpiry: models.MustMakeDuration(3 * time.Minute),
		},
	}
	return commitTestHarness{
		ccipPluginTestHarness: th,
		plugin:                &plugin,
		mockFeeEstimator:      sourceFeeEstimator,
	}
}

func TestCommitReportSize(t *testing.T) {
	testParams := gopter.DefaultTestParameters()
	testParams.MinSuccessfulTests = 100
	p := gopter.NewProperties(testParams)
	p.Property("bounded commit report size", prop.ForAll(func(root []byte, min, max uint64) bool {
		var root32 [32]byte
		copy(root32[:], root)
		rep, err := EncodeCommitReport(&commit_store.CommitStoreCommitReport{MerkleRoot: root32, Interval: commit_store.CommitStoreInterval{Min: min, Max: max}, PriceUpdates: commit_store.InternalPriceUpdates{
			TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
			DestChainId:       1337,
			UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
		}})
		require.NoError(t, err)
		return len(rep) <= MaxCommitReportLength
	}, gen.SliceOfN(32, gen.UInt8()), gen.UInt64(), gen.UInt64()))
	p.TestingRun(t)
}

func TestCommitReportEncoding(t *testing.T) {
	th := setupCcipTestHarness(t)
	newTokenPrice := big.NewInt(9e18) // $9
	newGasPrice := big.NewInt(2000e9) // $2000 per eth * 1gwei

	// Send a report.
	mctx := hasher.NewKeccakCtx()
	tree, err := merklemulti.NewTree(mctx, [][32]byte{mctx.Hash([]byte{0xaa})})
	require.NoError(t, err)
	report := commit_store.CommitStoreCommitReport{
		PriceUpdates: commit_store.InternalPriceUpdates{
			TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{
				{
					SourceToken: th.feeTokenAddress,
					UsdPerToken: newTokenPrice,
				},
			},
			DestChainId:   th.sourceChainID,
			UsdPerUnitGas: newGasPrice,
		},
		MerkleRoot: tree.Root(),
		Interval:   commit_store.CommitStoreInterval{Min: 1, Max: 10},
	}
	out, err := EncodeCommitReport(&report)
	require.NoError(t, err)
	decodedReport, err := DecodeCommitReport(out)
	require.NoError(t, err)
	require.Equal(t, &report, decodedReport)

	tx, err := th.commitStoreHelper.Report(th.owner, out)
	require.NoError(t, err)
	th.flushLogs()
	res, err := th.client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)

	// Ensure root exists.
	ts, err := th.commitStore.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	require.NotEqual(t, ts.String(), "0")

	// Ensure price update went through
	destChainGasPrice, err := th.priceRegistry.GetDestinationChainGasPrice(nil, th.sourceChainID)
	require.NoError(t, err)
	assert.Equal(t, newGasPrice, destChainGasPrice.Value)

	linkTokenPrice, err := th.priceRegistry.GetTokenPrice(nil, th.feeTokenAddress)
	require.NoError(t, err)
	assert.Equal(t, newTokenPrice, linkTokenPrice.Value)
}

func TestCalculatePriceUpdates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		commitObservations []CommitObservation
		f                  int
		wantGas            *big.Int
	}{
		{"median", []CommitObservation{
			{SourceGasPriceUSD: big.NewInt(1)},
			{SourceGasPriceUSD: big.NewInt(2)},
			{SourceGasPriceUSD: big.NewInt(3)},
			{SourceGasPriceUSD: big.NewInt(4)},
		}, 2, big.NewInt(3)},
		{"insufficient", []CommitObservation{
			{SourceGasPriceUSD: nil}, {SourceGasPriceUSD: nil}, {SourceGasPriceUSD: big.NewInt(3)},
		}, 1, big.NewInt(0)},
		{"median including empties", []CommitObservation{
			{SourceGasPriceUSD: nil}, {SourceGasPriceUSD: big.NewInt(1)}, {SourceGasPriceUSD: big.NewInt(2)},
		}, 1, big.NewInt(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePriceUpdates(10, tt.commitObservations, tt.f)
			assert.Equal(t, tt.wantGas, got.UsdPerUnitGas)
		})
	}
	feeToken1 := common.HexToAddress("0xa")
	feeToken2 := common.HexToAddress("0xb")
	tokenPricesTests := []struct {
		name               string
		commitObservations []CommitObservation
		f                  int
		wantedUpdates      []commit_store.InternalTokenPriceUpdate
	}{
		{"median one token", []CommitObservation{
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(10)}},
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(12)}},
		}, 1, []commit_store.InternalTokenPriceUpdate{
			{SourceToken: feeToken1, UsdPerToken: big.NewInt(12)},
		}},
		{
			"median two tokens", []CommitObservation{
				{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(10), feeToken2: big.NewInt(13)}},
				{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(12), feeToken2: big.NewInt(7)}},
			}, 1, []commit_store.InternalTokenPriceUpdate{
				{SourceToken: feeToken1, UsdPerToken: big.NewInt(12)},
				{SourceToken: feeToken2, UsdPerToken: big.NewInt(13)},
			},
		},
		{
			"only one token with enough votes", []CommitObservation{
				{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(10)}},
				{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(12), feeToken2: big.NewInt(7)}},
			}, 1, []commit_store.InternalTokenPriceUpdate{
				{SourceToken: feeToken1, UsdPerToken: big.NewInt(12)},
			},
		},
	}
	for _, tt := range tokenPricesTests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePriceUpdates(10, tt.commitObservations, tt.f)
			assert.Equal(t, tt.wantedUpdates, got.TokenPriceUpdates)
		})
	}
}

func TestCalculateIntervalConsensus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		intervals []commit_store.CommitStoreInterval
		f         int
		wantMin   uint64
		wantMax   uint64
		wantErr   bool
	}{
		{"no obs", []commit_store.CommitStoreInterval{{Min: 0, Max: 0}}, 0, 0, 0, false},
		{"basic", []commit_store.CommitStoreInterval{
			{Min: 9, Max: 14},
			{Min: 10, Max: 12},
			{Min: 10, Max: 14},
		}, 1, 10, 14, false},
		{"not enough intervals", []commit_store.CommitStoreInterval{}, 1, 0, 0, true},
		{"min > max", []commit_store.CommitStoreInterval{
			{Min: 9, Max: 4},
			{Min: 10, Max: 4},
			{Min: 10, Max: 6},
		}, 1, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateIntervalConsensus(tt.intervals, tt.f)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.wantMin, got.Min)
			assert.Equal(t, tt.wantMax, got.Max)
		})
	}
}

func TestCommitReportToEthTxMeta(t *testing.T) {
	mctx := hasher.NewKeccakCtx()
	tree, err := merklemulti.NewTree(mctx, [][32]byte{mctx.Hash([]byte{0xaa})})
	require.NoError(t, err)

	tests := []struct {
		name          string
		min, max      uint64
		expectedRange []uint64
	}{
		{
			"happy flow",
			1, 10,
			[]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"same sequence",
			1, 1,
			[]uint64{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			report := commit_store.CommitStoreCommitReport{
				PriceUpdates: commit_store.InternalPriceUpdates{
					TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
					DestChainId:       uint64(1337),
					UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
				},
				MerkleRoot: tree.Root(),
				Interval:   commit_store.CommitStoreInterval{Min: tc.min, Max: tc.max},
			}
			out, err := EncodeCommitReport(&report)
			require.NoError(t, err)

			txMeta, err := CommitReportToEthTxMeta(out)
			require.NoError(t, err)
			require.NotNil(t, txMeta)
			require.EqualValues(t, tc.expectedRange, txMeta.SeqNumbers)
		})
	}
}

func TestGeneratePriceUpdates(t *testing.T) {
	th := setupCommitTestHarness(t)

	gasPrice := big.NewInt(3e9)
	mockedGetFee := th.mockFeeEstimator.On(
		"GetFee",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(gas.EvmFee{Legacy: assets.NewWei(gasPrice)}, uint32(150e3), nil)

	fakePrices, err := th.plugin.config.priceGetter.TokenPricesUSD(testutils.Context(t), []common.Address{th.feeTokenAddress}) // 2e20 hardcoded in fakePriceGetter
	require.NoError(t, err)
	fakePrice := fakePrices[th.feeTokenAddress]

	expectedGasPrice := big.NewInt(0).Mul(fakePrice, gasPrice)
	expectedGasPrice.Div(expectedGasPrice, big.NewInt(1e18))

	newGasPrice := big.NewInt(106) // +6%, just outside the default deviation margin of ±5%
	newGasPrice.Mul(newGasPrice, fakePrice)
	newGasPrice.Div(newGasPrice, big.NewInt(100))

	newExpectedGasPriceUSD := big.NewInt(0).Mul(newGasPrice, fakePrice)
	newExpectedGasPriceUSD.Div(newExpectedGasPriceUSD, big.NewInt(1e18))

	newFeeToken := testutils.NewAddress()

	tests := []struct {
		name                   string
		addFeeTokens           []common.Address
		updateTokenPricesUSD   map[common.Address]*big.Int
		updateGasPriceUSD      *big.Int
		updateGasPrice         *big.Int
		expectedGasPriceUSD    *big.Int
		expectedTokenPricesUSD map[common.Address]*big.Int
	}{
		{
			name:                   "first update",
			expectedGasPriceUSD:    expectedGasPrice,
			expectedTokenPricesUSD: map[common.Address]*big.Int{th.feeTokenAddress: fakePrice},
		},
		{
			name:                   "gasPrice up-to-date",
			updateGasPriceUSD:      expectedGasPrice,
			expectedGasPriceUSD:    nil,
			expectedTokenPricesUSD: map[common.Address]*big.Int{th.feeTokenAddress: fakePrice},
		},
		{
			name:                   "tokenPrice up-to-date, gasPrice deviated",
			updateTokenPricesUSD:   map[common.Address]*big.Int{th.feeTokenAddress: fakePrice},
			updateGasPrice:         newGasPrice,
			expectedGasPriceUSD:    newExpectedGasPriceUSD,
			expectedTokenPricesUSD: map[common.Address]*big.Int{},
		},
		{
			name:                   "new feeToken, getLatestPriceUpdates returns nil",
			addFeeTokens:           []common.Address{newFeeToken},
			expectedGasPriceUSD:    newExpectedGasPriceUSD,
			expectedTokenPricesUSD: map[common.Address]*big.Int{newFeeToken: fakePrice},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.addFeeTokens) > 0 {
				_, err = th.priceRegistry.ApplyFeeTokensUpdates(th.owner, tt.addFeeTokens, []common.Address{})
				require.NoError(t, err)
				th.flushLogs()
			}
			if len(tt.updateTokenPricesUSD) > 0 || tt.updateGasPriceUSD != nil {
				destChainID := uint64(0)
				if tt.updateGasPriceUSD != nil {
					destChainID = th.sourceChainID
				} else {
					tt.updateGasPriceUSD = big.NewInt(0)
				}
				tokenPriceUpdates := []price_registry.InternalTokenPriceUpdate{}
				if len(tt.updateTokenPricesUSD) > 0 {
					for token, value := range tt.updateTokenPricesUSD {
						tokenPriceUpdates = append(tokenPriceUpdates, price_registry.InternalTokenPriceUpdate{SourceToken: token, UsdPerToken: value})
					}
				}
				// update gasPrice in priceRegistry
				_, err = th.priceRegistry.UpdatePrices(th.owner, price_registry.InternalPriceUpdates{
					TokenPriceUpdates: tokenPriceUpdates,
					DestChainId:       destChainID,
					UsdPerUnitGas:     tt.updateGasPriceUSD,
				})
				require.NoError(t, err)
				th.flushLogs()
			}
			if tt.updateGasPrice != nil {
				mockedGetFee = mockedGetFee.Return(gas.EvmFee{Legacy: assets.NewWei(tt.updateGasPrice)}, uint32(200e3), nil)
			}

			gotGasPriceUSD, gotTokenPricesUSD, err := th.plugin.generatePriceUpdates(testutils.Context(t), time.Unix(int64((i+1)*10), 0))

			require.NoError(t, err)
			assert.Equal(t, tt.expectedGasPriceUSD, gotGasPriceUSD)
			assert.Equal(t, tt.expectedTokenPricesUSD, gotTokenPricesUSD)
		})
	}
}

func TestShouldTransmitAcceptedReport(t *testing.T) {
	th := setupCommitTestHarness(t)

	nextMinSeqNr := uint64(10)
	_, err := th.commitStore.SetMinSeqNr(th.owner, nextMinSeqNr)
	require.NoError(t, err)
	th.flushLogs()

	tests := []struct {
		name            string
		commitStoreDown bool
		seq             uint64
		expected        bool
	}{
		{"base", false, nextMinSeqNr, true},
		{"future", false, nextMinSeqNr * 2, true},
		{"empty", false, 0, true}, // empty should have priceUpdates
		{"stale", false, nextMinSeqNr - 1, false},
		{"commitStore down", true, nextMinSeqNr, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.commitStoreDown && !isCommitStoreDownNow(testutils.Context(t), th.lggr, th.commitStore) {
				_, err := th.commitStore.Pause(th.owner)
				require.NoError(t, err)
			} else if !tt.commitStoreDown && isCommitStoreDownNow(testutils.Context(t), th.lggr, th.commitStore) {
				_, err := th.commitStore.Unpause(th.owner)
				require.NoError(t, err)
			}
			th.flushLogs()

			var root [32]byte
			if tt.seq > 0 {
				root = testutils.Random32Byte()
			}

			report, err := EncodeCommitReport(&commit_store.CommitStoreCommitReport{
				PriceUpdates: commit_store.InternalPriceUpdates{
					TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
					DestChainId:       0,
					UsdPerUnitGas:     new(big.Int),
				},
				MerkleRoot: root,
				Interval:   commit_store.CommitStoreInterval{Min: tt.seq, Max: tt.seq},
			})
			require.NoError(t, err)

			got, err := th.plugin.ShouldTransmitAcceptedReport(testutils.Context(t), types.ReportTimestamp{}, report)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestShouldAcceptFinalizedReport(t *testing.T) {
	th := setupCommitTestHarness(t)

	nextMinSeqNr := uint64(10)
	_, err := th.commitStore.SetMinSeqNr(th.owner, nextMinSeqNr)
	require.NoError(t, err)
	th.flushLogs()

	tests := []struct {
		name     string
		seq      uint64
		expected bool
		err      bool
	}{
		{"future", nextMinSeqNr * 2, false, true},
		{"empty", 0, false, false},
		{"stale", nextMinSeqNr - 1, false, true},
		{"base", nextMinSeqNr, true, false},          // accepted is not side-effects-free: it adds to inFlight cache
		{"base inFlight", nextMinSeqNr, false, true}, // this one is like future, but caused by previous test added to inFlight
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root [32]byte
			if tt.seq > 0 {
				root = testutils.Random32Byte()
			}

			report, err := EncodeCommitReport(&commit_store.CommitStoreCommitReport{
				PriceUpdates: commit_store.InternalPriceUpdates{
					TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
					DestChainId:       0,
					UsdPerUnitGas:     new(big.Int),
				},
				MerkleRoot: root,
				Interval:   commit_store.CommitStoreInterval{Min: tt.seq, Max: tt.seq},
			})
			require.NoError(t, err)

			got, err := th.plugin.ShouldAcceptFinalizedReport(testutils.Context(t), types.ReportTimestamp{}, report)
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}

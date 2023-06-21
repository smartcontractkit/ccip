package ccip

import (
	"context"
	"encoding/json"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	mock_contracts "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cache"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"

	plugintesthelpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers/plugins"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var defaultGasPrice = big.NewInt(3e9)

type commitTestHarness = struct {
	plugintesthelpers.CCIPPluginTestHarness
	plugin       *CommitReportingPlugin
	mockedGetFee *mock.Call
}

func setupCommitTestHarness(t *testing.T) commitTestHarness {
	th := plugintesthelpers.SetupCCIPTestHarness(t)

	sourceFeeEstimator := mocks.NewEvmFeeEstimator(t)

	mockedGetFee := sourceFeeEstimator.On(
		"GetFee",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Maybe().Return(gas.EvmFee{Legacy: assets.NewWei(defaultGasPrice)}, uint32(200e3), nil)

	backendClient := client.NewSimulatedBackendClient(t, th.Dest.Chain, new(big.Int).SetUint64(th.Dest.ChainID))
	plugin := CommitReportingPlugin{
		config: CommitPluginConfig{
			lggr:                th.Lggr,
			sourceLP:            th.SourceLP,
			destLP:              th.DestLP,
			offRamp:             th.Dest.OffRamp,
			onRampAddress:       th.Source.OnRamp.Address(),
			commitStore:         th.Dest.CommitStore,
			priceGetter:         fakePriceGetter{},
			sourceNative:        utils.RandomAddress(),
			sourceFeeEstimator:  sourceFeeEstimator,
			sourceChainSelector: th.Source.ChainID,
			destClient:          backendClient,
			leafHasher:          hasher.NewLeafHasher(th.Source.ChainID, th.Dest.ChainID, th.Source.OnRamp.Address(), hasher.NewKeccakCtx()),
			getSeqNumFromLog:    getSeqNumFromLog(th.Source.OnRamp),
		},
		inflightReports: newInflightCommitReportsContainer(time.Hour),
		onchainConfig:   th.CommitOnchainConfig,
		offchainConfig: ccipconfig.CommitOffchainConfig{
			SourceFinalityDepth:   0,
			DestFinalityDepth:     0,
			FeeUpdateDeviationPPB: 5e7,
			FeeUpdateHeartBeat:    models.MustMakeDuration(12 * time.Hour),
			MaxGasPrice:           200e9,
		},
		lggr:                  th.Lggr,
		destPriceRegistry:     th.Dest.PriceRegistry,
		tokenToDecimalMapping: cache.NewTokenToDecimals(th.Lggr, th.DestLP, th.Dest.OffRamp, th.Dest.PriceRegistry, backendClient, 0),
	}
	return commitTestHarness{
		CCIPPluginTestHarness: th,
		plugin:                &plugin,
		mockedGetFee:          mockedGetFee,
	}
}

func TestCommitReportSize(t *testing.T) {
	testParams := gopter.DefaultTestParameters()
	testParams.MinSuccessfulTests = 100
	p := gopter.NewProperties(testParams)
	p.Property("bounded commit report size", prop.ForAll(func(root []byte, min, max uint64) bool {
		var root32 [32]byte
		copy(root32[:], root)
		rep, err := abihelpers.EncodeCommitReport(commit_store.CommitStoreCommitReport{
			MerkleRoot: root32,
			Interval:   commit_store.CommitStoreInterval{Min: min, Max: max},
			PriceUpdates: commit_store.InternalPriceUpdates{
				TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
				DestChainSelector: 1337,
				UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
			},
		})
		require.NoError(t, err)
		return len(rep) <= MaxCommitReportLength
	}, gen.SliceOfN(32, gen.UInt8()), gen.UInt64(), gen.UInt64()))
	p.TestingRun(t)
}

func TestCommitReportEncoding(t *testing.T) {
	th := plugintesthelpers.SetupCCIPTestHarness(t)
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
					SourceToken: th.Dest.LinkToken.Address(),
					UsdPerToken: newTokenPrice,
				},
			},
			DestChainSelector: th.Source.ChainID,
			UsdPerUnitGas:     newGasPrice,
		},
		MerkleRoot: tree.Root(),
		Interval:   commit_store.CommitStoreInterval{Min: 1, Max: 10},
	}
	out, err := abihelpers.EncodeCommitReport(report)
	require.NoError(t, err)
	decodedReport, err := abihelpers.DecodeCommitReport(out)
	require.NoError(t, err)
	require.Equal(t, report, decodedReport)

	latestEpocAndRound, err := th.Dest.CommitStoreHelper.GetLatestPriceEpochAndRound(nil)
	require.NoError(t, err)

	tx, err := th.Dest.CommitStoreHelper.Report(th.Dest.User, out, big.NewInt(int64(latestEpocAndRound+1)))
	require.NoError(t, err)
	th.CommitAndPollLogs(t)
	res, err := th.Dest.Chain.TransactionReceipt(testutils.Context(t), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)

	// Ensure root exists.
	ts, err := th.Dest.CommitStore.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	require.NotEqual(t, ts.String(), "0")

	// Ensure price update went through
	destChainGasPrice, err := th.Dest.PriceRegistry.GetDestinationChainGasPrice(nil, th.Source.ChainID)
	require.NoError(t, err)
	assert.Equal(t, newGasPrice, destChainGasPrice.Value)

	linkTokenPrice, err := th.Dest.PriceRegistry.GetTokenPrice(nil, th.Dest.LinkToken.Address())
	require.NoError(t, err)
	assert.Equal(t, newTokenPrice, linkTokenPrice.Value)
}

func TestCommitObservation(t *testing.T) {
	th := setupCommitTestHarness(t)
	th.plugin.F = 1

	mb := th.GenerateAndSendMessageBatch(t, 1, 0, 0)

	tests := []struct {
		name            string
		commitStoreDown bool
		expected        *CommitObservation
		expectedError   bool
	}{
		{
			"base",
			false,
			&CommitObservation{
				Interval:          mb.Interval,
				SourceGasPriceUSD: new(big.Int).Mul(defaultGasPrice, big.NewInt(200)),
				TokenPricesUSD: map[common.Address]*big.Int{
					th.Dest.LinkToken.Address(): new(big.Int).Mul(big.NewInt(200), big.NewInt(1e18)),
				},
			},
			false,
		},
		{
			"commitStore down",
			true,
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.commitStoreDown && !isCommitStoreDownNow(testutils.Context(t), th.Lggr, th.Dest.CommitStore) {
				_, err := th.Dest.CommitStore.Pause(th.Dest.User)
				require.NoError(t, err)
				th.CommitAndPollLogs(t)
			} else if !tt.commitStoreDown && isCommitStoreDownNow(testutils.Context(t), th.Lggr, th.Dest.CommitStore) {
				_, err := th.Dest.CommitStore.Unpause(th.Dest.User)
				require.NoError(t, err)
				th.CommitAndPollLogs(t)
			}

			gotObs, err := th.plugin.Observation(testutils.Context(t), types.ReportTimestamp{}, types.Query{})

			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			var decodedObservation *CommitObservation
			if gotObs != nil {
				decodedObservation = new(CommitObservation)
				err = json.Unmarshal(gotObs, decodedObservation)
				require.NoError(t, err)

			}
			assert.Equal(t, tt.expected, decodedObservation)
		})
	}
}

func TestCommitReport(t *testing.T) {
	th := setupCommitTestHarness(t)
	th.plugin.F = 1

	mb := th.GenerateAndSendMessageBatch(t, 1, 0, 0)

	tests := []struct {
		name          string
		observations  []CommitObservation
		shouldReport  bool
		commitReport  *commit_store.CommitStoreCommitReport
		expectedError bool
	}{
		{
			"base",
			[]CommitObservation{
				{Interval: commit_store.CommitStoreInterval{Min: 1, Max: 1}},
				{Interval: commit_store.CommitStoreInterval{Min: 1, Max: 1}},
			},
			true,
			&commit_store.CommitStoreCommitReport{
				MerkleRoot: mb.Root,
				Interval:   commit_store.CommitStoreInterval{Min: 1, Max: 1},
				PriceUpdates: commit_store.InternalPriceUpdates{
					TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
					DestChainSelector: 0,
					UsdPerUnitGas:     new(big.Int),
				},
			},
			false,
		},
		{
			"not enough observations",
			[]CommitObservation{
				{Interval: commit_store.CommitStoreInterval{Min: 1, Max: 1}},
			},
			false,
			nil,
			true,
		},
		{
			"empty",
			[]CommitObservation{
				{Interval: commit_store.CommitStoreInterval{Min: 0, Max: 0}},
				{Interval: commit_store.CommitStoreInterval{Min: 0, Max: 0}},
			},
			false,
			nil,
			false,
		},
		{
			"no leaves",
			[]CommitObservation{
				{Interval: commit_store.CommitStoreInterval{Min: 2, Max: 2}},
				{Interval: commit_store.CommitStoreInterval{Min: 2, Max: 2}},
			},
			false,
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aos := make([]types.AttributedObservation, 0, len(tt.observations))
			for _, o := range tt.observations {
				obs, err := o.Marshal()
				require.NoError(t, err)
				aos = append(aos, types.AttributedObservation{Observation: obs})
			}
			gotShouldReport, gotReport, err := th.plugin.Report(testutils.Context(t), types.ReportTimestamp{}, types.Query{}, aos)

			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.shouldReport, gotShouldReport)

			var expectedReport types.Report
			if tt.commitReport != nil {
				expectedReport, err = abihelpers.EncodeCommitReport(*tt.commitReport)
				require.NoError(t, err)
			}
			assert.Equal(t, expectedReport, gotReport)
		})
	}
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
		name       string
		intervals  []commit_store.CommitStoreInterval
		rangeLimit uint64
		f          int
		wantMin    uint64
		wantMax    uint64
		wantErr    bool
	}{
		{"no obs", []commit_store.CommitStoreInterval{{Min: 0, Max: 0}}, 0, 0, 0, 0, false},
		{"basic", []commit_store.CommitStoreInterval{
			{Min: 9, Max: 14},
			{Min: 10, Max: 12},
			{Min: 10, Max: 14},
		}, 0, 1, 10, 14, false},
		{"not enough intervals", []commit_store.CommitStoreInterval{}, 0, 1, 0, 0, true},
		{"min > max", []commit_store.CommitStoreInterval{
			{Min: 9, Max: 4},
			{Min: 10, Max: 4},
			{Min: 10, Max: 6},
		}, 0, 1, 0, 0, true},
		{
			"range limit", []commit_store.CommitStoreInterval{
				{Min: 10, Max: 100},
				{Min: 1, Max: 1000},
			}, 256, 1, 10, 265, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateIntervalConsensus(tt.intervals, tt.f, tt.rangeLimit)
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

func TestGeneratePriceUpdates(t *testing.T) {
	th := setupCommitTestHarness(t)

	fakePrices, err := th.plugin.config.priceGetter.TokenPricesUSD(testutils.Context(t), []common.Address{th.Dest.LinkToken.Address()}) // 2e20 hardcoded in fakePriceGetter
	require.NoError(t, err)
	fakePrice := fakePrices[th.Dest.LinkToken.Address()]

	expectedGasPrice := new(big.Int).Mul(fakePrice, defaultGasPrice)
	expectedGasPrice.Div(expectedGasPrice, big.NewInt(1e18))

	newGasPrice := big.NewInt(106) // +6%, just outside the default deviation margin of ±5%
	newGasPrice.Mul(newGasPrice, fakePrice)
	newGasPrice.Div(newGasPrice, big.NewInt(100))

	newExpectedGasPriceUSD := new(big.Int).Mul(newGasPrice, fakePrice)
	newExpectedGasPriceUSD.Div(newExpectedGasPriceUSD, big.NewInt(1e18))

	newFeeToken, _, _, err := link_token_interface.DeployLinkToken(th.Dest.User, th.Dest.Chain)
	require.NoError(t, err)

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
			expectedTokenPricesUSD: map[common.Address]*big.Int{th.Dest.LinkToken.Address(): fakePrice},
		},
		{
			name:                   "gasPrice up-to-date",
			updateGasPriceUSD:      expectedGasPrice,
			expectedGasPriceUSD:    nil,
			expectedTokenPricesUSD: map[common.Address]*big.Int{th.Dest.LinkToken.Address(): fakePrice},
		},
		{
			name:                   "tokenPrice up-to-date, gasPrice deviated",
			updateTokenPricesUSD:   map[common.Address]*big.Int{th.Dest.LinkToken.Address(): fakePrice},
			updateGasPrice:         newGasPrice,
			expectedGasPriceUSD:    newExpectedGasPriceUSD,
			expectedTokenPricesUSD: map[common.Address]*big.Int{},
		},
		{
			name:                   "new feeToken, getLatestTokenPriceUpdates returns nil",
			addFeeTokens:           []common.Address{newFeeToken},
			expectedGasPriceUSD:    newExpectedGasPriceUSD,
			expectedTokenPricesUSD: map[common.Address]*big.Int{newFeeToken: fakePrice},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.addFeeTokens) > 0 {
				_, err = th.Dest.PriceRegistry.ApplyFeeTokensUpdates(th.Dest.User, tt.addFeeTokens, []common.Address{})
				require.NoError(t, err)
				th.CommitAndPollLogs(t)
			}
			if len(tt.updateTokenPricesUSD) > 0 || tt.updateGasPriceUSD != nil {
				destChainID := uint64(0)
				if tt.updateGasPriceUSD != nil {
					destChainID = th.Source.ChainID
				} else {
					tt.updateGasPriceUSD = big.NewInt(0)
				}
				tokenPriceUpdates := []price_registry.InternalTokenPriceUpdate{}
				if len(tt.updateTokenPricesUSD) > 0 {
					for token, value := range tt.updateTokenPricesUSD {
						tokenPriceUpdates = append(tokenPriceUpdates, price_registry.InternalTokenPriceUpdate{SourceToken: token, UsdPerToken: value})
					}
				}
				// update gasPrice in destPriceRegistry
				_, err = th.Dest.PriceRegistry.UpdatePrices(th.Dest.User, price_registry.InternalPriceUpdates{
					TokenPriceUpdates: tokenPriceUpdates,
					DestChainSelector: destChainID,
					UsdPerUnitGas:     tt.updateGasPriceUSD,
				})
				require.NoError(t, err)
				th.CommitAndPollLogs(t)
			}
			if tt.updateGasPrice != nil {
				th.mockedGetFee.Return(gas.EvmFee{Legacy: assets.NewWei(tt.updateGasPrice)}, uint32(200e3), nil)
			}

			gotGasPriceUSD, gotTokenPricesUSD, err := th.plugin.generatePriceUpdates(testutils.Context(t), th.Lggr, time.Unix(int64((i+1)*10), 0))

			require.NoError(t, err)
			assert.Equal(t, tt.expectedGasPriceUSD, gotGasPriceUSD)
			assert.Equal(t, tt.expectedTokenPricesUSD, gotTokenPricesUSD)
		})
	}
}

func TestUpdateTokenToDecimalMapping(t *testing.T) {
	th := plugintesthelpers.SetupCCIPTestHarness(t)

	destToken, _, _, err := link_token_interface.DeployLinkToken(th.Dest.User, th.Dest.Chain)
	require.NoError(t, err)

	feeToken, _, _, err := link_token_interface.DeployLinkToken(th.Dest.User, th.Dest.Chain)
	require.NoError(t, err)
	th.CommitAndPollLogs(t)

	tokens := []common.Address{}
	tokens = append(tokens, destToken)
	tokens = append(tokens, feeToken)

	mockOffRamp := &mock_contracts.EVM2EVMOffRampInterface{}
	mockOffRamp.On("GetDestinationTokens", mock.Anything).Return([]common.Address{destToken}, nil)
	mockOffRamp.On("Address").Return(common.Address{})

	mockPriceRegistry := &mock_contracts.PriceRegistryInterface{}
	mockPriceRegistry.On("GetFeeTokens", mock.Anything).Return([]common.Address{feeToken}, nil)
	mockPriceRegistry.On("Address").Return(common.Address{})

	backendClient := client.NewSimulatedBackendClient(t, th.Dest.Chain, new(big.Int).SetUint64(th.Dest.ChainID))
	plugin := CommitReportingPlugin{
		config: CommitPluginConfig{
			offRamp:    mockOffRamp,
			destClient: backendClient,
		},
		destPriceRegistry:     mockPriceRegistry,
		tokenToDecimalMapping: cache.NewTokenToDecimals(th.Lggr, th.DestLP, mockOffRamp, mockPriceRegistry, backendClient, 0),
	}

	tokenMapping, err := plugin.tokenToDecimalMapping.Get(testutils.Context(t))
	require.NoError(t, err)
	assert.Equal(t, len(tokens), len(tokenMapping))
	assert.Equal(t, uint8(18), tokenMapping[destToken])
	assert.Equal(t, uint8(18), tokenMapping[feeToken])
}

func TestCalculateUsdPer1e18TokenAmount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		price      *big.Int
		decimal    uint8
		wantResult *big.Int
	}{
		{
			name:       "18-decimal token, $6.5 per token",
			price:      big.NewInt(65e17),
			decimal:    18,
			wantResult: big.NewInt(65e17),
		},
		{
			name:       "6-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    6,
			wantResult: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e12)), // 1e30
		},
		{
			name:       "0-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    0,
			wantResult: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e18)), // 1e36
		},
		{
			name:       "36-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    36,
			wantResult: big.NewInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateUsdPer1e18TokenAmount(tt.price, tt.decimal)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestShouldTransmitAcceptedReport(t *testing.T) {
	th := setupCommitTestHarness(t)
	tokenPrice := big.NewInt(9e18) // $9
	gasPrice := big.NewInt(1500e9) // $1500 per eth * 1gwei

	nextMinSeqNr := uint64(10)
	_, err := th.Dest.CommitStore.SetMinSeqNr(th.Dest.User, nextMinSeqNr)
	require.NoError(t, err)
	_, err = th.Dest.PriceRegistry.UpdatePrices(th.Dest.User, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{SourceToken: th.Dest.LinkToken.Address(), UsdPerToken: tokenPrice},
		},
		DestChainSelector: th.Source.ChainID,
		UsdPerUnitGas:     gasPrice,
	})
	require.NoError(t, err)
	th.CommitAndPollLogs(t)
	round := uint8(1)

	tests := []struct {
		name       string
		seq        uint64
		gasPrice   *big.Int
		tokenPrice *big.Int
		expected   bool
	}{
		{"base", nextMinSeqNr, nil, nil, true},
		{"future", nextMinSeqNr + 10, nil, nil, true},
		{"empty", 0, nil, nil, false},
		{"gasPrice update", 0, big.NewInt(10), nil, true},
		{"gasPrice stale", 0, gasPrice, nil, false},
		{"tokenPrice update", 0, nil, big.NewInt(20), true},
		{"tokenPrice stale", 0, nil, tokenPrice, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var destChainID uint64
			gasPrice := new(big.Int)
			if tt.gasPrice != nil {
				destChainID = th.Source.ChainID
				gasPrice = tt.gasPrice
			}

			var tokenPrices []commit_store.InternalTokenPriceUpdate
			if tt.tokenPrice != nil {
				tokenPrices = []commit_store.InternalTokenPriceUpdate{
					{SourceToken: th.Dest.LinkToken.Address(), UsdPerToken: tt.tokenPrice},
				}
			} else {
				tokenPrices = []commit_store.InternalTokenPriceUpdate{}
			}

			var root [32]byte
			if tt.seq > 0 {
				root = testutils.Random32Byte()
			}

			report, err := abihelpers.EncodeCommitReport(commit_store.CommitStoreCommitReport{
				PriceUpdates: commit_store.InternalPriceUpdates{
					TokenPriceUpdates: tokenPrices,
					DestChainSelector: destChainID,
					UsdPerUnitGas:     gasPrice,
				},
				MerkleRoot: root,
				Interval:   commit_store.CommitStoreInterval{Min: tt.seq, Max: tt.seq},
			})
			require.NoError(t, err)

			got, err := th.plugin.ShouldTransmitAcceptedReport(testutils.Context(t), types.ReportTimestamp{Epoch: 1, Round: round}, report)
			round++
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestShouldAcceptFinalizedReport(t *testing.T) {
	nextMinSeqNr := uint64(10)

	tests := []struct {
		name                     string
		seq                      uint64
		latestPriceEpochAndRound int64
		epoch                    uint32
		round                    uint8
		destChainSelector        int
		skipRoot                 bool
		expected                 bool
		err                      bool
	}{
		{
			name:  "future",
			seq:   nextMinSeqNr * 2,
			epoch: 1,
			round: 1,
		},
		{
			name:  "empty",
			epoch: 1,
			round: 2,
		},
		{
			name:  "stale",
			seq:   nextMinSeqNr - 1,
			epoch: 1,
			round: 3,
		},
		{
			name:     "base",
			seq:      nextMinSeqNr,
			epoch:    1,
			round:    4,
			expected: true,
		},
		{
			name:                     "price update - epoch and round is ok",
			seq:                      nextMinSeqNr,
			latestPriceEpochAndRound: int64(mergeEpochAndRound(2, 10)),
			epoch:                    2,
			round:                    11,
			destChainSelector:        rand.Int(),
			skipRoot:                 true,
			expected:                 true,
		},
		{
			name:                     "price update - epoch and round is behind",
			seq:                      nextMinSeqNr,
			latestPriceEpochAndRound: int64(mergeEpochAndRound(2, 10)),
			epoch:                    2,
			round:                    9,
			destChainSelector:        rand.Int(),
			skipRoot:                 true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := setupCommitTestHarness(t)
			_, err := th.Dest.CommitStore.SetMinSeqNr(th.Dest.User, nextMinSeqNr)
			require.NoError(t, err)

			_, err = th.Dest.CommitStoreHelper.SetLatestPriceEpochAndRound(th.Dest.User, big.NewInt(tt.latestPriceEpochAndRound))
			require.NoError(t, err)

			th.CommitAndPollLogs(t)

			var root [32]byte
			if tt.seq > 0 {
				root = testutils.Random32Byte()
			}

			r := commit_store.CommitStoreCommitReport{
				PriceUpdates: commit_store.InternalPriceUpdates{
					TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
					DestChainSelector: uint64(tt.destChainSelector),
					UsdPerUnitGas:     new(big.Int),
				},
				Interval: commit_store.CommitStoreInterval{Min: tt.seq, Max: tt.seq},
			}
			if !tt.skipRoot {
				r.MerkleRoot = root
			}
			report, err := abihelpers.EncodeCommitReport(r)
			require.NoError(t, err)

			got, err := th.plugin.ShouldAcceptFinalizedReport(
				testutils.Context(t),
				types.ReportTimestamp{Epoch: tt.epoch, Round: tt.round}, report)
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.expected, got)

			if got { // already added to inflight, should not be accepted again
				got, err = th.plugin.ShouldAcceptFinalizedReport(
					testutils.Context(t),
					types.ReportTimestamp{Epoch: tt.epoch, Round: tt.round}, report)
				require.NoError(t, err)
				assert.False(t, got)
			}
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
					DestChainSelector: uint64(1337),
					UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
				},
				MerkleRoot: tree.Root(),
				Interval:   commit_store.CommitStoreInterval{Min: tc.min, Max: tc.max},
			}
			out, err := abihelpers.EncodeCommitReport(report)
			require.NoError(t, err)

			txMeta, err := CommitReportToEthTxMeta(out)
			require.NoError(t, err)
			require.NotNil(t, txMeta)
			require.EqualValues(t, tc.expectedRange, txMeta.SeqNumbers)
		})
	}
}

func TestNextMin(t *testing.T) {
	lggr := logger.TestLogger(t)
	commitStore := mock_contracts.CommitStoreInterface{}
	cp := CommitReportingPlugin{config: CommitPluginConfig{commitStore: &commitStore}, inflightReports: newInflightCommitReportsContainer(time.Hour)}
	root1 := utils.Keccak256Fixed(hexutil.MustDecode("0xaa"))
	var tt = []struct {
		onChainMin          uint64
		inflight            []commit_store.CommitStoreCommitReport
		expectedOnChainMin  uint64
		expectedInflightMin uint64
	}{
		{
			onChainMin:          uint64(1),
			inflight:            nil,
			expectedInflightMin: uint64(1),
			expectedOnChainMin:  uint64(1),
		},
		{
			onChainMin: uint64(1),
			inflight: []commit_store.CommitStoreCommitReport{
				{Interval: commit_store.CommitStoreInterval{Min: uint64(1), Max: uint64(2)}, MerkleRoot: root1}},
			expectedInflightMin: uint64(3),
			expectedOnChainMin:  uint64(1),
		},
		{
			onChainMin: uint64(1),
			inflight: []commit_store.CommitStoreCommitReport{
				{Interval: commit_store.CommitStoreInterval{Min: uint64(3), Max: uint64(4)}, MerkleRoot: root1}},
			expectedInflightMin: uint64(5),
			expectedOnChainMin:  uint64(1),
		},
		{
			onChainMin: uint64(1),
			inflight: []commit_store.CommitStoreCommitReport{
				{Interval: commit_store.CommitStoreInterval{Min: uint64(1), Max: uint64(MaxInflightSeqNumGap + 2)}, MerkleRoot: root1}},
			expectedInflightMin: uint64(1),
			expectedOnChainMin:  uint64(1),
		},
	}
	for _, tc := range tt {
		commitStore.On("GetExpectedNextSequenceNumber", mock.Anything).Return(tc.onChainMin, nil)
		epochAndRound := uint64(1)
		for _, rep := range tc.inflight {
			rc := rep
			require.NoError(t, cp.inflightReports.add(lggr, rc, epochAndRound))
			epochAndRound++
		}
		t.Log("inflight", cp.inflightReports.maxInflightSeqNr())
		inflightMin, onchainMin, err := cp.nextMinSeqNum(context.Background(), lggr)
		require.NoError(t, err)
		assert.Equal(t, tc.expectedInflightMin, inflightMin)
		assert.Equal(t, tc.expectedOnChainMin, onchainMin)
		cp.inflightReports.reset(lggr)
	}
}

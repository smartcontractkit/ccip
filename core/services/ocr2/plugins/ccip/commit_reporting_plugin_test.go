package ccip

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/common/txmgr/types/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store_helper"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

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
	// Set up a user.
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destChainId := uint64(1337)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(int64(destChainId)))
	require.NoError(t, err)
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18))}},
		ethconfig.Defaults.Miner.GasCeil)

	// Deploy link token.
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)

	// Deploy link token pool.
	destPoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	_, err = lock_release_token_pool.NewLockReleaseTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)

	// Deploy AFN.
	afnAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		destUser,
		destChain,
	)
	require.NoError(t, err)

	priceRegistry, _, _, err := price_registry.DeployPriceRegistry(destUser, destChain, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: nil,
		DestChainId:       0,
		UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
	}, []common.Address{}, []common.Address{}, uint32(time.Hour.Seconds()))
	require.NoError(t, err)

	// Deploy commitStore.
	onRampAddress := common.HexToAddress("0x01BE23585060835E02B77ef475b0Cc51aA1e0709")
	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		destUser,  // user
		destChain, // client
		commit_store_helper.CommitStoreStaticConfig{
			ChainId:       destChainId,
			SourceChainId: 1337,
			OnRamp:        onRampAddress,
		},
		commit_store_helper.CommitStoreDynamicConfig{
			PriceRegistry: priceRegistry,
			Afn:           afnAddress, // AFN address
		},
	)
	require.NoError(t, err)
	commitStore, err := commit_store_helper.NewCommitStoreHelper(commitStoreAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	newPriceRegistry, err := price_registry.NewPriceRegistry(priceRegistry, destChain)
	require.NoError(t, err)

	_, err = newPriceRegistry.ApplyPriceUpdatersUpdates(destUser, []common.Address{commitStoreAddress}, []common.Address{})
	require.NoError(t, err)
	destChain.Commit()

	// Send a report.
	mctx := hasher.NewKeccakCtx()
	tree, err := merklemulti.NewTree(mctx, [][32]byte{mctx.Hash([]byte{0xaa})})
	require.NoError(t, err)
	report := commit_store.CommitStoreCommitReport{
		PriceUpdates: commit_store.InternalPriceUpdates{
			TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{
				{
					SourceToken: destLinkTokenAddress,
					UsdPerToken: big.NewInt(8e18), // 8usd
				},
			},
			DestChainId:   destChainId,
			UsdPerUnitGas: big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
		},
		MerkleRoot: tree.Root(),
		Interval:   commit_store.CommitStoreInterval{Min: 1, Max: 10},
	}
	out, err := EncodeCommitReport(&report)
	require.NoError(t, err)
	decodedReport, err := DecodeCommitReport(out)
	require.NoError(t, err)
	require.Equal(t, &report, decodedReport)

	tx, err := commitStore.Report(destUser, out)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)

	// Ensure root exists.
	ts, err := commitStore.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	require.NotEqual(t, ts.String(), "0")

	// Ensure price update went through
	destChainGasPrice, err := newPriceRegistry.GetDestinationChainGasPrice(nil, destChainId)
	require.NoError(t, err)
	assert.Equal(t, "2000000000000", destChainGasPrice.Value.String())

	linkTokenPrice, err := newPriceRegistry.GetTokenPrice(nil, destLinkTokenAddress)
	require.NoError(t, err)
	assert.Equal(t, "8000000000000000000", linkTokenPrice.Value.String())
}

func TestCalculatePriceUpdates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		commitObservations []CommitObservation
		wantGas            *big.Int
	}{
		{"median", []CommitObservation{
			{SourceGasPriceUSD: big.NewInt(1)}, {SourceGasPriceUSD: big.NewInt(2)}, {SourceGasPriceUSD: big.NewInt(3)},
			{SourceGasPriceUSD: big.NewInt(4)},
		}, big.NewInt(3)},
		{"insufficient", []CommitObservation{
			{SourceGasPriceUSD: nil}, {SourceGasPriceUSD: nil}, {SourceGasPriceUSD: big.NewInt(3)},
		}, big.NewInt(0)},
		{"median including empties", []CommitObservation{
			{SourceGasPriceUSD: nil}, {SourceGasPriceUSD: big.NewInt(1)}, {SourceGasPriceUSD: big.NewInt(2)},
		}, big.NewInt(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePriceUpdates(10, tt.commitObservations)
			assert.Equal(t, tt.wantGas, got.UsdPerUnitGas)
		})
	}
	feeToken1 := common.HexToAddress("0xa")
	feeToken2 := common.HexToAddress("0xb")
	tokenPricesTests := []struct {
		name               string
		commitObservations []CommitObservation
		tokenPricesUpdates []commit_store.InternalTokenPriceUpdate
	}{
		{"median one token", []CommitObservation{
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(10)}},
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(12)}},
		}, []commit_store.InternalTokenPriceUpdate{
			{SourceToken: feeToken1, UsdPerToken: big.NewInt(12)}}},
		{"median two tokens", []CommitObservation{
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(10), feeToken2: big.NewInt(13)}},
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(12), feeToken2: big.NewInt(7)}},
		}, []commit_store.InternalTokenPriceUpdate{
			{SourceToken: feeToken1, UsdPerToken: big.NewInt(12)},
			{SourceToken: feeToken2, UsdPerToken: big.NewInt(13)}},
		},
		{"only one token with enough votes", []CommitObservation{
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(10)}},
			{TokenPricesUSD: map[common.Address]*big.Int{feeToken1: big.NewInt(12), feeToken2: big.NewInt(7)}},
		}, []commit_store.InternalTokenPriceUpdate{
			{SourceToken: feeToken1, UsdPerToken: big.NewInt(12)}},
		},
	}
	for _, tt := range tokenPricesTests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePriceUpdates(10, tt.commitObservations)
			assert.Equal(t, tt.tokenPricesUpdates, got.TokenPriceUpdates)
		})
	}
}

func TestCalculateIntervalConsensus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                    string
		intervals               []commit_store.CommitStoreInterval
		f                       int
		nextMinSeqNumForOffRamp uint64
		wantMin                 uint64
		wantMax                 uint64
		wantErr                 bool
	}{
		{"no obs", []commit_store.CommitStoreInterval{{Min: 0, Max: 0}}, 0, 100, 0, 0, false},
		{"basic", []commit_store.CommitStoreInterval{
			{Min: 9, Max: 14},
			{Min: 10, Max: 12},
			{Min: 10, Max: 14},
		}, 1, 10, 10, 14, false},
		{"not enough intervals", []commit_store.CommitStoreInterval{}, 1, 0, 0, 0, true},
		{"wrong next min", []commit_store.CommitStoreInterval{
			{Min: 9, Max: 14},
			{Min: 10, Max: 12},
			{Min: 10, Max: 14},
		}, 1, 11, 0, 0, true},
		{"min > max", []commit_store.CommitStoreInterval{
			{Min: 9, Max: 5},
			{Min: 10, Max: 4},
			{Min: 10, Max: 6},
		}, 1, 10, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateIntervalConsensus(context.Background(), tt.intervals, tt.f, func(ctx context.Context) (uint64, error) { return tt.nextMinSeqNumForOffRamp, nil })
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

type testPluginHarness = struct {
	plugin           *CommitReportingPlugin
	client           *backends.SimulatedBackend
	owner            *bind.TransactOpts
	lggr             logger.Logger
	logPoller        logpoller.LogPollerTest
	mockFeeEstimator *mocks.FeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, common.Hash]
	feeToken         common.Address
	flushLogs        func()
}

func setupTestPlugin(t *testing.T) testPluginHarness {
	chainID := testutils.NewRandomEVMChainID()
	lggr := logger.TestLogger(t)
	owner := testutils.MustNewSimTransactor(t)
	db := pgtest.NewSqlxDB(t)
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_log_poller_blocks_evm_chain_id_fkey DEFERRED`)))
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_log_poller_filters_evm_chain_id_fkey DEFERRED`)))
	require.NoError(t, utils.JustError(db.Exec(`SET CONSTRAINTS evm_logs_evm_chain_id_fkey DEFERRED`)))
	orm := logpoller.NewORM(chainID, db, lggr, pgtest.NewQConfig(true))
	client := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		owner.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10e6)
	var logPoller logpoller.LogPollerTest = logpoller.NewLogPoller(orm, evmclient.NewSimulatedBackendClient(t, client, chainID), lggr, 1*time.Hour, 2, 3, 2, 1000)

	prevStart := int64(1)
	flushLogs := func() {
		client.Commit()
		logPoller.PollAndSaveLogs(testutils.Context(t), prevStart)
		latestBlock, err := orm.SelectLatestBlock()
		require.NoError(t, err)
		prevStart = latestBlock.BlockNumber + 1
	}

	sourceChainId := testutils.NewRandomEVMChainID().Uint64()

	// Deploy link token.
	feeToken, _, _, err := link_token_interface.DeployLinkToken(owner, client)
	require.NoError(t, err)
	flushLogs()

	// Deploy native
	sourceNative := utils.RandomAddress()

	priceRegistryAddress, _, _, err := price_registry.DeployPriceRegistry(owner, client, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{SourceToken: feeToken, UsdPerToken: big.NewInt(2)},
		},
		DestChainId:   sourceChainId,
		UsdPerUnitGas: big.NewInt(1),
	}, []common.Address{}, []common.Address{feeToken}, uint32(time.Hour.Seconds()))
	require.NoError(t, err)
	flushLogs()

	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddress, client)
	require.NoError(t, err)

	err = logPoller.RegisterFilter(logpoller.Filter{Name: logpoller.FilterName(COMMIT_PRICE_UPDATES, priceRegistryAddress),
		EventSigs: []common.Hash{UsdPerUnitGasUpdated, UsdPerTokenUpdated}, Addresses: []common.Address{priceRegistryAddress}})
	require.NoError(t, err)

	sourceFeeEstimator := mocks.NewFeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, common.Hash](t)

	plugin := CommitReportingPlugin{
		config: CommitPluginConfig{
			dest:               logPoller,
			priceRegistry:      priceRegistry,
			priceGetter:        fakePriceGetter{},
			sourceNative:       sourceNative,
			sourceFeeEstimator: sourceFeeEstimator,
			sourceChainID:      sourceChainId,
		},
		offchainConfig: OffchainConfig{
			FeeUpdateHeartBeat: models.MustMakeDuration(12 * time.Hour),
		},
	}
	return testPluginHarness{
		plugin:           &plugin,
		client:           client,
		owner:            owner,
		lggr:             lggr,
		logPoller:        logPoller,
		mockFeeEstimator: sourceFeeEstimator,
		feeToken:         feeToken,
		flushLogs:        flushLogs,
	}
}

func TestGeneratePriceUpdates(t *testing.T) {
	th := setupTestPlugin(t)

	gasPrice := big.NewInt(3e9)
	mockedGetFee := th.mockFeeEstimator.On(
		"GetFee",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(gas.EvmFee{Legacy: assets.NewWei(gasPrice)}, uint32(150e3), nil)

	fakePrices, err := th.plugin.config.priceGetter.TokenPricesUSD(testutils.Context(t), []common.Address{th.feeToken}) // 2e20 hardcoded in fakePriceGetter
	require.NoError(t, err)
	fakePrice := fakePrices[th.feeToken]

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
			expectedTokenPricesUSD: map[common.Address]*big.Int{th.feeToken: fakePrice},
		},
		{
			name:                   "gasPrice up-to-date",
			updateGasPriceUSD:      expectedGasPrice,
			expectedGasPriceUSD:    nil,
			expectedTokenPricesUSD: map[common.Address]*big.Int{th.feeToken: fakePrice},
		},
		{
			name:                   "tokenPrice up-to-date, gasPrice deviated",
			updateTokenPricesUSD:   map[common.Address]*big.Int{th.feeToken: fakePrice},
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
				_, err = th.plugin.config.priceRegistry.ApplyFeeTokensUpdates(th.owner, tt.addFeeTokens, []common.Address{})
				require.NoError(t, err)
				th.flushLogs()
			}
			if len(tt.updateTokenPricesUSD) > 0 || tt.updateGasPriceUSD != nil {
				destChainId := uint64(0)
				if tt.updateGasPriceUSD != nil {
					destChainId = th.plugin.config.sourceChainID
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
				_, err = th.plugin.config.priceRegistry.UpdatePrices(th.owner, price_registry.InternalPriceUpdates{
					TokenPriceUpdates: tokenPriceUpdates,
					DestChainId:       destChainId,
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

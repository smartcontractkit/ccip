package legacygasstation_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/test-go/testify/mock"
	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	forwarder_wrapper "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/forwarder"
	mock_contracts "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation"
	forwarder_mocks "github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/types"
)

type request struct {
	tx        legacygasstation.TestLegacyGaslessTx
	confirmed bool
	failed    bool
}

type testcase struct {
	name                 string
	latestBlock          int64
	lookbackBlock        int64
	chainID              uint64
	requestData          []request
	forwardSucceededLogs []*forwarder_wrapper.ForwarderForwardSucceeded
	offrampExecutionLogs []*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged
	resultData           []legacygasstation.TestLegacyGaslessTx
}

var (
	tests = []testcase{
		{
			name:          "submitted transaction confirmed",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       testutils.SimulatedChainID.Uint64(),
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
					},
					confirmed: true,
				},
			},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
					Status:             types.Confirmed,
				},
			},
		},
		{
			name:          "submitted transaction failed",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       testutils.SimulatedChainID.Uint64(),
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
					},
					failed: true,
				},
			},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
					Status:             types.Failure,
				},
			},
		},
		{
			name:          "confirmed transaction finalized",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       testutils.SimulatedChainID.Uint64(),
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
						Status:             types.Confirmed,
					},
				},
			},
			forwardSucceededLogs: []*forwarder_wrapper.ForwarderForwardSucceeded{
				{
					From:  legacygasstation.FromAddress,
					Nonce: big.NewInt(0),
					Raw: geth_types.Log{
						Address: legacygasstation.ForwarderAddress,
					},
				},
			},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
					Status:             types.Finalized,
				},
			},
		},
		{
			name:          "confirmed transaction failed",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       testutils.SimulatedChainID.Uint64(),
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
						Status:             types.Confirmed,
					},
					failed: true,
				},
			},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
					Status:             types.Failure,
				},
			},
		},
		{
			name:          "multiple submitted txs finalized",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       testutils.SimulatedChainID.Uint64(),
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
						Status:             types.Confirmed,
					},
				},
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						From:               common.HexToAddress("0x780b3102c62d5DfDCc658B3480B93041Ba46F499"),
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
						Status:             types.Confirmed,
					},
				},
			},
			forwardSucceededLogs: []*forwarder_wrapper.ForwarderForwardSucceeded{
				{
					From:  legacygasstation.FromAddress,
					Nonce: big.NewInt(0),
					Raw: geth_types.Log{
						Address: legacygasstation.ForwarderAddress,
					},
				},
				{
					From:  common.HexToAddress("0x780b3102c62d5DfDCc658B3480B93041Ba46F499"),
					Nonce: big.NewInt(0),
					Raw: geth_types.Log{
						Address: legacygasstation.ForwarderAddress,
					},
				},
			},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
					Status:             types.Finalized,
				},
				{
					From:               common.HexToAddress("0x780b3102c62d5DfDCc658B3480B93041Ba46F499"),
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
					Status:             types.Finalized,
				},
			},
		},
		{
			name:          "no forwarder logs",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       testutils.SimulatedChainID.Uint64(),
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
					},
				},
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						From:               common.HexToAddress("0x780b3102c62d5DfDCc658B3480B93041Ba46F499"),
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
					},
				},
			},
			forwardSucceededLogs: []*forwarder_wrapper.ForwarderForwardSucceeded{},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
				},
				{
					From:               common.HexToAddress("0x780b3102c62d5DfDCc658B3480B93041Ba46F499"),
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
				},
			},
		},
		{
			name:          "cross chain submitted to source finalized log",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       testutils.SimulatedChainID.Uint64(),
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      testutils.SimulatedChainID.Uint64(),
						DestinationChainID: 1000,
						Status:             types.Confirmed,
					},
				},
			},
			forwardSucceededLogs: []*forwarder_wrapper.ForwarderForwardSucceeded{
				{
					From:  legacygasstation.FromAddress,
					Nonce: big.NewInt(0),
					Raw: geth_types.Log{
						Address: legacygasstation.ForwarderAddress,
					},
					ReturnValue: common.HexToHash("0x30").Bytes(),
				},
			},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      testutils.SimulatedChainID.Uint64(),
					DestinationChainID: 1000,
					Status:             types.SourceFinalized,
				},
			},
		},
		{
			name:          "cross chain source finalized to finalized",
			latestBlock:   100,
			lookbackBlock: 50,
			chainID:       1000,
			requestData: []request{
				{
					tx: legacygasstation.TestLegacyGaslessTx{
						Nonce:              big.NewInt(0),
						Amount:             big.NewInt(1e18),
						SourceChainID:      1000,
						DestinationChainID: testutils.SimulatedChainID.Uint64(),
						CCIPMessageID:      ptr(common.HexToHash("0x30")),
						Status:             types.SourceFinalized,
					},
				},
			},
			offrampExecutionLogs: []*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged{
				{
					MessageId: common.HexToHash("0x30"),
					Raw: geth_types.Log{
						Address: legacygasstation.OfframpAddress,
					},
				},
			},
			resultData: []legacygasstation.TestLegacyGaslessTx{
				{
					Nonce:              big.NewInt(0),
					Amount:             big.NewInt(1e18),
					SourceChainID:      1000,
					DestinationChainID: testutils.SimulatedChainID.Uint64(),
					Status:             types.Finalized,
				},
			},
		},
	}
)

func TestSidecar(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sc, orm := setUp(t, test)
			err := sc.Run(testutils.Context(t))
			require.NoError(t, err)
			assertAfterSidecarRun(t, test, orm)
		})
	}
}

func setUp(t *testing.T, test testcase) (*legacygasstation.Sidecar, legacygasstation.ORM) {
	cfg, db := heavyweight.FullTestDBV2(t, "legacy_gas_station_sidecar_test", func(c *chainlink.Config, s *chainlink.Secrets) {
		require.Zero(t, testutils.SimulatedChainID.Cmp(c.EVM[0].ChainID.ToInt()))
	})
	backend := cltest.NewSimulatedBackend(t, core.GenesisAlloc{}, uint32(ethconfig.Defaults.Miner.GasCeil))
	app := cltest.NewApplicationWithConfigV2AndKeyOnSimulatedBlockchain(t, cfg, backend)
	forwarder := forwarder_mocks.NewForwarderInterface(t)
	offramp := mock_contracts.NewEVM2EVMOffRampInterface(t)
	lggr := logger.TestLogger(t)
	orm := legacygasstation.NewORM(db, lggr, cfg)
	chain, err := app.Chains.EVM.Get(testutils.SimulatedChainID)
	require.NoError(t, err)
	lp := mocks.NewLogPoller(t)
	lp.On("RegisterFilter", mock.Anything).Return(nil)
	lp.On("LatestBlock", mock.Anything).Return(test.latestBlock, nil)
	forwarder.On("Address").Return(legacygasstation.ForwarderAddress)
	offramp.On("Address").Return(legacygasstation.OfframpAddress)
	var (
		fsLpLogs  []logpoller.Log
		oelLpLogs []logpoller.Log
	)

	for _, fl := range test.forwardSucceededLogs {
		forwarder.On("ParseLog", mock.Anything).Return(fl, nil).Once()
		fsLpLogs = append(fsLpLogs, logpoller.Log{
			EventSig: forwarder_wrapper.ForwarderForwardSucceeded{}.Topic(),
		})
	}
	for _, oel := range test.offrampExecutionLogs {
		offramp.On("ParseLog", mock.Anything).Return(oel, nil).Once()
		oelLpLogs = append(oelLpLogs, logpoller.Log{
			EventSig: evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged{}.Topic(),
		})
	}
	lp.On("IndexedLogsByBlockRange",
		mock.Anything,
		mock.Anything,
		forwarder_wrapper.ForwarderForwardSucceeded{}.Topic(),
		legacygasstation.ForwarderAddress,
		1,
		mock.Anything,
		mock.Anything,
	).Return(fsLpLogs, nil).Maybe()
	lp.On("IndexedLogsByBlockRange",
		mock.Anything,
		mock.Anything,
		evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged{}.Topic(),
		legacygasstation.OfframpAddress,
		2,
		mock.Anything,
		mock.Anything,
	).Return(oelLpLogs, nil).Maybe()

	sc, err := legacygasstation.NewSidecar(
		lggr,
		lp,
		forwarder,
		offramp,
		chain.Config(),
		testutils.SimulatedChainID.Uint64(),
		uint32(test.lookbackBlock),
		orm,
	)
	require.NoError(t, err)
	for i, r := range test.requestData {
		chainID := cltest.FixtureChainID
		blockNumber := int64(75)
		_, fromAddress := cltest.MustInsertRandomKeyReturningState(t, app.KeyStore.Eth(), chainID)
		var ethTx txmgr.EvmTx
		if r.confirmed {
			ethTx = cltest.MustInsertConfirmedEthTxBySaveFetchedReceipts(t, app.TxmStorageService(), fromAddress, int64(i), blockNumber, chainID)
		} else if r.failed {
			ethTx = cltest.MustInsertFatalErrorEthTx(t, app.TxmStorageService(), fromAddress)
		} else {
			ethTx = cltest.MustInsertInProgressEthTxWithAttempt(t, app.TxmStorageService(), int64(i), fromAddress)
		}
		r.tx.EthTxID = ethTx.GetID()
		tx := legacygasstation.LegacyGaslessTx(t, r.tx)
		err = orm.InsertLegacyGaslessTx(tx)
		require.NoError(t, err)
		err = orm.UpdateLegacyGaslessTx(tx) // update populates ccipMessageID and failureReason
		require.NoError(t, err)
	}
	return sc, orm
}

func assertAfterSidecarRun(t *testing.T, test testcase, orm legacygasstation.ORM) {
	confirmedTxs, submittedTxs, finalizedTxs, sourceFinalizedTxs, failedTxs := categorizeTestTxs(t, test.resultData)

	txs, err := orm.SelectBySourceChainIDAndStatus(test.chainID, types.Confirmed)
	require.NoError(t, err)
	require.Equal(t, len(confirmedTxs), len(txs))

	txs, err = orm.SelectBySourceChainIDAndStatus(test.chainID, types.Submitted)
	require.NoError(t, err)
	require.Equal(t, len(submittedTxs), len(txs))

	txs, err = orm.SelectBySourceChainIDAndStatus(test.chainID, types.Finalized)
	require.NoError(t, err)
	require.Equal(t, len(finalizedTxs), len(txs))

	txs, err = orm.SelectBySourceChainIDAndStatus(test.chainID, types.SourceFinalized)
	require.NoError(t, err)
	require.Equal(t, len(sourceFinalizedTxs), len(txs))

	txs, err = orm.SelectBySourceChainIDAndStatus(test.chainID, types.Failure)
	require.NoError(t, err)
	require.Equal(t, len(failedTxs), len(txs))
}

func categorizeTestTxs(t *testing.T, testTxs []legacygasstation.TestLegacyGaslessTx) (
	confirmedTxs,
	submittedTxs,
	finalizedTxs,
	sourceFinalizedTxs,
	failedTxs []types.LegacyGaslessTx,
) {
	for _, testTx := range testTxs {
		tx := legacygasstation.LegacyGaslessTx(t, testTx)
		switch tx.Status {
		case types.Confirmed:
			confirmedTxs = append(confirmedTxs, tx)
		case types.Submitted:
			submittedTxs = append(submittedTxs, tx)
		case types.SourceFinalized:
			sourceFinalizedTxs = append(sourceFinalizedTxs, tx)
		case types.Finalized:
			finalizedTxs = append(finalizedTxs, tx)
		case types.Failure:
			failedTxs = append(failedTxs, tx)
		default:
			t.Errorf("unexpected status: %s", tx.Status)
		}
	}
	return
}

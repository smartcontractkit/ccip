package ccip

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpMocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	mock_contracts "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	plugintesthelpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers/plugins"

	txmgrMocks "github.com/smartcontractkit/chainlink/v2/common/txmgr/types/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

var MaxPayloadLength = 100_000

type execTestHarness = struct {
	plugintesthelpers.CCIPPluginTestHarness
	plugin *ExecutionReportingPlugin
}

func setupExecTestHarness(t *testing.T) execTestHarness {
	th := plugintesthelpers.SetupCCIPTestHarness(t)

	destFeeEstimator := txmgrMocks.NewFeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, common.Hash](t)

	destFeeEstimator.On(
		"GetFee",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Maybe().Return(gas.EvmFee{Legacy: assets.NewWei(defaultGasPrice)}, uint32(200e3), nil)

	offchainConfig := ccipconfig.ExecOffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   0,
		MaxGasPrice:                 200e9,
		BatchGasLimit:               5e6,
		RootSnoozeTime:              models.MustMakeDuration(10 * time.Minute),
		InflightCacheExpiry:         models.MustMakeDuration(3 * time.Minute),
		RelativeBoostPerWaitHour:    0.07,
	}
	plugin := ExecutionReportingPlugin{
		config: ExecutionPluginConfig{
			lggr:                  th.Lggr,
			sourceLP:              th.SourceLP,
			destLP:                th.DestLP,
			srcPriceRegistry:      th.Source.PriceRegistry,
			onRamp:                th.Source.OnRamp,
			commitStore:           th.Dest.CommitStore,
			offRamp:               th.Dest.OffRamp,
			srcWrappedNativeToken: th.Source.WrappedNative.Address(),
			leafHasher:            hasher.NewLeafHasher(th.Source.ChainID, th.Dest.ChainID, th.Source.OnRamp.Address(), hasher.NewKeccakCtx()),
			destGasEstimator:      destFeeEstimator,
		},
		onchainConfig:        th.ExecOnchainConfig,
		offchainConfig:       offchainConfig,
		lggr:                 th.Lggr.Named("ExecutionReportingPlugin"),
		snoozedRoots:         make(map[[32]byte]time.Time),
		inflightReports:      newInflightReportsContainer(offchainConfig.InflightCacheExpiry.Duration()),
		destPriceRegistry:    th.Dest.PriceRegistry,
		destWrappedNative:    th.Dest.WrappedNative.Address(),
		srcToDstTokenMapping: map[common.Address]common.Address{},
	}
	return execTestHarness{
		CCIPPluginTestHarness: th,
		plugin:                &plugin,
	}
}

func TestMaxExecutionReportSize(t *testing.T) {
	// Ensure that given max payload size and max num tokens,
	// Our report size is under the tx size limit.
	th := setupExecTestHarness(t)
	th.plugin.F = 1
	mb := th.GenerateAndSendMessageBatch(t, 50, MaxPayloadLength, MaxTokensPerMessage)

	// commit root
	encoded, err := abihelpers.EncodeCommitReport(commit_store.CommitStoreCommitReport{
		Interval:   mb.Interval,
		MerkleRoot: mb.Root,
		PriceUpdates: commit_store.InternalPriceUpdates{
			TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
			DestChainSelector: 0,
			UsdPerUnitGas:     big.NewInt(0),
		},
	})
	require.NoError(t, err)
	_, err = th.Dest.CommitStoreHelper.Report(th.Dest.User, encoded)
	require.NoError(t, err)
	// double commit to ensure enough confirmations
	th.CommitAndPollLogs(t)
	th.CommitAndPollLogs(t)

	fullReport, err := abihelpers.EncodeExecutionReport(evm_2_evm_offramp.InternalExecutionReport{
		SequenceNumbers:   mb.SeqNums,
		EncodedMessages:   mb.AllMsgBytes,
		OffchainTokenData: mb.TokenData,
		Proofs:            mb.Proof.Hashes,
		ProofFlagBits:     mb.ProofBits,
	})
	require.NoError(t, err)
	// ensure "naive" full report would be bigger than limit
	require.Greater(t, len(fullReport), MaxExecutionReportLength, "full execution report length")

	observations := make([]ObservedMessage, len(mb.SeqNums))
	for i, seqNr := range mb.SeqNums {
		observations[i] = ObservedMessage{SeqNr: seqNr, TokenData: mb.TokenData[i]}
	}

	// buildReport should cap the built report to fit in MaxExecutionReportLength
	execReport, err := th.plugin.buildReport(testutils.Context(t), th.Lggr, observations)
	require.NoError(t, err)
	require.LessOrEqual(t, len(execReport), MaxExecutionReportLength, "built execution report length")
}

func TestExecutionReportToEthTxMetadata(t *testing.T) {
	c := plugintesthelpers.SetupCCIPTestHarness(t)
	tests := []struct {
		name     string
		msgBatch plugintesthelpers.MessageBatch
		err      error
	}{
		{
			"happy flow",
			c.GenerateAndSendMessageBatch(t, 5, MaxPayloadLength, MaxTokensPerMessage),
			nil,
		},
		{
			"invalid msgs",
			func() plugintesthelpers.MessageBatch {
				mb := c.GenerateAndSendMessageBatch(t, 5, MaxPayloadLength, MaxTokensPerMessage)
				mb.AllMsgBytes[0] = []byte{1, 1, 1, 1}
				return mb
			}(),
			errors.New("abi: cannot marshal in to go type: length insufficient 4 require 32"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			encExecReport, err := abihelpers.EncodeExecutionReport(evm_2_evm_offramp.InternalExecutionReport{
				SequenceNumbers:   tc.msgBatch.SeqNums,
				EncodedMessages:   tc.msgBatch.AllMsgBytes,
				OffchainTokenData: tc.msgBatch.TokenData,
				Proofs:            tc.msgBatch.Proof.Hashes,
				ProofFlagBits:     tc.msgBatch.ProofBits,
			})
			require.NoError(t, err)
			txMeta, err := ExecutionReportToEthTxMeta(encExecReport)
			if tc.err != nil {
				require.Equal(t, tc.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotNil(t, txMeta)
			require.Len(t, txMeta.MessageIDs, len(tc.msgBatch.AllMsgBytes))
		})
	}
}

func TestUpdateSourceToDestTokenMapping(t *testing.T) {
	expectedNewBlockNumber := int64(10000)
	logs := []logpoller.Log{{BlockNumber: expectedNewBlockNumber}}
	mockDestLP := &lpMocks.LogPoller{}

	mockDestLP.On("LatestLogEventSigsAddrsWithConfs", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(logs, nil)

	sourceToken, destToken := common.HexToAddress("111111"), common.HexToAddress("222222")

	mockOffRamp := &mock_contracts.EVM2EVMOffRampInterface{}
	mockOffRamp.On("Address").Return(common.HexToAddress("0x01"))
	mockOffRamp.On("GetSupportedTokens", mock.Anything).Return([]common.Address{sourceToken}, nil)
	mockOffRamp.On("GetDestinationToken", mock.Anything, sourceToken).Return(destToken, nil)

	plugin := ExecutionReportingPlugin{
		config: ExecutionPluginConfig{
			destLP:  mockDestLP,
			offRamp: mockOffRamp,
		},
		srcToDstTokenMappingBlock: 0,
	}

	require.NoError(t, plugin.updateSourceToDestTokenMapping(context.Background()))
	assert.Equal(t, expectedNewBlockNumber+1, plugin.srcToDstTokenMappingBlock)

	gotDestToken, ok := plugin.srcToDstTokenMapping[sourceToken]
	require.True(t, ok)
	require.Equal(t, destToken, gotDestToken)
}

func TestExecObservation(t *testing.T) {
	th := setupExecTestHarness(t)
	th.plugin.F = 1
	mb := th.GenerateAndSendMessageBatch(t, 2, 10, 1)

	// commit root
	encoded, err := abihelpers.EncodeCommitReport(commit_store.CommitStoreCommitReport{
		Interval:   mb.Interval,
		MerkleRoot: mb.Root,
		PriceUpdates: commit_store.InternalPriceUpdates{
			TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
			DestChainSelector: 0,
			UsdPerUnitGas:     big.NewInt(0),
		},
	})
	require.NoError(t, err)
	_, err = th.Dest.CommitStoreHelper.Report(th.Dest.User, encoded)
	require.NoError(t, err)
	// double commit to ensure enough confirmations
	th.CommitAndPollLogs(t)
	th.CommitAndPollLogs(t)

	tests := []struct {
		name            string
		commitStoreDown bool
		expected        *ExecutionObservation
		expectedError   bool
	}{
		{
			"base",
			false,
			&ExecutionObservation{Messages: []ObservedMessage{
				{SeqNr: 1, TokenData: [][]byte{{}}},
				{SeqNr: 2, TokenData: [][]byte{{}}},
			}},
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

			gotObs, err := th.plugin.Observation(testutils.Context(t), ocrtypes.ReportTimestamp{}, ocrtypes.Query{})

			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			var decodedObservation *ExecutionObservation
			if gotObs != nil {
				decodedObservation = new(ExecutionObservation)
				err = json.Unmarshal(gotObs, decodedObservation)
				require.NoError(t, err)

			}
			assert.Equal(t, tt.expected, decodedObservation)
		})
	}
}

func TestExecReport(t *testing.T) {
	th := setupExecTestHarness(t)
	th.plugin.F = 1
	mb := th.GenerateAndSendMessageBatch(t, 2, 10, 1)

	// commit root
	encoded, err := abihelpers.EncodeCommitReport(commit_store.CommitStoreCommitReport{
		Interval:   mb.Interval,
		MerkleRoot: mb.Root,
		PriceUpdates: commit_store.InternalPriceUpdates{
			TokenPriceUpdates: []commit_store.InternalTokenPriceUpdate{},
			DestChainSelector: 0,
			UsdPerUnitGas:     big.NewInt(0),
		},
	})
	require.NoError(t, err)
	execReport := mb.ToExecutionReport()

	_, err = th.Dest.CommitStoreHelper.Report(th.Dest.User, encoded)
	require.NoError(t, err)
	// double commit to ensure enough confirmations
	th.CommitAndPollLogs(t)
	th.CommitAndPollLogs(t)

	tests := []struct {
		name                 string
		commitStoreDown      bool
		observations         [][]ObservedMessage
		expectedShouldReport bool
		expectedReport       *evm_2_evm_offramp.InternalExecutionReport
		expectedError        bool
	}{
		{
			"base",
			false,
			[][]ObservedMessage{
				{{SeqNr: 1, TokenData: [][]byte{{}}}, {SeqNr: 2, TokenData: [][]byte{{}}}},
				{{SeqNr: 1, TokenData: [][]byte{{}}}, {SeqNr: 2, TokenData: [][]byte{{}}}},
			},
			true,
			&execReport,
			false,
		},
		{
			"partial observation",
			false,
			[][]ObservedMessage{
				{{SeqNr: 1, TokenData: [][]byte{{}}}, {SeqNr: 2, TokenData: [][]byte{{}}}},
				{{SeqNr: 1, TokenData: [][]byte{{}}}},
			},
			true,
			func() *evm_2_evm_offramp.InternalExecutionReport {
				mb2 := mb
				mb2.Messages = mb.Messages[:1]
				mb2.SeqNums = mb.SeqNums[:1]
				mb2.AllMsgBytes = mb.AllMsgBytes[:1]
				mb2.TokenData = mb.TokenData[:1]
				mb2.Interval = commit_store.CommitStoreInterval{Min: 1, Max: 1}
				mb2.Proof = mb2.Tree.Prove([]int{0})
				mb2.ProofBits = abihelpers.ProofFlagsToBits(mb2.Proof.SourceFlags)
				report := mb2.ToExecutionReport()
				return &report
			}(),
			false,
		},
		{
			"empty",
			false,
			[][]ObservedMessage{
				{{SeqNr: 1, TokenData: [][]byte{{}}}, {SeqNr: 2, TokenData: [][]byte{{}}}},
				{},
			},
			false,
			nil,
			false,
		},
		{
			"unknown seqNr",
			false,
			[][]ObservedMessage{
				{{SeqNr: 1, TokenData: [][]byte{{}}}, {SeqNr: 2, TokenData: [][]byte{{}}}, {SeqNr: 3, TokenData: [][]byte{{}}}},
				{{SeqNr: 1, TokenData: [][]byte{{}}}, {SeqNr: 2, TokenData: [][]byte{{}}}, {SeqNr: 3, TokenData: [][]byte{{}}}},
			},
			false,
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obs []ocrtypes.AttributedObservation
			for _, o := range tt.observations {
				encoded, err := ExecutionObservation{Messages: o}.Marshal()
				require.NoError(t, err)
				obs = append(obs, ocrtypes.AttributedObservation{Observation: encoded})
			}
			gotShouldReport, gotReport, err := th.plugin.Report(testutils.Context(t), ocrtypes.ReportTimestamp{}, ocrtypes.Query{}, obs)

			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expectedShouldReport, gotShouldReport)

			var encodedReport ocrtypes.Report
			if tt.expectedReport != nil {
				encodedReport, err = abihelpers.EncodeExecutionReport(*tt.expectedReport)
				require.NoError(t, err)
			}
			assert.Equal(t, encodedReport, gotReport)
		})
	}
}

func TestExecShouldAcceptFinalizedReport(t *testing.T) {
	report := evm_2_evm_offramp.InternalExecutionReport{
		SequenceNumbers:   []uint64{12},
		EncodedMessages:   [][]byte{[]byte(`msg12`)},
		OffchainTokenData: [][][]byte{{}},
		Proofs:            [][32]byte{{}},
		ProofFlagBits:     big.NewInt(1),
	}
	encodedReport, err := abihelpers.EncodeExecutionReport(report)
	require.NoError(t, err)

	mockOffRamp := &mock_contracts.EVM2EVMOffRampInterface{}
	plugin := ExecutionReportingPlugin{
		config: ExecutionPluginConfig{
			offRamp: mockOffRamp,
		},
		lggr:            logger.TestLogger(t),
		inflightReports: newInflightReportsContainer(models.MustMakeDuration(1 * time.Hour).Duration()),
	}

	mockedExecState := mockOffRamp.On("GetExecutionState", mock.Anything, uint64(12)).Return(uint8(abihelpers.ExecutionStateUntouched), nil).Once()

	should, err := plugin.ShouldAcceptFinalizedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, true, should)

	mockedExecState.Return(uint8(abihelpers.ExecutionStateSuccess), nil).Once()

	should, err = plugin.ShouldAcceptFinalizedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, false, should)
}

func TestExecShouldTransmitAcceptedReport(t *testing.T) {
	report := evm_2_evm_offramp.InternalExecutionReport{
		SequenceNumbers:   []uint64{12},
		EncodedMessages:   [][]byte{[]byte(`msg12`)},
		OffchainTokenData: [][][]byte{{}},
		Proofs:            [][32]byte{{}},
		ProofFlagBits:     big.NewInt(1),
	}
	encodedReport, err := abihelpers.EncodeExecutionReport(report)
	require.NoError(t, err)

	mockOffRamp := &mock_contracts.EVM2EVMOffRampInterface{}
	mockCommitStore := &mock_contracts.CommitStoreInterface{}

	plugin := ExecutionReportingPlugin{
		config: ExecutionPluginConfig{
			offRamp:     mockOffRamp,
			commitStore: mockCommitStore,
		},
		lggr:            logger.TestLogger(t),
		inflightReports: newInflightReportsContainer(models.MustMakeDuration(1 * time.Hour).Duration()),
	}

	mockCommitStore.On("Paused", mock.Anything).Return(false, nil)
	mockCommitStore.On("IsAFNHealthy", mock.Anything).Return(true, nil)
	mockedExecState := mockOffRamp.On("GetExecutionState", mock.Anything, uint64(12)).Return(uint8(abihelpers.ExecutionStateUntouched), nil).Once()

	should, err := plugin.ShouldTransmitAcceptedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, true, should)

	mockedExecState.Return(uint8(abihelpers.ExecutionStateFailure), nil).Once()
	should, err = plugin.ShouldTransmitAcceptedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, false, should)
}

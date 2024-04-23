package ccipexec

import (
	"bytes"
	"context"
	"encoding/json"
	"math"
	"math/big"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/cometbft/cometbft/libs/rand"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	lpMocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	ccipcachemocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/batchreader"
	ccipdataprovidermocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/ccipdataprovider/mocks"
	ccipdatamocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

func TestExecutionReportingPlugin_Observation(t *testing.T) {
	testCases := []struct {
		name               string
		commitStorePaused  bool
		sourceChainCursed  bool
		inflightReports    []InflightInternalExecutionReport
		unexpiredReports   []cciptypes.CommitStoreReportWithTxMeta
		sendRequests       []cciptypes.EVM2EVMMessageWithTxMeta
		executedSeqNums    []uint64
		tokenPoolsMapping  map[common.Address]common.Address
		blessedRoots       map[[32]byte]bool
		senderNonce        uint64
		rateLimiterState   cciptypes.TokenBucketRateLimit
		expErr             bool
		sourceChainHealthy bool
		destChainHealthy   bool
	}{
		{
			name:               "commit store is down",
			commitStorePaused:  true,
			sourceChainCursed:  false,
			sourceChainHealthy: true,
			destChainHealthy:   true,
			expErr:             true,
		},
		{
			name:               "source chain is cursed",
			commitStorePaused:  false,
			sourceChainCursed:  true,
			sourceChainHealthy: true,
			destChainHealthy:   true,
			expErr:             true,
		},
		{
			name:               "source chain not healthy",
			commitStorePaused:  false,
			sourceChainCursed:  false,
			sourceChainHealthy: false,
			destChainHealthy:   true,
			expErr:             true,
		},
		{
			name:               "dest chain not healthy",
			commitStorePaused:  false,
			sourceChainCursed:  false,
			sourceChainHealthy: true,
			destChainHealthy:   false,
			expErr:             true,
		},
		{
			name:               "happy flow",
			commitStorePaused:  false,
			sourceChainCursed:  false,
			sourceChainHealthy: true,
			destChainHealthy:   true,
			inflightReports:    []InflightInternalExecutionReport{},
			unexpiredReports: []cciptypes.CommitStoreReportWithTxMeta{
				{
					CommitStoreReport: cciptypes.CommitStoreReport{
						Interval:   cciptypes.CommitStoreInterval{Min: 10, Max: 12},
						MerkleRoot: [32]byte{123},
					},
				},
			},
			blessedRoots: map[[32]byte]bool{
				{123}: true,
			},
			rateLimiterState: cciptypes.TokenBucketRateLimit{
				IsEnabled: false,
			},
			tokenPoolsMapping: map[common.Address]common.Address{},
			senderNonce:       9,
			sendRequests: []cciptypes.EVM2EVMMessageWithTxMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 10, GasLimit: big.NewInt(0)},
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 11, GasLimit: big.NewInt(0)},
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 12, GasLimit: big.NewInt(0)},
				},
			},
		},
	}

	ctx := testutils.Context(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := &ExecutionReportingPlugin{}
			p.inflightReports = newInflightExecReportsContainer(time.Minute)
			p.inflightReports.reports = tc.inflightReports
			p.lggr = logger.TestLogger(t)
			p.tokenDataWorker = tokendata.NewBackgroundWorker(
				make(map[cciptypes.Address]tokendata.Reader), 10, 5*time.Second, time.Hour)
			p.metricsCollector = ccip.NoopMetricsCollector

			commitStoreReader := ccipdatamocks.NewCommitStoreReader(t)
			commitStoreReader.On("IsDown", mock.Anything).Return(tc.commitStorePaused, nil).Maybe()
			commitStoreReader.On("IsDestChainHealthy", mock.Anything).Return(tc.destChainHealthy, nil).Maybe()
			// Blessed roots return true
			for root, blessed := range tc.blessedRoots {
				commitStoreReader.On("IsBlessed", mock.Anything, root).Return(blessed, nil).Maybe()
			}
			commitStoreReader.On("GetAcceptedCommitReportsGteTimestamp", ctx, mock.Anything, 0).
				Return(tc.unexpiredReports, nil).Maybe()
			p.commitStoreReader = commitStoreReader

			var executionEvents []cciptypes.ExecutionStateChangedWithTxMeta
			for _, seqNum := range tc.executedSeqNums {
				executionEvents = append(executionEvents, cciptypes.ExecutionStateChangedWithTxMeta{
					ExecutionStateChanged: cciptypes.ExecutionStateChanged{SequenceNumber: seqNum},
				})
			}

			offRamp, _ := testhelpers.NewFakeOffRamp(t)
			offRamp.SetRateLimiterState(tc.rateLimiterState)

			tokenPoolBatchedReader, err := batchreader.NewEVMTokenPoolBatchedReader(p.lggr, 0, ccipcalc.EvmAddrToGeneric(offRamp.Address()), nil)
			assert.NoError(t, err)
			p.tokenPoolBatchedReader = tokenPoolBatchedReader

			mockOffRampReader := ccipdatamocks.NewOffRampReader(t)
			mockOffRampReader.On("GetExecutionStateChangesBetweenSeqNums", ctx, mock.Anything, mock.Anything, 0).
				Return(executionEvents, nil).Maybe()
			mockOffRampReader.On("CurrentRateLimiterState", mock.Anything).Return(tc.rateLimiterState, nil).Maybe()
			mockOffRampReader.On("Address", ctx).Return(cciptypes.Address(offRamp.Address().String()), nil).Maybe()
			senderNonces := map[cciptypes.Address]uint64{
				cciptypes.Address(utils.RandomAddress().String()): tc.senderNonce,
			}
			mockOffRampReader.On("GetSendersNonce", mock.Anything, mock.Anything).Return(senderNonces, nil).Maybe()
			mockOffRampReader.On("GetTokenPoolsRateLimits", ctx, []ccipdata.TokenPoolReader{}).
				Return([]cciptypes.TokenBucketRateLimit{}, nil).Maybe()

			mockOffRampReader.On("GetSourceToDestTokensMapping", ctx).Return(nil, nil).Maybe()
			mockOffRampReader.On("GetTokens", ctx).Return(cciptypes.OffRampTokens{
				DestinationTokens: []cciptypes.Address{},
				SourceTokens:      []cciptypes.Address{},
			}, nil).Maybe()
			p.offRampReader = mockOffRampReader

			mockOnRampReader := ccipdatamocks.NewOnRampReader(t)
			mockOnRampReader.On("IsSourceCursed", ctx).Return(tc.sourceChainCursed, nil).Maybe()
			mockOnRampReader.On("IsSourceChainHealthy", ctx).Return(tc.sourceChainHealthy, nil).Maybe()
			mockOnRampReader.On("GetSendRequestsBetweenSeqNums", ctx, mock.Anything, mock.Anything, false).
				Return(tc.sendRequests, nil).Maybe()
			sourcePriceRegistryAddress := cciptypes.Address(utils.RandomAddress().String())
			mockOnRampReader.On("SourcePriceRegistryAddress", ctx).Return(sourcePriceRegistryAddress, nil).Maybe()
			p.onRampReader = mockOnRampReader

			mockGasPriceEstimator := prices.NewMockGasPriceEstimatorExec(t)
			mockGasPriceEstimator.On("GetGasPrice", ctx).Return(big.NewInt(1), nil).Maybe()
			p.gasPriceEstimator = mockGasPriceEstimator

			destPriceRegReader := ccipdatamocks.NewPriceRegistryReader(t)
			destPriceRegReader.On("GetTokenPrices", ctx, mock.Anything).Return(
				[]cciptypes.TokenPriceUpdate{{TokenPrice: cciptypes.TokenPrice{Token: ccipcalc.HexToAddress("0x1"), Value: big.NewInt(123)}, TimestampUnixSec: big.NewInt(time.Now().Unix())}}, nil).Maybe()
			destPriceRegReader.On("Address", ctx).Return(cciptypes.Address(utils.RandomAddress().String()), nil).Maybe()
			destPriceRegReader.On("GetFeeTokens", ctx).Return([]cciptypes.Address{}, nil).Maybe()
			sourcePriceRegReader := ccipdatamocks.NewPriceRegistryReader(t)
			sourcePriceRegReader.On("Address", ctx).Return(sourcePriceRegistryAddress, nil).Maybe()
			sourcePriceRegReader.On("GetFeeTokens", ctx).Return([]cciptypes.Address{}, nil).Maybe()
			sourcePriceRegReader.On("GetTokenPrices", ctx, mock.Anything).Return(
				[]cciptypes.TokenPriceUpdate{{TokenPrice: cciptypes.TokenPrice{Token: ccipcalc.HexToAddress("0x1"), Value: big.NewInt(123)}, TimestampUnixSec: big.NewInt(time.Now().Unix())}}, nil).Maybe()
			p.destPriceRegistry = destPriceRegReader

			mockOnRampPriceRegistryProvider := ccipdataprovidermocks.NewPriceRegistry(t)
			mockOnRampPriceRegistryProvider.On("NewPriceRegistryReader", ctx, sourcePriceRegistryAddress).Return(sourcePriceRegReader, nil).Maybe()
			p.sourcePriceRegistryProvider = mockOnRampPriceRegistryProvider

			p.commitRootsCache = cache.NewCommitRootsCache(logger.TestLogger(t), time.Minute, time.Minute)
			p.chainHealthcheck = cache.NewChainHealthcheck(p.lggr, mockOnRampReader, commitStoreReader)

			_, err = p.Observation(ctx, types.ReportTimestamp{}, types.Query{})
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestExecutionReportingPlugin_Report(t *testing.T) {
	testCases := []struct {
		name            string
		f               int
		committedSeqNum uint64
		observations    []ccip.ExecutionObservation

		expectingSomeReport bool
		expectedReport      cciptypes.ExecReport
		expectingSomeErr    bool
	}{
		{
			name:            "not enough observations to form consensus",
			f:               5,
			committedSeqNum: 5,
			observations: []ccip.ExecutionObservation{
				{Messages: map[uint64]ccip.MsgData{3: {}, 4: {}}},
				{Messages: map[uint64]ccip.MsgData{3: {}, 4: {}}},
			},
			expectingSomeErr:    false,
			expectingSomeReport: false,
		},
		{
			name:                "zero observations",
			f:                   0,
			committedSeqNum:     5,
			observations:        []ccip.ExecutionObservation{},
			expectingSomeErr:    false,
			expectingSomeReport: false,
		},
	}

	ctx := testutils.Context(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := ExecutionReportingPlugin{}
			p.lggr = logger.TestLogger(t)
			p.F = tc.f

			p.commitStoreReader = ccipdatamocks.NewCommitStoreReader(t)
			chainHealthcheck := ccipcachemocks.NewChainHealthcheck(t)
			chainHealthcheck.On("IsHealthy", ctx).Return(true, nil)
			p.chainHealthcheck = chainHealthcheck

			observations := make([]types.AttributedObservation, len(tc.observations))
			for i := range observations {
				b, err := json.Marshal(tc.observations[i])
				assert.NoError(t, err)
				observations[i] = types.AttributedObservation{Observation: b, Observer: commontypes.OracleID(i + 1)}
			}

			_, _, err := p.Report(ctx, types.ReportTimestamp{}, types.Query{}, observations)
			if tc.expectingSomeErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}

}

func TestExecutionReportingPlugin_ShouldAcceptFinalizedReport(t *testing.T) {
	msg := cciptypes.EVM2EVMMessage{
		SequenceNumber: 12,
		FeeTokenAmount: big.NewInt(1e9),
		Sender:         cciptypes.Address(utils.RandomAddress().String()),
		Nonce:          1,
		GasLimit:       big.NewInt(1),
		Strict:         false,
		Receiver:       cciptypes.Address(utils.RandomAddress().String()),
		Data:           nil,
		TokenAmounts:   nil,
		FeeToken:       cciptypes.Address(utils.RandomAddress().String()),
		MessageID:      [32]byte{},
	}
	report := cciptypes.ExecReport{
		Messages:          []cciptypes.EVM2EVMMessage{msg},
		OffchainTokenData: [][][]byte{{}},
		Proofs:            [][32]byte{{}},
		ProofFlagBits:     big.NewInt(1),
	}

	encodedReport := encodeExecutionReport(t, report)
	mockOffRampReader := ccipdatamocks.NewOffRampReader(t)
	mockOffRampReader.On("DecodeExecutionReport", mock.Anything, encodedReport).Return(report, nil)

	chainHealthcheck := ccipcachemocks.NewChainHealthcheck(t)
	chainHealthcheck.On("IsHealthy", mock.Anything).Return(true, nil)

	plugin := ExecutionReportingPlugin{
		offRampReader:    mockOffRampReader,
		lggr:             logger.TestLogger(t),
		inflightReports:  newInflightExecReportsContainer(1 * time.Hour),
		chainHealthcheck: chainHealthcheck,
		metricsCollector: ccip.NoopMetricsCollector,
	}

	mockedExecState := mockOffRampReader.On("GetExecutionState", mock.Anything, uint64(12)).Return(uint8(cciptypes.ExecutionStateUntouched), nil).Once()

	should, err := plugin.ShouldAcceptFinalizedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, true, should)

	mockedExecState.Return(uint8(cciptypes.ExecutionStateSuccess), nil).Once()

	should, err = plugin.ShouldAcceptFinalizedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, false, should)
}

func TestExecutionReportingPlugin_ShouldTransmitAcceptedReport(t *testing.T) {
	msg := cciptypes.EVM2EVMMessage{
		SequenceNumber: 12,
		FeeTokenAmount: big.NewInt(1e9),
		Sender:         cciptypes.Address(utils.RandomAddress().String()),
		Nonce:          1,
		GasLimit:       big.NewInt(1),
		Strict:         false,
		Receiver:       cciptypes.Address(utils.RandomAddress().String()),
		Data:           nil,
		TokenAmounts:   nil,
		FeeToken:       cciptypes.Address(utils.RandomAddress().String()),
		MessageID:      [32]byte{},
	}
	report := cciptypes.ExecReport{
		Messages:          []cciptypes.EVM2EVMMessage{msg},
		OffchainTokenData: [][][]byte{{}},
		Proofs:            [][32]byte{{}},
		ProofFlagBits:     big.NewInt(1),
	}
	encodedReport := encodeExecutionReport(t, report)

	mockCommitStoreReader := ccipdatamocks.NewCommitStoreReader(t)
	mockOffRampReader := ccipdatamocks.NewOffRampReader(t)
	mockOffRampReader.On("DecodeExecutionReport", mock.Anything, encodedReport).Return(report, nil)
	mockedExecState := mockOffRampReader.On("GetExecutionState", mock.Anything, uint64(12)).Return(uint8(cciptypes.ExecutionStateUntouched), nil).Once()

	chainHealthcheck := ccipcachemocks.NewChainHealthcheck(t)
	chainHealthcheck.On("IsHealthy", mock.Anything).Return(true, nil)

	plugin := ExecutionReportingPlugin{
		commitStoreReader: mockCommitStoreReader,
		offRampReader:     mockOffRampReader,
		lggr:              logger.TestLogger(t),
		inflightReports:   newInflightExecReportsContainer(1 * time.Hour),
		chainHealthcheck:  chainHealthcheck,
	}

	should, err := plugin.ShouldTransmitAcceptedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, true, should)

	mockedExecState.Return(uint8(cciptypes.ExecutionStateFailure), nil).Once()
	should, err = plugin.ShouldTransmitAcceptedReport(testutils.Context(t), ocrtypes.ReportTimestamp{}, encodedReport)
	require.NoError(t, err)
	assert.Equal(t, false, should)
}

func TestExecutionReportingPlugin_buildReport(t *testing.T) {
	ctx := testutils.Context(t)

	const numMessages = 100
	const tokensPerMessage = 20
	const bytesPerMessage = 1000

	executionReport := generateExecutionReport(t, numMessages, tokensPerMessage, bytesPerMessage)
	encodedReport := encodeExecutionReport(t, executionReport)
	// ensure "naive" full report would be bigger than limit
	assert.Greater(t, len(encodedReport), MaxExecutionReportLength, "full execution report length")

	observations := make([]ccip.ObservedMessage, len(executionReport.Messages))
	for i, msg := range executionReport.Messages {
		observations[i] = ccip.NewObservedMessage(msg.SequenceNumber, executionReport.OffchainTokenData[i])
	}

	// ensure that buildReport should cap the built report to fit in MaxExecutionReportLength
	p := &ExecutionReportingPlugin{}
	p.lggr = logger.TestLogger(t)

	commitStore := ccipdatamocks.NewCommitStoreReader(t)
	commitStore.On("VerifyExecutionReport", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
	commitStore.On("GetExpectedNextSequenceNumber", mock.Anything).
		Return(executionReport.Messages[len(executionReport.Messages)-1].SequenceNumber+1, nil)
	commitStore.On("GetCommitReportMatchingSeqNum", ctx, observations[0].SeqNr, 0).
		Return([]cciptypes.CommitStoreReportWithTxMeta{
			{
				CommitStoreReport: cciptypes.CommitStoreReport{
					Interval: cciptypes.CommitStoreInterval{
						Min: observations[0].SeqNr,
						Max: observations[len(observations)-1].SeqNr,
					},
				},
			},
		}, nil)
	p.metricsCollector = ccip.NoopMetricsCollector
	p.commitStoreReader = commitStore

	lp := lpMocks.NewLogPoller(t)
	offRampReader, err := v1_0_0.NewOffRamp(logger.TestLogger(t), utils.RandomAddress(), nil, lp, nil, nil)
	assert.NoError(t, err)
	p.offRampReader = offRampReader

	sendReqs := make([]cciptypes.EVM2EVMMessageWithTxMeta, len(observations))
	sourceReader := ccipdatamocks.NewOnRampReader(t)
	for i := range observations {
		msg := cciptypes.EVM2EVMMessage{
			SourceChainSelector: math.MaxUint64,
			SequenceNumber:      uint64(i + 1),
			FeeTokenAmount:      big.NewInt(math.MaxInt64),
			Sender:              cciptypes.Address(utils.RandomAddress().String()),
			Nonce:               math.MaxUint64,
			GasLimit:            big.NewInt(math.MaxInt64),
			Strict:              false,
			Receiver:            cciptypes.Address(utils.RandomAddress().String()),
			Data:                bytes.Repeat([]byte{0}, bytesPerMessage),
			TokenAmounts:        nil,
			FeeToken:            cciptypes.Address(utils.RandomAddress().String()),
			MessageID:           [32]byte{12},
		}
		sendReqs[i] = cciptypes.EVM2EVMMessageWithTxMeta{EVM2EVMMessage: msg}
	}
	sourceReader.On("GetSendRequestsBetweenSeqNums",
		ctx, observations[0].SeqNr, observations[len(observations)-1].SeqNr, false).Return(sendReqs, nil)
	p.onRampReader = sourceReader

	execReport, err := p.buildReport(ctx, p.lggr, observations)
	assert.NoError(t, err)
	assert.LessOrEqual(t, len(execReport), MaxExecutionReportLength, "built execution report length")
}

func TestExecutionReportingPlugin_buildBatch(t *testing.T) {
	offRamp, _ := testhelpers.NewFakeOffRamp(t)
	lggr := logger.TestLogger(t)

	sender1 := ccipcalc.HexToAddress("0xa")
	destNative := ccipcalc.HexToAddress("0xb")
	srcNative := ccipcalc.HexToAddress("0xc")

	msg1 := cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
		EVM2EVMMessage: cciptypes.EVM2EVMMessage{
			SequenceNumber: 1,
			FeeTokenAmount: big.NewInt(1e9),
			Sender:         sender1,
			Nonce:          1,
			GasLimit:       big.NewInt(1),
			Strict:         false,
			Receiver:       "",
			Data:           nil,
			TokenAmounts:   nil,
			FeeToken:       srcNative,
			MessageID:      [32]byte{},
		},
		BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
	}

	msg2 := msg1
	msg2.Executed = true

	msg3 := msg1
	msg3.Executed = true
	msg3.Finalized = true

	msg4 := msg1
	msg4.TokenAmounts = []cciptypes.TokenAmount{
		{Token: srcNative, Amount: big.NewInt(100)},
	}

	msg5 := msg4
	msg5.SequenceNumber = msg5.SequenceNumber + 1
	msg5.Nonce = msg5.Nonce + 1

	var tt = []struct {
		name                     string
		reqs                     []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta
		inflight                 *big.Int
		tokenLimit, destGasPrice *big.Int
		srcPrices, dstPrices     map[cciptypes.Address]*big.Int
		offRampNoncesBySender    map[cciptypes.Address]uint64
		srcToDestTokens          map[cciptypes.Address]cciptypes.Address
		expectedSeqNrs           []ccip.ObservedMessage
		expectedStates           []messageExecStatus
	}{
		{
			name:                  "single message no tokens",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(1)}},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, AddedToBatch)},
		},
		{
			name:                  "executed non finalized messages should be skipped",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg2},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg2.SequenceNumber, msg2.MessageID, AlreadyExecuted)},
		},
		{
			name:                  "finalized executed log",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg3},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg3.SequenceNumber, msg3.MessageID, AlreadyExecuted)},
		},
		{
			name:                  "dst token price does not exist",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, TokenNotInDestTokenPrices)},
		},
		{
			name:                  "src token price does not exist",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, TokenNotInSrcTokenPrices)},
		},
		{
			name:         "message with tokens is not executed if limit is reached",
			reqs:         []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg4},
			inflight:     big.NewInt(0),
			tokenLimit:   big.NewInt(2),
			destGasPrice: big.NewInt(10),
			srcPrices:    map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1e18)},
			dstPrices:    map[cciptypes.Address]*big.Int{destNative: big.NewInt(1e18)},
			srcToDestTokens: map[cciptypes.Address]cciptypes.Address{
				srcNative: destNative,
			},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg4.SequenceNumber, msg4.MessageID, AggregateTokenLimitExceeded)},
		},
		{
			name:         "message with tokens is not executed if limit is reached when inflight is full",
			reqs:         []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg5},
			inflight:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100)),
			tokenLimit:   big.NewInt(19),
			destGasPrice: big.NewInt(10),
			srcPrices:    map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1e18)},
			dstPrices:    map[cciptypes.Address]*big.Int{destNative: big.NewInt(1e18)},
			srcToDestTokens: map[cciptypes.Address]cciptypes.Address{
				srcNative: destNative,
			},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 1},
			expectedStates:        []messageExecStatus{newMessageExecState(msg5.SequenceNumber, msg5.MessageID, AggregateTokenLimitExceeded)},
		},
		{
			name:                  "skip when nonce doesn't match chain value",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 123},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, InvalidNonce)},
		},
		{
			name:                  "skip when nonce not found",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, MissingNonce)},
		},
		{
			name: "skip when batch gas limit is reached",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 10,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          1,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 11,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          2,
						GasLimit:       big.NewInt(math.MaxInt64),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 12,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          3,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
			},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(10)}},
			expectedStates: []messageExecStatus{
				newMessageExecState(10, [32]byte{}, AddedToBatch),
				newMessageExecState(11, [32]byte{}, InsufficientRemainingBatchGas),
				newMessageExecState(12, [32]byte{}, InvalidNonce),
			},
		},
		{
			name: "some messages skipped after hitting max batch data len",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 10,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          1,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 11,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          2,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, MaxDataLenPerBatch-500), // skipped from batch
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 12,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          3,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
			},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(10)}},
			expectedStates: []messageExecStatus{
				newMessageExecState(10, [32]byte{}, AddedToBatch),
				newMessageExecState(11, [32]byte{}, InsufficientRemainingBatchDataLength),
				newMessageExecState(12, [32]byte{}, InvalidNonce),
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			offRamp.SetSenderNonces(tc.offRampNoncesBySender)

			gasPriceEstimator := prices.NewMockGasPriceEstimatorExec(t)
			if tc.expectedSeqNrs != nil {
				gasPriceEstimator.On("EstimateMsgCostUSD", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(big.NewInt(0), nil)
			}

			// Mock calls to reader.
			mockOffRampReader := ccipdatamocks.NewOffRampReader(t)
			mockOffRampReader.On("GetSendersNonce", mock.Anything, mock.Anything).Return(tc.offRampNoncesBySender, nil).Maybe()

			plugin := ExecutionReportingPlugin{
				tokenDataWorker:   tokendata.NewBackgroundWorker(map[cciptypes.Address]tokendata.Reader{}, 10, 5*time.Second, time.Hour),
				offRampReader:     mockOffRampReader,
				destWrappedNative: destNative,
				offchainConfig: cciptypes.ExecOffchainConfig{
					DestOptimisticConfirmations: 1,
					BatchGasLimit:               300_000,
					RelativeBoostPerWaitHour:    1,
				},
				lggr:              logger.TestLogger(t),
				gasPriceEstimator: gasPriceEstimator,
			}

			seqNrs, execStates := plugin.buildBatch(
				context.Background(),
				lggr,
				commitReportWithSendRequests{sendRequestsWithMeta: tc.reqs},
				tc.inflight,
				tc.tokenLimit,
				tc.srcPrices,
				tc.dstPrices,
				tc.destGasPrice,
				tc.srcToDestTokens,
			)
			if tc.expectedSeqNrs == nil {
				assert.Len(t, seqNrs, 0)
			} else {
				assert.Equal(t, tc.expectedSeqNrs, seqNrs)
			}

			if tc.expectedStates == nil {
				assert.Len(t, execStates, 0)
			} else {
				assert.Equal(t, tc.expectedStates, execStates)
			}
		})
	}
}

func TestExecutionReportingPlugin_getReportsWithSendRequests(t *testing.T) {
	testCases := []struct {
		name                string
		reports             []cciptypes.CommitStoreReport
		expQueryMin         uint64 // expected min/max used in the query to get ccipevents
		expQueryMax         uint64
		onchainEvents       []cciptypes.EVM2EVMMessageWithTxMeta
		destExecutedSeqNums []uint64

		expReports []commitReportWithSendRequests
		expErr     bool
	}{
		{
			name:       "no reports",
			reports:    nil,
			expReports: nil,
			expErr:     false,
		},
		{
			name: "two reports happy flow",
			reports: []cciptypes.CommitStoreReport{
				{
					Interval:   cciptypes.CommitStoreInterval{Min: 1, Max: 2},
					MerkleRoot: [32]byte{100},
				},
				{
					Interval:   cciptypes.CommitStoreInterval{Min: 3, Max: 3},
					MerkleRoot: [32]byte{200},
				},
			},
			expQueryMin: 1,
			expQueryMax: 3,
			onchainEvents: []cciptypes.EVM2EVMMessageWithTxMeta{
				{EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 1}},
				{EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 2}},
				{EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 3}},
			},
			destExecutedSeqNums: []uint64{1},
			expReports: []commitReportWithSendRequests{
				{
					commitReport: cciptypes.CommitStoreReport{
						Interval:   cciptypes.CommitStoreInterval{Min: 1, Max: 2},
						MerkleRoot: [32]byte{100},
					},
					sendRequestsWithMeta: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
						{
							EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 1},
							Executed:       true,
							Finalized:      true,
						},
						{
							EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 2},
							Executed:       false,
							Finalized:      false,
						},
					},
				},
				{
					commitReport: cciptypes.CommitStoreReport{
						Interval:   cciptypes.CommitStoreInterval{Min: 3, Max: 3},
						MerkleRoot: [32]byte{200},
					},
					sendRequestsWithMeta: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
						{
							EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 3},
							Executed:       false,
							Finalized:      false,
						},
					},
				},
			},
			expErr: false,
		},
	}

	ctx := testutils.Context(t)
	lggr := logger.TestLogger(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := &ExecutionReportingPlugin{}
			p.lggr = lggr

			offRampReader := ccipdatamocks.NewOffRampReader(t)
			p.offRampReader = offRampReader

			sourceReader := ccipdatamocks.NewOnRampReader(t)
			sourceReader.On("GetSendRequestsBetweenSeqNums", ctx, tc.expQueryMin, tc.expQueryMax, false).
				Return(tc.onchainEvents, nil).Maybe()
			p.onRampReader = sourceReader

			finalized := make(map[uint64]bool)
			for _, r := range tc.expReports {
				for _, s := range r.sendRequestsWithMeta {
					finalized[s.SequenceNumber] = s.Finalized
				}
			}

			var executedEvents []cciptypes.ExecutionStateChangedWithTxMeta
			for _, executedSeqNum := range tc.destExecutedSeqNums {
				executedEvents = append(executedEvents, cciptypes.ExecutionStateChangedWithTxMeta{
					ExecutionStateChanged: cciptypes.ExecutionStateChanged{
						SequenceNumber: executedSeqNum,
						Finalized:      finalized[executedSeqNum],
					},
				})
			}
			offRampReader.On("GetExecutionStateChangesBetweenSeqNums", ctx, tc.expQueryMin, tc.expQueryMax, 0).Return(executedEvents, nil).Maybe()

			populatedReports, err := p.getReportsWithSendRequests(ctx, tc.reports)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, len(tc.expReports), len(populatedReports))
			for i, expReport := range tc.expReports {
				assert.Equal(t, len(expReport.sendRequestsWithMeta), len(populatedReports[i].sendRequestsWithMeta))
				for j, expReq := range expReport.sendRequestsWithMeta {
					assert.Equal(t, expReq.Executed, populatedReports[i].sendRequestsWithMeta[j].Executed)
					assert.Equal(t, expReq.Finalized, populatedReports[i].sendRequestsWithMeta[j].Finalized)
					assert.Equal(t, expReq.SequenceNumber, populatedReports[i].sendRequestsWithMeta[j].SequenceNumber)
				}
			}
		})
	}
}

type delayedTokenDataWorker struct {
	delay time.Duration
	tokendata.Worker
}

func (m delayedTokenDataWorker) GetMsgTokenData(ctx context.Context, msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([][]byte, error) {
	time.Sleep(m.delay)
	return nil, ctx.Err()
}

func TestExecutionReportingPlugin_getTokenDataWithCappedLatency(t *testing.T) {
	testCases := []struct {
		name               string
		allowedWaitingTime time.Duration
		workerLatency      time.Duration
		expErr             bool
	}{
		{
			name:               "happy flow",
			allowedWaitingTime: 10 * time.Millisecond,
			workerLatency:      time.Nanosecond,
			expErr:             false,
		},
		{
			name:               "worker takes long to reply",
			allowedWaitingTime: 10 * time.Millisecond,
			workerLatency:      20 * time.Millisecond,
			expErr:             true,
		},
	}

	ctx := testutils.Context(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := &ExecutionReportingPlugin{}
			p.tokenDataWorker = delayedTokenDataWorker{delay: tc.workerLatency}

			msg := cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				EVM2EVMMessage: cciptypes.EVM2EVMMessage{TokenAmounts: make([]cciptypes.TokenAmount, 1)},
			}

			_, _, err := p.getTokenDataWithTimeout(ctx, msg, tc.allowedWaitingTime)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_calculateObservedMessagesConsensus(t *testing.T) {
	type args struct {
		observations []ccip.ExecutionObservation
		f            int
	}
	tests := []struct {
		name string
		args args
		want []ccip.ObservedMessage
	}{
		{
			name: "no observations",
			args: args{
				observations: nil,
				f:            0,
			},
			want: []ccip.ObservedMessage{},
		},
		{
			name: "common path",
			args: args{
				observations: []ccip.ExecutionObservation{
					{
						Messages: map[uint64]ccip.MsgData{
							1: {TokenData: [][]byte{{0x1}, {0x1}, {0x1}}},
							2: {TokenData: [][]byte{{0x2}, {0x2}, {0x2}}},
						},
					},
					{
						Messages: map[uint64]ccip.MsgData{
							1: {TokenData: [][]byte{{0x1}, {0x1}, {0xff}}}, // different token data - should not be picked
							2: {TokenData: [][]byte{{0x2}, {0x2}, {0x2}}},
							3: {TokenData: [][]byte{{0x3}, {0x3}, {0x3}}},
						},
					},
					{
						Messages: map[uint64]ccip.MsgData{
							1: {TokenData: [][]byte{{0x1}, {0x1}, {0x1}}},
							2: {TokenData: [][]byte{{0x2}, {0x2}, {0x2}}},
						},
					},
				},
				f: 1,
			},
			want: []ccip.ObservedMessage{
				{SeqNr: 1, MsgData: ccip.MsgData{TokenData: [][]byte{{0x1}, {0x1}, {0x1}}}},
				{SeqNr: 2, MsgData: ccip.MsgData{TokenData: [][]byte{{0x2}, {0x2}, {0x2}}}},
			},
		},
		{
			name: "similar token data",
			args: args{
				observations: []ccip.ExecutionObservation{
					{
						Messages: map[uint64]ccip.MsgData{
							1: {TokenData: [][]byte{{0x1}, {0x1}, {0x1}}},
						},
					},
					{
						Messages: map[uint64]ccip.MsgData{
							1: {TokenData: [][]byte{{0x1}, {0x1, 0x1}}},
						},
					},
					{
						Messages: map[uint64]ccip.MsgData{
							1: {TokenData: [][]byte{{0x1}, {0x1, 0x1}}},
						},
					},
				},
				f: 1,
			},
			want: []ccip.ObservedMessage{
				{SeqNr: 1, MsgData: ccip.MsgData{TokenData: [][]byte{{0x1}, {0x1, 0x1}}}},
			},
		},
		{
			name: "results should be deterministic",
			args: args{
				observations: []ccip.ExecutionObservation{
					{Messages: map[uint64]ccip.MsgData{1: {TokenData: [][]byte{{0x2}}}}},
					{Messages: map[uint64]ccip.MsgData{1: {TokenData: [][]byte{{0x2}}}}},
					{Messages: map[uint64]ccip.MsgData{1: {TokenData: [][]byte{{0x1}}}}},
					{Messages: map[uint64]ccip.MsgData{1: {TokenData: [][]byte{{0x3}}}}},
					{Messages: map[uint64]ccip.MsgData{1: {TokenData: [][]byte{{0x3}}}}},
					{Messages: map[uint64]ccip.MsgData{1: {TokenData: [][]byte{{0x1}}}}},
				},
				f: 1,
			},
			want: []ccip.ObservedMessage{
				{SeqNr: 1, MsgData: ccip.MsgData{TokenData: [][]byte{{0x3}}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := calculateObservedMessagesConsensus(
				tt.args.observations,
				tt.args.f,
			)
			assert.NoError(t, err)
			sort.Slice(res, func(i, j int) bool {
				return res[i].SeqNr < res[j].SeqNr
			})
			assert.Equalf(t, tt.want, res, "calculateObservedMessagesConsensus(%v, %v)", tt.args.observations, tt.args.f)
		})
	}
}

func Test_getTokensPrices(t *testing.T) {
	tk1 := ccipcalc.HexToAddress("1")
	tk2 := ccipcalc.HexToAddress("2")
	tk3 := ccipcalc.HexToAddress("3")

	testCases := []struct {
		name      string
		feeTokens []cciptypes.Address
		tokens    []cciptypes.Address
		retPrices []cciptypes.TokenPriceUpdate
		expPrices map[cciptypes.Address]*big.Int
		expErr    bool
	}{
		{
			name:      "base",
			feeTokens: []cciptypes.Address{tk1, tk2},
			tokens:    []cciptypes.Address{tk3},
			retPrices: []cciptypes.TokenPriceUpdate{
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(10)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(20)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(30)}},
			},
			expPrices: map[cciptypes.Address]*big.Int{
				tk1: big.NewInt(10),
				tk2: big.NewInt(20),
				tk3: big.NewInt(30),
			},
			expErr: false,
		},
		{
			name:      "token is both fee token and normal token",
			feeTokens: []cciptypes.Address{tk1, tk2},
			tokens:    []cciptypes.Address{tk3, tk1},
			retPrices: []cciptypes.TokenPriceUpdate{
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(10)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(20)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(30)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(10)}},
			},
			expPrices: map[cciptypes.Address]*big.Int{
				tk1: big.NewInt(10),
				tk2: big.NewInt(20),
				tk3: big.NewInt(30),
			},
			expErr: false,
		},
		{
			name:      "token is both fee token and normal token and price registry gave different price",
			feeTokens: []cciptypes.Address{tk1, tk2},
			tokens:    []cciptypes.Address{tk3, tk1},
			retPrices: []cciptypes.TokenPriceUpdate{
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(10)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(20)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(30)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(1000)}},
			},
			expErr: true,
		},
		{
			name:      "contract returns less prices than requested",
			feeTokens: []cciptypes.Address{tk1, tk2},
			tokens:    []cciptypes.Address{tk3},
			retPrices: []cciptypes.TokenPriceUpdate{
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(10)}},
				{TokenPrice: cciptypes.TokenPrice{Value: big.NewInt(20)}},
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			priceReg := ccipdatamocks.NewPriceRegistryReader(t)
			priceReg.On("GetTokenPrices", mock.Anything, mock.Anything).Return(tc.retPrices, nil)
			priceReg.On("Address", mock.Anything).Return(cciptypes.Address(utils.RandomAddress().String()), nil).Maybe()

			tokenPrices, err := getTokensPrices(context.Background(), priceReg, append(tc.feeTokens, tc.tokens...))
			if tc.expErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			for tk, price := range tc.expPrices {
				assert.Equal(t, price, tokenPrices[tk])
			}
		})
	}
}

func Test_calculateMessageMaxGas(t *testing.T) {
	type args struct {
		gasLimit    *big.Int
		numRequests int
		dataLen     int
		numTokens   int
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "base",
			args:    args{gasLimit: big.NewInt(1000), numRequests: 5, dataLen: 5, numTokens: 2},
			want:    203836,
			wantErr: false,
		},
		{
			name:    "large",
			args:    args{gasLimit: big.NewInt(1000), numRequests: 1000, dataLen: 1000, numTokens: 1000},
			want:    36482676,
			wantErr: false,
		},
		{
			name:    "gas limit overflow",
			args:    args{gasLimit: big.NewInt(0).Mul(big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64))},
			want:    36391540,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateMessageMaxGas(tt.args.gasLimit, tt.args.numRequests, tt.args.dataLen, tt.args.numTokens)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "calculateMessageMaxGas(%v, %v, %v, %v)", tt.args.gasLimit, tt.args.numRequests, tt.args.dataLen, tt.args.numTokens)
		})
	}
}

func Test_inflightAggregates(t *testing.T) {
	const n = 10
	addrs := make([]cciptypes.Address, n)
	tokenAddrs := make([]cciptypes.Address, n)
	for i := range addrs {
		addrs[i] = cciptypes.Address(utils.RandomAddress().String())
		tokenAddrs[i] = cciptypes.Address(utils.RandomAddress().String())
	}
	lggr := logger.TestLogger(t)

	testCases := []struct {
		name            string
		inflight        []InflightInternalExecutionReport
		destTokenPrices map[cciptypes.Address]*big.Int
		sourceToDest    map[cciptypes.Address]cciptypes.Address

		expInflightSeqNrs          mapset.Set[uint64]
		expInflightAggrVal         *big.Int
		expMaxInflightSenderNonces map[cciptypes.Address]uint64
		expInflightTokenAmounts    map[cciptypes.Address]*big.Int
		expErr                     bool
	}{
		{
			name: "base",
			inflight: []InflightInternalExecutionReport{
				{
					messages: []cciptypes.EVM2EVMMessage{
						{
							Sender:         addrs[0],
							SequenceNumber: 100,
							Nonce:          2,
							TokenAmounts: []cciptypes.TokenAmount{
								{Token: tokenAddrs[0], Amount: big.NewInt(1e18)},
								{Token: tokenAddrs[0], Amount: big.NewInt(2e18)},
							},
						},
						{
							Sender:         addrs[0],
							SequenceNumber: 106,
							Nonce:          4,
							TokenAmounts: []cciptypes.TokenAmount{
								{Token: tokenAddrs[0], Amount: big.NewInt(1e18)},
								{Token: tokenAddrs[0], Amount: big.NewInt(5e18)},
								{Token: tokenAddrs[2], Amount: big.NewInt(5e18)},
							},
						},
					},
				},
			},
			destTokenPrices: map[cciptypes.Address]*big.Int{
				tokenAddrs[1]: big.NewInt(1000),
				tokenAddrs[3]: big.NewInt(500),
			},
			sourceToDest: map[cciptypes.Address]cciptypes.Address{
				tokenAddrs[0]: tokenAddrs[1],
				tokenAddrs[2]: tokenAddrs[3],
			},
			expInflightSeqNrs:  mapset.NewSet[uint64](100, 106),
			expInflightAggrVal: big.NewInt(9*1000 + 5*500),
			expMaxInflightSenderNonces: map[cciptypes.Address]uint64{
				addrs[0]: 4,
			},
			expInflightTokenAmounts: map[cciptypes.Address]*big.Int{
				tokenAddrs[0]: big.NewInt(9e18),
				tokenAddrs[2]: big.NewInt(5e18),
			},
			expErr: false,
		},
		{
			name: "missing price should be 0",
			inflight: []InflightInternalExecutionReport{
				{
					messages: []cciptypes.EVM2EVMMessage{
						{
							Sender:         addrs[0],
							SequenceNumber: 100,
							Nonce:          2,
							TokenAmounts: []cciptypes.TokenAmount{
								{Token: tokenAddrs[0], Amount: big.NewInt(1e18)},
							},
						},
					},
				},
			},
			destTokenPrices: map[cciptypes.Address]*big.Int{
				tokenAddrs[3]: big.NewInt(500),
			},
			sourceToDest: map[cciptypes.Address]cciptypes.Address{
				tokenAddrs[2]: tokenAddrs[3],
			},
			expInflightAggrVal: big.NewInt(0),
			expErr:             false,
		},
		{
			name:                       "nothing inflight",
			inflight:                   []InflightInternalExecutionReport{},
			expInflightSeqNrs:          mapset.NewSet[uint64](),
			expInflightAggrVal:         big.NewInt(0),
			expMaxInflightSenderNonces: map[cciptypes.Address]uint64{},
			expInflightTokenAmounts:    map[cciptypes.Address]*big.Int{},
			expErr:                     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inflightAggrVal, err := getInflightAggregateRateLimit(
				lggr,
				tc.inflight,
				tc.destTokenPrices,
				tc.sourceToDest,
			)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.True(t, reflect.DeepEqual(tc.expInflightAggrVal, inflightAggrVal))
		})
	}
}

func Test_commitReportWithSendRequests_validate(t *testing.T) {
	testCases := []struct {
		name           string
		reportInterval cciptypes.CommitStoreInterval
		numReqs        int
		expValid       bool
	}{
		{
			name:           "valid report",
			reportInterval: cciptypes.CommitStoreInterval{Min: 10, Max: 20},
			numReqs:        11,
			expValid:       true,
		},
		{
			name:           "report with one request",
			reportInterval: cciptypes.CommitStoreInterval{Min: 1234, Max: 1234},
			numReqs:        1,
			expValid:       true,
		},
		{
			name:           "request is missing",
			reportInterval: cciptypes.CommitStoreInterval{Min: 1234, Max: 1234},
			numReqs:        0,
			expValid:       false,
		},
		{
			name:           "requests are missing",
			reportInterval: cciptypes.CommitStoreInterval{Min: 1, Max: 10},
			numReqs:        5,
			expValid:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rep := commitReportWithSendRequests{
				commitReport: cciptypes.CommitStoreReport{
					Interval: tc.reportInterval,
				},
				sendRequestsWithMeta: make([]cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta, tc.numReqs),
			}
			err := rep.validate()
			isValid := err == nil
			assert.Equal(t, tc.expValid, isValid)
		})
	}
}

func Test_commitReportWithSendRequests_allRequestsAreExecutedAndFinalized(t *testing.T) {
	testCases := []struct {
		name   string
		reqs   []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta
		expRes bool
	}{
		{
			name: "all requests executed and finalized",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{Executed: true, Finalized: true},
				{Executed: true, Finalized: true},
				{Executed: true, Finalized: true},
			},
			expRes: true,
		},
		{
			name:   "true when there are zero requests",
			reqs:   []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{},
			expRes: true,
		},
		{
			name: "some request not executed",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{Executed: true, Finalized: true},
				{Executed: true, Finalized: true},
				{Executed: false, Finalized: true},
			},
			expRes: false,
		},
		{
			name: "some request not finalized",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{Executed: true, Finalized: true},
				{Executed: true, Finalized: true},
				{Executed: true, Finalized: false},
			},
			expRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rep := commitReportWithSendRequests{sendRequestsWithMeta: tc.reqs}
			res := rep.allRequestsAreExecutedAndFinalized()
			assert.Equal(t, tc.expRes, res)
		})
	}
}

func Test_commitReportWithSendRequests_sendReqFits(t *testing.T) {
	testCases := []struct {
		name   string
		req    cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta
		report cciptypes.CommitStoreReport
		expRes bool
	}{
		{
			name: "all requests executed and finalized",
			req: cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 1},
			},
			report: cciptypes.CommitStoreReport{
				Interval: cciptypes.CommitStoreInterval{Min: 1, Max: 10},
			},
			expRes: true,
		},
		{
			name: "all requests executed and finalized",
			req: cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 10},
			},
			report: cciptypes.CommitStoreReport{
				Interval: cciptypes.CommitStoreInterval{Min: 1, Max: 10},
			},
			expRes: true,
		},
		{
			name: "all requests executed and finalized",
			req: cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 11},
			},
			report: cciptypes.CommitStoreReport{
				Interval: cciptypes.CommitStoreInterval{Min: 1, Max: 10},
			},
			expRes: false,
		},
		{
			name: "all requests executed and finalized",
			req: cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: 10},
			},
			report: cciptypes.CommitStoreReport{
				Interval: cciptypes.CommitStoreInterval{Min: 10, Max: 10},
			},
			expRes: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &commitReportWithSendRequests{commitReport: tc.report}
			assert.Equal(t, tc.expRes, r.sendReqFits(tc.req))
		})
	}
}

// generateExecutionReport generates an execution report that can be used in tests
func generateExecutionReport(t *testing.T, numMsgs, tokensPerMsg, bytesPerMsg int) cciptypes.ExecReport {
	messages := make([]cciptypes.EVM2EVMMessage, numMsgs)

	randAddr := func() cciptypes.Address {
		return cciptypes.Address(utils.RandomAddress().String())
	}

	offChainTokenData := make([][][]byte, numMsgs)
	for i := range messages {
		tokenAmounts := make([]cciptypes.TokenAmount, tokensPerMsg)
		for j := range tokenAmounts {
			tokenAmounts[j] = cciptypes.TokenAmount{
				Token:  randAddr(),
				Amount: big.NewInt(math.MaxInt64),
			}
		}

		messages[i] = cciptypes.EVM2EVMMessage{
			SourceChainSelector: rand.Uint64(),
			SequenceNumber:      uint64(i + 1),
			FeeTokenAmount:      big.NewInt(rand.Int64()),
			Sender:              randAddr(),
			Nonce:               rand.Uint64(),
			GasLimit:            big.NewInt(rand.Int64()),
			Strict:              false,
			Receiver:            randAddr(),
			Data:                bytes.Repeat([]byte{1}, bytesPerMsg),
			TokenAmounts:        tokenAmounts,
			FeeToken:            randAddr(),
			MessageID:           utils.RandomBytes32(),
		}

		data := []byte(`{"foo": "bar"}`)
		offChainTokenData[i] = [][]byte{data, data, data}
	}

	return cciptypes.ExecReport{
		Messages:          messages,
		OffchainTokenData: offChainTokenData,
		Proofs:            make([][32]byte, numMsgs),
		ProofFlagBits:     big.NewInt(rand.Int64()),
	}
}

func Test_selectReportsToFillBatch(t *testing.T) {
	reports := []cciptypes.CommitStoreReport{
		{Interval: cciptypes.CommitStoreInterval{Min: 1, Max: 10}},
		{Interval: cciptypes.CommitStoreInterval{Min: 11, Max: 20}},
		{Interval: cciptypes.CommitStoreInterval{Min: 21, Max: 25}},
		{Interval: cciptypes.CommitStoreInterval{Min: 26, Max: math.MaxUint64}},
	}

	tests := []struct {
		name            string
		step            uint64
		numberOfBatches int
	}{
		{
			name:            "pick all at once when step size is high",
			step:            100,
			numberOfBatches: 1,
		},
		{
			name:            "pick one by one when step size is 1",
			step:            1,
			numberOfBatches: 4,
		},
		{
			name:            "pick two when step size doesn't match report",
			step:            15,
			numberOfBatches: 2,
		},
		{
			name:            "pick one by one when step size is smaller then reports",
			step:            4,
			numberOfBatches: 4,
		},
		{
			name:            "batch some reports together",
			step:            7,
			numberOfBatches: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var unexpiredReportsBatches [][]cciptypes.CommitStoreReport
			for i := 0; i < len(reports); {
				unexpiredReports, step := selectReportsToFillBatch(reports[i:], tt.step)
				unexpiredReportsBatches = append(unexpiredReportsBatches, unexpiredReports)
				i += step
			}
			assert.Len(t, unexpiredReportsBatches, tt.numberOfBatches)

			var flatten []cciptypes.CommitStoreReport
			for _, r := range unexpiredReportsBatches {
				flatten = append(flatten, r...)
			}
			assert.Len(t, flatten, len(reports))
			assert.Equal(t, reports, flatten)
		})
	}
}

func Test_prepareTokenExecData(t *testing.T) {
	ctx := testutils.Context(t)

	weth := cciptypes.Address(utils.RandomAddress().String())
	wavax := cciptypes.Address(utils.RandomAddress().String())
	link := cciptypes.Address(utils.RandomAddress().String())
	usdc := cciptypes.Address(utils.RandomAddress().String())

	wethPriceUpdate := cciptypes.TokenPriceUpdate{TokenPrice: cciptypes.TokenPrice{Token: weth, Value: big.NewInt(2e18)}}
	wavaxPriceUpdate := cciptypes.TokenPriceUpdate{TokenPrice: cciptypes.TokenPrice{Token: wavax, Value: big.NewInt(3e18)}}
	linkPriceUpdate := cciptypes.TokenPriceUpdate{TokenPrice: cciptypes.TokenPrice{Token: link, Value: big.NewInt(4e18)}}
	usdcPriceUpdate := cciptypes.TokenPriceUpdate{TokenPrice: cciptypes.TokenPrice{Token: usdc, Value: big.NewInt(5e18)}}

	tokenPrices := map[cciptypes.Address]cciptypes.TokenPriceUpdate{weth: wethPriceUpdate, wavax: wavaxPriceUpdate, link: linkPriceUpdate, usdc: usdcPriceUpdate}

	tests := []struct {
		name               string
		sourceFeeTokens    []cciptypes.Address
		sourceFeeTokensErr error
		destTokens         []cciptypes.Address
		destTokensErr      error
		destFeeTokens      []cciptypes.Address
		destFeeTokensErr   error
		sourcePrices       []cciptypes.TokenPriceUpdate
		destPrices         []cciptypes.TokenPriceUpdate
	}{
		{
			name:         "only native token",
			sourcePrices: []cciptypes.TokenPriceUpdate{wethPriceUpdate},
			destPrices:   []cciptypes.TokenPriceUpdate{wavaxPriceUpdate},
		},
		{
			name:          "additional dest fee token",
			destFeeTokens: []cciptypes.Address{link},
			sourcePrices:  []cciptypes.TokenPriceUpdate{wethPriceUpdate},
			destPrices:    []cciptypes.TokenPriceUpdate{linkPriceUpdate, wavaxPriceUpdate},
		},
		{
			name:         "dest tokens",
			destTokens:   []cciptypes.Address{link, usdc},
			sourcePrices: []cciptypes.TokenPriceUpdate{wethPriceUpdate},
			destPrices:   []cciptypes.TokenPriceUpdate{linkPriceUpdate, usdcPriceUpdate, wavaxPriceUpdate},
		},
		{
			name:            "source fee tokens",
			sourceFeeTokens: []cciptypes.Address{usdc},
			sourcePrices:    []cciptypes.TokenPriceUpdate{usdcPriceUpdate, wethPriceUpdate},
			destPrices:      []cciptypes.TokenPriceUpdate{wavaxPriceUpdate},
		},
		{
			name:            "source, dest and fee tokens",
			sourceFeeTokens: []cciptypes.Address{usdc},
			destTokens:      []cciptypes.Address{link},
			destFeeTokens:   []cciptypes.Address{usdc},
			sourcePrices:    []cciptypes.TokenPriceUpdate{usdcPriceUpdate, wethPriceUpdate},
			destPrices:      []cciptypes.TokenPriceUpdate{usdcPriceUpdate, linkPriceUpdate, wavaxPriceUpdate},
		},
		{
			name:            "source, dest and fee tokens with duplicates",
			sourceFeeTokens: []cciptypes.Address{link, weth},
			destTokens:      []cciptypes.Address{link, wavax},
			destFeeTokens:   []cciptypes.Address{link, wavax},
			sourcePrices:    []cciptypes.TokenPriceUpdate{linkPriceUpdate, wethPriceUpdate},
			destPrices:      []cciptypes.TokenPriceUpdate{linkPriceUpdate, wavaxPriceUpdate},
		},
		{
			name:               "everything fails when source fails",
			sourceFeeTokensErr: errors.New("source error"),
		},
		{
			name:             "everything fails when dest fee fails",
			destFeeTokensErr: errors.New("dest fee error"),
		},
		{
			name:          "everything fails when dest  fails",
			destTokensErr: errors.New("dest error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			onrampReader := ccipdatamocks.NewOnRampReader(t)
			offrampReader := ccipdatamocks.NewOffRampReader(t)
			sourcePriceRegistry := ccipdatamocks.NewPriceRegistryReader(t)
			destPriceRegistry := ccipdatamocks.NewPriceRegistryReader(t)
			gasPriceEstimator := prices.NewMockGasPriceEstimatorExec(t)
			sourcePriceRegistryProvider := ccipdataprovidermocks.NewPriceRegistry(t)

			sourcePriceRegistryAddress := cciptypes.Address(utils.RandomAddress().String())
			onrampReader.On("SourcePriceRegistryAddress", ctx).Return(sourcePriceRegistryAddress, nil).Maybe()
			offrampReader.On("CurrentRateLimiterState", ctx).Return(cciptypes.TokenBucketRateLimit{}, nil).Maybe()
			offrampReader.On("GetSourceToDestTokensMapping", ctx).Return(map[cciptypes.Address]cciptypes.Address{}, nil).Maybe()
			gasPriceEstimator.On("GetGasPrice", ctx).Return(big.NewInt(1e9), nil).Maybe()

			offrampReader.On("GetTokens", ctx).Return(cciptypes.OffRampTokens{DestinationTokens: tt.destTokens}, tt.destTokensErr).Maybe()
			sourcePriceRegistry.On("Address", mock.Anything).Return(sourcePriceRegistryAddress, nil).Maybe()
			sourcePriceRegistry.On("GetFeeTokens", ctx).Return(tt.sourceFeeTokens, tt.sourceFeeTokensErr).Maybe()
			sourcePriceRegistry.On("GetTokenPrices", ctx, mock.Anything).Return(tt.sourcePrices, nil).Maybe()
			destPriceRegistry.On("GetFeeTokens", ctx).Return(tt.destFeeTokens, tt.destFeeTokensErr).Maybe()
			destPriceRegistry.On("GetTokenPrices", ctx, mock.Anything).Return(tt.destPrices, nil).Maybe()

			sourcePriceRegistryProvider.On("NewPriceRegistryReader", ctx, sourcePriceRegistryAddress).Return(sourcePriceRegistry, nil).Maybe()

			reportingPlugin := ExecutionReportingPlugin{
				onRampReader:                onrampReader,
				offRampReader:               offrampReader,
				sourcePriceRegistry:         sourcePriceRegistry,
				sourcePriceRegistryProvider: sourcePriceRegistryProvider,
				destPriceRegistry:           destPriceRegistry,
				gasPriceEstimator:           gasPriceEstimator,
				sourceWrappedNativeToken:    weth,
				destWrappedNative:           wavax,
			}

			tokenData, err := reportingPlugin.prepareTokenExecData(ctx)
			if tt.destFeeTokensErr != nil || tt.sourceFeeTokensErr != nil || tt.destTokensErr != nil {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Len(t, tokenData.sourceTokenPrices, len(tt.sourcePrices))
			assert.Len(t, tokenData.destTokenPrices, len(tt.destPrices))

			for token, price := range tokenData.sourceTokenPrices {
				assert.Equal(t, tokenPrices[token].Value, price)
			}

			for token, price := range tokenData.destTokenPrices {
				assert.Equal(t, tokenPrices[token].Value, price)
			}
		})
	}
}

func encodeExecutionReport(t *testing.T, report cciptypes.ExecReport) []byte {
	reader, err := v1_2_0.NewOffRamp(logger.TestLogger(t), utils.RandomAddress(), nil, nil, nil, nil)
	require.NoError(t, err)
	ctx := testutils.Context(t)
	encodedReport, err := reader.EncodeExecutionReport(ctx, report)
	require.NoError(t, err)
	return encodedReport
}

// Verify the price registry update mechanism in case of configuration change on the source onRamp.
func TestExecutionReportingPlugin_ensurePriceRegistrySynchronization(t *testing.T) {
	p := &ExecutionReportingPlugin{}
	p.lggr = logger.TestLogger(t)
	p.sourcePriceRegistryLock = sync.RWMutex{}

	sourcePriceRegistryAddress1 := cciptypes.Address(utils.RandomAddress().String())
	sourcePriceRegistryAddress2 := cciptypes.Address(utils.RandomAddress().String())

	mockPriceRegistryReader1 := ccipdatamocks.NewPriceRegistryReader(t)
	mockPriceRegistryReader2 := ccipdatamocks.NewPriceRegistryReader(t)
	mockPriceRegistryReader1.On("Address", mock.Anything).Return(sourcePriceRegistryAddress1, nil)
	mockPriceRegistryReader2.On("Address", mock.Anything).Return(sourcePriceRegistryAddress2, nil).Maybe()
	mockPriceRegistryReader1.On("Close", mock.Anything).Return(nil)
	mockPriceRegistryReader2.On("Close", mock.Anything).Return(nil).Maybe()

	mockSourcePriceRegistryProvider := ccipdataprovidermocks.NewPriceRegistry(t)
	mockSourcePriceRegistryProvider.On("NewPriceRegistryReader", mock.Anything, sourcePriceRegistryAddress1).Return(mockPriceRegistryReader1, nil)
	mockSourcePriceRegistryProvider.On("NewPriceRegistryReader", mock.Anything, sourcePriceRegistryAddress2).Return(mockPriceRegistryReader2, nil)
	p.sourcePriceRegistryProvider = mockSourcePriceRegistryProvider

	mockOnRampReader := ccipdatamocks.NewOnRampReader(t)
	p.onRampReader = mockOnRampReader

	mockOnRampReader.On("SourcePriceRegistryAddress", mock.Anything).Return(sourcePriceRegistryAddress1, nil).Once()
	require.Equal(t, nil, p.sourcePriceRegistry)
	err := p.ensurePriceRegistrySynchronization(context.Background())
	require.NoError(t, err)
	require.Equal(t, mockPriceRegistryReader1, p.sourcePriceRegistry)

	mockOnRampReader.On("SourcePriceRegistryAddress", mock.Anything).Return(sourcePriceRegistryAddress2, nil).Once()
	err = p.ensurePriceRegistrySynchronization(context.Background())
	require.NoError(t, err)
	require.Equal(t, mockPriceRegistryReader2, p.sourcePriceRegistry)
}

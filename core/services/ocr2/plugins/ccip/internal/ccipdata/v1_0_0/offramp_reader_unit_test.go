package v1_0_0

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib/rpclibmocks"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	evmclimocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp_1_0_0"
	mock_contracts "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/mocks/v1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
)

func TestOffRampGetDestinationTokensFromSourceTokens(t *testing.T) {
	ctx := testutils.Context(t)
	const numSrcTokens = 20

	testCases := []struct {
		name           string
		outputChangeFn func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr
		expErr         bool
	}{
		{
			name:           "happy path",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr { return outputs },
			expErr:         false,
		},
		{
			name: "rpc error",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr {
				outputs[2].Err = fmt.Errorf("some error")
				return outputs
			},
			expErr: true,
		},
		{
			name: "unexpected outputs length should be fine if the type is correct",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr {
				outputs[0].Outputs = append(outputs[0].Outputs, "unexpected", 123)
				return outputs
			},
			expErr: false,
		},
		{
			name: "different compatible type",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr {
				outputs[0].Outputs = []any{outputs[0].Outputs[0].(common.Address)}
				return outputs
			},
			expErr: false,
		},
		{
			name: "different incompatible type",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr {
				outputs[0].Outputs = []any{outputs[0].Outputs[0].(common.Address).Bytes()}
				return outputs
			},
			expErr: true,
		},
	}

	lp := mocks.NewLogPoller(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			batchCaller := rpclibmocks.NewEvmBatchCaller(t)
			o := &OffRamp{evmBatchCaller: batchCaller, lp: lp}
			srcTks, dstTks, outputs := generateTokensAndOutputs(numSrcTokens)
			outputs = tc.outputChangeFn(outputs)
			batchCaller.On("BatchCall", mock.Anything, mock.Anything, mock.Anything).Return(outputs, nil)
			genericAddrs := ccipcalc.EvmAddrsToGeneric(srcTks...)
			actualDstTokens, err := o.getDestinationTokensFromSourceTokens(ctx, genericAddrs)

			if tc.expErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, ccipcalc.EvmAddrsToGeneric(dstTks...), actualDstTokens)
		})
	}
}

func TestCachedOffRampTokens(t *testing.T) {
	// Test data.
	srcTks, dstTks, _ := generateTokensAndOutputs(3)

	// Mock contract wrapper.
	mockOffRamp := mock_contracts.NewEVM2EVMOffRampInterface(t)
	mockOffRamp.On("GetDestinationTokens", mock.Anything).Return(dstTks, nil)
	mockOffRamp.On("GetSupportedTokens", mock.Anything).Return(srcTks, nil)
	mockOffRamp.On("Address").Return(utils.RandomAddress())

	lp := mocks.NewLogPoller(t)
	lp.On("LatestBlock", mock.Anything).Return(logpoller.LogPollerBlock{BlockNumber: rand.Int63()}, nil)

	offRamp := OffRamp{
		offRampV100:    mockOffRamp,
		lp:             lp,
		Logger:         logger.TestLogger(t),
		Client:         evmclimocks.NewClient(t),
		evmBatchCaller: rpclibmocks.NewEvmBatchCaller(t),
		cachedOffRampTokens: cache.NewLogpollerEventsBased[cciptypes.OffRampTokens](
			lp,
			offRamp_poolAddedPoolRemovedEvents,
			mockOffRamp.Address(),
		),
	}

	ctx := testutils.Context(t)
	tokens, err := offRamp.GetTokens(ctx)
	require.NoError(t, err)

	// Verify data is properly loaded in the cache.
	expectedPools := make(map[cciptypes.Address]cciptypes.Address)
	for i := range dstTks {
		expectedPools[cciptypes.Address(dstTks[i].String())] = cciptypes.Address(dstTks[i].String())
	}
	require.Equal(t, cciptypes.OffRampTokens{
		DestinationTokens: ccipcalc.EvmAddrsToGeneric(dstTks...),
		SourceTokens:      ccipcalc.EvmAddrsToGeneric(srcTks...),
	}, tokens)
}

func generateTokensAndOutputs(nbTokens uint) ([]common.Address, []common.Address, []rpclib.DataAndErr) {
	srcTks := make([]common.Address, nbTokens)
	dstTks := make([]common.Address, nbTokens)
	outputs := make([]rpclib.DataAndErr, nbTokens)
	for i := range srcTks {
		srcTks[i] = utils.RandomAddress()
		dstTks[i] = utils.RandomAddress()
		outputs[i] = rpclib.DataAndErr{
			Outputs: []any{dstTks[i]}, Err: nil,
		}
	}
	return srcTks, dstTks, outputs
}

func Test_GetExecutionStateChangesForSeqNums(t *testing.T) {
	ctx := testutils.Context(t)
	chainID := testutils.NewRandomEVMChainID()
	orm := logpoller.NewORM(chainID, pgtest.NewSqlxDB(t), logger.TestLogger(t))
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Hour,
		FinalityDepth:            2,
		BackfillBatchSize:        20,
		RpcBatchSize:             10,
		KeepFinalizedBlocksDepth: 1000,
	}
	lp := logpoller.NewLogPoller(orm, nil, logger.TestLogger(t), lpOpts)

	offrampAddress := utils.RandomAddress()
	inputLogs := []logpoller.Log{
		CreateExecutionStateChangeEventLog(t, chainID, offrampAddress, 10, 2, 1, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, chainID, offrampAddress, 11, 3, 1, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, chainID, offrampAddress, 12, 5, 1, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, chainID, offrampAddress, 13, 5, 2, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, chainID, offrampAddress, 14, 5, 3, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, chainID, offrampAddress, 15, 8, 1, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, chainID, offrampAddress, 16, 9, 1, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, chainID, utils.RandomAddress(), 16, 9, 1, utils.RandomBytes32()),
	}
	require.NoError(t, orm.InsertLogsWithBlock(ctx, inputLogs, logpoller.NewLogPollerBlock(utils.RandomBytes32(), 20, time.Now(), 20)))

	tests := []struct {
		name               string
		seqNums            []cciptypes.SequenceNumberRange
		expectedLogsSeqNrs []uint64
	}{
		{
			name: "no logs are returned",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 1, Max: 9},
			},
			expectedLogsSeqNrs: []uint64{},
		},
		{
			name: "all logs are returned",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 10, Max: 16},
			},
			expectedLogsSeqNrs: []uint64{10, 11, 12, 13, 14, 15, 16},
		},
		{
			name: "all logs are returned for wider range",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 8, Max: 17},
			},
			expectedLogsSeqNrs: []uint64{10, 11, 12, 13, 14, 15, 16},
		},
		{
			name: "some logs are returned for tighter range",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 11, Max: 14},
			},
			expectedLogsSeqNrs: []uint64{11, 12, 13, 14},
		},
		{
			name: "multiple smaller ranges",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 10, Max: 11},
				{Min: 13, Max: 14},
			},
			expectedLogsSeqNrs: []uint64{10, 11, 13, 14},
		},
		{
			name: "single element ranges",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 10, Max: 10},
				{Min: 14, Max: 14},
				{Min: 15, Max: 15},
			},
			expectedLogsSeqNrs: []uint64{10, 14, 15},
		},
		{
			name: "out of order ranges returns logs in proper order",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 14, Max: 14},
				{Min: 10, Max: 11},
				{Min: 15, Max: 16},
			},
			expectedLogsSeqNrs: []uint64{10, 11, 14, 15, 16},
		},
		{
			name: "overlapping ranges returns logs only once",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 10, Max: 14},
				{Min: 13, Max: 15},
				{Min: 11, Max: 12},
			},
			expectedLogsSeqNrs: []uint64{10, 11, 12, 13, 14, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offRamp, err1 := NewOffRamp(logger.TestLogger(t), offrampAddress, evmclimocks.NewClient(t), lp, nil, nil)
			require.NoError(t, err1)

			msgs, err1 := offRamp.GetExecutionStateChangesForSeqNums(ctx, tt.seqNums, 0)
			require.NoError(t, err1)

			assert.Len(t, msgs, len(tt.expectedLogsSeqNrs))
			for i, msg := range msgs {
				assert.Equal(t, tt.expectedLogsSeqNrs[i], msg.SequenceNumber)
			}
		})
	}
}

func Test_LogsAreProperlyMarkedAsFinalized(t *testing.T) {
	minSeqNr := uint64(10)
	maxSeqNr := uint64(14)
	inputLogs := []logpoller.Log{
		CreateExecutionStateChangeEventLog(t, nil, utils.RandomAddress(), 10, 2, 0, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, nil, utils.RandomAddress(), 11, 3, 0, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, nil, utils.RandomAddress(), 12, 5, 0, utils.RandomBytes32()),
		CreateExecutionStateChangeEventLog(t, nil, utils.RandomAddress(), 14, 7, 0, utils.RandomBytes32()),
	}

	tests := []struct {
		name                        string
		lastFinalizedBlock          uint64
		expectedFinalizedSequenceNr []uint64
	}{
		{
			"all logs are finalized",
			10,
			[]uint64{10, 11, 12, 14},
		},
		{
			"some logs are finalized",
			5,
			[]uint64{10, 11, 12},
		},
		{
			"no logs are finalized",
			1,
			[]uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offrampAddress := utils.RandomAddress()

			lp := mocks.NewLogPoller(t)
			lp.On("LatestBlock", mock.Anything).
				Return(logpoller.LogPollerBlock{FinalizedBlockNumber: int64(tt.lastFinalizedBlock)}, nil)
			lp.On("IndexedLogsTopicRange", mock.Anything, ExecutionStateChangedEvent, offrampAddress, 1, logpoller.EvmWord(minSeqNr), logpoller.EvmWord(maxSeqNr), evmtypes.Confirmations(0)).
				Return(inputLogs, nil)

			offRamp, err := NewOffRamp(logger.TestLogger(t), offrampAddress, evmclimocks.NewClient(t), lp, nil, nil)
			require.NoError(t, err)
			logs, err := offRamp.GetExecutionStateChangesBetweenSeqNums(testutils.Context(t), minSeqNr, maxSeqNr, 0)
			require.NoError(t, err)
			assert.Len(t, logs, len(inputLogs))

			for _, log := range logs {
				assert.Equal(t, slices.Contains(tt.expectedFinalizedSequenceNr, log.SequenceNumber), log.Finalized)
			}
		})
	}
}

func TestGetRouter(t *testing.T) {
	routerAddr := utils.RandomAddress()

	mockOffRamp := mock_contracts.NewEVM2EVMOffRampInterface(t)
	mockOffRamp.On("GetDynamicConfig", mock.Anything).Return(evm_2_evm_offramp_1_0_0.EVM2EVMOffRampDynamicConfig{
		Router: routerAddr,
	}, nil)

	offRamp := OffRamp{
		offRampV100: mockOffRamp,
	}

	ctx := testutils.Context(t)
	gotRouterAddr, err := offRamp.GetRouter(ctx)
	require.NoError(t, err)

	gotRouterEvmAddr, err := ccipcalc.GenericAddrToEvm(gotRouterAddr)
	require.NoError(t, err)
	assert.Equal(t, routerAddr, gotRouterEvmAddr)
}

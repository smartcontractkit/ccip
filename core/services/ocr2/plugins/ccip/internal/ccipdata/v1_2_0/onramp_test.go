package v1_2_0_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	evmclimocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
)

func TestOnRamp_GetSendRequestsForSeqNums(t *testing.T) {
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

	onrampAddress := utils.RandomAddress()
	inputLogs := []logpoller.Log{
		createCCIPSenRequestedLog(t, chainID, onrampAddress, 10, 2, 1, utils.RandomBytes32()),
		createCCIPSenRequestedLog(t, chainID, onrampAddress, 11, 3, 1, utils.RandomBytes32()),
		createCCIPSenRequestedLog(t, chainID, onrampAddress, 12, 5, 1, utils.RandomBytes32()),
		createCCIPSenRequestedLog(t, chainID, onrampAddress, 13, 5, 2, utils.RandomBytes32()),
		createCCIPSenRequestedLog(t, chainID, onrampAddress, 14, 5, 3, utils.RandomBytes32()),
		createCCIPSenRequestedLog(t, chainID, onrampAddress, 15, 8, 1, utils.RandomBytes32()),
		createCCIPSenRequestedLog(t, chainID, onrampAddress, 16, 9, 1, utils.RandomBytes32()),
		createCCIPSenRequestedLog(t, chainID, utils.RandomAddress(), 16, 9, 1, utils.RandomBytes32()),
	}
	require.NoError(t, orm.InsertLogsWithBlock(ctx, inputLogs, logpoller.NewLogPollerBlock(utils.RandomBytes32(), 20, time.Now(), 5)))

	tests := []struct {
		name               string
		seqNums            []cciptypes.SequenceNumberRange
		expectedLogsSeqNrs []uint64
		finalized          bool
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
		{
			name: "only finalized logs are returned",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 10, Max: 16},
			},
			expectedLogsSeqNrs: []uint64{10, 11, 12, 13, 14},
			finalized:          true,
		},
		{
			name: "finalized logs works with ranges",
			seqNums: []cciptypes.SequenceNumberRange{
				{Min: 10, Max: 11},
				{Min: 13, Max: 15},
				{Min: 16, Max: 16},
			},
			expectedLogsSeqNrs: []uint64{10, 11, 13, 14},
			finalized:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			onRamp, err1 := v1_2_0.NewOnRamp(logger.TestLogger(t), uint64(123), uint64(123), onrampAddress, lp, nil)
			require.NoError(t, err1)

			msgs, err1 := onRamp.GetSendRequestsForSeqNums(ctx, tt.seqNums, tt.finalized)
			require.NoError(t, err1)

			require.Len(t, msgs, len(tt.expectedLogsSeqNrs))
			for i, msg := range msgs {
				assert.Equal(t, tt.expectedLogsSeqNrs[i], msg.SequenceNumber)
			}
		})
	}
}

// Scenario 1
// Benchmark_FilteredLogsQuery
// Benchmark_FilteredLogsQuery-12    	      50	  24661219 ns/op
// Benchmark_FilteredLogsQuery-12    	      49	  24392641 ns/op
// Benchmark_FilteredLogsQuery-12    	      50	  25084857 ns/op
// Benchmark_FilteredLogsQuery-12    	      54	  28100956 ns/op
//
// Scenario 2
// Benchmark_FilteredLogsQuery
// Benchmark_FilteredLogsQuery-12    	      51	  24117666 ns/op
// Benchmark_FilteredLogsQuery-12    	      51	  24520130 ns/op
// Benchmark_FilteredLogsQuery-12    	      52	  23826061 ns/op
// Benchmark_FilteredLogsQuery-12    	      56	  22766220 ns/op
//
// Scenario 3
// Benchmark_FilteredLogsQuery
// Benchmark_FilteredLogsQuery-12    	      79	  17122976 ns/op
// Benchmark_FilteredLogsQuery-12    	      78	  17512757 ns/op
// Benchmark_FilteredLogsQuery-12    	      80	  83212702 ns/op
// Benchmark_FilteredLogsQuery-12    	      76	  16195331 ns/op
// Benchmark_FilteredLogsQuery-12    	      79	  15901089 ns/op
func Benchmark_FilteredLogsQuery(b *testing.B) {
	ctx := testutils.Context(b)
	_, db := heavyweight.FullTestDBV2(b, nil)
	chainID := testutils.NewRandomEVMChainID()
	orm := logpoller.NewORM(chainID, db, logger.TestLogger(b))
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Hour,
		FinalityDepth:            2,
		BackfillBatchSize:        20,
		RpcBatchSize:             10,
		KeepFinalizedBlocksDepth: 1000,
	}
	lp := logpoller.NewLogPoller(orm, nil, logger.TestLogger(b), lpOpts)

	onrampAddress := utils.RandomAddress()

	for j := 1; j <= 100; j++ {
		var logs []logpoller.Log
		for i := 0; i < 1_000; i++ {
			logs = append(
				logs,
				createCCIPSenRequestedLog(
					b,
					chainID,
					onrampAddress,
					uint64(j*1000+i),
					int64(j*1000+i),
					int64(j),
					utils.RandomBytes32(),
				),
			)
		}
		require.NoError(b, orm.InsertLogs(ctx, logs))
		require.NoError(b, orm.InsertBlock(ctx, utils.RandomHash(), int64((j+1)*1000-1), time.Now(), 0))
	}

	onRamp, err := v1_2_0.NewOnRamp(logger.TestLogger(b), uint64(1), uint64(2), onrampAddress, lp, evmclimocks.NewClient(b))
	require.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Scenario 1
		//logs, err1 := onRamp.GetSendRequestsBetweenSeqNums(testutils.Context(b), 1024, 2047, false)
		// Scenario 2
		//logs, err1 := onRamp.GetSendRequestsForSeqNums(testutils.Context(b), []cciptypes.SequenceNumberRange{{Min: 1024, Max: 2047}}, false)
		//require.NoError(b, err1)
		//assert.Len(b, logs, 1024)
		// Scenario 3
		logs, err1 := onRamp.GetSendRequestsForSeqNums(testutils.Context(b), []cciptypes.SequenceNumberRange{
			{Min: 1000, Max: 1099},
			{Min: 1200, Max: 1299},
			{Min: 1800, Max: 1999},
			{Min: 2200, Max: 2499},
		}, false)
		require.NoError(b, err1)
		assert.Len(b, logs, 700)
	}
}

func createCCIPSenRequestedLog(t testing.TB, chainID *big.Int, address common.Address, seqNr uint64, blockNumber int64, logIndex int64, messageID common.Hash) logpoller.Log {
	tAbi, err := evm_2_evm_onramp_1_2_0.EVM2EVMOnRampMetaData.GetAbi()
	require.NoError(t, err)
	eseEvent, ok := tAbi.Events["CCIPSendRequested"]
	require.True(t, ok)

	message := evm_2_evm_onramp_1_2_0.InternalEVM2EVMMessage{
		SourceChainSelector: 123,
		Sender:              utils.RandomAddress(),
		Receiver:            utils.RandomAddress(),
		SequenceNumber:      seqNr,
		GasLimit:            big.NewInt(100),
		Strict:              false,
		Nonce:               1337,
		FeeToken:            utils.RandomAddress(),
		FeeTokenAmount:      big.NewInt(1),
		Data:                []byte{},
		TokenAmounts:        []evm_2_evm_onramp_1_2_0.ClientEVMTokenAmount{},
		MessageId:           messageID,
	}

	logData, err := eseEvent.Inputs.Pack(message)
	require.NoError(t, err)

	topic0 := evm_2_evm_onramp_1_2_0.EVM2EVMOnRampCCIPSendRequested{}.Topic()

	return logpoller.Log{
		Topics: [][]byte{
			topic0[:],
		},
		Data:        logData,
		LogIndex:    logIndex,
		BlockHash:   utils.RandomBytes32(),
		BlockNumber: blockNumber,
		EventSig:    topic0,
		Address:     address,
		TxHash:      utils.RandomBytes32(),
		EvmChainId:  ubig.New(chainID),
	}
}

package v1_0_0

import (
	"testing"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
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
		CreateCCIPSenRequestedLog(t, chainID, onrampAddress, 10, 2, 1, utils.RandomBytes32()),
		CreateCCIPSenRequestedLog(t, chainID, onrampAddress, 11, 3, 1, utils.RandomBytes32()),
		CreateCCIPSenRequestedLog(t, chainID, onrampAddress, 12, 5, 1, utils.RandomBytes32()),
		CreateCCIPSenRequestedLog(t, chainID, onrampAddress, 13, 5, 2, utils.RandomBytes32()),
		CreateCCIPSenRequestedLog(t, chainID, onrampAddress, 14, 5, 3, utils.RandomBytes32()),
		CreateCCIPSenRequestedLog(t, chainID, onrampAddress, 15, 8, 1, utils.RandomBytes32()),
		CreateCCIPSenRequestedLog(t, chainID, onrampAddress, 16, 9, 1, utils.RandomBytes32()),
		CreateCCIPSenRequestedLog(t, chainID, utils.RandomAddress(), 16, 9, 1, utils.RandomBytes32()),
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
			onRamp, err1 := NewOnRamp(logger.TestLogger(t), uint64(123), uint64(123), onrampAddress, lp, nil)
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

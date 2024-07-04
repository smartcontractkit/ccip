package v1_0_0_test

import (
	"encoding/binary"
	"math/big"
	"slices"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	evmclimocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_0_0"
)

func TestExecutionReportEncodingV100(t *testing.T) {
	// Note could consider some fancier testing here (fuzz/property)
	// but I think that would essentially be testing geth's abi library
	// as our encode/decode is a thin wrapper around that.
	report := cciptypes.ExecReport{
		Messages:          []cciptypes.EVM2EVMMessage{},
		OffchainTokenData: [][][]byte{{}},
		Proofs:            [][32]byte{testutils.Random32Byte()},
		ProofFlagBits:     big.NewInt(133),
	}

	offRamp, err := v1_0_0.NewOffRamp(logger.TestLogger(t), utils.RandomAddress(), nil, lpmocks.NewLogPoller(t), nil, nil)
	require.NoError(t, err)

	ctx := testutils.Context(t)
	encodeExecutionReport, err := offRamp.EncodeExecutionReport(ctx, report)
	require.NoError(t, err)
	decodeCommitReport, err := offRamp.DecodeExecutionReport(ctx, encodeExecutionReport)
	require.NoError(t, err)
	require.Equal(t, report.Proofs, decodeCommitReport.Proofs)
	require.Equal(t, report, decodeCommitReport)
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
		createExecutionStateChangeEventLog(t, chainID, offrampAddress, 10, 2, 1, utils.RandomBytes32()),
		createExecutionStateChangeEventLog(t, chainID, offrampAddress, 11, 3, 1, utils.RandomBytes32()),
		createExecutionStateChangeEventLog(t, chainID, offrampAddress, 12, 5, 1, utils.RandomBytes32()),
		createExecutionStateChangeEventLog(t, chainID, offrampAddress, 13, 5, 2, utils.RandomBytes32()),
		createExecutionStateChangeEventLog(t, chainID, offrampAddress, 14, 5, 3, utils.RandomBytes32()),
		createExecutionStateChangeEventLog(t, chainID, offrampAddress, 15, 8, 1, utils.RandomBytes32()),
		createExecutionStateChangeEventLog(t, chainID, offrampAddress, 16, 9, 1, utils.RandomBytes32()),
		createExecutionStateChangeEventLog(t, chainID, utils.RandomAddress(), 16, 9, 1, utils.RandomBytes32()),
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
			offRamp, err1 := v1_0_0.NewOffRamp(logger.TestLogger(t), offrampAddress, evmclimocks.NewClient(t), lp, nil, nil)
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
	ctx := testutils.Context(t)

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
				createExecutionStateChangeEventLog(t, chainID, offrampAddress, 10, 2, 0, utils.RandomBytes32()),
				createExecutionStateChangeEventLog(t, chainID, offrampAddress, 11, 3, 0, utils.RandomBytes32()),
				createExecutionStateChangeEventLog(t, chainID, offrampAddress, 12, 5, 0, utils.RandomBytes32()),
				createExecutionStateChangeEventLog(t, chainID, offrampAddress, 14, 7, 0, utils.RandomBytes32()),
			}
			err := orm.InsertLogsWithBlock(ctx, inputLogs, logpoller.NewLogPollerBlock(utils.RandomBytes32(), 100, time.Now(), int64(tt.lastFinalizedBlock)))
			require.NoError(t, err)

			offRamp, err := v1_0_0.NewOffRamp(logger.TestLogger(t), offrampAddress, evmclimocks.NewClient(t), lp, nil, nil)
			require.NoError(t, err)
			logs, err := offRamp.GetExecutionStateChangesForSeqNums(testutils.Context(t), []cciptypes.SequenceNumberRange{{Min: 0, Max: 100}}, 0)
			require.NoError(t, err)
			assert.Len(t, logs, len(inputLogs))

			for _, log := range logs {
				assert.Equal(t, slices.Contains(tt.expectedFinalizedSequenceNr, log.SequenceNumber), log.Finalized)
			}
		})
	}
}

func createExecutionStateChangeEventLog(t *testing.T, chainID *big.Int, address common.Address, seqNr uint64, blockNumber int64, logIndex int64, messageID common.Hash) logpoller.Log {
	tAbi, err := evm_2_evm_offramp_1_0_0.EVM2EVMOffRampMetaData.GetAbi()
	require.NoError(t, err)
	eseEvent, ok := tAbi.Events["ExecutionStateChanged"]
	require.True(t, ok)

	logData, err := eseEvent.Inputs.NonIndexed().Pack(uint8(1), []byte("some return data"))
	require.NoError(t, err)
	seqNrBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(seqNrBytes, seqNr)
	seqNrTopic := common.BytesToHash(seqNrBytes)
	topic0 := evm_2_evm_offramp_1_0_0.EVM2EVMOffRampExecutionStateChanged{}.Topic()

	return logpoller.Log{
		Topics: [][]byte{
			topic0[:],
			seqNrTopic[:],
			messageID[:],
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

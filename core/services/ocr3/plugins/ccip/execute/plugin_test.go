package execute

import (
	"context"
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/cometbft/cometbft/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/ccipocr3/internal/libs/slicelib"
	"github.com/smartcontractkit/ccipocr3/internal/mocks"
	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestSomethingCool(t *testing.T) {

	foo := map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1):              1,
		cciptypes.ChainSelector(math.MaxUint64): 1,
	}

	js, _ := json.Marshal(foo)
	t.Log(string(js))

	b := []byte(`{"1":1,"18446744073709551615":1}`)
	var bar map[cciptypes.ChainSelector]int
	assert.NoError(t, json.Unmarshal(b, &bar))
	t.Log(bar)
}

func Test_getPendingExecutedReports(t *testing.T) {
	tests := []struct {
		name    string
		reports []cciptypes.CommitPluginReportWithMeta
		ranges  map[cciptypes.ChainSelector][]cciptypes.SeqNumRange
		want    cciptypes.ExecutePluginCommitObservations
		want1   time.Time
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:    "empty",
			reports: nil,
			ranges:  nil,
			want:    cciptypes.ExecutePluginCommitObservations{},
			want1:   time.Time{},
			wantErr: assert.NoError,
		},
		{
			name: "single non-executed report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				1: nil,
			},
			want: cciptypes.ExecutePluginCommitObservations{
				1: []cciptypes.ExecutePluginCommitDataWithMessages{
					{ExecutePluginCommitData: cciptypes.ExecutePluginCommitData{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						ExecutedMessages:    nil,
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            999,
					}},
				},
			},
			want1:   time.UnixMilli(10101010101),
			wantErr: assert.NoError,
		},
		{
			name: "single half-executed report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				1: {
					cciptypes.NewSeqNumRange(1, 3),
					cciptypes.NewSeqNumRange(7, 8),
				},
			},
			want: cciptypes.ExecutePluginCommitObservations{
				1: []cciptypes.ExecutePluginCommitDataWithMessages{
					{ExecutePluginCommitData: cciptypes.ExecutePluginCommitData{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            999,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3, 7, 8},
					}},
				},
			},
			want1:   time.UnixMilli(10101010101),
			wantErr: assert.NoError,
		},
		{
			name: "last timestamp",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report:    cciptypes.CommitPluginReport{},
				},
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(9999999999999999),
					Report:    cciptypes.CommitPluginReport{},
				},
			},
			ranges:  map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{},
			want:    cciptypes.ExecutePluginCommitObservations{},
			want1:   time.UnixMilli(9999999999999999),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockReader := mocks.NewCCIPReader()
			mockReader.On(
				"CommitReportsGTETimestamp", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			).Return(tt.reports, nil)
			for k, v := range tt.ranges {
				mockReader.On("ExecutedMessageRanges", mock.Anything, k, mock.Anything, mock.Anything).Return(v, nil)
			}

			// CCIP Reader mocks:
			// once:
			//      CommitReportsGTETimestamp(ctx, dest, ts, 1000) -> ([]cciptypes.CommitPluginReportWithMeta, error)
			// for each chain selector:
			//      ExecutedMessageRanges(ctx, selector, dest, seqRange) -> ([]cciptypes.SeqNumRange, error)

			got, got1, err := getPendingExecutedReports(context.Background(), mockReader, 123, time.Now())
			if !tt.wantErr(t, err, "getPendingExecutedReports(...)") {
				return
			}
			assert.Equalf(t, tt.want, got, "getPendingExecutedReports(...)")
			assert.Equalf(t, tt.want1, got1, "getPendingExecutedReports(...)")
		})
	}
}

// TODO: better than this
type tdr int

func (t tdr) ReadTokenData(ctx context.Context, srcChain cciptypes.ChainSelector, num cciptypes.SeqNum) ([][]byte, error) {
	return nil, nil
}

func breakCommitReport(commitReport cciptypes.ExecutePluginCommitDataWithMessages) cciptypes.ExecutePluginCommitDataWithMessages {
	commitReport.Messages = append(commitReport.Messages, cciptypes.CCIPMsg{
		CCIPMsgBaseDetails: cciptypes.CCIPMsgBaseDetails{
			ID:          crypto.CRandHex(32),
			SourceChain: cciptypes.ChainSelector(1),
			SeqNum:      cciptypes.SeqNum(999),
			MsgHash:     cciptypes.Bytes32{},
		},
	})
	return commitReport
}

func makeTestCommitReport(numMessages, srcChain, firstSeqNum, block int, timestamp int64, executed []cciptypes.SeqNum) cciptypes.ExecutePluginCommitDataWithMessages {
	for _, e := range executed {
		if e < cciptypes.SeqNum(firstSeqNum) || e > cciptypes.SeqNum(firstSeqNum+numMessages-1) {
			panic("executed message out of range")
		}
	}
	var messages []cciptypes.CCIPMsg
	for i := 0; i < numMessages; i++ {
		messages = append(messages, cciptypes.CCIPMsg{
			CCIPMsgBaseDetails: cciptypes.CCIPMsgBaseDetails{
				ID:          crypto.CRandHex(32),
				SourceChain: cciptypes.ChainSelector(srcChain),
				SeqNum:      cciptypes.SeqNum(i + firstSeqNum),
				MsgHash:     cciptypes.Bytes32{},
			},
			//ChainFeeLimit:      cciptypes.BigInt{},
			Nonce: uint64(i),
			//Sender:             "0xGrovor",
			//Receiver:           "0xOscar",
			//Strict:             false,
			//FeeToken:           "",
			//FeeTokenAmount:     cciptypes.BigInt{},
			//Data:               nil,
			//TokenAmounts:       nil,
			//SourceTokenData:    nil,
			//Metadata:           cciptypes.CCIPMsgMetadata{},
		})
	}

	return cciptypes.ExecutePluginCommitDataWithMessages{
		ExecutePluginCommitData: cciptypes.ExecutePluginCommitData{
			SourceChain:         cciptypes.ChainSelector(srcChain),
			SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(firstSeqNum), cciptypes.SeqNum(firstSeqNum+numMessages-1)),
			Timestamp:           time.UnixMilli(timestamp),
			BlockNum:            uint64(block),
			ExecutedMessages:    executed,
		},
		Messages: messages,
	}

}

// assertMerkleRoot computes the source messages merkle root, then computes a verification with the proof, then compares the roots.
func assertMerkleRoot(t *testing.T, hasher cciptypes.MessageHasher, execReport cciptypes.ExecutePluginReportSingleChain, commitReport cciptypes.ExecutePluginCommitDataWithMessages) {
	keccak := hashutil.NewKeccak()
	// Generate merkle root from commit report messages
	var leafHashes [][32]byte
	for _, msg := range commitReport.Messages {
		hash, err := hasher.Hash(context.Background(), msg)
		require.NoError(t, err)
		leafHashes = append(leafHashes, hash)
	}
	tree, err := merklemulti.NewTree(keccak, leafHashes)
	require.NoError(t, err)
	merkleRoot := tree.Root()

	// Generate merkle root from exec report messages and proofj
	ctx := context.Background()
	var leaves [][32]byte
	for _, msg := range execReport.Messages {
		hash, err := hasher.Hash(ctx, msg)
		require.NoError(t, err)
		leaves = append(leaves, hash)
	}
	proofCast := make([][32]byte, len(execReport.Proofs))
	for i, p := range execReport.Proofs {
		copy(proofCast[i][:], p[:32])
		proofCast[i][2] = proofCast[i][2]
	}
	var proof merklemulti.Proof[[32]byte]
	proof.Hashes = proofCast
	proof.SourceFlags = slicelib.BitFlagsToBools(execReport.ProofFlagBits.Int, len(leaves)+len(proofCast)-1)
	recomputedMerkleRoot, err := merklemulti.VerifyComputeRoot(hashutil.NewKeccak(),
		leaves,
		proof)
	assert.NoError(t, err)
	assert.NotNil(t, recomputedMerkleRoot)

	// Compare them
	assert.Equal(t, merkleRoot, recomputedMerkleRoot)
}

func Test_selectReport(t *testing.T) {
	hasher := mocks.NewMessageHasher()
	codec := mocks.NewExecutePluginJSONReportCodec()
	lggr := logger.Test(t)
	var tokenDataReader tdr

	type args struct {
		reports       []cciptypes.ExecutePluginCommitDataWithMessages
		maxReportSize int
	}
	tests := []struct {
		name                  string
		args                  args
		expectedExecReports   int
		expectedCommitReports int
		expectedExecThings    []int
		lastReportExecuted    []cciptypes.SeqNum
		wantErr               string
	}{
		{
			name: "empty report",
			args: args{
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{},
			},
			expectedExecReports:   0,
			expectedCommitReports: 0,
		},
		{
			name: "half report",
			args: args{
				maxReportSize: 2200,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, nil),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{5},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104},
		},
		{
			name: "full report",
			args: args{
				maxReportSize: 10000,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, nil),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{10},
		},
		{
			name: "two reports",
			args: args{
				maxReportSize: 15000,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, nil),
					makeTestCommitReport(20, 2, 100, 999, 10101010101, nil),
				},
			},
			expectedExecReports:   2,
			expectedCommitReports: 0,
			expectedExecThings:    []int{10, 20},
		},
		{
			name: "one and half reports",
			args: args{
				maxReportSize: 8000,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, nil),
					makeTestCommitReport(20, 2, 100, 999, 10101010101, nil),
				},
			},
			expectedExecReports:   2,
			expectedCommitReports: 1,
			expectedExecThings:    []int{10, 10},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		},
		{
			name: "exactly one report",
			args: args{
				maxReportSize: 3900,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, nil),
					makeTestCommitReport(20, 2, 100, 999, 10101010101, nil),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{10},
			lastReportExecuted:    []cciptypes.SeqNum{},
		},
		{
			name: "execute remainder of partially executed report",
			args: args{
				maxReportSize: 2500,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, []cciptypes.SeqNum{100, 101, 102, 103, 104}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{5},
		},
		{
			name: "partially execute remainder of partially executed report",
			args: args{
				maxReportSize: 2000,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, []cciptypes.SeqNum{100, 101, 102, 103, 104}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{4},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 105, 106, 107, 108},
		},
		{
			name: "execute remainder of sparsely executed report",
			args: args{
				maxReportSize: 2500,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, []cciptypes.SeqNum{100, 102, 104, 106, 108}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 0,
			expectedExecThings:    []int{5},
		},
		{
			name: "partially execute remainder of partially executed sparse report",
			args: args{
				maxReportSize: 2000,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					makeTestCommitReport(10, 1, 100, 999, 10101010101, []cciptypes.SeqNum{100, 102, 104, 106, 108}),
				},
			},
			expectedExecReports:   1,
			expectedCommitReports: 1,
			expectedExecThings:    []int{4},
			lastReportExecuted:    []cciptypes.SeqNum{100, 101, 102, 103, 104, 105, 106, 107, 108},
		},
		{
			name: "broken report",
			args: args{
				maxReportSize: 10000,
				reports: []cciptypes.ExecutePluginCommitDataWithMessages{
					breakCommitReport(makeTestCommitReport(10, 1, 100, 999, 10101010101, nil)),
				},
			},
			wantErr: "unable to build a single chain report",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			execReports, commitReports, err := selectReport(ctx, lggr, hasher, codec, tokenDataReader, tt.args.reports, tt.args.maxReportSize)
			if tt.wantErr != "" {
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Len(t, execReports, tt.expectedExecReports)
			require.Len(t, commitReports, tt.expectedCommitReports)
			for i, execReport := range execReports {
				assert.Len(t, execReport.Messages, tt.expectedExecThings[i])
				assert.Len(t, execReport.OffchainTokenData, tt.expectedExecThings[i])
				assert.NotEmptyf(t, execReport.Proofs, "Proof should not be empty.")
				assertMerkleRoot(t, hasher, execReport, tt.args.reports[i])
			}
			// If the last report is partially executed, the executed messages can be checked.
			if len(execReports) > 0 && len(tt.lastReportExecuted) > 0 {
				lastReport := commitReports[len(commitReports)-1]
				assert.ElementsMatch(t, tt.lastReportExecuted, lastReport.ExecutedMessages)
			}
		})
	}
}

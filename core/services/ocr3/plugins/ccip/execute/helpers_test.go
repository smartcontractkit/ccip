package commit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/ccipocr3/internal/model"
)

func Test_computeRanges(t *testing.T) {
	type args struct {
		reports []model.ExecutePluginCommitData
	}

	tests := []struct {
		name string
		args args
		want []model.SeqNumRange
		err  error
	}{
		{
			name: "empty",
			args: args{reports: []model.ExecutePluginCommitData{}},
			want: nil,
		},
		{
			name: "overlapping ranges",
			args: args{reports: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(15, 25)}},
			},
			err: ErrOverlappingRanges,
		},
		{
			name: "simple ranges collapsed",
			args: args{reports: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(21, 40)},
				{SequenceNumberRange: model.NewSeqNumRange(41, 60)}},
			},
			want: []model.SeqNumRange{{10, 60}},
		},
		{
			name: "non-contiguous ranges",
			args: args{reports: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: model.NewSeqNumRange(50, 60)}},
			},
			want: []model.SeqNumRange{{10, 20}, {30, 40}, {50, 60}},
		},
		{
			name: "contiguous and non-contiguous ranges",
			args: args{reports: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(21, 40)},
				{SequenceNumberRange: model.NewSeqNumRange(50, 60)}},
			},
			want: []model.SeqNumRange{{10, 40}, {50, 60}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := computeRanges(tt.args.reports)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_groupByChainSelector(t *testing.T) {
	type args struct {
		reports []model.CommitPluginReport
	}
	tests := []struct {
		name string
		args args
		want model.ExecutePluginCommitObservations
	}{
		{
			name: "empty",
			args: args{reports: []model.CommitPluginReport{}},
			want: model.ExecutePluginCommitObservations{},
		},
		{
			name: "reports",
			args: args{reports: []model.CommitPluginReport{{MerkleRoots: []model.MerkleRootChain{
				{ChainSel: 1, SeqNumsRange: model.NewSeqNumRange(10, 20), MerkleRoot: model.Bytes32{1}},
				{ChainSel: 2, SeqNumsRange: model.NewSeqNumRange(30, 40), MerkleRoot: model.Bytes32{2}},
			}}}},
			want: model.ExecutePluginCommitObservations{
				1: {
					{
						MerkleRoot:          model.Bytes32{1},
						SequenceNumberRange: model.NewSeqNumRange(10, 20),
						ExecutedMessages:    nil,
					},
				},
				2: {
					{

						MerkleRoot:          model.Bytes32{2},
						SequenceNumberRange: model.NewSeqNumRange(30, 40),
						ExecutedMessages:    nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, groupByChainSelector(tt.args.reports), "groupByChainSelector(%v)", tt.args.reports)
		})
	}
}

func Test_filterOutFullyExecutedMessages(t *testing.T) {
	type args struct {
		reports          []model.ExecutePluginCommitData
		executedMessages []model.SeqNumRange
	}
	tests := []struct {
		name    string
		args    args
		want    []model.ExecutePluginCommitData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				reports:          nil,
				executedMessages: nil,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "empty2",
			args: args{
				reports:          []model.ExecutePluginCommitData{},
				executedMessages: nil,
			},
			want:    []model.ExecutePluginCommitData{},
			wantErr: assert.NoError,
		},
		{
			name: "no executed messages",
			args: args{
				reports: []model.ExecutePluginCommitData{
					{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
				},
				executedMessages: nil,
			},
			want: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
			},
			wantErr: assert.NoError,
		},
		{
			name: "executed messages",
			args: args{
				reports: []model.ExecutePluginCommitData{
					{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
				},
				executedMessages: []model.SeqNumRange{
					model.NewSeqNumRange(0, 100),
				},
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "2 partially executed",
			args: args{
				reports: []model.ExecutePluginCommitData{
					{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
				},
				executedMessages: []model.SeqNumRange{
					model.NewSeqNumRange(15, 35),
				},
			},
			want: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
			},
			wantErr: assert.NoError,
		},
		{
			name: "2 partially executed 1 fully executed",
			args: args{
				reports: []model.ExecutePluginCommitData{
					{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
				},
				executedMessages: []model.SeqNumRange{
					model.NewSeqNumRange(15, 55),
				},
			},
			want: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
			},
			wantErr: assert.NoError,
		},
		{
			name: "first report executed",
			args: args{
				reports: []model.ExecutePluginCommitData{
					{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
				},
				executedMessages: []model.SeqNumRange{
					model.NewSeqNumRange(10, 20),
				},
			},
			want: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
			},
			wantErr: assert.NoError,
		},
		{
			name: "last report executed",
			args: args{
				reports: []model.ExecutePluginCommitData{
					{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: model.NewSeqNumRange(50, 60)},
				},
				executedMessages: []model.SeqNumRange{
					model.NewSeqNumRange(50, 60),
				},
			},
			want: []model.ExecutePluginCommitData{
				{SequenceNumberRange: model.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: model.NewSeqNumRange(30, 40)},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filterOutFullyExecutedMessages(tt.args.reports, tt.args.executedMessages)
			if !tt.wantErr(t, err, fmt.Sprintf("filterOutFullyExecutedMessages(%v, %v)", tt.args.reports, tt.args.executedMessages)) {
				return
			}
			assert.Equalf(t, tt.want, got, "filterOutFullyExecutedMessages(%v, %v)", tt.args.reports, tt.args.executedMessages)
		})
	}
}

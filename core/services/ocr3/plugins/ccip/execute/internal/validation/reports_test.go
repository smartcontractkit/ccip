package validation

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func Test_CommitReportValidator(t *testing.T) {
	tests := []struct {
		name    string
		min     int
		reports []cciptypes.ExecutePluginCommitData
		want    []cciptypes.ExecutePluginCommitData
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "single report, enough observations",
			min:  1,
			reports: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
			},
			want: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
			},
			wantErr: assert.NoError,
		},
		{
			name: "single report, not enough observations",
			min:  2,
			reports: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "multiple reports, partial observations",
			min:  2,
			reports: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{3}},
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
			},
			want: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
			},
			wantErr: assert.NoError,
		},
		{
			name: "multiple reports for same root",
			min:  2,
			reports: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
				{MerkleRoot: [32]byte{1}, BlockNum: 2},
				{MerkleRoot: [32]byte{1}, BlockNum: 3},
				{MerkleRoot: [32]byte{1}, BlockNum: 4},
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
			},
			want: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "different executed messages same root",
			min:  2,
			reports: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{2, 3}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{3, 4}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{4, 5}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{5, 6}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
			},
			want: []cciptypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Initialize the validator
			idFunc := func(data cciptypes.ExecutePluginCommitData) [32]byte {
				return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
			}
			validator := NewValidator[cciptypes.ExecutePluginCommitData](tt.min, idFunc)
			for _, report := range tt.reports {
				err := validator.AddReport(report)
				require.NoError(t, err)
			}

			// Test the results
			got, err := validator.GetValidatedReports()
			if !tt.wantErr(t, err, "GetValidatedReports()") {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValidatedReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

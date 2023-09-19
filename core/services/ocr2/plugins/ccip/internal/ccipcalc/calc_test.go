package ccipcalc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestMergeEpochAndRound(t *testing.T) {
	type args struct {
		epoch uint32
		round uint8
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "zero round and epoch",
			args: args{epoch: 0, round: 0},
			want: 0,
		},
		{
			name: "avg case",
			args: args{
				epoch: 243,
				round: 15,
			},
			want: 62223,
		},
		{
			name: "largest epoch and round",
			args: args{
				epoch: math.MaxUint32,
				round: math.MaxUint8,
			},
			want: 1099511627775,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want,
				MergeEpochAndRound(tt.args.epoch, tt.args.round),
				"mergeEpochAndRound(%v, %v)", tt.args.epoch, tt.args.round)
		})
	}
}

func TestContiguousReqs(t *testing.T) {
	testCases := []struct {
		min    uint64
		max    uint64
		seqNrs []uint64
		exp    bool
	}{
		{min: 5, max: 10, seqNrs: []uint64{5, 6, 7, 8, 9, 10}, exp: true},
		{min: 5, max: 10, seqNrs: []uint64{5, 7, 8, 9, 10}, exp: false},
		{min: 5, max: 10, seqNrs: []uint64{5, 6, 7, 8, 9, 10, 11}, exp: false},
		{min: 5, max: 10, seqNrs: []uint64{}, exp: false},
		{min: 1, max: 1, seqNrs: []uint64{1}, exp: true},
	}

	for _, tc := range testCases {
		res := ContiguousReqs(logger.NullLogger, tc.min, tc.max, tc.seqNrs)
		assert.Equal(t, tc.exp, res)
	}
}

package ccip

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/logger"
)

func TestObservationFilter(t *testing.T) {
	lggr := logger.TestLogger(t)
	obs1 := RelayObservation{IntervalsByOnRamp: map[common.Address]blob_verifier.CCIPInterval{
		common.HexToAddress("0x5431F5F973781809D18643b87B44921b11355d81"): blob_verifier.CCIPInterval{Min: 1, Max: 10},
	}}
	b1, err := obs1.Marshal()
	require.NoError(t, err)
	nonEmpty := getNonEmptyObservations[RelayObservation](lggr, []types.AttributedObservation{{Observation: b1}, {Observation: []byte{}}})
	require.NoError(t, err)
	require.Equal(t, 1, len(nonEmpty))
	assert.Equal(t, nonEmpty[0].IntervalsByOnRamp, obs1.IntervalsByOnRamp)
}

func TestObservationSize(t *testing.T) {
	testParams := gopter.DefaultTestParameters()
	testParams.MinSuccessfulTests = 100
	p := gopter.NewProperties(testParams)
	p.Property("bounded observation size", prop.ForAll(func(min, max uint64) bool {
		o := ExecutionObservation{SeqNrs: []uint64{min, max}}
		b, err := o.Marshal()
		require.NoError(t, err)
		return len(b) <= MaxObservationLength
	}, gen.UInt64(), gen.UInt64()))
	p.TestingRun(t)
}

func TestGetMinMaxSequenceNumbers(t *testing.T) {
	tests := []struct {
		input     []ExecutionObservation
		f         int
		minSeqNum uint64
		maxSeqNum uint64
		err       error
	}{
		{[]ExecutionObservation{{SeqNrs: []uint64{1, 1}}}, 0, 1, 1, nil},
		{[]ExecutionObservation{{SeqNrs: []uint64{0, 0}}}, 1, 0, 0,
			fmt.Errorf("number of observations (%d) too low for given F (%d)", 1, 1)},
		{[]ExecutionObservation{{SeqNrs: []uint64{10, 9}}}, 0, 0, 0,
			errors.New("max seq num smaller than min")},
		{[]ExecutionObservation{{SeqNrs: []uint64{5, 6}}, {SeqNrs: []uint64{4, 6}}}, 1, 5, 6, nil},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("getMinMaxSequenceNumbers=%d", i), func(t *testing.T) {
			minSeqNum, maxSeqNum, err := getMinMaxSequenceNumbers(tc.input, tc.f)
			if tc.err == nil {
				if err != nil {
					t.Fatalf("got %v; want %v", err, tc.err)
				}
			} else {
				if err == nil {
					t.Fatalf("got %v; want %v", err, tc.err)
				} else {
					if err.Error() != tc.err.Error() {
						t.Fatalf("got %v; want %v", err, tc.err)
					}
				}
			}
			if minSeqNum != tc.minSeqNum {
				t.Fatalf("got %v; want %v", minSeqNum, tc.minSeqNum)
			} else if maxSeqNum != tc.maxSeqNum {
				t.Fatalf("got %v; want %v", maxSeqNum, tc.maxSeqNum)
			}
		})
	}
}

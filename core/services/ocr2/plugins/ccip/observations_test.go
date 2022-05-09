package ccip

import (
	"fmt"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestObservationSize(t *testing.T) {
	testParams := gopter.DefaultTestParameters()
	testParams.MinSuccessfulTests = 100
	p := gopter.NewProperties(testParams)
	p.Property("bounded observation size", prop.ForAll(func(min, max uint64) bool {
		o := Observation{MinSeqNum: min, MaxSeqNum: max}
		b, err := o.Marshal()
		require.NoError(t, err)
		return len(b) <= MaxObservationLength
	}, gen.UInt64(), gen.UInt64()))
	p.TestingRun(t)
}

func TestGetMinMaxSequenceNumbers(t *testing.T) {
	tests := []struct {
		input     []Observation
		f         int
		minSeqNum uint64
		maxSeqNum uint64
		err       error
	}{
		{[]Observation{{MinSeqNum: 1, MaxSeqNum: 1}}, 0, 1, 1, nil},
		{[]Observation{{MinSeqNum: 0, MaxSeqNum: 0}}, 1, 0, 0,
			fmt.Errorf("number of observations (%d) too low for given F (%d)", 1, 1)},
		{[]Observation{{MinSeqNum: 10, MaxSeqNum: 9}}, 0, 0, 0,
			errors.New("max seq num smaller than min")},
		{[]Observation{{MinSeqNum: 5, MaxSeqNum: 6}, {MinSeqNum: 4, MaxSeqNum: 6}}, 1, 5, 6, nil},
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

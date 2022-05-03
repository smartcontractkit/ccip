package ccip

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

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

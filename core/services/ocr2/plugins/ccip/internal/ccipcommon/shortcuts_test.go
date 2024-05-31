package ccipcommon

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/stretchr/testify/assert"
)

func TestGetMessageIDsAsHexString(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		hashes := make([]cciptypes.Hash, 10)
		for i := range hashes {
			hashes[i] = cciptypes.Hash(common.HexToHash(strconv.Itoa(rand.Intn(100000))))
		}

		msgs := make([]cciptypes.EVM2EVMMessage, len(hashes))
		for i := range msgs {
			msgs[i] = cciptypes.EVM2EVMMessage{MessageID: hashes[i]}
		}

		messageIDs := GetMessageIDsAsHexString(msgs)
		for i := range messageIDs {
			assert.Equal(t, hashes[i].String(), messageIDs[i])
		}
	})

	t.Run("empty", func(t *testing.T) {
		messageIDs := GetMessageIDsAsHexString(nil)
		assert.Empty(t, messageIDs)
	})
}

func TestFlattenUniqueSlice(t *testing.T) {
	testCases := []struct {
		name           string
		inputSlices    [][]int
		expectedOutput []int
	}{
		{name: "empty", inputSlices: nil, expectedOutput: []int{}},
		{name: "empty 2", inputSlices: [][]int{}, expectedOutput: []int{}},
		{name: "single", inputSlices: [][]int{{1, 2, 3, 3, 3, 4}}, expectedOutput: []int{1, 2, 3, 4}},
		{name: "simple", inputSlices: [][]int{{1, 2, 3}, {2, 3, 4}}, expectedOutput: []int{1, 2, 3, 4}},
		{
			name:           "more complex case",
			inputSlices:    [][]int{{1, 3}, {2, 4, 3}, {5, 2, -1, 7, 10}},
			expectedOutput: []int{1, 3, 2, 4, 5, -1, 7, 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FlattenUniqueSlice(tc.inputSlices...)
			assert.Equal(t, tc.expectedOutput, res)
		})
	}
}

func TestIsTxRevertError(t *testing.T) {
	testCases := []struct {
		name           string
		inputError     error
		expectedOutput bool
	}{
		{name: "empty", inputError: nil, expectedOutput: false},
		{name: "non-revert error", inputError: fmt.Errorf("nothing"), expectedOutput: false},
		{name: "geth error", inputError: fmt.Errorf("execution reverted"), expectedOutput: true},
		{name: "nethermind error", inputError: fmt.Errorf("VM execution error"), expectedOutput: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, IsTxRevertError(tc.inputError))
		})
	}
}

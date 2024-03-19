package ccipcommon

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciptypes"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	ccipdatamocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
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

func TestGetChainTokens(t *testing.T) {
	tokens := ccipcalc.EvmAddrsToGeneric(
		utils.RandomAddress(),
		utils.RandomAddress(),
		utils.RandomAddress(),
		utils.RandomAddress(),
		utils.RandomAddress(),
		utils.RandomAddress(),
	)

	testCases := []struct {
		name                  string
		feeTokens             []cciptypes.Address
		destTokens            [][]cciptypes.Address
		expectedFeeTokens     []cciptypes.Address
		expectedBridgedTokens []cciptypes.Address
	}{
		{
			name:                  "empty",
			feeTokens:             []cciptypes.Address{},
			destTokens:            [][]cciptypes.Address{{}},
			expectedFeeTokens:     []cciptypes.Address{},
			expectedBridgedTokens: []cciptypes.Address{},
		},
		{
			name:      "single offRamp",
			feeTokens: []cciptypes.Address{tokens[0]},
			destTokens: [][]cciptypes.Address{
				{tokens[1], tokens[2], tokens[3]},
			},
			expectedFeeTokens:     []cciptypes.Address{tokens[0]},
			expectedBridgedTokens: []cciptypes.Address{tokens[1], tokens[2], tokens[3]},
		},
		{
			name:      "multiple offRamps with distinct tokens",
			feeTokens: []cciptypes.Address{tokens[0]},
			destTokens: [][]cciptypes.Address{
				{tokens[1], tokens[2]},
				{tokens[3], tokens[4]},
				{tokens[5]},
			},
			expectedFeeTokens:     []cciptypes.Address{tokens[0]},
			expectedBridgedTokens: []cciptypes.Address{tokens[1], tokens[2], tokens[3], tokens[4], tokens[5]},
		},
		{
			name:      "overlapping tokens",
			feeTokens: []cciptypes.Address{tokens[0]},
			destTokens: [][]cciptypes.Address{
				{tokens[0], tokens[1], tokens[2], tokens[3]},
				{tokens[0], tokens[2], tokens[3], tokens[4], tokens[5]},
				{tokens[5]},
			},
			expectedFeeTokens:     []cciptypes.Address{tokens[0]},
			expectedBridgedTokens: []cciptypes.Address{tokens[0], tokens[1], tokens[2], tokens[3], tokens[4], tokens[5]},
		},
	}

	ctx := testutils.Context(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			priceRegistry := ccipdatamocks.NewPriceRegistryReader(t)
			priceRegistry.On("GetFeeTokens", ctx).Return(tc.feeTokens, nil).Maybe()

			var offRamps []ccipdata.OffRampReader
			for _, destTokens := range tc.destTokens {
				offRamp := ccipdatamocks.NewOffRampReader(t)
				offRamp.On("GetTokens", ctx).Return(cciptypes.OffRampTokens{DestinationTokens: destTokens}, nil).Once()
				offRamps = append(offRamps, offRamp)
			}

			feeTokens, destTokens, err := GetChainTokens(ctx, offRamps, priceRegistry)
			assert.NoError(t, err)

			assert.ElementsMatch(t, tc.expectedFeeTokens, feeTokens)
			assert.ElementsMatch(t, tc.expectedBridgedTokens, destTokens)
		})
	}
}

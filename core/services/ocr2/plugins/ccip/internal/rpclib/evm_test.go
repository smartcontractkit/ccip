package rpclib

import (
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

func TestDefaultEvmBatchCaller_BatchCallDynamicLimit(t *testing.T) {
	testCases := []struct {
		name                          string
		maxBatchSize                  int
		backOffMultiplier             int
		numCalls                      int
		expectedBatchSizesOnEachRetry []int
	}{
		{
			name:                          "defaults",
			maxBatchSize:                  DefaultRpcBatchSizeLimit,
			backOffMultiplier:             DefaultRpcBatchBackOffMultiplier,
			numCalls:                      200,
			expectedBatchSizesOnEachRetry: []int{100, 20, 4, 1},
		},
		{
			name:                          "base simple scenario",
			maxBatchSize:                  20,
			backOffMultiplier:             2,
			numCalls:                      100,
			expectedBatchSizesOnEachRetry: []int{20, 10, 5, 2, 1},
		},
		{
			name:                          "remainder",
			maxBatchSize:                  99,
			backOffMultiplier:             5,
			numCalls:                      100,
			expectedBatchSizesOnEachRetry: []int{99, 19, 3, 1},
		},
		{
			name:                          "large back off multiplier",
			maxBatchSize:                  20,
			backOffMultiplier:             18,
			numCalls:                      100,
			expectedBatchSizesOnEachRetry: []int{20, 1},
		},
		{
			name:                          "back off equal to batch size",
			maxBatchSize:                  20,
			backOffMultiplier:             20,
			numCalls:                      100,
			expectedBatchSizesOnEachRetry: []int{20, 1},
		},
		{
			name:                          "back off larger than batch size",
			maxBatchSize:                  20,
			backOffMultiplier:             220,
			numCalls:                      100,
			expectedBatchSizesOnEachRetry: []int{20, 1},
		},
		{
			name:                          "back off 1",
			maxBatchSize:                  20,
			backOffMultiplier:             1,
			numCalls:                      100,
			expectedBatchSizesOnEachRetry: []int{20, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			batchSizes := make([]int, 0)

			ec := mocks.NewClient(t)
			bc := NewDefaultEvmBatchCaller(logger.TestLogger(t), ec, tc.maxBatchSize, tc.backOffMultiplier)
			ctx := testutils.Context(t)
			calls := make([]EvmCall, tc.numCalls)
			for i := range calls {
				calls[i].abi = abihelpers.MustParseABI("[]")
			}
			ec.On("BatchCallContext", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
				evmCalls := args.Get(1).([]rpc.BatchElem)
				batchSizes = append(batchSizes, len(evmCalls))
			}).Return(errors.New("some error"))
			_, _ = bc.BatchCallDynamicLimitRetries(ctx, 123, calls)

			assert.Equal(t, tc.expectedBatchSizesOnEachRetry, batchSizes)
		})
	}

}

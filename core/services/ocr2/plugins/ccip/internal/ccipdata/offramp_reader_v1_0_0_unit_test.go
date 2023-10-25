package ccipdata

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestOffRampGetDestinationTokensFromSourceTokens(t *testing.T) {
	ctx := testutils.Context(t)
	const numSrcTokens = 10

	generateTokensAndOutputs := func() ([]common.Address, []common.Address, []rpclib.DataAndErr) {
		srcTks := make([]common.Address, numSrcTokens)
		dstTks := make([]common.Address, numSrcTokens)
		outputs := make([]rpclib.DataAndErr, numSrcTokens)
		for i := range srcTks {
			srcTks[i] = utils.RandomAddress()
			dstTks[i] = utils.RandomAddress()
			outputs[i] = rpclib.DataAndErr{
				Outputs: []any{dstTks[i]}, Err: nil,
			}
		}
		return srcTks, dstTks, outputs
	}

	t.Run("happy path", func(t *testing.T) {
		batchCaller := rpclib.NewMockEvmBatchCaller(t)
		o := &OffRampV1_0_0{evmBatchCaller: batchCaller}
		srcTks, dstTks, outputs := generateTokensAndOutputs()
		batchCaller.On("BatchCallLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(outputs, nil)
		actualDstTokens, err := o.GetDestinationTokensFromSourceTokens(ctx, srcTks)
		assert.NoError(t, err)
		assert.Equal(t, dstTks, actualDstTokens)
	})

	t.Run("rpc error", func(t *testing.T) {
		batchCaller := rpclib.NewMockEvmBatchCaller(t)
		o := &OffRampV1_0_0{evmBatchCaller: batchCaller}
		srcTks, _, outputs := generateTokensAndOutputs()
		outputs[0].Err = fmt.Errorf("some error")
		batchCaller.On("BatchCallLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(outputs, nil)
		_, err := o.GetDestinationTokensFromSourceTokens(ctx, srcTks)
		assert.Error(t, err)
	})

	t.Run("unexpected outputs count", func(t *testing.T) {
		batchCaller := rpclib.NewMockEvmBatchCaller(t)
		o := &OffRampV1_0_0{evmBatchCaller: batchCaller}

		srcTks, _, outputs := generateTokensAndOutputs()
		outputs[0].Outputs = append(outputs[0].Outputs, "unexpected", 123)
		batchCaller.On("BatchCallLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(outputs, nil)
		_, err := o.GetDestinationTokensFromSourceTokens(ctx, srcTks)
		assert.Error(t, err)
	})

	t.Run("unexpected output type", func(t *testing.T) {
		batchCaller := rpclib.NewMockEvmBatchCaller(t)
		o := &OffRampV1_0_0{evmBatchCaller: batchCaller}
		srcTks, _, outputs := generateTokensAndOutputs()
		outputs[0].Outputs[0] = "0x123"
		batchCaller.On("BatchCallLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(outputs, nil)
		_, err := o.GetDestinationTokensFromSourceTokens(ctx, srcTks)
		assert.Error(t, err)
	})
}

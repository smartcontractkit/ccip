package ccipdata

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestOffRampGetDestinationTokensFromSourceTokens(t *testing.T) {
	ctx := testutils.Context(t)
	const numSrcTokens = 20

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

	testCases := []struct {
		name           string
		outputChangeFn func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr
		expErr         bool
	}{
		{
			name:           "happy path",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr { return outputs },
			expErr:         false,
		},
		{
			name: "rpc error",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr {
				outputs[2].Err = fmt.Errorf("some error")
				return outputs
			},
			expErr: true,
		},
		{
			name: "unexpected outputs",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr {
				outputs[0].Outputs = append(outputs[0].Outputs, "unexpected", 123)
				return outputs
			},
			expErr: true,
		},
		{
			name: "unexpected output type",
			outputChangeFn: func(outputs []rpclib.DataAndErr) []rpclib.DataAndErr {
				outputs[0].Outputs = []any{utils.RandomAddress().String()}
				return outputs
			},
			expErr: true,
		},
	}

	lp := mocks.NewLogPoller(t)
	lp.On("LatestBlock", mock.Anything).Return(int64(rand.Uint64()), nil)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			batchCaller := rpclib.NewMockEvmBatchCaller(t)
			o := &OffRampV1_0_0{evmBatchCaller: batchCaller, lp: lp}
			srcTks, dstTks, outputs := generateTokensAndOutputs()
			outputs = tc.outputChangeFn(outputs)
			batchCaller.On("BatchCallDynamicLimitRetries", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(outputs, nil)
			actualDstTokens, err := o.GetDestinationTokensFromSourceTokens(ctx, srcTks)

			if tc.expErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, dstTks, actualDstTokens)
		})
	}
}

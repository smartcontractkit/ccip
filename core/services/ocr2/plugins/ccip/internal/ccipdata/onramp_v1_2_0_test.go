package ccipdata

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	evmClientMocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestLogPollerClient_GetSendRequestsGteSeqNum(t *testing.T) {
	onRampAddr := utils.RandomAddress()
	seqNum := uint64(100)
	confs := 4
	lggr := logger.TestLogger(t)
	t.Run("using confs", func(t *testing.T) {
		lp := mocks.NewLogPoller(t)
		lp.On("RegisterFilter", mock.Anything).Return(nil)
		onRampV2, err := NewOnRampV1_2_0(lggr, 1, 1, onRampAddr, lp, nil, false)
		require.NoError(t, err)
		lp.On("LogsDataWordGreaterThan",
			onRampV2.sendRequestedEventSig,
			onRampAddr,
			onRampV2.sendRequestedSeqNumberWord,
			abihelpers.EvmWord(seqNum),
			confs,
			mock.Anything,
		).Return([]logpoller.Log{}, nil)

		//c := &LogPollerReader{lp: lp}
		events, err := onRampV2.GetSendRequestsGteSeqNum(
			context.Background(),
			seqNum,
			confs,
		)
		assert.NoError(t, err)
		assert.Empty(t, events)
		lp.AssertExpectations(t)
	})

	t.Run("using latest confirmed block", func(t *testing.T) {
		h := &types.Header{Number: big.NewInt(100000)}
		cl := evmClientMocks.NewClient(t)
		cl.On("HeaderByNumber", mock.Anything, mock.Anything).Return(h, nil)
		lp := mocks.NewLogPoller(t)
		lp.On("RegisterFilter", mock.Anything).Return(nil)
		onRampV2, err := NewOnRampV1_2_0(lggr, 1, 1, onRampAddr, lp, cl, true)
		require.NoError(t, err)
		lp.On("LogsUntilBlockHashDataWordGreaterThan",
			onRampV2.sendRequestedEventSig,
			onRampAddr,
			onRampV2.sendRequestedSeqNumberWord,
			abihelpers.EvmWord(seqNum),
			h.Hash(),
			mock.Anything,
		).Return([]logpoller.Log{}, nil)

		events, err := onRampV2.GetSendRequestsGteSeqNum(
			context.Background(),
			seqNum,
			confs,
		)
		assert.NoError(t, err)
		assert.Empty(t, events)
		lp.AssertExpectations(t)
		cl.AssertExpectations(t)
	})
}

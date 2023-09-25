package ccipdata

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestLogPollerClient_GetLastUSDCMessagePriorToLogIndexInTx(t *testing.T) {
	txHash := utils.RandomAddress().Hash()
	ccipLogIndex := int64(100)

	expectedData := []byte("-1")

	t.Run("multiple found", func(t *testing.T) {
		lp := mocks.NewLogPoller(t)
		u, err := NewUSDCReader(utils.RandomAddress(), lp)
		require.NoError(t, err)
		lp.On("IndexedLogsByTxHash",
			u.usdcMessageSent,
			txHash,
			mock.Anything,
		).Return([]logpoller.Log{
			{LogIndex: ccipLogIndex - 2, Data: []byte("-2")},
			{LogIndex: ccipLogIndex - 1, Data: expectedData},
			{LogIndex: ccipLogIndex, Data: []byte("0")},
			{LogIndex: ccipLogIndex + 1, Data: []byte("1")},
		}, nil)

		usdcMessageData, err := u.GetLastUSDCMessagePriorToLogIndexInTx(context.Background(), ccipLogIndex, txHash)
		assert.NoError(t, err)
		assert.Equal(t, expectedData, usdcMessageData)

		lp.AssertExpectations(t)
	})

	t.Run("none found", func(t *testing.T) {
		lp := mocks.NewLogPoller(t)
		u, err := NewUSDCReader(utils.RandomAddress(), lp)
		require.NoError(t, err)
		lp.On("IndexedLogsByTxHash",
			u.usdcMessageSent,
			txHash,
			mock.Anything,
		).Return([]logpoller.Log{}, nil)

		usdcMessageData, err := u.GetLastUSDCMessagePriorToLogIndexInTx(context.Background(), ccipLogIndex, txHash)
		assert.Errorf(t, err, fmt.Sprintf("no USDC message found prior to log index %d in tx %s", ccipLogIndex, txHash.Hex()))
		assert.Nil(t, usdcMessageData)

		lp.AssertExpectations(t)
	})
}

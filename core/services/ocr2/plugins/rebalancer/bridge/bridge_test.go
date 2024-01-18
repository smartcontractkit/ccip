package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

func TestContainer(t *testing.T) {
	c := NewContainer()
	mockLP := mocks.NewLogPoller(t)
	mockLP.On("RegisterFilter", mock.Anything).Return(nil)
	b, err := NewEthereumToOptimism(mockLP, utils.RandomAddress())
	assert.NoError(t, err)

	source := models.NetworkSelector(1)
	dest := models.NetworkSelector(2)

	c.AddBridge(b, source, dest)
	gotB, exists := c.GetBridge(source, dest)
	assert.Equal(t, b, gotB)
	assert.True(t, exists)

	_, exists = c.GetBridge(source, models.NetworkSelector(3))
	assert.False(t, exists)

	_, exists = c.GetBridge(models.NetworkSelector(4), models.NetworkSelector(3))
	assert.False(t, exists)
}

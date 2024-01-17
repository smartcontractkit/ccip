package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

func TestContainer(t *testing.T) {
	c := NewContainer()
	b := NewEthereumToOptimism()

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

package ccipevents

import (
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogPollerClient_loadDependency(t *testing.T) {
	c := &LogPollerClient{}

	someAddr := utils.RandomAddress()

	onRamp, err := c.loadOnRamp(someAddr)
	assert.NoError(t, err)
	onRamp2, err := c.loadOnRamp(someAddr)
	// the objects should have the same pointer
	// since the second time the dependency should've been loaded from cache instead of initializing a new instance.
	assert.True(t, onRamp == onRamp2)

	offRamp, err := c.loadOffRamp(someAddr)
	assert.NoError(t, err)
	offRamp2, err := c.loadOffRamp(someAddr)
	assert.True(t, offRamp == offRamp2)

	priceReg, err := c.loadPriceRegistry(someAddr)
	assert.NoError(t, err)
	priceReg2, err := c.loadPriceRegistry(someAddr)
	assert.True(t, priceReg == priceReg2)
}

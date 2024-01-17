package bridge

import "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"

type Bridge interface {
	// PopulateStatusOfTransfers determines the transfer status and returns the transfers
	// with all the statuses being set.
	PopulateStatusOfTransfers(transfers []models.Transfer) ([]models.PendingTransfer, error)
}

type Container struct {
	bridges map[models.NetworkSelector]map[models.NetworkSelector]Bridge
}

func NewContainer() *Container {
	return &Container{
		bridges: make(map[models.NetworkSelector]map[models.NetworkSelector]Bridge),
	}
}

func (c *Container) AddBridge(b Bridge, source, dest models.NetworkSelector) {
	if _, exists := c.bridges[source]; !exists {
		c.bridges[source] = make(map[models.NetworkSelector]Bridge)
	}
	c.bridges[source][dest] = b
}

func (c *Container) GetBridge(source, dest models.NetworkSelector) (Bridge, bool) {
	if _, exists := c.bridges[source]; !exists {
		return nil, false
	}
	b, exists := c.bridges[source][dest]
	return b, exists
}

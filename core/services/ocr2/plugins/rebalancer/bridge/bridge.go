package bridge

import (
	"context"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type Bridge interface {
	// PopulateStatusOfTransfers determines the transfer status and returns the transfers
	// with all the statuses being set.
	PopulateStatusOfTransfers(ctx context.Context, token, sender models.Address, transfers []models.Transfer) ([]models.PendingTransfer, error)

	// Close should be called to release any resources the implementation holds.
	Close(ctx context.Context) error
	// todo: figure out where to call Close()
}

type Container interface {
	GetBridge(source, dest models.NetworkSelector) (Bridge, bool)
}

type BaseContainer struct {
	bridges map[models.NetworkSelector]map[models.NetworkSelector]Bridge
}

func NewContainer() *BaseContainer {
	return &BaseContainer{
		bridges: make(map[models.NetworkSelector]map[models.NetworkSelector]Bridge),
	}
}

func (c *BaseContainer) AddBridge(b Bridge, source, dest models.NetworkSelector) {
	if _, exists := c.bridges[source]; !exists {
		c.bridges[source] = make(map[models.NetworkSelector]Bridge)
	}
	c.bridges[source][dest] = b
}

func (c *BaseContainer) GetBridge(source, dest models.NetworkSelector) (Bridge, bool) {
	if _, exists := c.bridges[source]; !exists {
		return nil, false
	}
	b, exists := c.bridges[source][dest]
	return b, exists
}

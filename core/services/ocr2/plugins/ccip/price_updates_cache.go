package ccip

import (
	"time"
)

type priceUpdatesCache struct {
	lastUpdate update
}

func newPriceUpdatesCache() *priceUpdatesCache {
	return &priceUpdatesCache{
		lastUpdate: update{},
	}
}

func (c *priceUpdatesCache) containsData() bool {
	return c.lastUpdate.timestamp != time.Time{}
}

func (c *priceUpdatesCache) lastCheckpoint() time.Time {
	return c.lastUpdate.timestamp
}

func (c *priceUpdatesCache) get() update {
	return c.lastUpdate
}

func (c *priceUpdatesCache) updateCache(update update) {
	if update.timestamp.After(c.lastUpdate.timestamp) {
		c.lastUpdate = update
	}
}

package ccip

type priceUpdatesCache struct {
	lastUpdate update
}

func newPriceUpdatesCache() *priceUpdatesCache {
	return &priceUpdatesCache{
		lastUpdate: update{},
	}
}

func (c *priceUpdatesCache) get() update {
	return c.lastUpdate
}

func (c *priceUpdatesCache) updateCache(update update) {
	if update.timestamp.After(c.lastUpdate.timestamp) {
		c.lastUpdate = update
	}
}

package ccip

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type priceUpdatesCache struct {
	tokenPriceUpdates map[common.Address]update
	gasPriceUpdate    update
	mu                *sync.RWMutex
}

func newPriceUpdatesCache() *priceUpdatesCache {
	return &priceUpdatesCache{
		tokenPriceUpdates: make(map[common.Address]update),
		mu:                &sync.RWMutex{},
	}
}

func (c *priceUpdatesCache) mostRecentTokenPriceUpdate() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ts := time.Time{}
	for _, upd := range c.tokenPriceUpdates {
		if upd.timestamp.After(ts) {
			ts = upd.timestamp
		}
	}
	return ts
}

func (c *priceUpdatesCache) updateTokenPriceIfMoreRecent(ts time.Time, tk common.Address, val *big.Int) bool {
	c.mu.RLock()
	v, exists := c.tokenPriceUpdates[tk]
	c.mu.RUnlock()

	if !exists || v.timestamp.Before(ts) {
		c.mu.Lock()
		c.tokenPriceUpdates[tk] = update{timestamp: ts, value: val}
		c.mu.Unlock()
		return true
	}

	return false
}

// getTokenPriceUpdates returns all the price updates with timestamp greater than or equal to the provided
func (c *priceUpdatesCache) getTokenPriceUpdates(minTs time.Time) map[common.Address]update {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cp := make(map[common.Address]update, len(c.tokenPriceUpdates))
	for k, v := range c.tokenPriceUpdates {
		if v.timestamp.Before(minTs) {
			continue
		}
		cp[k] = v
	}
	return cp
}

func (c *priceUpdatesCache) getGasPriceUpdate() update {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.gasPriceUpdate
}

func (c *priceUpdatesCache) updateGasPriceIfMoreRecent(update update) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if update.timestamp.After(c.gasPriceUpdate.timestamp) {
		c.gasPriceUpdate = update
		return true
	}

	return false
}

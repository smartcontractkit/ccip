package ccip

import (
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type tokenPriceUpdatesCache struct {
	mem map[common.Address]update
	mu  *sync.RWMutex
}

func newTokenPriceUpdatesCache() *tokenPriceUpdatesCache {
	c := &tokenPriceUpdatesCache{
		mem: make(map[common.Address]update),
		mu:  &sync.RWMutex{},
	}
	return c
}

func (c *tokenPriceUpdatesCache) mostRecentTs() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ts := time.Time{}
	for _, upd := range c.mem {
		if upd.timestamp.After(ts) {
			ts = upd.timestamp
		}
	}
	return ts
}

func (c *tokenPriceUpdatesCache) updateIfMoreRecent(ts time.Time, tk common.Address, val *big.Int) bool {
	c.mu.RLock()
	v, exists := c.mem[tk]
	c.mu.RUnlock()

	if !exists || v.timestamp.Before(ts) {
		c.mu.Lock()
		c.mem[tk] = update{timestamp: ts, value: val}
		c.mu.Unlock()
		return true
	}

	return false
}

// get returns all the price updates with timestamp greater than or equal to the provided
func (c *tokenPriceUpdatesCache) get(minTs time.Time) map[common.Address]update {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cp := make(map[common.Address]update, len(c.mem))
	for k, v := range c.mem {
		if v.timestamp.Before(minTs) {
			continue
		}
		cp[k] = v
	}
	return cp
}

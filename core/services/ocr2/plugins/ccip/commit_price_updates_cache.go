package ccip

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type tokenPriceUpdatesCache struct {
	mem    map[common.Address]update
	mu     *sync.RWMutex
	expiry time.Duration
}

func newTokenPriceUpdatesCache(ctx context.Context, expiry time.Duration) *tokenPriceUpdatesCache {
	c := &tokenPriceUpdatesCache{
		mem:    make(map[common.Address]update),
		mu:     &sync.RWMutex{},
		expiry: expiry,
	}
	go c.expirationWorker(ctx)
	return c
}

func (c *tokenPriceUpdatesCache) expirationWorker(ctx context.Context) {
	tick := time.NewTicker(c.expiry)

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			c.mu.Lock()
			c.mem = make(map[common.Address]update)
			c.mu.Unlock()
		}
	}
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

func (c *tokenPriceUpdatesCache) get() map[common.Address]update {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cp := make(map[common.Address]update, len(c.mem))
	for k, v := range c.mem {
		cp[k] = v
	}
	return cp
}

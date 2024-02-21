package rebalancer

import (
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type InflightCache interface {
	Add(lggr logger.Logger, trs []models.Transfer)
	Expire(lggr logger.Logger)
	Get() []models.Transfer
}

type timedTransfer struct {
	transfer  models.Transfer
	createdAt time.Time
}

// entryKey is a non-perfect way to identify an in-flight transfer.
// it's probably accurate enough for very small timescales, but it's not
// perfect. It's possible that two transfers with the same from, to, and
// amount could be different transfers over all time.
type entryKey struct {
	From   models.NetworkSelector
	To     models.NetworkSelector
	Amount string
}

type inflightCache struct {
	mu         sync.RWMutex
	entries    map[entryKey]timedTransfer
	expiryTime time.Duration
}

func NewInflightCache(expiryTime time.Duration) InflightCache {
	return &inflightCache{
		mu:      sync.RWMutex{},
		entries: make(map[entryKey]timedTransfer),
	}
}

func (c *inflightCache) Add(lggr logger.Logger, trs []models.Transfer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, tr := range trs {
		_, exists := c.entries[entryKey{
			From:   tr.From,
			To:     tr.To,
			Amount: tr.Amount.String(),
		}]
		if exists {
			lggr.Infow("Skipping adding transfer to inflight cache, already exists", "transfer", tr)
			continue
		}
		c.entries[entryKey{
			From:   tr.From,
			To:     tr.To,
			Amount: tr.Amount.String(),
		}] = timedTransfer{
			transfer:  tr,
			createdAt: time.Now(),
		}
	}
	lggr.Infow("Added send transfers to inflight cache", "transfers", trs)
}

func (c *inflightCache) Expire(lggr logger.Logger) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, entry := range c.entries {
		if time.Since(entry.createdAt) >= c.expiryTime {
			lggr.Infow("Pruning expired transfer", "transfer", entry.transfer)
			delete(c.entries, k)
		}
	}
}

func (c *inflightCache) Get() []models.Transfer {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var transfers []models.Transfer
	for _, entry := range c.entries {
		transfers = append(transfers, entry.transfer)
	}
	return transfers
}

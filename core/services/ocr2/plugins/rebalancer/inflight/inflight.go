package inflight

import (
	"context"
	"sync"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type Container interface {
	// Add adds a transfer to the inflight container.
	Add(ctx context.Context, t models.Transfer)
	// Expire removes any transfers from the inflight container that are in the pending list.
	Expire(ctx context.Context, pending []models.PendingTransfer)
	// GetAll returns all transfers in the inflight container.
	GetAll(ctx context.Context) []models.Transfer
	// IsInflight returns true if the transfer is in the inflight container.
	IsInflight(ctx context.Context, t models.Transfer) bool
}

type mapKey struct {
	From   models.NetworkSelector
	To     models.NetworkSelector
	Amount string
}

type inflight struct {
	items map[mapKey]models.Transfer
	mu    sync.RWMutex
	lggr  logger.Logger
}

func New(lggr logger.Logger) Container {
	return &inflight{
		items: make(map[mapKey]models.Transfer),
		lggr:  lggr,
	}
}

func (i *inflight) Add(ctx context.Context, t models.Transfer) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.items[mapKey{
		From:   t.From,
		To:     t.To,
		Amount: t.Amount.String(),
	}] = t
}

func (i *inflight) Expire(ctx context.Context, pending []models.PendingTransfer) {
	i.mu.Lock()
	defer i.mu.Unlock()

	var numExpired int
	for _, p := range pending {
		k := mapKey{
			From:   p.From,
			To:     p.To,
			Amount: p.Amount.String(),
		}
		_, ok := i.items[k]
		if ok {
			numExpired++
			delete(i.items, k)
		}
	}

	if numExpired > 0 {
		i.lggr.Debugw("Expired inflight transfers", "numExpired", numExpired)
	}
}

func (i *inflight) GetAll(ctx context.Context) []models.Transfer {
	i.mu.RLock()
	defer i.mu.RUnlock()

	var transfers []models.Transfer
	for k := range i.items {
		transfers = append(transfers, i.items[k])
	}
	return transfers
}

func (i *inflight) IsInflight(ctx context.Context, t models.Transfer) bool {
	i.mu.RLock()
	defer i.mu.RUnlock()

	_, ok := i.items[mapKey{
		From:   t.From,
		To:     t.To,
		Amount: t.Amount.String(),
	}]
	return ok
}

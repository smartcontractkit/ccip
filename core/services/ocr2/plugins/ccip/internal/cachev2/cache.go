package cachev2

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type Cache[T any] interface {
	Get(ctx context.Context, syncFunc func(ctx context.Context) (T, error)) (T, error)
}

func NewLPCache[T any](
	lp logpoller.LogPoller,
	observedEvents []common.Hash,
	contractAddress common.Address,
	optimisticConfirmations int64,
) *LPCache[T] {
	var emptyValue T
	return &LPCache[T]{
		logPoller:               lp,
		observedEvents:          observedEvents,
		address:                 contractAddress,
		optimisticConfirmations: optimisticConfirmations,

		lock:            &sync.RWMutex{},
		value:           emptyValue,
		lastChangeBlock: 0,
	}
}

func (c *LPCache[T]) Get(ctx context.Context, syncFunc func(ctx context.Context) (T, error)) (T, error) {
	var empty T

	hasExpired, newEventBlockNum, err := c.hasExpired(ctx)
	if err != nil {
		return empty, fmt.Errorf("check cache expiration: %w", err)
	}

	if hasExpired {
		latestValue, err := syncFunc(ctx)
		if err != nil {
			return empty, fmt.Errorf("sync func: %w", err)
		}

		c.set(ctx, latestValue, newEventBlockNum)
		return latestValue, nil
	}

	cachedValue, _, err := c.get(ctx)
	if err != nil {
		return empty, fmt.Errorf("get cached value: %w", err)
	}
	return cachedValue, nil
}

type LPCache[T any] struct {
	logPoller               logpoller.LogPoller
	observedEvents          []common.Hash
	address                 common.Address
	optimisticConfirmations int64

	lock            *sync.RWMutex
	value           T
	lastChangeBlock int64
}

func (c *LPCache[T]) hasExpired(ctx context.Context) (expired bool, blockOfLatestEvent int64, err error) {
	c.lock.RLock()
	blockOfCurrentValue := c.lastChangeBlock
	c.lock.RUnlock()

	if blockOfCurrentValue == 0 {
		return true, 0, nil
	}

	blockOfLatestEvent, err = c.logPoller.LatestBlockByEventSigsAddrsWithConfs(
		c.lastChangeBlock,
		c.observedEvents,
		[]common.Address{c.address},
		logpoller.Confirmations(c.optimisticConfirmations),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return false, 0, fmt.Errorf("get latest events form lp: %w", err)
	}
	return blockOfLatestEvent > blockOfCurrentValue, blockOfLatestEvent, nil
}

func (c *LPCache[T]) set(_ context.Context, value T, blockNum int64) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.lastChangeBlock > blockNum {
		return false
	}

	c.value = value
	c.lastChangeBlock = blockNum
	return true
}

func (c *LPCache[T]) get(_ context.Context) (T, int64, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.value, c.lastChangeBlock, nil
}

package cache

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type AutoSync[T any] interface {
	Get(ctx context.Context, syncFunc func(ctx context.Context) (T, error)) (T, error)
	SetOptimisticConfirmations(confs int64)
}

type LogpollerEventsBased[T any] struct {
	logPoller               logpoller.LogPoller
	observedEvents          []common.Hash
	address                 common.Address
	optimisticConfirmations int64

	lock            *sync.RWMutex
	value           T
	lastChangeBlock int64
}

func NewLogpollerEventsBased[T any](
	lp logpoller.LogPoller,
	observedEvents []common.Hash,
	contractAddress common.Address,
	optimisticConfirmations int64,
) *LogpollerEventsBased[T] {
	var emptyValue T
	return &LogpollerEventsBased[T]{
		logPoller:               lp,
		observedEvents:          observedEvents,
		address:                 contractAddress,
		optimisticConfirmations: optimisticConfirmations,

		lock:            &sync.RWMutex{},
		value:           emptyValue,
		lastChangeBlock: 0,
	}
}

func (c *LogpollerEventsBased[T]) Get(ctx context.Context, syncFunc func(ctx context.Context) (T, error)) (T, error) {
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

	c.lock.Lock()
	if newEventBlockNum > c.lastChangeBlock {
		// update the most recent block number
		// that way the scanning window is shorter in the next run
		c.lastChangeBlock = newEventBlockNum
	}
	c.lock.Unlock()

	return cachedValue, nil
}

func (c *LogpollerEventsBased[T]) SetOptimisticConfirmations(confs int64) {
	c.lock.Lock()
	c.optimisticConfirmations = confs
	c.lock.Unlock()
}

func (c *LogpollerEventsBased[T]) hasExpired(ctx context.Context) (expired bool, blockOfLatestEvent int64, err error) {
	c.lock.RLock()
	blockOfCurrentValue := c.lastChangeBlock
	c.lock.RUnlock()

	latestBlock, err := c.logPoller.LatestBlock(pg.WithParentCtx(ctx))
	latestBlockNumber := int64(0)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return false, 0, fmt.Errorf("get latest log poller block: %w", err)
		}
	} else {
		latestBlockNumber = latestBlock.BlockNumber
	}

	if blockOfCurrentValue == 0 {
		return true, latestBlockNumber, nil
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

	if blockOfLatestEvent > latestBlockNumber {
		latestBlockNumber = blockOfLatestEvent
	}
	return blockOfLatestEvent > blockOfCurrentValue, latestBlockNumber, nil
}

func (c *LogpollerEventsBased[T]) set(_ context.Context, value T, blockNum int64) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.lastChangeBlock > blockNum {
		return false
	}

	c.value = value
	c.lastChangeBlock = blockNum
	return true
}

func (c *LogpollerEventsBased[T]) get(_ context.Context) (T, int64, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.value, c.lastChangeBlock, nil
}

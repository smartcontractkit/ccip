package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	cciporm "github.com/smartcontractkit/chainlink/v2/core/services/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

// PriceCleanup deletes old gas and token prices recorded in db.
// CCIP plugin only uses the latest prices. To make debugging and updating easier, new prices are always inserted
// to the DB, as opposed to overwriting the old ones.
// As a result, we need to clean up the old prices periodically. This is the service that does the cleanup.
type PriceCleanup interface {
	job.ServiceCtx
}

var _ PriceCleanup = (*priceCleanup)(nil)

const (
	// Prices should be cleaned up after 8 hours
	priceExpireSec = 28800
	// Run this cleanup service every 8 hours
	cleanupInterval = priceExpireSec * time.Second
)

type priceCleanup struct {
	priceExpireSec  int
	cleanupInterval time.Duration

	lggr              logger.Logger
	orm               cciporm.ORM
	destChainSelector uint64

	services.StateMachine
	wg               *sync.WaitGroup
	backgroundCtx    context.Context
	backgroundCancel context.CancelFunc
}

func NewPriceCleanup(lggr logger.Logger, orm cciporm.ORM, destChainSelector uint64) PriceCleanup {
	ctx, cancel := context.WithCancel(context.Background())

	pc := &priceCleanup{
		priceExpireSec:  priceExpireSec,
		cleanupInterval: cleanupInterval,

		lggr:              lggr,
		orm:               orm,
		destChainSelector: destChainSelector,

		wg:               new(sync.WaitGroup),
		backgroundCtx:    ctx,
		backgroundCancel: cancel,
	}
	return pc
}

func (c *priceCleanup) Start(context.Context) error {
	return c.StateMachine.StartOnce("PriceCleanup", func() error {
		c.lggr.Info("Starting PriceCleanup")
		c.wg.Add(1)
		c.run()
		return nil
	})
}

func (c *priceCleanup) Close() error {
	return c.StateMachine.StopOnce("PriceCleanup", func() error {
		c.lggr.Info("Closing PriceCleanup")
		c.backgroundCancel()
		c.wg.Wait()
		return nil
	})
}

func (c *priceCleanup) run() {
	ticker := time.NewTicker(c.cleanupInterval)
	go func() {
		defer c.wg.Done()
		// Do the first cleanup immediately upon starting the service
		_ = c.clean(c.backgroundCtx)

		for {
			select {
			case <-c.backgroundCtx.Done():
				return
			case <-ticker.C:
				err := c.clean(c.backgroundCtx)
				if err != nil {
					c.lggr.Errorw("Failed to cleanup in-db prices in the background", "err", err)
				}
			}
		}
	}()
}

func (c *priceCleanup) clean(ctx context.Context) error {
	eg := new(errgroup.Group)

	eg.Go(func() error {
		err := c.orm.ClearGasPricesByDestChain(ctx, c.destChainSelector, c.priceExpireSec)
		if err != nil {
			return fmt.Errorf("error clearing gas prices: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		err := c.orm.ClearTokenPricesByDestChain(ctx, c.destChainSelector, c.priceExpireSec)
		if err != nil {
			return fmt.Errorf("error clearing token prices: %w", err)
		}
		return nil
	})

	return eg.Wait()
}

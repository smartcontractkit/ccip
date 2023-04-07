package ccip

import (
	"context"
	"sync"
	"time"

	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

type BackfilledOracle struct {
	srcStartBlock, dstStartBlock int64
	src, dst                     logpoller.LogPoller
	oracle                       job.ServiceCtx
	lggr                         logger.Logger
}

func NewBackfilledOracle(lggr logger.Logger, src, dst logpoller.LogPoller, srcStartBlock, dstStartBlock int64, oracle job.ServiceCtx) *BackfilledOracle {
	return &BackfilledOracle{
		srcStartBlock: srcStartBlock,
		dstStartBlock: dstStartBlock,
		src:           src,
		dst:           dst,
		oracle:        oracle,
		lggr:          lggr,
	}
}

func (r *BackfilledOracle) Start(ctx context.Context) error {
	go func() {
		var err error
		var errMu sync.Mutex
		var wg sync.WaitGroup
		// Replay in parallel if both requested.
		if r.srcStartBlock != 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s := time.Now()
				r.lggr.Infow("start replaying src chain", "fromBlock", r.dstStartBlock)
				srcReplayErr := r.src.Replay(context.Background(), r.srcStartBlock)
				errMu.Lock()
				err = multierr.Combine(err, srcReplayErr)
				errMu.Unlock()
				r.lggr.Infow("finished replaying src chain", "time", time.Since(s))
			}()
		}
		if r.dstStartBlock != 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				s := time.Now()
				r.lggr.Infow("start replaying dst chain", "fromBlock", r.dstStartBlock)
				dstReplayErr := r.dst.Replay(context.Background(), r.dstStartBlock)
				errMu.Lock()
				err = multierr.Combine(err, dstReplayErr)
				errMu.Unlock()
				r.lggr.Infow("finished replaying  dst chain", "time", time.Since(s))
			}()
		}
		wg.Wait()
		if err != nil {
			r.lggr.Errorw("unexpected error replaying", "err", err)
			return
		}
		// Start oracle with all logs present from dstStartBlock on dst and
		// all logs from srcStartBlock on src.
		if err := r.oracle.Start(ctx); err != nil {
			// Should never happen.
			r.lggr.Errorw("unexpected error starting oracle", "err", err)
		}
	}()
	return nil
}

func (r *BackfilledOracle) Close() error {
	return r.oracle.Close()
}

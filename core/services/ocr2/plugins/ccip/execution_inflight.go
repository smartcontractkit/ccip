package ccip

import (
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/logger"
)

type InflightInternalExecutionReport struct {
	createdAt   time.Time
	seqNrs      []uint64
	encMessages [][]byte
}

// inflightReportsContainer holds existing inflight reports.
// it provides a thread-safe access as it is called from multiple goroutines,
// e.g. reporting and transmission protocols.
type inflightReportsContainer struct {
	locker  *sync.RWMutex
	reports []InflightInternalExecutionReport

	cacheExpiry time.Duration
}

func newInflightReportsContainer(inflightCacheExpiry time.Duration) *inflightReportsContainer {
	return &inflightReportsContainer{
		locker:      &sync.RWMutex{},
		reports:     make([]InflightInternalExecutionReport, 0),
		cacheExpiry: inflightCacheExpiry,
	}
}

func (container *inflightReportsContainer) getAll() []InflightInternalExecutionReport {
	container.locker.RLock()
	defer container.locker.RUnlock()

	reports := make([]InflightInternalExecutionReport, len(container.reports))
	copy(reports[:], container.reports[:])

	return reports
}

func (container *inflightReportsContainer) expire(lggr logger.Logger) {
	container.locker.Lock()
	defer container.locker.Unlock()
	// Reap old inflight txs and check if any messages in the report are inflight.
	var stillInFlight []InflightInternalExecutionReport
	for _, report := range container.reports {
		if time.Since(report.createdAt) > container.cacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the change so we retry.
			lggr.Infow("Inflight report expired", "seqNums", report.seqNrs)
		} else {
			stillInFlight = append(stillInFlight, report)
		}
	}
	container.reports = stillInFlight
}

func (container *inflightReportsContainer) add(lggr logger.Logger, seqNrs []uint64, encMsgs [][]byte) error {
	container.locker.Lock()
	defer container.locker.Unlock()

	for _, report := range container.reports {
		// TODO: Think about if this fails in reorgs
		if (len(report.seqNrs) > 0 && len(seqNrs) > 0) && (report.seqNrs[0] == seqNrs[0]) {
			return errors.Errorf("report is already in flight")
		}
	}
	// Otherwise not already in flight, add it.
	lggr.Infow("Added report to inflight",
		"seqNums", seqNrs)
	container.reports = append(container.reports, InflightInternalExecutionReport{
		createdAt:   time.Now(),
		seqNrs:      seqNrs,
		encMessages: encMsgs,
	})
	return nil
}

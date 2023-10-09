package executionplugin

import (
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

// InflightInternalExecutionReport serves the same purpose as InflightCommitReport
// see the comment on that struct for context.
type inflightInternalReport struct {
	createdAt time.Time
	messages  []internal.EVM2EVMMessage
}

// inflightExecReportsContainer holds existing inflight reports.
// it provides a thread-safe access as it is called from multiple goroutines,
// e.g. reporting and transmission protocols.
type inflightReportsContainer struct {
	locker  sync.RWMutex
	reports []inflightInternalReport

	cacheExpiry time.Duration
}

func newInflightExecReportsContainer(inflightCacheExpiry time.Duration) *inflightReportsContainer {
	return &inflightReportsContainer{
		locker:      sync.RWMutex{},
		reports:     make([]inflightInternalReport, 0),
		cacheExpiry: inflightCacheExpiry,
	}
}

func (container *inflightReportsContainer) getAll() []inflightInternalReport {
	container.locker.RLock()
	defer container.locker.RUnlock()

	reports := make([]inflightInternalReport, len(container.reports))
	copy(reports[:], container.reports[:])

	return reports
}

func (container *inflightReportsContainer) expire(lggr logger.Logger) {
	container.locker.Lock()
	defer container.locker.Unlock()
	// Reap old inflight txs and check if any messages in the report are inflight.
	var stillInFlight []inflightInternalReport
	for _, report := range container.reports {
		if time.Since(report.createdAt) > container.cacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the change so we retry.
			lggr.Infow("Inflight report expired", "messages", report.messages)
		} else {
			stillInFlight = append(stillInFlight, report)
		}
	}
	container.reports = stillInFlight
}

func (container *inflightReportsContainer) add(lggr logger.Logger, messages []internal.EVM2EVMMessage) error {
	container.locker.Lock()
	defer container.locker.Unlock()

	for _, report := range container.reports {
		if (len(report.messages) > 0) && (report.messages[0].SequenceNumber == messages[0].SequenceNumber) {
			return errors.Errorf("report is already in flight")
		}
	}

	// Otherwise not already in flight, add it.
	lggr.Info("Inflight report added")
	container.reports = append(container.reports, inflightInternalReport{
		createdAt: time.Now(),
		messages:  messages,
	})
	return nil
}

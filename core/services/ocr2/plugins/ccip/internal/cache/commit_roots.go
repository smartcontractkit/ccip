package cache

import (
	"context"
	"encoding/hex"
	"slices"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	orderedmap "github.com/wk8/go-ordered-map/v2"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

const (
	// CleanupInterval defines how often roots cache is scanned to evict stale roots
	CleanupInterval = 30 * time.Minute
)

type CommitsRootsCache interface {
	RootsEligibleForExecution(ctx context.Context) ([]ccip.CommitStoreReport, error)
	MarkAsExecuted(merkleRoot [32]byte)
	Snooze(merkleRoot [32]byte)
}

func NewCommitRootsCache(
	lggr logger.Logger,
	reader ccip.CommitStoreReader,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
) CommitsRootsCache {
	return newCommitRootsCache(
		lggr,
		reader,
		messageVisibilityInterval,
		rootSnoozeTime,
		CleanupInterval,
	)
}

func newCommitRootsCache(
	lggr logger.Logger,
	reader ccip.CommitStoreReader,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
	cleanupInterval time.Duration,
) *commitRootsCache {
	snoozedRoots := cache.New(rootSnoozeTime, cleanupInterval)

	return &commitRootsCache{
		lggr:                        lggr,
		reader:                      reader,
		finalizedRoots:              orderedmap.New[string, ccip.CommitStoreReportWithTxMeta](),
		snoozedRoots:                snoozedRoots,
		messageVisibilityInterval:   messageVisibilityInterval,
		latestFinalizedCommitRootTs: time.Now().Add(-messageVisibilityInterval),
		cacheMu:                     sync.RWMutex{},
	}
}

type commitRootsCache struct {
	lggr                      logger.Logger
	reader                    ccip.CommitStoreReader
	messageVisibilityInterval time.Duration

	// Mutable state
	cacheMu                     sync.RWMutex
	finalizedRoots              *orderedmap.OrderedMap[string, ccip.CommitStoreReportWithTxMeta]
	snoozedRoots                *cache.Cache
	latestFinalizedCommitRootTs time.Time
}

func (r *commitRootsCache) RootsEligibleForExecution(ctx context.Context) ([]ccip.CommitStoreReport, error) {
	// 1. Fetch all the logs from the database after the latest finalized commit root timestamp
	logs, err := r.fetchLogsFromCommitStore(ctx)
	if err != nil {
		return nil, err
	}

	// 2. Iterate over the logs and check if the root is finalized or not. Return finalized and unfinalized reports
	finalizedReports, unfinalizedReports := r.updateFinalizedRoots(logs)

	// 3. Join finalized commit reports with unfinalized reports and outfilter snoozed roots.
	return r.pickReadyToExecute(finalizedReports, unfinalizedReports), nil

}

func (r *commitRootsCache) MarkAsExecuted(merkleRoot [32]byte) {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()

	prettyMerkleRoot := merkleRootToString(merkleRoot)
	r.finalizedRoots.Delete(prettyMerkleRoot)
}

func (r *commitRootsCache) Snooze(merkleRoot [32]byte) {
	r.snoozedRoots.SetDefault(merkleRootToString(merkleRoot), struct{}{})
}

func (r *commitRootsCache) isSnoozed(merkleRoot [32]byte) bool {
	_, snoozed := r.snoozedRoots.Get(merkleRootToString(merkleRoot))
	return snoozed
}

func (r *commitRootsCache) fetchLogsFromCommitStore(ctx context.Context) ([]ccip.CommitStoreReportWithTxMeta, error) {
	r.cacheMu.Lock()
	messageVisibilityWindow := time.Now().Add(-r.messageVisibilityInterval)
	if r.latestFinalizedCommitRootTs.Before(messageVisibilityWindow) {
		r.latestFinalizedCommitRootTs = messageVisibilityWindow
	}
	commitRootsFilterTimestamp := r.latestFinalizedCommitRootTs
	r.cacheMu.Unlock()

	// IO operation, release lock before!
	return r.reader.GetAcceptedCommitReportsGteTimestamp(ctx, commitRootsFilterTimestamp, 0)
}

func (r *commitRootsCache) updateFinalizedRoots(logs []ccip.CommitStoreReportWithTxMeta) ([]ccip.CommitStoreReportWithTxMeta, []ccip.CommitStoreReportWithTxMeta) {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()

	// Assuming logs are properly ordered by block_timestamp, log_index
	var unfinalizedReports []ccip.CommitStoreReportWithTxMeta
	for _, log := range logs {
		if _, finalized := r.finalizedRoots.Get(merkleRootToString(log.MerkleRoot)); finalized {
			r.finalizedRoots.Store(merkleRootToString(log.MerkleRoot), log)
		} else {
			unfinalizedReports = append(unfinalizedReports, log)
		}
	}

	if r.finalizedRoots.Newest() != nil {
		r.latestFinalizedCommitRootTs = time.UnixMilli(r.finalizedRoots.Newest().Value.BlockTimestampUnixMilli)
	}

	var finalizedRoots []ccip.CommitStoreReportWithTxMeta
	for pair := r.finalizedRoots.Oldest(); pair != nil; pair = pair.Next() {
		// Evict stale items
		if time.UnixMilli(pair.Value.BlockTimestampUnixMilli).Before(time.Now().Add(-r.messageVisibilityInterval)) {
			r.finalizedRoots.Delete(pair.Key)
			continue
		}
		finalizedRoots = append(finalizedRoots, pair.Value)
	}
	return finalizedRoots, unfinalizedReports
}

func (r *commitRootsCache) pickReadyToExecute(r1 []ccip.CommitStoreReportWithTxMeta, r2 []ccip.CommitStoreReportWithTxMeta) []ccip.CommitStoreReport {
	allReports := append(r1, r2...)
	eligibleReports := make([]ccip.CommitStoreReport, 0, len(allReports))
	for _, report := range allReports {
		if r.isSnoozed(report.MerkleRoot) {
			continue
		}
		eligibleReports = append(eligibleReports, report.CommitStoreReport)
	}
	// safety check
	slices.SortFunc(eligibleReports, func(i, j ccip.CommitStoreReport) int {
		return int(i.Interval.Min - j.Interval.Min)
	})
	return eligibleReports
}

func merkleRootToString(merkleRoot [32]byte) string {
	return hex.EncodeToString(merkleRoot[:])
}

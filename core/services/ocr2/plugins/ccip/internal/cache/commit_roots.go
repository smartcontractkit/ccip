package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

const (
	// EvictionGracePeriod defines how long after the permissionless execution threshold a root is still kept in the cache
	EvictionGracePeriod = 1 * time.Hour
	// CleanupInterval defines how often roots have to be evicted
	CleanupInterval = 30 * time.Minute
)

type CommitsRootsCache interface {
	IsSnoozed(merkleRoot [32]byte) bool
	MarkAsExecuted(merkleRoot [32]byte)
	Snooze(merkleRoot [32]byte)

	CommitSearchTimestamp() time.Time
	AppendUnexecutedRoot(merkleRoot [32]byte, blockTimestamp time.Time)
}

type commitRootsCache struct {
	// snoozedRoots is used to keep track of the roots that are snoozed. Roots that are executed and finalized are considered
	// marked kept in the cache for permissionLessExecutionThresholdDuration + EvictionGracePeriod
	snoozedRoots *cache.Cache
	// unexecutedRootsQueue is used to keep track of the unexecuted roots in the order they are fetched from database (should be ordered by block_number, log_index)
	// First run of Exec will fill the queue with all the roots that are not executed yet within the [now-permissionlessExecThreshold, now] window.
	// When a root is executed, it is removed from the fifo and its timestamp is stored under rootsSearchFilter. Next database query instead of using entire permissionlessExecWindow
	// will use rootsSearchFilter as the lower bound filter. This way we can reduce the number of database rows fetched with every OCR round.
	// We can do that because roots for most of the cases are executed in sequentially. Instead of skipping snoozed roots after we fetch them from the database,
	// we do that on the db level by narrowing the search window.
	//
	// Example
	// permissionLessExecThresholds - 10 days, now - 2010-10-15
	// We fetch all the roots that within the [2010-10-05, 2010-10-15] window and load them to the queue
	// [0xA - 2010-10-10, 0xB - 2010-10-11, 0xC - 2010-10-12] -> 0xA is the oldest root
	// We executed 0xA and a couple of rounds later, we mark 0xA as executed and snoozed that forever which removes it from the queue.
	// [0xB - 2010-10-11, 0xC - 2010-10-12]
	// Now the search filter wil be 0xA timestamp -> [2010-10-10, 20-10-15]
	// If roots are executed out of order, it's not going to change anything. However, for most of the cases we have sequential root execution and that is
	// a huge improvement because we don't need to fetch all the roots from the database in every round.
	unexecutedRootsQueue *orderedmap.OrderedMap[[32]byte, time.Time]
	rootSearchFilter     time.Time
	rootsQueueMu         sync.RWMutex

	// Both rootSnoozedTime and permissionLessExecutionThresholdDuration can be kept in the commitRootsCache without need to be updated.
	// Those config properties are populates via onchain/offchain config. When changed, OCR plugin will be restarted and cache initialized with new config.
	rootSnoozedTime                          time.Duration
	permissionLessExecutionThresholdDuration time.Duration
}

func newCommitRootsCache(
	permissionLessExecutionThresholdDuration time.Duration,
	rootSnoozeTime time.Duration,
	evictionGracePeriod time.Duration,
	cleanupInterval time.Duration,
) *commitRootsCache {
	evictionTime := permissionLessExecutionThresholdDuration + evictionGracePeriod
	internalCache := cache.New(evictionTime, cleanupInterval)

	return &commitRootsCache{
		snoozedRoots:                             internalCache,
		unexecutedRootsQueue:                     orderedmap.New[[32]byte, time.Time](),
		rootSnoozedTime:                          rootSnoozeTime,
		permissionLessExecutionThresholdDuration: permissionLessExecutionThresholdDuration,
	}
}

func NewSnoozedRoots(permissionLessExecutionThresholdDuration time.Duration, rootSnoozeTime time.Duration) *commitRootsCache {
	return newCommitRootsCache(permissionLessExecutionThresholdDuration, rootSnoozeTime, EvictionGracePeriod, CleanupInterval)
}

func (s *commitRootsCache) IsSnoozed(merkleRoot [32]byte) bool {
	rawValue, found := s.snoozedRoots.Get(merkleRootToString(merkleRoot))
	return found && time.Now().Before(rawValue.(time.Time))
}

func (s *commitRootsCache) MarkAsExecuted(merkleRoot [32]byte) {
	s.snoozedRoots.SetDefault(merkleRootToString(merkleRoot), time.Now().Add(s.permissionLessExecutionThresholdDuration))

	// if there is only one root in the queue, we put it as a search filter
	if s.unexecutedRootsQueue.Len() == 1 {
		s.rootSearchFilter = s.unexecutedRootsQueue.Oldest().Value
	}
	s.unexecutedRootsQueue.Delete(merkleRoot)
	if head := s.unexecutedRootsQueue.Oldest(); head != nil {
		s.rootSearchFilter = head.Value
	}
}

func (s *commitRootsCache) Snooze(merkleRoot [32]byte) {
	s.snoozedRoots.SetDefault(merkleRootToString(merkleRoot), time.Now().Add(s.rootSnoozedTime))
}

func (s *commitRootsCache) CommitSearchTimestamp() time.Time {
	permissionlessExecWindow := time.Now().Add(-s.permissionLessExecutionThresholdDuration)

	timestamp, ok := func() (time.Time, bool) {
		s.rootsQueueMu.RLock()
		defer s.rootsQueueMu.RUnlock()

		// If there are no roots in the queue, we can return the permissionlessExecWindow
		if s.rootSearchFilter.IsZero() {
			return permissionlessExecWindow, true
		}

		if s.rootSearchFilter.After(permissionlessExecWindow) {
			// Query used for fetching roots from the database is exclusive (block_timestamp > :timestamp)
			// so we need to subtract 1 second from the head timestamp to make sure that this root is included in the results
			return s.rootSearchFilter.Add(-time.Second), true
		}
		return time.Time{}, false
	}()

	if ok {
		return timestamp
	}

	s.rootsQueueMu.Lock()
	defer s.rootsQueueMu.Unlock()

	// If rootsSearchFilter is before permissionlessExecWindow, it means that we have roots that are stuck and will never be executed
	// In that case, we wipe out the queue. Next round should start from the permissionlessExecThreshold and rebuild that cache from scratch.
	s.unexecutedRootsQueue = orderedmap.New[[32]byte, time.Time]()
	return permissionlessExecWindow
}
func (s *commitRootsCache) AppendUnexecutedRoot(merkleRoot [32]byte, blockTimestamp time.Time) {
	s.rootsQueueMu.Lock()
	defer s.rootsQueueMu.Unlock()

	// If the root is already in the queue, we don't need to add it again
	if _, found := s.unexecutedRootsQueue.Get(merkleRoot); found {
		return
	}
	// Initialize the search filter with the first root that is added to the queue
	if s.unexecutedRootsQueue.Len() == 0 {
		s.rootSearchFilter = blockTimestamp
	}
	s.unexecutedRootsQueue.Set(merkleRoot, blockTimestamp)
}

func merkleRootToString(merkleRoot [32]byte) string {
	return string(merkleRoot[:])
}

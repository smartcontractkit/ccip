package cache

import (
	"encoding/hex"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

const (
	// EvictionGracePeriod defines how long after the messageVisibilityInterval a root is still kept in the cache
	EvictionGracePeriod = 1 * time.Hour
	// CleanupInterval defines how often roots cache is scanned to evict stale roots
	CleanupInterval = 30 * time.Minute
)

type CommitsRootsCache interface {
	MarkAsExecuted(merkleRoot [32]byte)
	Snooze(merkleRoot [32]byte)
	GetSnoozedRoots() [][32]byte
}

type commitRootsCache struct {
	lggr logger.Logger
	// executedRoots is used to keep track of the roots that are executed. Roots that are considered as executed
	// when all messages are executed on the dest and matching execution state change logs are finalized
	executedRoots *cache.Cache
	// snoozedRoots is used to keep track of the roots that are temporary snoozed
	snoozedRoots *cache.Cache

	// Both rootSnoozedTime and messageVisibilityInterval can be kept in the commitRootsCache without need to be updated.
	// Those config properties are populates via onchain/offchain config. When changed, OCR plugin will be restarted and cache initialized with new config.
	rootSnoozedTime           time.Duration
	messageVisibilityInterval time.Duration
}

func newCommitRootsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
	evictionGracePeriod time.Duration,
	cleanupInterval time.Duration,
) *commitRootsCache {
	executedRoots := cache.New(messageVisibilityInterval+evictionGracePeriod, cleanupInterval)
	snoozedRoots := cache.New(rootSnoozeTime, cleanupInterval)

	return &commitRootsCache{
		lggr:                      lggr,
		executedRoots:             executedRoots,
		snoozedRoots:              snoozedRoots,
		rootSnoozedTime:           rootSnoozeTime,
		messageVisibilityInterval: messageVisibilityInterval,
	}
}

func NewCommitRootsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
) *commitRootsCache {
	return newCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
		EvictionGracePeriod,
		CleanupInterval,
	)
}

func (s *commitRootsCache) GetSnoozedRoots() [][32]byte {
	snoozedRootsItems := s.snoozedRoots.Items()
	executedRootsItems := s.executedRoots.Items()

	merkleRootsBytes := make([][32]byte, 0, len(snoozedRootsItems)+len(executedRootsItems))

	for _, v := range snoozedRootsItems {
		merkleRootsBytes = append(merkleRootsBytes, v.Object.([32]byte))
	}
	for _, v := range executedRootsItems {
		merkleRootsBytes = append(merkleRootsBytes, v.Object.([32]byte))
	}
	return merkleRootsBytes
}

func (s *commitRootsCache) IsSkipped(merkleRoot [32]byte) bool {
	_, snoozed := s.snoozedRoots.Get(merkleRootToString(merkleRoot))
	_, executed := s.executedRoots.Get(merkleRootToString(merkleRoot))
	return snoozed || executed
}

func (s *commitRootsCache) MarkAsExecuted(merkleRoot [32]byte) {
	prettyMerkleRoot := merkleRootToString(merkleRoot)
	s.executedRoots.SetDefault(prettyMerkleRoot, merkleRoot)

	s.lggr.Debugw("Marking root as executed, it's gonna be permanently skipped",
		"merkleRoot", prettyMerkleRoot,
	)
}

func (s *commitRootsCache) Snooze(merkleRoot [32]byte) {
	s.snoozedRoots.SetDefault(merkleRootToString(merkleRoot), struct{}{})
}

func merkleRootToString(merkleRoot [32]byte) string {
	return hex.EncodeToString(merkleRoot[:])
}

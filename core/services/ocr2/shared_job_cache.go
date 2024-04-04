package ocr2

import (
	"maps"
	"sync"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

// SharedJobCache is a cache in delegate that can be share job specs across ReportingPlugins.
type SharedJobCache struct {
	jobs map[int32]job.Job
	mu   sync.RWMutex
}

// NewSharedJobCache creates a new instance of SharedCache.
func NewSharedJobCache() *SharedJobCache {
	return &SharedJobCache{
		jobs: make(map[int32]job.Job),
	}
}

func (s *SharedJobCache) set(job job.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = job
}

func (s *SharedJobCache) delete(job job.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.jobs, job.ID)
}

func (s *SharedJobCache) Get() map[int32]job.Job {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// deep copy the map
	copyJobs := make(map[int32]job.Job, len(s.jobs))
	maps.Copy(copyJobs, s.jobs)
	
	return copyJobs
}

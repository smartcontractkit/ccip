package ocr2

import (
	"context"
	"maps"
	"sync"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

type JobSubscriber interface {
	ID() int32
	OnNewJob(ctx context.Context, jb job.Job) error
	OnDeleteJob(ctx context.Context, jb job.Job) error
}

// SharedJobCache is a cache in delegate that can be share job specs across ReportingPlugins.
type SharedJobCache struct {
	jobs        map[int32]job.Job
	subscribers map[types.OCR2PluginType][]JobSubscriber

	jobMu sync.RWMutex
	subMu sync.RWMutex
}

// NewSharedJobCache creates a new instance of SharedCache.
func NewSharedJobCache() *SharedJobCache {
	return &SharedJobCache{
		jobs:        make(map[int32]job.Job),
		subscribers: make(map[types.OCR2PluginType][]JobSubscriber),
	}
}

func (s *SharedJobCache) addJob(ctx context.Context, job job.Job) error {
	s.jobMu.Lock()
	s.jobs[job.ID] = job
	s.jobMu.Unlock()

	for _, subscriber := range s.subscribers[job.Type] {
		err := subscriber.OnNewJob(ctx, job)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SharedJobCache) deleteJob(ctx context.Context, job job.Job) error {
	s.jobMu.Lock()
	delete(s.jobs, job.ID)
	s.jobMu.Unlock()

	for _, subscriber := range s.subscribers[job.Type] {
		err := subscriber.OnDeleteJob(ctx, job)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SharedJobCache) Get() map[int32]job.Job {
	s.jobMu.RLock()
	defer s.jobMu.RUnlock()

	// deep copy the map
	copyJobs := make(map[int32]job.Job, len(s.jobs))
	maps.Copy(copyJobs, s.jobs)

	return copyJobs
}

func (s *SharedJobCache) Subscribe(jobType types.OCR2PluginType, subscriber JobSubscriber) {
	s.subMu.Lock()
	defer s.subMu.Unlock()

	s.subscribers[jobType] = append(s.subscribers[jobType], subscriber)
}

func (s *SharedJobCache) Unsubscribe(jobType types.OCR2PluginType, subscriber JobSubscriber) {
	s.subMu.Lock()
	defer s.subMu.Unlock()

	var subs []JobSubscriber
	for _, sub := range s.subscribers[jobType] {
		if sub.ID() != subscriber.ID() {
			subs = append(subs, sub)
		}
	}

	s.subscribers[jobType] = subs
}

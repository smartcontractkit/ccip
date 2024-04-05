package ocr2

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

func TestSharedJobCache_AddJobs(t *testing.T) {
	cache := NewSharedJobCache()

	require.Equal(t, 0, len(cache.Get()))

	numJobs := 100
	for i := 0; i < numJobs; i++ {
		err := cache.addJob(job.Job{ID: int32(i)})
		require.NoError(t, err)
	}

	cachedJobs := cache.Get()
	require.Equal(t, numJobs, len(cachedJobs))
	for i := 0; i < numJobs; i++ {
		require.Equal(t, job.Job{ID: int32(i)}, cachedJobs[int32(i)])
	}

	// set duplicate jobs
	for i := 0; i < numJobs; i++ {
		err := cache.addJob(job.Job{ID: int32(i)})
		require.NoError(t, err)
	}

	cachedJobs = cache.Get()
	require.Equal(t, numJobs, len(cachedJobs))
	for i := 0; i < numJobs; i++ {
		require.Equal(t, job.Job{ID: int32(i)}, cachedJobs[int32(i)])
	}

	// modifying returned mapping does not modify cache
	cachedJobs[int32(numJobs+1)] = job.Job{ID: int32(numJobs + 1)}
	require.Equal(t, numJobs, len(cache.Get()))
}

func TestSharedJobCache_DeleteJobs(t *testing.T) {
	cache := NewSharedJobCache()

	numJobs := 100
	for i := 0; i < numJobs; i++ {
		err := cache.addJob(job.Job{ID: int32(i)})
		require.NoError(t, err)
	}

	for i := 0; i < numJobs/2; i++ {
		err := cache.deleteJob(job.Job{ID: int32(i)})
		require.NoError(t, err)
	}

	cachedJobs := cache.Get()
	require.Equal(t, numJobs/2, len(cachedJobs))
	for i := numJobs / 2; i < numJobs; i++ {
		require.Equal(t, job.Job{ID: int32(i)}, cachedJobs[int32(i)])
	}

	// can delete non-existent job when cache is empty
	err := cache.deleteJob(job.Job{ID: int32(numJobs + 1)})
	require.NoError(t, err)
	require.Equal(t, cachedJobs, cache.Get())

	// can remove all jobs
	for i := numJobs / 2; i < numJobs; i++ {
		err = cache.deleteJob(job.Job{ID: int32(i)})
		require.NoError(t, err)
	}

	cachedJobs = cache.Get()
	require.Equal(t, 0, len(cachedJobs))

	// can delete non-existent job when cache is empty
	err = cache.deleteJob(job.Job{ID: int32(0)})
	require.NoError(t, err)
}

// TODO add test for pubsub

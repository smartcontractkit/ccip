package cache

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCacheFuncNoErrorCachingSuccess tests caching behavior when the function succeeds.
func TestCacheFuncNoErrorCachingSuccess(t *testing.T) {
	callCount := 0
	testFunc := func(ctx context.Context) (string, error) {
		callCount++
		return "test result", nil
	}

	cachedFunc := CallOnceOnNoError(testFunc)

	// Call the function twice.
	_, err := cachedFunc(tests.Context(t))
	assert.NoError(t, err, "Expected no error on the first call")

	_, err = cachedFunc(tests.Context(t))
	assert.NoError(t, err, "Expected no error on the second call")

	assert.Equal(t, 1, callCount, "Function should be called exactly once")
}

// TestCacheFuncNoErrorCachingError tests that the function is retried after an error.
func TestCacheFuncNoErrorCachingError(t *testing.T) {
	callCount := 0
	testFunc := func(ctx context.Context) (string, error) {
		callCount++
		if callCount == 1 {
			return "", errors.New("test error")
		}
		return "test result", nil
	}

	cachedFunc := CallOnceOnNoError(testFunc)

	// First call should fail.
	_, err := cachedFunc(tests.Context(t))
	require.Error(t, err, "Expected an error on the first call")

	// Second call should succeed.
	r, err := cachedFunc(tests.Context(t))
	assert.NoError(t, err, "Expected no error on the second call")
	assert.Equal(t, "test result", r)
	assert.Equal(t, 2, callCount, "Function should be called exactly twice")
}

// TestCacheFuncNoErrorCachingConcurrency tests that the function works correctly under concurrent access.
func TestCacheFuncNoErrorCachingConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	callCount := 0
	testFunc := func(ctx context.Context) (string, error) {
		callCount++
		return "test result", nil
	}

	cachedFunc := CallOnceOnNoError(testFunc)

	// Simulate concurrent calls.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := cachedFunc(tests.Context(t))
			assert.NoError(t, err, "Expected no error in concurrent execution")
		}()
	}

	wg.Wait()

	assert.Equal(t, 1, callCount, "Function should be called exactly once despite concurrent calls")
}

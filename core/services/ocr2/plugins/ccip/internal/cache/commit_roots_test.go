package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestSnoozedRoots(t *testing.T) {
	c := NewCommitRootsCache(logger.TestLogger(t), 1*time.Minute, 1*time.Minute)

	k1 := [32]byte{1}
	k2 := [32]byte{2}

	// return false for non existing element
	snoozed := c.IsSkipped(k1)
	assert.False(t, snoozed)

	// after an element is marked as executed it should be snoozed
	c.MarkAsExecuted(k1)
	snoozed = c.IsSkipped(k1)
	assert.True(t, snoozed)

	// after snoozing an element it should be snoozed
	c.Snooze(k2)
	snoozed = c.IsSkipped(k2)
	assert.True(t, snoozed)
}

func TestEvictingElements(t *testing.T) {
	c := newCommitRootsCache(logger.TestLogger(t), 1*time.Hour, 1*time.Millisecond, 1*time.Millisecond, 1*time.Millisecond)

	k1 := [32]byte{1}
	c.Snooze(k1)

	time.Sleep(10 * time.Millisecond)

	assert.False(t, c.IsSkipped(k1))
}

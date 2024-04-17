package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSnoozedRoots(t *testing.T) {
	c := NewSnoozedRoots(1*time.Minute, 1*time.Minute)

	k1 := [32]byte{1}
	k2 := [32]byte{2}

	// return false for non existing element
	snoozed := c.IsSnoozed(k1)
	assert.False(t, snoozed)

	// after an element is marked as executed it should be snoozed
	c.MarkAsExecuted(k1)
	snoozed = c.IsSnoozed(k1)
	assert.True(t, snoozed)

	// after snoozing an element it should be snoozed
	c.Snooze(k2)
	snoozed = c.IsSnoozed(k2)
	assert.True(t, snoozed)
}

func TestEvictingElements(t *testing.T) {
	c := newCommitRootsCache(1*time.Millisecond, 1*time.Hour, 1*time.Millisecond, 1*time.Millisecond)

	k1 := [32]byte{1}
	c.Snooze(k1)

	time.Sleep(10 * time.Millisecond)

	assert.False(t, c.IsSnoozed(k1))
}

func Test_UnexecutedRootsTracking(t *testing.T) {
	permissionLessThreshold := 10 * time.Hour
	c := newCommitRootsCache(permissionLessThreshold, 1*time.Hour, 1*time.Millisecond, 1*time.Millisecond)

	k1 := [32]byte{1}
	k2 := [32]byte{2}
	k3 := [32]byte{3}
	k4 := [32]byte{4}

	t1 := time.Now().Add(-4 * time.Hour)
	t2 := time.Now().Add(-3 * time.Hour)
	t3 := time.Now().Add(-2 * time.Hour)
	t4 := time.Now().Add(-1 * time.Hour)

	// First check should return permissionLessThreshold window
	commitTs := c.CommitSearchTimestamp()
	assert.True(t, commitTs.Before(time.Now().Add(-permissionLessThreshold)))

	c.AppendUnexecutedRoot(k1, t1)
	c.AppendUnexecutedRoot(k2, t2)
	c.AppendUnexecutedRoot(k3, t3)

	// After loading roots it should return the first one
	commitTs = c.CommitSearchTimestamp()
	assert.Equal(t, t1.Add(-time.Second), commitTs)

	// Marking root in the middle as executed shouldn't change the commitTs
	c.MarkAsExecuted(k2)
	commitTs = c.CommitSearchTimestamp()
	assert.Equal(t, t1.Add(-time.Second), commitTs)

	// Marking k1 as executed when k2 is already executed should return timestamp of k3
	c.MarkAsExecuted(k1)
	commitTs = c.CommitSearchTimestamp()
	assert.Equal(t, t3.Add(-time.Second), commitTs)

	// Marking all as executed should return timestamp of the latest
	c.MarkAsExecuted(k3)
	commitTs = c.CommitSearchTimestamp()
	assert.Equal(t, t3.Add(-time.Second), commitTs)

	// Adding k4 should return timestamp of k4
	c.AppendUnexecutedRoot(k4, t4)
	commitTs = c.CommitSearchTimestamp()
	assert.Equal(t, t4.Add(-time.Second), commitTs)

	c.MarkAsExecuted(k4)
	commitTs = c.CommitSearchTimestamp()
	assert.Equal(t, t4.Add(-time.Second), commitTs)
}

func Test_UnexecutedRootsStaleQueue(t *testing.T) {
	permissionLessThreshold := 5 * time.Hour
	c := newCommitRootsCache(permissionLessThreshold, 1*time.Hour, 1*time.Millisecond, 1*time.Millisecond)

	k1 := [32]byte{1}
	k2 := [32]byte{2}
	k3 := [32]byte{3}

	t1 := time.Now().Add(-4 * time.Hour)
	t2 := time.Now().Add(-3 * time.Hour)
	t3 := time.Now().Add(-2 * time.Hour)

	c.AppendUnexecutedRoot(k1, t1)
	c.AppendUnexecutedRoot(k2, t2)
	c.AppendUnexecutedRoot(k3, t3)

	// First check should return permissionLessThreshold window
	commitTs := c.CommitSearchTimestamp()
	assert.Equal(t, t1.Add(-time.Second), commitTs)

	// Reducing permissionLessExecutionThreshold works as speeding the clock
	c.permissionLessExecutionThresholdDuration = 1 * time.Hour

	commitTs = c.CommitSearchTimestamp()
	assert.True(t, commitTs.Before(time.Now().Add(-1*time.Hour)))
	assert.True(t, commitTs.After(t1))
	assert.True(t, commitTs.After(t2))
	assert.True(t, commitTs.After(t3))
}

package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestBlessedRoots(t *testing.T) {
	m1 := utils.RandomBytes32()
	m2 := utils.RandomBytes32()
	m3 := utils.RandomBytes32()

	t.Run("base", func(t *testing.T) {
		c := NewBlessedRoots(time.Minute, time.Minute)

		c.Set(m1, true)
		c.Set(m2, false)

		val, exists := c.Get(m1)
		assert.True(t, val)
		assert.True(t, exists)

		val, exists = c.Get(m2)
		assert.False(t, val)
		assert.True(t, exists)

		// unset key
		val, exists = c.Get(m3)
		assert.False(t, val)
		assert.False(t, exists)

		// over-write m2
		c.Set(m2, true)
		val, exists = c.Get(m2)
		assert.True(t, val)
		assert.True(t, exists)
	})

	t.Run("test expiry", func(t *testing.T) {
		c := NewBlessedRoots(time.Nanosecond, time.Minute)
		c.Set(m1, true)
		time.Sleep(10 * time.Nanosecond) // after 10ns it should've been expired
		_, exists := c.Get(m1)
		assert.False(t, exists)
	})

	t.Run("internal issue", func(t *testing.T) {
		c := NewBlessedRoots(time.Minute, time.Minute)
		c.mem.Set(c.merkleRootKey(m1), 1234, 0) // non-bool
		_, exists := c.Get(m1)
		assert.False(t, exists)
	})
}

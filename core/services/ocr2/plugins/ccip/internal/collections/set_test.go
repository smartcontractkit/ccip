package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := NewSet[string]()

	assert.False(t, s.Contains("foo"))
	assert.True(t, s.Add("foo"))
	assert.False(t, s.Add("foo"))
	assert.True(t, s.Contains("foo"))
	assert.False(t, s.Contains("bar"))

	s2 := NewSet[string]("bar")
	assert.False(t, s.Equal(&s2))
	s2.Add("foo")
	assert.False(t, s.Equal(&s2))
	s2.Remove("bar")
	assert.True(t, s.Equal(&s2))
	assert.False(t, s2.Remove("zoo"))

	s2 = NewSet[string]("foo")
	assert.True(t, s.Equal(&s2))
}

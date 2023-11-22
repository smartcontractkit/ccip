package collections

import (
	"fmt"
	"math/rand"
	"sync"
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
	assert.Equal(t, 2, s2.Size())
	assert.False(t, s.Equal(&s2))
	s2.Remove("bar")
	assert.True(t, s.Equal(&s2))
	assert.False(t, s2.Remove("zoo"))

	s2 = NewSet[string]("foo")
	assert.True(t, s.Equal(&s2))
}

func TestSetConcurrency(t *testing.T) {
	// this test tries to add numAdditions elements
	// and remove numRemovals items from the set from multiple goroutines
	const numAdditions = 50
	const numRemovals = 20

	const numUnqElems = 500
	elems := make([]string, numUnqElems)
	for i := range elems {
		elems[i] = fmt.Sprintf("elem[%d]", i)
	}
	getRandomElem := func() string {
		idx := rand.Intn(numUnqElems)
		return elems[idx]
	}

	s := NewSet[string]()

	wg := sync.WaitGroup{}
	wg.Add(numAdditions + numRemovals)

	for i := 0; i < numRemovals; i++ {
		go func() {
			defer wg.Done()
			for {
				el := getRandomElem()
				if s.Remove(el) {
					break
				}
			}
		}()
	}

	for i := 0; i < numAdditions; i++ {
		go func() {
			defer wg.Done()
			for {
				el := getRandomElem()
				if !s.Contains(el) {
					s.Add(el)
					break
				}
			}
		}()
	}

	wg.Wait()
	assert.True(t, len(s.mem) >= numAdditions-numRemovals)
}

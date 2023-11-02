package hashlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesOfBytesKeccak(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		h, err := BytesOfBytesKeccak(nil)
		assert.NoError(t, err)
		assert.Equal(t, [32]byte{}, h)

		h1, err := BytesOfBytesKeccak([][]byte{{0x1}, {0x1}})
		assert.NoError(t, err)
		h2, err := BytesOfBytesKeccak([][]byte{{0x1, 0x1}})
		assert.NoError(t, err)
		assert.NotEqual(t, h1, h2)
	})

	t.Run("should not have collision", func(t *testing.T) {
		s1 := []byte("abc")
		s2 := []byte("efg")
		hs1, err := BytesOfBytesKeccak([][]byte{s1}) // xyz
		assert.NoError(t, err)

		h1, err := BytesOfBytesKeccak([][]byte{s1, s2})
		assert.NoError(t, err)

		h2, err := BytesOfBytesKeccak([][]byte{append(hs1[:], s2...)})
		assert.NoError(t, err)

		assert.NotEqual(t, h1, h2)
	})

}

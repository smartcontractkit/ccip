package cache

import (
	"encoding/hex"
	"time"

	goc "github.com/patrickmn/go-cache"
)

type BlessedRootsCache interface {
	Get(m [32]byte) (val bool, exists bool)
	Set(m [32]byte, val bool)
}

var _ BlessedRootsCache = &BlessedRoots{}

type BlessedRoots struct {
	mem *goc.Cache
}

func NewBlessedRoots(expiry, cleanup time.Duration) *BlessedRoots {
	return &BlessedRoots{
		mem: goc.New(expiry, cleanup),
	}
}

func (c *BlessedRoots) Get(m [32]byte) (val bool, exists bool) {
	rawVal, exists := c.mem.Get(c.merkleRootKey(m))
	if !exists {
		return false, false
	}

	boolVal, is := rawVal.(bool)
	if !is {
		return false, false
	}

	return boolVal, true
}

func (c *BlessedRoots) Set(m [32]byte, val bool) {
	c.mem.Set(c.merkleRootKey(m), val, 0)
}

func (c *BlessedRoots) merkleRootKey(m [32]byte) string {
	return hex.EncodeToString(m[:])
}

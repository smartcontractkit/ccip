package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/rs/zerolog/log"
	"go.uber.org/atomic"
)

const DEFAULTMAXSIZE = 1000000

type Cache struct {
	maxSize    int64
	filePath   string
	resetCount *atomic.Int64
	cache      *fastcache.Cache
}

func (c *Cache) SetMaxSize(size int64) {
	c.maxSize = size
}

func (c *Cache) SaveCurrentStateAndReset() error {
	filePath := fmt.Sprintf("%s_%d", c.filePath, c.resetCount.Load())
	log.Info().
		Int64("Reset Count", c.resetCount.Load()).
		Str("filepath", c.filePath).
		Msgf("Resetting cache, dumping to back up")
	err := c.cache.SaveToFile(c.filePath)
	if err != nil {
		return fmt.Errorf("error %w saving the cache into file %s", err, filePath)
	}
	c.cache.Reset()
	c.resetCount.Inc()
	return nil
}

func (c *Cache) Store(key []byte, value any) error {
	vBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	newC := c.cache

	if err != nil {
		return fmt.Errorf("error %w in converting %v to byte array", err, value)
	}
	// check if the new entry will cause cache bytesSize to be more than MaxBytesSize
	// if yes, store the cache into file and reset the cache
	var s fastcache.Stats
	newC.UpdateStats(&s)
	if s.BytesSize+uint64(binary.Size(vBytes)) >= uint64(c.maxSize) {
		err := c.SaveCurrentStateAndReset()
		if err != nil {
			return err
		}
	}
	newC.Set(key, vBytes)
	return nil
}

func (c *Cache) Delete(key []byte) {
	c.cache.Del(key)
}

func (c *Cache) Load(key []byte, value any) (bool, error) {
	var dstBytes []byte
	var exists bool
	cache := c.cache
	dstBytes, exists = cache.HasGet(dstBytes, key)
	// if the cache has already been reset, check if the key exists in previous versions
	if !exists && c.resetCount.Load() > 0 {
		for i := int64(1); i < c.resetCount.Load(); i++ {
			filePath := fmt.Sprintf("%s_%d", c.filePath, i)
			c1, err := fastcache.LoadFromFile(filePath)
			if err != nil {
				return false, fmt.Errorf("error %w loading cache from back up file %s", err, filePath)
			}
			dstBytes, exists = c1.HasGet(dstBytes, key)
			if exists {
				log.Info().Msgf("found key in back up %s", filePath)
				break
			}
		}
	}
	if exists {
		err := json.Unmarshal(dstBytes, value)
		if err != nil {
			return exists, fmt.Errorf("error %w converting cache value to struct %v", err, value)
		}
		return exists, nil
	}
	return false, nil
}

func NewCache(maxBytes int, cacheName string) *Cache {
	tmpDir, err := os.MkdirTemp("", "tmp_")
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(tmpDir, fmt.Sprintf("%s.fastcache", cacheName))
	return &Cache{
		maxSize:    DEFAULTMAXSIZE,
		filePath:   filePath,
		resetCount: atomic.NewInt64(0),
		cache:      fastcache.New(maxBytes),
	}
}

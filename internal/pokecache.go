package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	cache map[string]cacheEntry
	mu    sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for tickTime := range ticker.C {
		c.mu.Lock()
		cutoff := tickTime.Add(-interval)
		for key, entry := range c.cache {
			if entry.createdAt.Before(cutoff) {
				delete(c.cache, key)
			}

		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
	}

	go cache.reapLoop(interval)
	return &cache
}

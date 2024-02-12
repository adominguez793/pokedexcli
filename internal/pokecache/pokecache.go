package pokecache

import (
	// "sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	// mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		// mu: *sync.Mutex,
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cache, ok := c.cache[key]
	return cache.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)

	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	timeBomb := time.NewTicker(interval)
	for range timeBomb.C {
		c.reap(interval)
	}
}

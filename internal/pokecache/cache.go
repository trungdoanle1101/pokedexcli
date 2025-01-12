package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux          *sync.Mutex
	cacheEntries map[string]cacheEntry
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mux:          &sync.Mutex{},
		interval:     interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if entry, ok := c.cacheEntries[key]; ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for t := range ticker.C {
		for k, entry := range c.cacheEntries {
			c.mux.Lock()
			if entry.createdAt.Before(t.Add(-c.interval)) {
				delete(c.cacheEntries, k)
			}
			c.mux.Unlock()
		}
	}
}

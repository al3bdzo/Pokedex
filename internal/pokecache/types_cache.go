package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	CacheMap map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(t time.Duration) *Cache {
	c := &Cache{
		CacheMap: map[string]cacheEntry{},
		mu: sync.Mutex{},
	}
	go c.reapLoop(t)
	return c
}
package pokecache


import (
	"time"
)


func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.CacheMap[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.CacheMap {
			if time.Now().Sub(val.createdAt) > t {
				delete(c.CacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
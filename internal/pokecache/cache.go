package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache{
	
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	
	go cache.reapLoop(interval)

	return cache
}


func (c *Cache) Add(key string, val []byte)  {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt : time.Now().UTC(),
		val : val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mutex.Lock()

	defer c.mutex.Unlock()

	entry, exists := c.entries[key]

	if !exists{
		return nil, false
	}

	return entry.val, true
}


func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}
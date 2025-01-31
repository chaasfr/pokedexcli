package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.RWMutex
	interval     time.Duration
}


func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	var ce cacheEntry
	ce.val = val
	ce.createdAt = time.Now()
	cache.cacheEntries[key] = ce
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RLocker().Unlock()
	ce, found := cache.cacheEntries[key]
	if !found {
		return nil, false
	}
	return ce.val, true
}

func (cache *Cache) reap() {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for key, ce := range cache.cacheEntries {
		if time.Since(ce.createdAt) > cache.interval {
			delete(cache.cacheEntries, key)
		}
	}
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	for {
		<- ticker.C
		cache.reap()
	}
}


func NewCache(interval time.Duration) *Cache {
	var cache Cache
	cache.interval = interval
	cache.cacheEntries = make(map[string]cacheEntry)
	go cache.reapLoop()
	return &cache
}
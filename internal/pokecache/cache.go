package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func New(interval time.Duration) *PokeCache {
	cache := PokeCache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return &cache
}

func (pc *PokeCache) Add(key string, val []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	pc.cache[key] = newEntry
}

func (pc *PokeCache) Get(key string) ([]byte, bool) {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	if cache, exists := pc.cache[key]; exists {
		return cache.val, true
	}

	return nil, false
}

func (pc *PokeCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		pc.mu.Lock()
		for key, entry := range pc.cache {
			if entry.createdAt.Before(time.Now().Add(-interval)) {
				delete(pc.cache, key)
			}
		}
		pc.mu.Unlock()
	}
}

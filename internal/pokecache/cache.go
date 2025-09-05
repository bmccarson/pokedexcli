package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func New(interval time.Duration) *PokeCache {
	cache := PokeCache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
		mu:       sync.Mutex{},
	}

	go cache.reapLoop()

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

func (pc *PokeCache) reapLoop() {
	for {
		time.Sleep(pc.interval)
		pc.mu.Lock()
		for key, entry := range pc.cache {
			if entry.createdAt.After(time.Now().Add(-pc.interval)) {
				delete(pc.cache, key)

			}
		}
		pc.mu.Unlock()
	}
}

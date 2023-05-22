package cache

import (
	"sync"
	"time"
)

//Cache entry holds a timestamp of when it was created and a value in bytes
type CacheEntry struct {
	timestamp time.Time
	val []byte
}

//Cache struct holds a map of entries
type Cache struct {
	cache map[string] CacheEntry
	mux *sync.Mutex
}

//----- Cache Methods -----

//Gets an item from the cache 
func (c* Cache) Get(key string) ([]byte, bool){
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.cache[key].val != nil {
		return c.cache[key].val, true
	}

	return nil, false
}

//Adds an entry to the cache
func (c* Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = CacheEntry{time.Now(), val}
}

//Delete a cache value that has reached the current interval
func (c* Cache) reap(currentTime time.Time, interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, value := range c.cache {
		if value.timestamp.Before(currentTime.Add(-interval)) {
			delete(c.cache, key)
		}
	}
}

//Generates a new ticker and executes the reaping function after every tick
func (c* Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

//-------------------------

//Generates a new empty cache with a specified reaping interval
func NewCache(interval time.Duration) Cache{
	cache := Cache{
		map[string]CacheEntry{}, 
		&sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
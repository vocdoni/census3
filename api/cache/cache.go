package cache

import (
	"context"
	"crypto/md5"
	"sync"
	"time"
)

var (
	DefaultSanityInterval    = 6 * time.Hour
	DefaultMinViewsToSurvive = 1
)

// Cache is a simple in-memory key-value storage with mutexes
type Cache struct {
	views             map[[16]byte]int
	storage           map[[16]byte]interface{}
	mtx               *sync.RWMutex
	ctx               context.Context
	cancel            context.CancelFunc
	sanityInterval    time.Duration
	minViewsToSurvive int
}

// NewCache creates a new Cache instance with the given sanity interval and
// min views to survive and starts the sanity loop.
func NewCache(sanityInterval time.Duration, minViewsToSurvive int) *Cache {
	ctx, cancel := context.WithCancel(context.Background())
	cache := &Cache{
		views:             make(map[[16]byte]int),
		storage:           make(map[[16]byte]interface{}),
		mtx:               &sync.RWMutex{},
		ctx:               ctx,
		cancel:            cancel,
		sanityInterval:    sanityInterval,
		minViewsToSurvive: minViewsToSurvive,
	}
	go cache.sanityLoop()
	return cache
}

// DefaultCache creates a new Cache instance with default values for sanity
// interval and min views to survive
func DefaultCache() *Cache {
	return NewCache(DefaultSanityInterval, DefaultMinViewsToSurvive)
}

// Set sets a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.views[encKey(key)] = 0
	c.storage[encKey(key)] = value
}

// Get gets a value from the cache
func (c *Cache) Get(rawKey string) (interface{}, bool) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	key := encKey(rawKey)
	if val, ok := c.storage[key]; ok {
		c.views[key]++
		return val, true
	}
	return nil, false
}

// Delete deletes a value from the cache
func (c *Cache) Delete(rawKey string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	key := encKey(rawKey)
	delete(c.storage, key)
	delete(c.views, key)
}

// Has checks if a key exists in the cache
func (c *Cache) Has(rawKey string) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	key := encKey(rawKey)
	if _, ok := c.storage[key]; ok {
		c.views[key]++
		return true
	}
	return false
}

// Clear clears the cache
func (c *Cache) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.storage = make(map[[16]byte]interface{})
}

// Destroy clears the cache and stops the sanity loop
func (c *Cache) Destroy() {
	c.cancel()
	c.Clear()
}

// sanityLoop is a loop that runs every sanityInterval and removes the keys that
// have less than minViewsToSurvive views. This is to prevent the cache from
// growing indefinitely. If a key have more than minViewsToSurvive views, it
// will be reduced by 1 every sanityInterval.
func (c *Cache) sanityLoop() {
	ticker := time.NewTicker(c.sanityInterval)
	for {
		select {
		case <-ticker.C:
			c.mtx.Lock()
			for key, views := range c.views {
				if views < c.minViewsToSurvive {
					delete(c.storage, key)
					delete(c.views, key)
				} else {
					c.views[key] = views - 1
				}
			}
			c.mtx.Unlock()
		case <-c.ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// encKey encodes the key to a string to be used as a map key, it uses md5 to
// ensure the key is always the same length
func encKey(key string) [16]byte {
	return md5.Sum([]byte(key))
}

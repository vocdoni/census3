package cache

import (
	"crypto/md5"
	"sync"
)

// Cache is a simple in-memory key-value storage with mutexes
type Cache struct {
	storage map[[16]byte]interface{}
	mtx     *sync.RWMutex
}

// NewCache creates a new Cache instance
func NewCache() *Cache {
	return &Cache{
		storage: make(map[[16]byte]interface{}),
		mtx:     &sync.RWMutex{},
	}
}

// encKey encodes the key to a string to be used as a map key, it uses md5 to
// ensure the key is always the same length
func encKey(key string) [16]byte {
	return md5.Sum([]byte(key))
}

// Set sets a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.storage[encKey(key)] = value
}

// Get gets a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	val, ok := c.storage[encKey(key)]
	return val, ok
}

// Delete deletes a value from the cache
func (c *Cache) Delete(key string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	delete(c.storage, encKey(key))
}

// Has checks if a key exists in the cache
func (c *Cache) Has(key string) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	_, ok := c.storage[encKey(key)]
	return ok
}

// Clear clears the cache
func (c *Cache) Clear() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.storage = make(map[[16]byte]interface{})
}

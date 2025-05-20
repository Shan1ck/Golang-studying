package cache

import (
	"sync"
)

type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]int),
	}
}

func (c *Cache) Increase(key string, value int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	delete(c.storage, key)
}

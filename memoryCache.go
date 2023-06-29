package main

import (
	"fmt"
	"sync"
)

// MemoryCache is the implementation of the Cache interface
type MemoryCache struct {
	data map[string][]byte
	mtx  sync.RWMutex
}

// NewMemoryCache creates a new instance of the  MemoryCache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string][]byte), mtx: sync.RWMutex{}}
}

// Get returns the value stored wuth the given key, if the key does not exist
// the fucntion returns an error
func (c *MemoryCache) Get(key []byte) ([]byte, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	strKey := string(key)
	value, ok := c.data[strKey]
	if !ok {
		return nil, fmt.Errorf("key (%s) does not exist", strKey)
	}
	return value, nil
}

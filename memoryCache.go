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
func (c *MemoryCache) Get(key string) ([]byte, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	strKey := string(key)
	value, ok := c.data[strKey]
	if !ok {
		return nil, fmt.Errorf("key (%s) does not exist", strKey)
	}
	return value, nil
}

func (c *MemoryCache)Set(key string,value []byte) error{
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.data[key]=value
	return nil
}
func(c *MemoryCache)	Delete(key string) error{
	c.mtx.Lock()
	defer c.mtx.Unlock()
	delete(c.data,key)
	return nil
}
func(c *MemoryCache)	HasKey(key string) bool{
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	_,ok := c.data[key]
	return ok
}
func(c *MemoryCache)	GetKeys()[]string{
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	keysList :=make([]string,0)
	for key := range c.data {
		
		keysList= append(keysList,key )
        
    }
	return keysList
}

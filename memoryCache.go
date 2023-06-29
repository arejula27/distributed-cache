package main

// MemoryCache is the implementation of the Cache interface
type MemoryCache struct {
}

// NewMemoryCache creates a new instance of the  MemoryCache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

func (c *MemoryCache) Get(key []byte) ([]byte, error) {
	return []byte("OK"), nil
}

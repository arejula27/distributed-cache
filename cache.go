package main

// Cache defines the main funcionality of a cache
type Cache interface {
	Get([]byte) ([]byte, error)
}

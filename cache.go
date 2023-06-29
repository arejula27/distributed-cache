package main

// Cache defines the main funcionality of a cache
type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string,value []byte) error
	Delete(key string) error
	HasKey(key string) bool
	GetKeys()[]string
}

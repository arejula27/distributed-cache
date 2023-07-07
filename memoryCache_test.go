package main

import (
	"reflect"
	"testing"
)

func TestMemoryCache(t *testing.T) {
	cache := NewMemoryCache()

	t.Run("Test Get function", func(t *testing.T) {
		//clear the map
		clearMap(cache.data)

		//Add initial data
		key := "key1"
		initialValue := "value1"
		cache.data[key] = []byte(initialValue)

		result, err := cache.Get(key)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []byte(initialValue)) {
			t.Errorf("Expected value %v, but got %v", initialValue, result)
		}

	})
	t.Run("Test Set function", func(t *testing.T) {
		//clear the map
		clearMap(cache.data)
		//Add initial data
		key := "key1"
		initialValue := "value1"

		err := cache.Set(key, []byte(initialValue))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		result, ok := cache.data[key]
		if !ok {
			t.Errorf("Key %v does not found in the cache", key)
		}

		if !reflect.DeepEqual(result, []byte(initialValue)) {
			t.Errorf("Expected value %v, but got %v", initialValue, result)
		}

	})

	t.Run("Test Delete function", func(t *testing.T) {
		//clear the map
		clearMap(cache.data)
		key := "ketToBeDeleted"
		cache.data[key] = []byte("Value")
		err := cache.Delete(key)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		_, err = cache.Get(key)
		if err == nil {
			t.Error("Expected error, but got nil")
		}
	})

	t.Run("Test HasKey function", func(t *testing.T) {
		//clear the map
		clearMap(cache.data)
		key := "key"
		exists := cache.HasKey(key)
		if exists {
			t.Error("Expected key to not exist, but it does")
		}
		initialValue := "value"
		cache.data[key] = []byte(initialValue)
		exists = cache.HasKey(key)
		if !exists {
			t.Error("Expected key to  exist, but it does not")
		}
	})

	t.Run("Test GetKeys function", func(t *testing.T) {
		//clear the map
		clearMap(cache.data)
		keys := cache.GetKeys()
		if len(keys) != 0 {
			t.Errorf("Expected keys list to be empty, but got %v", keys)
		}
		//Add  data
		key1 := "key1"
		value1 := "value1"
		key2 := "key2"
		value2 := "value2"
		cache.data[key1] = []byte(value1)
		cache.data[key2] = []byte(value2)
		keys = cache.GetKeys()
		if len(keys) != 2 {
			t.Errorf("Expected keys length list to be two, but got %v", len(keys))
		}
	})
}

func clearMap(m map[string][]byte) {
	//clear the map
	for k := range m {
		delete(m, k)
	}
}

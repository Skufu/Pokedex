package pokeapi

import (
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	result, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected to find key %s in cache, but it was not found", key)
	}

	if string(result) != string(val) {
		t.Errorf("Expected value %s, got %s", string(val), string(result))
	}
}

func TestCacheReap(t *testing.T) {
	interval := 100 * time.Millisecond
	cache := NewCache(interval)

	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	// Verify the entry exists
	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected to find key %s in cache, but it was not found", key)
	}

	// Wait for the entry to be reaped
	time.Sleep(interval * 2)

	// Check that the entry has been removed
	_, ok = cache.Get(key)
	if ok {
		t.Errorf("Expected key %s to be reaped from cache, but it was found", key)
	}
}

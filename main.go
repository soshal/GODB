package main

import (
	
	
	"sync"
)

// to  hold our key-value pairs and synchronization mechanism:
type  KeyValueStore struct {
	data map[string]string // Map to hold key-value pairs
	mu sync.RWMutex // Mutex for synchronization
}

// Create adds a new key-value pair to the store
func (kv *KeyValueStore) Create(key, value string) {
    kv.mu.Lock() // Acquire write lock
    defer kv.mu.Unlock() // Release lock when function exits

    kv.data[key] = value // Add key-value pair to map
}


// Read retrieves the value for a given key
func (kv *KeyValueStore) Read(key string) string {
    kv.mu.RLock() // Acquire read lock
    defer kv.mu.RUnlock() // Release lock when function exits

    return kv.data[key]
}


// Update modifies the value for a given key
func (kv *KeyValueStore) Update(key, value string) {
    kv.mu.Lock() // Acquire write lock
    defer kv.mu.Unlock() // Release lock when function exits

    kv.data[key] = value
}

// Delete removes a key-value pair from the store
func (kv *KeyValueStore) Delete(key string) {
    kv.mu.Lock() // Acquire write lock
    defer kv.mu.Unlock() // Release lock when function exits

    delete(kv.data, key)
}

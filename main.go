package main

import (
    "fmt"
    "sync"
)

// KeyValueStore represents a simple in-memory key-value store
type KeyValueStore struct {
    data map[string]string // Map to hold key-value pairs
    mu   sync.RWMutex      // Mutex for synchronization
}

// NewKeyValueStore creates a new instance of KeyValueStore
func NewKeyValueStore() *KeyValueStore {
    return &KeyValueStore{
        data: make(map[string]string),
    }
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

func main() {
    kvStore := NewKeyValueStore()

    // Demonstrate Create
    kvStore.Create("exampleKey", "exampleValue")
    fmt.Printf("Created: exampleKey -> %s\n", kvStore.Read("exampleKey"))

    // Demonstrate Update
    kvStore.Update("exampleKey", "newValue")
    fmt.Printf("Updated: exampleKey -> %s\n", kvStore.Read("exampleKey"))

    // Demonstrate Delete
    kvStore.Delete("exampleKey")
    fmt.Printf("Deleted: exampleKey -> %s\n", kvStore.Read("exampleKey"))
}

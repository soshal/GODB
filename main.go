package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "sync"
)

type KeyValueStore struct {
    data map[string]string
    mu   sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
    return &KeyValueStore{
        data: make(map[string]string),
    }
}

func (kv *KeyValueStore) Create(key, value string) {
    kv.mu.Lock()
    defer kv.mu.Unlock()

    kv.data[key] = value
}

func (kv *KeyValueStore) Read(key string) string {
    kv.mu.RLock()
    defer kv.mu.RUnlock()

    return kv.data[key]
}

func (kv *KeyValueStore) Update(key, value string) {
    kv.mu.Lock()
    defer kv.mu.Unlock()

    kv.data[key] = value
}

func (kv *KeyValueStore) Delete(key string) {
    kv.mu.Lock()
    defer kv.mu.Unlock()

    delete(kv.data, key)
}

func saveToFile(kv *KeyValueStore, filename string) error {
    kv.mu.RLock()
    defer kv.mu.RUnlock()

    data, err := json.Marshal(kv.data)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(filename, data, 0644)
    if err != nil {
        return err
    }

    return nil
}

func loadFromFile(kv *KeyValueStore, filename string) error {
    kv.mu.Lock()
    defer kv.mu.Unlock()

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }

    err = json.Unmarshal(data, &kv.data)
    if err != nil {
        return err
    }

    return nil
}

func mai() {
    kvStore := NewKeyValueStore()
    filename := "kvstore.json"

    err := loadFromFile(kvStore, filename)
    if err != nil {
        fmt.Println("No existing data found, starting with an empty store")
    } else {
        fmt.Println("Loaded existing data from file")
    }

    
    kvStore.Create("key2", "value")
    kvStore.Create("key3", "value")
    kvStore.Create("key4", "value")

    err = saveToFile(kvStore, filename)
    if err != nil {
        fmt.Printf("Error saving data to file: %v\n", err)
    } else {
        fmt.Println("Data successfully saved to file")
    }
}

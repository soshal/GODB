package main

import (
    "fmt"
    "strings"
)

func (kv *KeyValueStore) QueryByKey(pattern string) map[string]string {
    kv.mu.RLock()
    defer kv.mu.RUnlock()

    results := make(map[string]string)
    for key, value := range kv.data {
        if strings.Contains(key, pattern) {
            results[key] = value
        }
    }
    return results
}

func (kv *KeyValueStore) QueryByValue(pattern string) map[string]string {
    kv.mu.RLock()
    defer kv.mu.RUnlock()

    results := make(map[string]string)
    for key, value := range kv.data {
        if strings.Contains(value, pattern) {
            results[key] = value
        }
    }
    return results
}

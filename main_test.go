package main

import (
    "testing"
    "fmt"
)

func TestKeyValueStore(t *testing.T) {
    kvStore := NewKeyValueStore()

    // Test Create
    kvStore.Create("key1", "value1")
    fmt.Printf("Key-value pair created: key1 -> value1\n")
    if val := kvStore.Read("key1"); val != "value1" {
        t.Errorf("Expected value1, got %s", val)
    }

    // Test Update
    kvStore.Update("key1", "value2")
    fmt.Printf("Key-value pair updated: key1 -> value2\n")
    if val := kvStore.Read("key1"); val != "value2" {
        t.Errorf("Expected value2, got %s", val)
    }

	// Test Delete
	kvStore.Delete("key1")
	fmt.Printf("Key-value pair deleted: key1\n")
	if val := kvStore.Read("key1"); val != "" {
		t.Errorf("Expected empty string, got %s", val)
	}	
}
// adding more 

package main

import (
	"fmt"
)

type KeyValue struct {
    Key   string
    Value string
}

type HashTable struct {
    data [][]KeyValue
}

func NewHashTable(size int) *HashTable {
    return &HashTable{make([][]KeyValue, size)}
}

func (h *HashTable) _hash(key string) int {
    hash := 0
    for i := 0; i < len(key); i++ {
        hash = (hash + int(key[i]) * i) % len(h.data)
    }
    return hash
}

func (h *HashTable) Set(key string, value string) [][]KeyValue {
    address := h._hash(key)
    if h.data[address] == nil {
        h.data[address] = make([]KeyValue, 0)
    }
    h.data[address] = append(h.data[address], KeyValue{key, value})
    return h.data
}

func (h *HashTable) Get(key string) string {
    address := h._hash(key)
    currentBucket := h.data[address]
    if currentBucket != nil {
        for i := 0; i < len(currentBucket); i++ {
            if currentBucket[i].Key == key {
                return currentBucket[i].Value
            }
        }
    }
    return ""
}

func main() {
    myHashTable := NewHashTable(100)
    myHashTable.Set("082124606606", "Rony Setyawan")
    fmt.Println(myHashTable.Get("082124606606"))
}

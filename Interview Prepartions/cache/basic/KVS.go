package basic

import (
	"fmt"
	"sync"
)

// type
// Constructor
// Methods

type KeyValueStore struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}

func (j *KeyValueStore) Set(Key string, Val string) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.store[Key] = Val
}

func (j *KeyValueStore) Get(Key string) (string, bool) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	val, ok := j.store[Key]
	return val, ok
}

func BasicRedisCache() {

	kvs := NewKeyValueStore() // instance 1
	jvm := NewKeyValueStore() // instance 2

	kvs.Set("Jayendra", "Go")
	kvs.Set("Harsh", "Cpp")
	kvs.Set("Jash", "Java")

	jvm.Set("P1", "Football")
	jvm.Set("P2", "VolleyBall")
	jvm.Set("P3", "Basketball")

	if val, ok := kvs.Get("Jash"); ok {
		fmt.Println(val)
	} else {
		fmt.Println("Key Not Found")
	}

}

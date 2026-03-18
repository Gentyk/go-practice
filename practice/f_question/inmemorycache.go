// in-memory cache.
package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type StorageCache struct {
	storage map[string]string
	mu      sync.RWMutex
}

func (S *StorageCache) Set(k, v string) {
	S.mu.Lock()
	defer S.mu.Unlock()
	if value, ok := S.storage[k]; ok {
		fmt.Println("WARN. update ", k, value)
	}
	S.storage[k] = v
}

func (S *StorageCache) Get(k string) (v string, ok bool) {
	S.mu.RLock()
	defer S.mu.RUnlock()
	if value, ok := S.storage[k]; ok {
		return value, true
	}
	return "", false
}

func main() {
	var newCache Cache = &StorageCache{
		storage: make(map[string]string),
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		newCache.Set("bla", "bla")
		wg.Done()
	}()

	go func() {
		v, ok := newCache.Get("tt")
		if !ok {
			fmt.Println("значения не сущ", v)
		}
		wg.Done()
	}()
	go func() {
		v, ok := newCache.Get("bla")
		if ok {
			fmt.Println("значение сущ", v)
		}
		wg.Done()
	}()

	wg.Wait()
}

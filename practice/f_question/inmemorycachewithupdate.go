package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type SimpleCache struct {
	data     map[string]string
	mu       sync.RWMutex
	maxCount int
}

func (S *SimpleCache) Set(k, v string) {
	S.mu.Lock()
	defer S.mu.Unlock()
	if len(S.data) >= S.maxCount { // удаление при переполнении случайного
		for key := range S.data {
			delete(S.data, key)
			break
		}
	}
	S.data[k] = v
}

func (S *SimpleCache) Get(k string) (v string, ok bool) {
	S.mu.RLock()
	defer S.mu.RUnlock()
	if v, ok := S.data[k]; ok {
		return v, true
	}
	return "", false
}

type ShardedCache struct {
	caches []*SimpleCache
	_count int
}

func NewSC(n int) *ShardedCache {
	result := ShardedCache{
		caches: make([]*SimpleCache, 0, n),
		_count: n,
	}
	for i := 0; i < n; i++ {
		result.caches = append(result.caches, &SimpleCache{
			maxCount: 10,
			data:     make(map[string]string),
		})
	}
	return &result
}

func (SC *ShardedCache) getShardId(k string) (i int) {
	if len(k) == 0 {
		return 0
	}
	return int(k[0]) % SC._count
}

func (SC *ShardedCache) Get(k string) (v string, ok bool) {
	shardId := SC.getShardId(k)
	return SC.caches[shardId].Get(k)
}

func (SC *ShardedCache) Set(k string, v string) {
	shardId := SC.getShardId(k)
	SC.caches[shardId].Set(k, v)
}

func main() {
	var c Cache = NewSC(3)
	var m sync.WaitGroup

	for _, i := range []string{"a", "b", "c"} {
		m.Add(1)
		go func() {
			defer m.Done()
			c.Set(i, i)
		}()
	}

	for _, i := range []string{"a", "b", "c"} {
		m.Add(1)
		go func() {
			defer m.Done()
			v, ok := c.Get(i)
			fmt.Println(i, v, ok)
		}()
	}
	m.Wait()
}

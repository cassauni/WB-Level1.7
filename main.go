package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[int]int
}

func NewCache(capacity int) *Cache {
	return &Cache{data: make(map[int]int, capacity)}
}

func (c *Cache) Get(key int) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	if !ok {
		fmt.Println("Cache miss")
		return 0
	}
	return value
}

func (c *Cache) Set(key int, value int) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func main() {
	cache := NewCache(30)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 30; i++ {
			cache.Set(i, i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 30; i++ {
			time.Sleep(100 * time.Millisecond)
			value := cache.Get(i)
			fmt.Printf("value of map[%d] is %d\n", i, value)
		}
	}()

	wg.Wait()
}

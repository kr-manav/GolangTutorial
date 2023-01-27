package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++

}

func (c *Counter) Value(key string) int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
		go c.Inc("somkey")
	}

	time.Sleep(time.Second)

	fmt.Println(c.Value("somkey"))
	fmt.Println(c.Value("somekey"))

}

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func (c *Counter) IncrCounter() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *Counter) Get() int {
	return c.counter
}

func main() {
	c := new(Counter)

	for i := 1; i <= 1010; i++ {
		c.wg.Add(1)
		go func(c *Counter) {
			c.IncrCounter()
			c.wg.Done()
		}(c)

	}
	c.wg.Wait()
	fmt.Println(c.Get())
}

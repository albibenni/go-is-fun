package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Release the lock when function returns
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := &Counter{}

	t := time.Now()
	// Start 100 goroutines that increment the counter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Final count: %d\n", counter.Value()) // Will be 100,000
	fmt.Printf("Final count - time: %v\n", time.Since(t))
}

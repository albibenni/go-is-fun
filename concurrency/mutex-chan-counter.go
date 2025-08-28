package main

import (
	"fmt"
	"sync"
	"time"
)

// With Mutex - natural fit
type Counterr struct {
	mu    sync.Mutex
	value int
}

func (c *Counterr) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// With Channel - awkward and inefficient
type ChannelCounter struct {
	ch    chan int
	value int
}

func (c *ChannelCounter) Increment() {
	c.ch <- 1 // Send increment request
}

func (c *ChannelCounter) run() {
	for delta := range c.ch {
		c.value += delta // Only one goroutine modifies value
	}
}

func main() {
	counterCM := &Counterr{}
	counterCC := &ChannelCounter{}

	t := time.Now()
	var waitGCM sync.WaitGroup

	waitingSomeCounterr(&waitGCM, counterCM)
	fmt.Printf("Time since 1st grp: %v\n", time.Since(t))

	var waitGCC sync.WaitGroup
	waitingSomeCCounter(&waitGCC, counterCC)
	fmt.Printf("Time since 1st grp: %v\n", time.Since(t))
}

func waitingSomeCounterr(wg *sync.WaitGroup, counter *Counterr) {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

}

func waitingSomeCCounter(wg *sync.WaitGroup, counter *ChannelCounter) {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

}

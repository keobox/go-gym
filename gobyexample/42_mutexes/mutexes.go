package main

import (
	"fmt"
	"sync"
)

// Concurrently change a map of counters

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// WARNING: since the mutex mu cannot be copied
// the above struct must be passed as pointer

func (c *Container) inc(name string) {
	// Lock the mutex to protect the
	// concurrency region
	c.mu.Lock()
	// Declare here but this unlock th mutex at the end
	// of the method
	defer c.mu.Unlock()
	// Protected region
	c.counters[name]++
}

func main() {
	// Note Container.mu initialized to 0 and is usable
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Function to increment a named counter in the map
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		// Signal the go routine is finished
		wg.Done()
	}

	// Fire 3 go routines
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Wait all routines to finish
	wg.Wait()
	fmt.Println(c.counters)
}

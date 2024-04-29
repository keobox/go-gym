package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Managing state when concurrency is in place

func main() {

	// Positive atomic counter
	var ops atomic.Uint64

	// Wait all concurrent operations are finished
	var wg sync.WaitGroup

	// Lauch 50 operations that increase the ops
	// counter 1000 times each
	// so we expect 50000 increases

	for i := 0; i < 50; i++ {
		// Add a go routine
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// Increase ops atomically
				ops.Add(1)
			}
			// Signal that go routine has finished
			// and decrease wg
			wg.Done()
		}()
	}

	// Wait all go routines has finished
	wg.Wait()
	// We expect 50000 ops
	fmt.Println("ops:", ops.Load())

	// If ops would be a normal counter (non Atomic) then ops will
	// be a different number since this is a race condition

}

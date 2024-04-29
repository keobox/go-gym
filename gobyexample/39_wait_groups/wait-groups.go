package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This wait group waits for all the
	// go routines to finish
	var wg sync.WaitGroup

	// Create 5 go routines
	// the worker function is wrapped
	// for being concurrency agnostic
	for i := 1; i <= 5; i++ {
		// Increase the wg's counter by 1
		wg.Add(1)

		// Next line is a fix for go < 1.21
		id := i
		// Closure for wg
		go func() {
			// Think of defer like a kind of
			// finally, it will execute at the end but
			// is evaluated instantly
			defer wg.Done() // decrease the counter
			// call the worker
			worker(id)
		}()
	}

	// Wait for all the routine to finish
	wg.Wait()

	// Note: is case wg is passed to functions
	// it should be passed as a pointer
}

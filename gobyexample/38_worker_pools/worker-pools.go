package main

import (
	"fmt"
	"time"
)

// It's possible to use channels and go routines to
// implement a worker pool

// This worker receive jobs from the jobs channel and send result to
// the result channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func main() {

	start := time.Now()

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Signal that jobs are finished
	close(jobs)

	// Collect the results
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// Note that the elapsed time should be less that 5 seconds
	// because of the parallelism
	fmt.Println("elapsed time", time.Since(start))
}

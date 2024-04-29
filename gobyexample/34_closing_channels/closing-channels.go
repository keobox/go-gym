package main

import "fmt"

func main() {

	// Closing a channel can be used as a completion signal
	// for channel's receivers

	// Buffered non blocking channel
	jobs := make(chan int, 5)
	// Unbuffered blocking channel
	done := make(chan bool)

	// worker
	go func() {
		// more will be false if the channel is closed
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Wait worker done
	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs", ok)
}

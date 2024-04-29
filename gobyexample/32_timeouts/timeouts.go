package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffered channel so a send is not blocking
	c1 := make(chan string, 1)

	go func() {
		// Execution time 2 seconds
		time.Sleep(2 * time.Second)
		// Not blocking send
		c1 <- "result 1"
	}()

	select {
	// Not blocking receive (my comment here)
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		// Timeout 1 second elapsed
		fmt.Println("timeout 1")
	}

	// Buffered channel so a send is not blocking
	c2 := make(chan string, 1)

	go func() {
		// Execution time 2 seconds
		time.Sleep(2 * time.Second)
		// Not blocking send
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		// Timeout 3 seconds elapsed
		fmt.Println("timeout 1")
	}
}

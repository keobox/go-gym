package main

import (
	"fmt"
	"time"
)

// Usa a channel to notify a go routine has completed a job

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	// channel for synchronization
	done := make(chan bool, 1)

	// start go routine
	go worker(done)

	// wait until is done
	<-done
}

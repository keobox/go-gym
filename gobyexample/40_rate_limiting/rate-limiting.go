package main

import (
	"fmt"
	"time"
)

// Go supports rate limiting with go routines
// channels and tickers

func main() {

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// simple rate limiter
	limiter := time.Tick(200 * time.Millisecond)

	// Consume a request every 200 ms
	for req := range requests {
		// Blocks on limiter
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Bursty limiter: a limiter that can consume
	// a fixed size burst of requests but then
	// is rate limited for requests up to the size
	// of the burst

	// This limiter accepts a burst of 3 requests
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the allowed burst
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Add values up to the allowed burst size
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// send 3 + 2 requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// Consume the bursty requests with rate limiting
	for req := range burstyRequests {
		// Blocks on limiter. 3 will be served quickly
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

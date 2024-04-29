package main

import (
	"fmt"
	"time"
)

func main() {
	// Tickers are recurring timers, they keep
	// firing until stopped, but the mechanism
	// it's the same as timers. A channel get
	// notified when a ticker fires.

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				// Exit the routine
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Give a chance to ticker to tick
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
	// Exit go routine
	done <- true
	fmt.Println("End")

}

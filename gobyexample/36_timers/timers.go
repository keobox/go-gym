package main

import (
	"fmt"
	"time"
)

func main() {

	// The timer built in gives you a timer with a channel
	// that gets notified when the timer expires

	timer1 := time.NewTimer(2 * time.Second)

	// The next line blocks the timer channel until
	// the timer expires
	<-timer1.C
	fmt.Println("timer1 fired")

	// All the above is equivalent to a time.Sleep, with
	// the difference that a timer can be stopped
	// (my comment: and is an object)

	// Example of a timer stop
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	stop := timer2.Stop()
	// stop := false
	if stop {
		fmt.Println("timer2 stopped")
	}

	// Give a chance to timer2 to fire
	time.Sleep(3 * time.Second)
}

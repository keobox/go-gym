package main

import "fmt"

// Accepts a channel argument just for sending msg
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Accepts a channel for receive (pings) and one for send (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

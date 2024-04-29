package main

import "fmt"

func main() {
	// Channels are pipes that connect concurrent goroutine

	// Create a channel to send strings
	messages := make(chan string)

	// Create a go routine from an anonymous function that sends a message
	// to the "messages" channel

	// This creates the go routine, from anonymous function
	// and call it in one line, see "()".
	// then go routine sends a message in the channel
	go func() { messages <- "ping" }()

	// The main routine (or thread) waits for the message,
	// hence the main routine and the go routine has a form
	// of syncronization
	msg := <-messages
	fmt.Println(msg)
}

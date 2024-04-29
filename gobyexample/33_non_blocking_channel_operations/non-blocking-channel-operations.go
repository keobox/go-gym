package main

import "fmt"

func main() {
	// It's possible to use a select with a default value
	// to implement not blocking operations on channels

	// Unbuffered channels, so it's blocking
	messages := make(chan string)
	signals := make(chan string)

	// Something like this will give a deadlock error
	// because there's no receiver defined (in a go routine)
	// msg := <-messages
	// fmt.Println(msg)

	// This is a not blocking receive, since the receiver
	// is not defined it will take the default case immediately
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// This is a not blocking send
	// since the receiver
	// is not defined it will take the default case immediately
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// This is a multi channel not blocking receive, since the receiver
	// is not defined it will take the default case immediately
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

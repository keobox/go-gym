package main

import "fmt"

func main() {

	// by default channels are unbuffered, meaning that is not
	// possible to send a message if there's no a concurrent
	// ready to receive the message.

	// But is possible to define a buffered channel that can
	// take some messages even if there's no receiver ready
	// to consume them.

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// there's no concurrent receive

	// we can later receive these messages
	// (same thread in this case?)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

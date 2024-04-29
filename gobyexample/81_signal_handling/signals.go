package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// UNIX signal handling
// For example for closing gracefully on SIGTERM or SIGINT

func main() {

	// Go signal notification works by sending os.Signal values on a channel.
	// We’ll create a channel to receive these notifications.
	// Note that this channel should be buffered.
	sigs := make(chan os.Signal, 1)

	// Register which signal to handle
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	// wait untile the signal is received
	<-done
	fmt.Println("exiting")
}

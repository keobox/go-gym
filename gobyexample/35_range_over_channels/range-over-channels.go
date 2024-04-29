package main

import "fmt"

func main() {

	// Is possible to range over channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// Is possible to close a buffered channel
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}

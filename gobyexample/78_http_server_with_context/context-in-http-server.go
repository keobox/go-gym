package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// context.Context is used for controlling cancellation in this example

	// Can be used for deadlines, cancellation signals, and other
	// http-request-scoped values across API boundaries and goroutines

	// context access is via the Context() method of a Request
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	/* Wait for a few seconds before sending a reply to the client.
	This could simulate some work the server is doing. While working,
	keep an eye on the context’s Done() channel for a signal that we should cancel the work
	and return as soon as possible */

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

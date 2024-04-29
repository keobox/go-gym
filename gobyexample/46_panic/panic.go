package main

import "os"

func main() {

	// The next line exit the program with an error and exit code 2
	// panic("error")

	// error handling with panic
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {

	// this defer will not run
	defer fmt.Println("!")

	// Immediate exit with return code 3
	// Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status.
	// If you’d like to exit with a non-zero status you should use os.Exit.
	os.Exit(3)
}

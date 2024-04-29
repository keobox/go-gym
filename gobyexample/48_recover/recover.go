package main

import "fmt"

// Recover must be called in a deferred function
func mayPanic() {
	panic("a problem")
	// fmt.Println("No panic")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered error:\n", r)
		}
	}()

	mayPanic()

	// Next code will not run
	fmt.Println("After mayPanic()")

}

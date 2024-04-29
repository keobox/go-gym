package main

import "fmt"

// intSeq returns a function that returns an int
func intSeq() func() int {
	i := 0
	// this is an anonymous function that close i
	// it's a closure around i
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	// test the closure
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// new closure restarts because there's a new closed i
	newInts := intSeq()
	fmt.Println(newInts())
}

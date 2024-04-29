package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// test fact
	fmt.Println(fact(7))

	// a closure can be recursive but this requires that
	// the closure is declared with a typed var before
	// definition
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	// test fib
	fmt.Println(fib(7))
}

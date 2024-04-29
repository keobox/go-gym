package main

import "fmt"

// simple function
func plus(a int, b int) int {

	// return is required
	return a + b
}

// if muntiple args have the same type, declare it once
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

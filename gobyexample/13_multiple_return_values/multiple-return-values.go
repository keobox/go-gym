package main

import "fmt"

// declare multiple return values
func vals() (int, int) {
	return 3, 7
}

func main() {

	// multiple assingment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// partial assignment
	_, c := vals()
	fmt.Println(c)
}

package main

import "fmt"

// A variadic function it's a function that takes any number of trailing arguments

// arbitrary int arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// calling from a slice: use func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

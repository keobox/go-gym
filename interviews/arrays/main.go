package main

import "fmt"

// This week, today in 2025-08-01, I saw a debate in linked in regarding
// the fact that 75% of programmers cannot write a simple program using
// data structure or cannot answer to simple data structure questions, like
// what is the difference between an array-list or a linked-list and
// what are advantages, disavantages.

// The question was, write an array (with 100 numbers) and return the sum of
// the even numbers in the array.

// Post by Marco Cecconi
// Antirez commented and reacted by writing this little program in C and Ruby in youtube,
// I'll try with go

// Note by me: for 100 consecutive numbers the sum can be done in O(1) using the Gauss sum
// sum = N*(N+1)/2
// but in this case the sum is just for even numbers, so can't be used.

func main() {
	const N = 100

	// Create array
	var a [N]int
	for i := range N {
		a[i] = i
	}

	// Computation of sum of the even values
	sum := 0
	for _, n := range a {
		if n%2 == 0 {
			sum += n
		}
	}
	fmt.Println(sum)
}

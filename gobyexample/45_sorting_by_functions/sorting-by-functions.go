package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Is possible to use functions to customize sorting
// instead of using natural ordering
func main() {

	// Sort by length

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))

	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println("fruits:", fruits)

	// Sort by age
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})

	fmt.Println("people:", people)
}

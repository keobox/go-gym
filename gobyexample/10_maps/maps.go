package main

import (
	"fmt"
	"maps"
)

func main() {

	// Create an empty map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Not exixting key gets 0 value
	v3 := m["k3"]
	fmt.Println("v3", v3)

	// NOTE: "k3" is not added
	fmt.Println("len:", len(m))

	// Delete
	delete(m, "k2")
	fmt.Println("map:", m)

	// Remove all
	clear(m)
	fmt.Println("map:", m)

	// Is a key present?
	// Use the second optional return value
	_, present := m["k2"]
	fmt.Println("present:", present)

	// Initialization
	// the order may not be respected
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Equality
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

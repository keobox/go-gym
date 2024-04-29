package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a "receiver" type of "*rect"
func (r *rect) area() int {
	return r.width * r.height
}

// This method has a value "receiver type"
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	// Go automatically handles conversion between values and pointers
	// for methods calls

	// use pointer receivers to avoid copying or mutate
	fmt.Println("area: ", rp.area())
	// use value receivers for immutability, or copy
	fmt.Println("perim:", rp.perim())
}

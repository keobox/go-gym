package main

import (
	"fmt"
	"math"
)

// This is an interface
type geometry interface {
	area() float64
	perim() float64
}

// We have 2 types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface you just need to implement the interface's methods
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// We can use an interface in a generic measure function
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

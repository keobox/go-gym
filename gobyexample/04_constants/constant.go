package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	// "const" can be where "var" can
	const n = 500000000

	// arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// a const has no type until specifically chosen
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

package main

// math/rand/v2 only from go1.22.2

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// int 0 <= n < 100
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// float 0 <= f < 1.0
	fmt.Println(rand.Float64())

	// float 5.0 <= f < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// seed usage
	// https://en.wikipedia.org/wiki/Permuted_congruential_generator
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}

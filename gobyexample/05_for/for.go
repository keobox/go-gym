package main

import "fmt"

func main() {

	// single condition
	i := 3
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// no condition
	for {
		fmt.Println("loop")
		break
	}

	// continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// loop on data
	data := []int{1, 2, 3, 4}
	for i, val := range data {
		fmt.Println(i, val)
	}
}

package main

import "fmt"

func main() {
	// basic
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// without else
	if 8%4 == 0 {
		fmt.Println("8 divisible by 4")
	}

	// logical operators
	if 8%2 == 0 && 7%2 == 0 {
		fmt.Println("either 8 and 7 are even")
	}

	// A statement can precede a conditional and is available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")

	}
}

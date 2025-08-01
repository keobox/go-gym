package main

import "fmt"

// This is something similar to a generator in Python
// Behing the scene the caller code inject a function,
// the "yield" paramenter that returns false if the
// stop condiftion is met.
func Sequence(yield func(int) bool) {
	i := 0
	for {
		if !yield(i) {
			return
		}
		i += 1
	}
}

func main() {
	for i := range Sequence {
		if i > 100 {
			break
		}
		fmt.Println(i)
	}
}

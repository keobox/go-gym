package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}

	// Range on values
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Range on index and values
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on a map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range on keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range in a string returns ascii
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

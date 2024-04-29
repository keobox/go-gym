package main

import "fmt"

// A person struct type
type person struct {
	name string
	age  int
}

// A constructor function
func newPerson(name string) *person {
	// Create a struct and assign a field
	p := person{name: name}
	p.age = 42
	// The local variable will survive the scope
	return &p // yields a pointer
}

func main() {

	// Create a struct
	fmt.Println(person{"Bob", 20})

	// Is possible to name the fields during creation
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields are zeroed
	fmt.Println(person{name: "Fred"})

	// Yield a pointer to a struct
	fmt.Println(&person{name: "Ann", age: 40})

	// Use constructor (idiomatic)
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	// Access fields with a dot
	fmt.Println(s.name)

	sp := &s
	// It's possible to access pointer to struct fields with a dot
	// fields are automatically deferenced
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// Anonymous struct: used in table drive tests
	// https://gobyexample.com/testing-and-benchmarking

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}

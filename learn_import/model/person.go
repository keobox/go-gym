package model

import (
	"fmt"
)

// struct name and fields name have to be capitalized to be exported
type Person struct {
	Name string
	Age  int
}

// An exportable method
func (p Person) Introduce() {
	p.privateIntroduce()
}

func (p Person) privateIntroduce() {
	fmt.Printf("My name is %s and I'm %d years old\n", p.Name, p.Age)
}

package main

import (
	"fmt"

	// is necessary to prefix the relative path of the package
	// with the first line in go.mod
	"example.com/go-gym/learn_import/model"
)

// FIXME is possible to run this only from terminal
// cd learn_import
// go run .
//
// Is not possible to run from vs-code (ctlr-F5)
func main() {
	fmt.Print("Hello ")
	fmt.Println(AGlobalVariableFromAnotherFile)
	if AFunctionFromAnotherFile() {
		fmt.Println("and from a function in another file")
	}
	if stillVisibleEvenIfIsNotCapitalized() {
		fmt.Println("and from a second function in another file")
	}
	p := model.Person{
		Name: "Paolo",
		Age:  32,
	}
	p.Introduce()
	// private is unreacheable
	// p.privateIntroduce()
}

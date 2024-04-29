package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	// embedding is a field without a name
	base
	str string
}

func main() {

	// Literal creation for base is explicit
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	// I can access base.num with co.num directly
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// I can access base.num with the full dot path
	fmt.Println("also num:", co.base.num)

	// I can access the base method from co
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// co implements the describer interface even if
	// is base embedded that implements it
	var d describer = co
	// Enbedded struct with methods is a way to confer
	// interface implementation into other structs
	fmt.Println("describer:", d.describe())
}

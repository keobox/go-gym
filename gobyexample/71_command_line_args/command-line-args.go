package main

import (
	"fmt"
	"os"
)

func main() {

	// sequential command line arguments

	// args[0] is the command's name
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

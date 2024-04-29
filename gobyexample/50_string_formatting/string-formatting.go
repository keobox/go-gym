package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// The formatting options are in printf style
	// The next are "verbs" formatting for general Go Values
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Includes field names
	fmt.Printf("struct2: %+v\n", p)

	// Go syntax representation
	fmt.Printf("struct3: %#v\n", p)

	// Print a type
	fmt.Printf("type: %T\n", p)

	// booleans
	fmt.Printf("bool: %t\n", true)

	// integer in decimal notation
	fmt.Printf("int: %d\n", 123)

	// binary representation
	fmt.Printf("bin: %b\n", 14)

	// char given and integer
	fmt.Printf("char: %c\n", 33)

	// hex
	fmt.Printf("hex: %x\n", 456)

	// decimal float
	fmt.Printf("float1: %f\n", 78.9)

	// scientific notation float
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// basic string
	fmt.Printf("str1: %s\n", "\"string\"")

	// Double quotes
	fmt.Printf("str2: %q\n", "\"string\"")

	// base16 of a string with 2 output chars per byte of input
	fmt.Printf("str3: %x\n", "hex this")

	// representation of a pointer
	fmt.Printf("pointer: %p\n", &p)

	// width of an integer
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// width and precision of a float
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// left justify and width and precision of a float
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// basic right justify a string
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// basic left justify a string
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// same rules for sprintf and fprintf
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice.
	// The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

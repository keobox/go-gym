package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// helper for error checking

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Read the intere file in memory
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	// more control on reading a file
	f, err := os.Open("/tmp/dat")
	check(err)

	// read 5 bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// seek to a known location in file
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// the same can be done with the io package
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// rewind a file
	_, err = f.Seek(0, 0)
	check(err)

	// buffered reader, more efficient in some cases
	r4 := bufio.NewReader(f)
	// read 5 bytes
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}

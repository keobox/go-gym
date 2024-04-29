package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// A pattern matches a string?
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Use Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")
	// Then use methods
	fmt.Println(r.MatchString("peach"))

	// find the first match
	fmt.Println(r.FindString("peach punch"))
	// find start and end indexes
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// find submatches so returns info for p([a-z]+)ch and ([a-z]+) the submatch
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// same for indexes
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The All variants find all matches not just the first
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// Just the first two
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// It's possible to use byte slices and drop String from methods
	fmt.Println(r.Match([]byte("peanch")))

	// Useful fro global variables, use MustCompile that panics
	// instead of returning error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// Replace
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Replace matched text with a func
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}

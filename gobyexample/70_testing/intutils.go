package main

// test target

// got test runs test
// Testing code typically lives in the same package as the code it tests.
// example: if source code is intutils.go test is in intutils_test.go, same package

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

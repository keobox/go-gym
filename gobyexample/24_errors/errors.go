package main

import (
	"errors"
	"fmt"
)

// No exceptions in Go, just return an explicit error

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New returns a basic error type with message
		return -1, errors.New("can't work with 42")
	}
	// by convention an error is the last returned value
	return arg + 3, nil
}

// A custom error must implement Error() string interface
type argError struct {
	arg  int
	prob string
}

// Error customization is here
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil

}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	// Same code for f2 with custom error
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	// error check with access to custom error fields
	// is necessary to get the error instance via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}

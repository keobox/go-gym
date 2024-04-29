package main

import (
	"fmt"
	"testing"
)

// go test -v

// A test function begins with Test.

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Idiom: table drive tests where test value and expected results are rows
// in a table.

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// a benchmark begins with "Benchmark"

// cd gobyexample/66_testing
// go test -bench=.

func BenchmarkIntMin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

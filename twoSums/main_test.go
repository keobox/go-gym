package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestingIsWorking(t *testing.T) {

	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")

}

func TestTwoSumExample1(t *testing.T) {
	result := TwoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	assert.Equal(t, expected, result, "The two arrays should be the same")

}
func TestTwoSumExample2(t *testing.T) {
	result := TwoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}
	assert.Equal(t, expected, result, "The two arrays should be the same.")

}

func TestTwoSumExample3(t *testing.T) {
	result := TwoSum([]int{3, 3}, 6)
	expected := []int{0, 1}
	assert.Equal(t, expected, result, "The two arrays should be the same.")

}

func TestTwoSumImprovedExample1(t *testing.T) {
	result := TwoSumImproved([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	assert.Equal(t, expected, result, "The two arrays should be the same")

}
func TestTwoSumImprovedExample2(t *testing.T) {
	result := TwoSumImproved([]int{3, 2, 4}, 6)
	expected := []int{1, 2}
	assert.Equal(t, expected, result, "The two arrays should be the same.")

}

func TestTwoSumImprovedExample3(t *testing.T) {
	result := TwoSumImproved([]int{3, 3}, 6)
	expected := []int{0, 1}
	assert.Equal(t, expected, result, "The two arrays should be the same.")

}

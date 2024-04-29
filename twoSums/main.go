package main

// leetcode problem

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Quadratic time
func TwoSum(nums []int, target int) []int {
	var result []int
	result = make([]int, 2)
	var found bool
	found = false
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i == j {
				continue
			}
			if v1+v2 == target {
				result[0] = i
				result[1] = j
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	return result
}

// Linear time
func TwoSumImproved(nums []int, target int) []int {
	var result []int
	result = make([]int, 2)
	var conj map[int]int
	conj = make(map[int]int)
	for i, v := range nums {
		// search for target - i was seen
		val, found := conj[v]
		if found {
			result[0] = val
			result[1] = i
			break
		} else {
			conj[target-v] = i
		}
	}
	return result
}

func main() {
	TwoSum([]int{2, 7, 11, 15}, 9)
	TwoSumImproved([]int{2, 7, 11, 15}, 9)
}

package main

import "strconv"

func findNumbers(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if len(strconv.Itoa(nums[i]))% 2 == 0 {
			result++
		}
	}
	return result
}

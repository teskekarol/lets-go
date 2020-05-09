package main

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	original := make([]int, len(nums))
	results := make([]int, len(nums))
	copy(original, nums)
	sort.Ints(nums)

	for oI, v := range original {
		for i, w := range nums {
			if v == w {
				results[oI] = i
				break
			}
		}
	}
	return results
}

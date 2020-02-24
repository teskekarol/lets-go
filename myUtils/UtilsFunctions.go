package myUtils

import "math/bits"

const (
	MaxInt  int  = (1<<bits.UintSize)/2 - 1
	MinInt  int  = (1 << bits.UintSize) / -2
)

func MinInArray(array []int) (int, int) {
	result := MaxInt
	index := 0
	for i, v := range array {
		if v < result {
			result = v
			index = i
		}
	}
	return index, result
}

func MaxInArray(array []int) (int, int) {
	result := MinInt
	index := 0
	for i, v := range array {
		if v > result {
			result = v
			index = i
		}
	}
	return index, result
}

func MinOf(left int, right int) int {
	if left <= right {
		return left
	} else {
		return right
	}
}

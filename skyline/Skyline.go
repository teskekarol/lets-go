package main

import (
	"fmt"
	"math/bits"
)

func main() {
	input := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0}}

	result := maxIncreaseKeepingSkyline(input)

	fmt.Print(result)

}

func maxIncreaseKeepingSkyline(skyline [][]int) int {
	result := 0
	var leftRight []int
	upDownMax := make(map[int]int)
	for _, verticales := range skyline {
		for indexValue, value := range verticales {
			if upDownMax[indexValue] < value {
				upDownMax[indexValue] = value
			}
		}
	}
	for _, verticales := range skyline {
		_, maxVerticalValue := MaxInArray(verticales)

		for vertIndex := range verticales {
			setTo := MinOf(upDownMax[vertIndex], maxVerticalValue)
			result += setTo - verticales[vertIndex]
			verticales[vertIndex] = setTo
		}

		leftRight = append(leftRight, maxVerticalValue)
	}

	return result
}

func MaxInArray(array []int) (int, int) {
	MinInt := (1 << bits.UintSize) / -2
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

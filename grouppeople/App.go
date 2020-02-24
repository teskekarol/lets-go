package main

import "fmt"

func main() {
	input := []int{2,1,3,3,3,2}
	result := groupThePeople(input)

	fmt.Print(result)
}

func groupThePeople(groupSizes []int) [][]int {
	mineGroupHas := make(map[int][]int)
	for i := 0; i < len(groupSizes); i++ {
		expectedGroupSize := groupSizes[i]
		mineGroupHas[expectedGroupSize] = append(mineGroupHas[expectedGroupSize], i)
	}

	var temp [][]int
	for key, val := range mineGroupHas {
		for i := 0; i < len(val); i += key {
			temp = append(temp, val[i:i+key])
		}
	}

	return temp
}

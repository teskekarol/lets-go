package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	input := [][]int{
		{3,0,8,4},
		{2,4,5,7},
		{9,2,6,3},
		{0,3,1,0}}

	result := maxIncreaseKeepingSkyline(input)
	fmt.Print()
	expected := 35
	if result != expected {
		t.Errorf("Result is %d, expected %d", result, expected)
	}

}
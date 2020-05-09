package main

import (
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	//given
	input := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	//when
	result := kidsWithCandies(input, extraCandies)

	//then
	expected := []bool{true, true, true, false, true}
	if NotEqual(result, expected) {
		t.Errorf("Result is %t, expected %t", result, expected)
	}
}

func Test_kidsWithCandies2(t *testing.T) {
	//given
	input := []int{4, 2, 1, 1, 2}
	extraCandies := 1

	//when
	result := kidsWithCandies(input, extraCandies)

	//then
	expected := []bool{true, false, false, false, false}
	if NotEqual(result, expected) {
		t.Errorf("Result is %t, expected %t", result, expected)
	}
}

func NotEqual(a, b []bool) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
}

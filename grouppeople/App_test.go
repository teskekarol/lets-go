package main

import (
	"reflect"
	"testing"
)

func TestGroupPeople(t *testing.T) {
	//given
	input := []int{12, 345, 2, 6, 7896}

	//when
	result := groupThePeople(input)

	//then
	expected := [][]int{
		{0, 1, 2},
		{3, 4, 6},
		{5},
	}
	if reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestSmallerThanCurrent(t *testing.T) {
	//given
	input := []int{8,1,2,2,3}

	//when
	result := smallerNumbersThanCurrent(input)

	//then
	expected := []int{4,0,1,1,3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}


func TestSmallerThanCurrent2(t *testing.T) {
	//given
	input := []int{6,5,4,8}

	//when
	result := smallerNumbersThanCurrent(input)

	//then
	expected := []int{2,1,0,3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestCreateTargetArray(t *testing.T) {
	//given
	nums := []int{0,1,2,3,4}
	index := []int{0,1,2,2,1}

	//when
	result := createTargetArray(nums, index)

	//then
	expected := []int{0,4,1,3,2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}
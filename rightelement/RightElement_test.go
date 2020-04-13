package main

import (
	"reflect"
	"testing"
)

func TestRightElement(t *testing.T) {
	//given
	input := []int{17, 18, 5, 4, 6, 1}

	//when
	result := replaceElements(input)

	//then
	expected := []int{18, 6, 6, 6, 1, -1}
	if reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}

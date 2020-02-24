package main

import "testing"

func TestEventDigit(t *testing.T) {
	//given
	input := []int{6,7,8,2,7,1,3,9,nil,1,4,nil,nil,nil,5}

	//when
	result := sumEvenGrandparent(input)

	//then
	expected := 2
	if result != expected {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}
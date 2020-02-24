package main

import "testing"

func TestEventDigit(t *testing.T) {
	//given
	input := []int{12,345,2,6,7896}

	//when
	result := findNumbers(input)

	//then
	expected := 2
	if result != expected {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}
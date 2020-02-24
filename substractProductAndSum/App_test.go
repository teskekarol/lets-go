package main

import "testing"

func TestProductAndSum(t *testing.T) {
	input := 234

	result := subtractProductAndSum(input)

	expected := 15
	if result != expected {
		t.Errorf("Result is %d, expected %d", result, expected)
	}
}
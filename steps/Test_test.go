package main

import "testing"

func TestCount14(t *testing.T) {
	const val = 14

	result := numberOfSteps(val)

	if result != 6 {
		t.Errorf("Result is %d", result)
	}
}

func TestCount8(t *testing.T) {
	const val = 8

	result := numberOfSteps(val)

	if result != 4 {
		t.Errorf("Result is %d", result)
	}
}

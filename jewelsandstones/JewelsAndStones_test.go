package main

import "testing"

func TestDecompressRLElistGowno(t *testing.T) {
	first, second := "aA", "aAAbbbb"
	result := numJewelsInStones(first, second)

	if result != 3 {
		t.Errorf("Result is %d", result)
	}
}
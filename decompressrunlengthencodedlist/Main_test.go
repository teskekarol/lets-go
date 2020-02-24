package main

import (
	"reflect"
	"testing"
)

func TestDecompressRLElist(t *testing.T) {
	var input = []int{1,2,3,4}
	var expected = []int{2,4,4,4}

	result := decompressRLElist(input)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Result is %d", result)
	}
}
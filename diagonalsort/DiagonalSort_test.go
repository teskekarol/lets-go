package main

import (
	"reflect"
	"testing"
)

func TestDiagonalSort(t *testing.T) {
	var input = [][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}}
	var expected = [][]int{{1,1,1,1},{1,2,2,2},{1,2,3,3}}

	result := diagonalSort(input)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Result is %d", result)
	}
}
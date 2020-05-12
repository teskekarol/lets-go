package main

import (
	"reflect"
	"testing"
)

func Test_destCity(t *testing.T) {
	//given
	input := [][]string{{"London","New York"},{"New York","Lima"},{"Lima","Sao Paulo"}}

	//when
	result := destCity(input)

	//then
	expected := "Sao Paulo"
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %s, expected %s", result, expected)
	}
}

func Test_destCity2(t *testing.T) {
	//given
	input := [][]string{{"B", "C"},{"D", "B"},{"C", "A"}}

	//when
	result := destCity(input)

	//then
	expected := "A"
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %s, expected %s", result, expected)
	}
}

func Test_destCity3(t *testing.T) {
	//given
	input := [][]string{{"A", "Z"}}

	//when
	result := destCity(input)

	//then
	expected := "Z"
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result is %s, expected %s", result, expected)
	}
}
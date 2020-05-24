package main

import (
	"reflect"
	"testing"
)

func Test_processQueries(t *testing.T) {
	type args struct {
		queries []int
		m       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{3, 1, 2, 1}, 5}, []int{2, 1, 2, 1}},
		{"test2", args{[]int{4,1,2,2}, 4}, []int{3,1,2,0}},
		{"test3", args{[]int{7,5,5,8,3}, 8}, []int{6,5,0,7,5}},
		{"test4", args{[]int{1,2,1}, 3}, []int{0,1,1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processQueries(tt.args.queries, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processQueries() = got %v, want %v", got, tt.want)
			}
		})
	}
}

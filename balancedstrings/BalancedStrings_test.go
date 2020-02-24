package main

import "testing"


func Test_balancedStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "fine short", args{"RLRRLLRLRL"}, 4},
		{ "fine short 2", args{"RLLLLRRRLR"}, 3},
		{ "fine short 3", args{"LLLLRRRR"}, 1},
		{ "fine short 4", args{"RLRRRLLRLL"}, 2},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
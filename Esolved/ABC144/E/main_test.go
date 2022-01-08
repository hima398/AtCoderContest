package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n int
		k int
		a []int
		f []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Example 1", args{3, 5, []int{4, 2, 1}, []int{2, 3, 1}}, 2},
		{"Example 2", args{3, 8, []int{4, 2, 1}, []int{2, 3, 1}}, 0},
		{"Example 3", args{11, 14, []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.n, tt.args.k, tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestSolveHonestly(t *testing.T) {
	type args struct {
		n int
		k int
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{3, 4, []int{1, 2, 7}, []int{1, 5, 4}}, 2},
		{"Example 2", args{2, 123, []int{4, 678}, []int{5, 901}}, 1},
		{"Example 3", args{7, 10, []int{20, 20, 20, 20, 30, 30, 40}, []int{20, 20, 30, 40, 20, 30, 20}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveHonestly(tt.args.n, tt.args.k, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SolveHonestly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		n int
		k int
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Example 1", args{3, 4, []int{1, 2, 7}, []int{1, 5, 4}}, 2},
		{"Example 2", args{2, 123, []int{4, 678}, []int{5, 901}}, 1},
		{"Example 3", args{7, 10, []int{20, 20, 20, 20, 30, 30, 40}, []int{20, 20, 30, 40, 20, 30, 20}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.n, tt.args.k, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

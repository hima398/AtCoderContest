package main

import (
	"math"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		n int
		x []int
		y []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns float64
	}{
		{"Example 1", args{3, []int{0, 0, 10}, []int{0, 10, 10}}, 90},
		{"Example 2", args{5, []int{8, 9, 2, 1, 0}, []int{6, 1, 0, 0, 1}}, 171.869897645844},
		{"Example 3", args{10, []int{0, 1, 2, 2, 3, 5, 6, 7, 7, 8}, []int{0, 7, 6, 8, 5, 5, 7, 1, 9, 8}}, 180},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := Solve(tt.args.n, tt.args.x, tt.args.y); math.Abs(gotAns-tt.wantAns) > 1e-7 {
				t.Errorf("Solve() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

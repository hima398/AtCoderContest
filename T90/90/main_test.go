package main

import (
	"testing"
)

func TestSubTask2(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 4", args{28593, 1}, 365728740},
		{"Case 1", args{int(1e5) + 1, 1}, 122792102},
		{"Case 2", args{int(1e11), 1}, 481668480},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveHonestly(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("SubTask2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubTask4(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Example 2", args{17, 29}, 263173793},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubTask4(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("SubTask4() = %v, want %v", got, tt.want)
			}
		})
	}
}

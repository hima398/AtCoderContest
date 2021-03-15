package main

import (
	"testing"
)

func TestCombination(t *testing.T) {
	type args struct {
		N int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{11, 1}, 11},
		{"Example 2", args{11, 11}, 1},
		{"Example 3", args{12, 11}, 12},
		{"Example 4", args{16, 11}, 4368},
		{"Example 5", args{199, 11}, 366461620334848584},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combination(tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("Combination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntQueue_Pop(t *testing.T) {
	q := NewIntQueue()
	v, e := q.Pop()
	if e == nil {
		t.Errorf("Pop() = %v, want %v, err", v, -1)
	}
}

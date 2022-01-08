package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestSolve(t *testing.T) {
	type args struct {
		n int
		s string
	}
	type TestCase struct {
		name string
		args args
		want int
	}
	var tests []TestCase
	tests = append(tests, TestCase{"E1", args{4, "<><"}, 5})
	tests = append(tests, TestCase{"E2", args{5, "<<<<"}, 1})
	tests = append(tests, TestCase{"E3", args{20, ">>>><>>><>><>>><<>>"}, 217136290})
	rand.Seed(time.Now().UnixNano())
	sign := []string{"<", ">"}
	for i := 0; i < 100; i++ {
		n := 2 + rand.Intn(10)
		var s string
		for i := 0; i < n; i++ {
			s += sign[rand.Intn(1)]
		}
		tests = append(tests, TestCase{"T" + strconv.Itoa(i+1), args{n, s}, SolveHonestly(n, s)})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
				t.Errorf("n = %v, s = %v", tt.args.n, tt.args.s)
			}
		})
	}
}

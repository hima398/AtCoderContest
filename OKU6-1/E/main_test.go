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
		a []int
	}
	const nt = 1000
	type TestCase struct {
		name    string
		args    args
		wantAns int
	}
	tests := make([]TestCase, nt)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nt; i++ {
		//n := rand.Intn(3*int(1e5)) + 2
		// int(1e3)より大きすぎるとSolveHonestlyがパンクするのでnの制約を下げておく
		n := rand.Intn(3*int(1e3)) + 2
		as := make([]int, n)
		for j := 0; j < n; j++ {
			as[j] = rand.Intn(1 << 31)
		}
		tests[i] = TestCase{strconv.Itoa(i), args{n, as}, SolveHonestly(n, as)}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := Solve(tt.args.n, tt.args.a); gotAns != tt.wantAns {
				t.Errorf("n = %v, a = %v, Solve() = %v, want %v", tt.args.n, tt.args.a, gotAns, tt.wantAns)
			}
		})
	}
}

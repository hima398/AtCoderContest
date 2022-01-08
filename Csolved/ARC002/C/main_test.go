package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestFakeSolve(t *testing.T) {
	type args struct {
		n int
		c string
	}
	type TestCase struct {
		name string
		args args
		want int
	}
	var tests []TestCase

	rand.Seed(time.Now().UnixNano())
	button := []string{"A", "B", "X", "Y"}
	nb := len(button)
	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		var c string
		for j := 0; j < n; j++ {
			c += button[rand.Intn(nb-1)]
		}
		want := Solve(n, c)
		tests = append(tests, TestCase{strconv.Itoa(i), args{n, c}, want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveHonestly(tt.args.n, tt.args.c)
			if got != tt.want {
				t.Errorf("n = %d, c = %s, Solve() = %v, want %v", tt.args.n, tt.args.c, got, tt.want)
			}
		})
	}
}

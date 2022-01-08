package main

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestSolveHonestly(t *testing.T) {
	type args struct {
		n int
		q int
		a []int
		b []int
		c []int
		d []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{6, 3, []int{8, 6, 9, 1, 2, 1}, []int{0, 1, 2, 0, 1, 2}, []int{3, 1, 0}, []int{2, 0, 1}}, []int{6, 2, 6}},
		{"Example 2", args{2, 2, []int{4208, 3056}, []int{1233, 5677}, []int{0, 1}, []int{2019, 2019}}, []int{3056, 4208}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveHonestly(tt.args.n, tt.args.q, tt.args.a, tt.args.b, tt.args.c, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveHonestly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		n int
		q int
		a []int
		b []int
		c []int
		d []int
	}
	type Tests struct {
		name string
		args args
		want []int
	}
	tests := []Tests{
		{"Example 1", args{6, 3, []int{8, 6, 9, 1, 2, 1}, []int{0, 1, 2, 0, 1, 2}, []int{3, 1, 0}, []int{2, 0, 1}}, []int{6, 2, 6}},
		{"Example 2", args{2, 2, []int{4208, 3056}, []int{1233, 5677}, []int{0, 1}, []int{2019, 2019}}, []int{3056, 4208}},
	}
	nc := 10000 // 追加のテストケース
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nc; i++ {
		var tt Tests
		var args args
		args.n = rand.Intn(99) + 1
		args.q = rand.Intn(99) + 1
		//var a, b, c, d []int
		for i := 0; i < args.n; i++ {
			args.a = append(args.a, rand.Intn(int(1e9)))
			args.b = append(args.b, rand.Intn(2*int(1e5))-1)
		}
		for i := 0; i < args.q; i++ {
			if args.n == 1 {
				args.c = append(args.c, 0)
			} else {
				args.c = append(args.c, rand.Intn(args.n-1))
			}
			args.d = append(args.d, rand.Intn(2*int(1e5))-1)
		}

		tt.name = "Ext Example " + strconv.Itoa(i+1)
		tt.args = args
		tt.want = SolveHonestly(args.n, args.q, args.a, args.b, args.c, args.d)
		tests = append(tests, tt)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.n, tt.args.q, tt.args.a, tt.args.b, tt.args.c, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestCase = %v, Solve() = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

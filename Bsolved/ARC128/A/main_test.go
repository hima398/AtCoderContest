package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestSolve(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	type TestCase struct {
		name    string
		args    args
		wantAns [][]int
	}
	var tests []TestCase
	tests = append(tests, TestCase{"E1", args{3, []int{3, 5, 2}}, [][]int{{0, 1, 1}}})
	tests = append(tests, TestCase{"E2", args{4, []int{1, 1, 1, 1}}, [][]int{{0, 0, 0, 0}, {1, 1, 1, 1}}})

	//e3n := 10
	//e3a := []int{426877385, 186049196, 624834740, 836880476, 19698398, 709113743, 436942115, 436942115, 436942115, 503843678}
	//tests = append(tests, TestCase{"E3", args{e3n, e3a}, SolveHonestly(e3n, e3a)})

	t1n := 6
	t1a := []int{3, 7, 3, 9, 2, 2}
	tests = append(tests, TestCase{"T1", args{t1n, t1a}, SolveHonestly(t1n, t1a)})

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 0; i++ {
		n := 2 + rand.Intn(10)
		var a []int
		for j := 0; j < n; j++ {
			//a = append(a, 1+rand.Intn(int(1e9)-1))
			a = append(a, 1+rand.Intn(9))
		}
		want := SolveHonestly(n, a)
		tests = append(tests, TestCase{"T" + strconv.Itoa(i+1), args{n, a}, want})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := FirstSolve(tt.args.n, tt.args.a)
			ok := false
			for _, expected := range tt.wantAns {
				ok = ok || reflect.DeepEqual(gotAns, expected)
			}

			if !ok {
				t.Errorf("Solve() = %v, want %v", gotAns, tt.wantAns)
				fmt.Println(tt.args.n)
				fmt.Println(tt.args.a)
			}
		})
	}
}

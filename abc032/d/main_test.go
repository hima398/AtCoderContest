package main

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	type args struct {
		n  int
		w  int
		sv []int
		sw []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{3, 10, []int{15, 10, 6}, []int{9, 6, 4}}, 16},
		{"Example 2", args{30, 499887702,
			[]int{128990795, 575374246, 471048785, 640066776, 819841327, 704171581, 536108301, 119980848, 117241527, 325850062, 623319578, 998395208, 475707585, 863910036, 340559411, 122579234, 696368935, 665665204, 958833732, 371084424, 463433600, 210508742, 685281136, 619500108, 88215377, 558193168, 475268130, 303022740, 122379996, 304092766},
			[]int{137274936, 989051853, 85168425, 856699603, 611065509, 22345022, 678298936, 616908153, 28801762, 478675378, 706900574, 738510039, 135746508, 599020879, 738084616, 545330137, 86797589, 592749599, 401229830, 523386474, 5310725, 907821957, 565237085, 730556272, 310581512, 136966252, 132739489, 12425915, 137199296, 23505143}}, 3673016420},
		{"Example 3", args{10, 2921, []int{981421680, 515936168, 17309336, 788067075, 104855562, 494541604, 32007355, 772339969, 55112800, 98577050}, []int{325, 845, 371, 112, 96, 960, 161, 581, 248, 22}}, 3657162058},
		{"Example 4", args{10, 936447862, []int{854, 691, 294, 333, 832, 642, 139, 101, 853, 369}, []int{810169801, 957981784, 687140254, 932608409, 42367415, 727293784, 870916042, 685539955, 243593312, 977358410}}, 1686},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve1(tt.args.n, tt.args.w, tt.args.sv, tt.args.sw); got != tt.want {
				t.Errorf("Solve1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve2(t *testing.T) {
	type args struct {
		n  int
		w  int
		sv []int
		sw []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{3, 10, []int{15, 10, 6}, []int{9, 6, 4}}, 16},
		{"Example 3", args{10, 2921, []int{981421680, 515936168, 17309336, 788067075, 104855562, 494541604, 32007355, 772339969, 55112800, 98577050}, []int{325, 845, 371, 112, 96, 960, 161, 581, 248, 22}}, 3657162058},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve2(tt.args.n, tt.args.w, tt.args.sv, tt.args.sw); got != tt.want {
				t.Errorf("Solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve3(t *testing.T) {
	type args struct {
		n  int
		w  int
		sv []int
		sw []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{3, 10, []int{15, 10, 6}, []int{9, 6, 4}}, 16},
		{"Example 4", args{10, 936447862, []int{854, 691, 294, 333, 832, 642, 139, 101, 853, 369}, []int{810169801, 957981784, 687140254, 932608409, 42367415, 727293784, 870916042, 685539955, 243593312, 977358410}}, 1686},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve3(tt.args.n, tt.args.w, tt.args.sv, tt.args.sw); got != tt.want {
				t.Errorf("Solve3() = %v, want %v", got, tt.want)
			}
		})
	}
}

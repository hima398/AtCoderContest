package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, k int, a, f []int) int {
	sort.Ints(a)
	sort.Slice(f, func(i, j int) bool {
		return f[i] > f[j]
	})
	af := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		af[i] = a[i] * f[i]
		max = Max(max, af[i])
	}
	//チーム全体の成績を、k回の修行以内でxにできるかを判定する
	Check := func(x int) bool {
		//必要な修行の回数
		nk := 0
		for i := 0; i < n; i++ {
			diff := af[i] - x
			if diff > 0 {
				nk += Ceil(diff, f[i])
			}
		}
		return nk <= k
	}
	//fmt.Println(a)
	//fmt.Println(f)
	//fmt.Println(af)
	ng, ok := -1, max
	for ok-ng > 1 {
		mid := (ng + ok) / 2
		if Check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	f := nextIntSlice(n)

	//ans := SolveHonestly(n, k, a, f)
	ans := Solve(n, k, a, f)

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

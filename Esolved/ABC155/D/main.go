package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n, k int, a []int) int {
	var res []int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			v := a[i] * a[j]
			res = append(res, v)
		}
	}
	sort.Ints(res)
	return res[k-1]
}

func Solve(n, k int, a []int) int {
	const INF = int(1e18) + 1
	revA := make([]int, n)
	copy(revA, a)
	sort.Ints(a)
	sort.Slice(revA, func(i, j int) bool {
		return revA[i] > revA[j]
	})
	//fmt.Println(a)
	//fmt.Println(revA)
	ok, ng := -INF+1, INF
	Check := func(x int) bool {
		s := 0
		for i := 0; i < n; i++ {
			if a[i] >= 0 {
				idx := sort.Search(n, func(j int) bool {
					return a[i]*a[j] >= x
				})
				s += idx
			} else {
				idx := sort.Search(n, func(j int) bool {
					return a[i]*revA[j] >= x
				})
				s += idx
			}
			if a[i]*a[i] < x {
				s--
			}
		}
		s /= 2
		//fmt.Println("x = ", x, " | s = ", s)
		return s < k
	}
	for ng-ok > 1 {
		mid := (ok + ng) / 2
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

	//ans := SolveHonestly(n, k, a)
	ans := Solve(n, k, a)
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
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

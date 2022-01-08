package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(s string, k int) int {
	n := len(s)

	sc := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			sc[i+1] = sc[i] + 1
		} else {
			sc[i+1] = sc[i]
		}
	}

	r := 0
	ans := 0
	for l := 0; l < n; l++ {
		for r < n && sc[r+1]-sc[l] <= k {
			r++
		}
		ans = Max(ans, r-l)
	}
	return ans
}

func SolveBinarySearch(s string, k int) int {
	n := len(s)
	sc := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			sc[i+1] = sc[i] + 1
		} else {
			sc[i+1] = sc[i]
		}
	}
	Check := func(x int) bool {
		for r := x; r <= n; r++ {
			l := r - x
			if sc[r]-sc[l] <= k {
				return true
			}
		}
		return false
	}
	ok, ng := 0, n+1
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

	s := nextString()
	k := nextInt()
	//fmt.Println(SolveCommentary(s, k))
	fmt.Println(SolveBinarySearch(s, k))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

/*
func FirstSolve(a, b, x int) int {
	ans := 0
	for i := 1; i <= 10; i++ {
		n := (x - b*i) / a
		fmt.Println(i, n)
		if n < 0 {
			continue
		}
		if D(n) != i {
			continue
		}
		if n > maxn {
			ans = maxn
			continue
		}
		ans = Max(ans, n)
	}
	return ans
}
*/

func ComputePrice(a, b, n int) int {
	var d func(n int) int
	d = func(n int) int {
		s := strconv.Itoa(n)
		return len(s)
	}

	return a*n + b*d(n)
}

func SolveCommentary(a, b, x int) int {
	const maxn = int(1e9)
	if ComputePrice(a, b, 1) > x {
		return 0
	}
	if ComputePrice(a, b, maxn) < x {
		return maxn
	}
	ok, ng := 1, maxn
	for ok < ng-1 {
		n := (ok + ng) / 2

		if ComputePrice(a, b, n) > x {
			ng = n
		} else {
			ok = n
		}
	}

	return ok
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, x := nextInt(), nextInt(), nextInt()
	fmt.Println(SolveCommentary(a, b, x))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

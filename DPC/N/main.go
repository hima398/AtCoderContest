package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 1 << 60

	n := nextInt()
	a := nextIntSlice(n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i-1]
	}
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}
	var F func(l, r int) int
	F = func(l, r int) int {
		if memo[l][r] > 0 {
			return memo[l][r]
		}
		if l == r {
			return 0
		}
		cost := INF
		for i := l; i < r; i++ {
			cost = Min(cost, F(l, i)+F(i+1, r))
		}
		memo[l][r] = cost + s[r+1] - s[l]
		return memo[l][r]
	}
	ans := F(0, n-1)

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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

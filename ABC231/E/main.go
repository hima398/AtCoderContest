package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, x int, a []int) int {
	memo := make(map[int]int)
	memo[0] = 0
	var Dfs func(x, i int) int
	Dfs = func(x, i int) int {
		if _, ok := memo[x]; ok {
			return memo[x]
		}
		if i == n-1 {
			memo[x] = x / a[n-1]
			return memo[x]
		}
		r := x % a[i+1] / a[i]
		res := Dfs(x/a[i+1]*a[i+1], i+1) + r
		if r > 0 {
			res = Min(res, Dfs(Ceil(x, a[i+1])*a[i+1], i+1)+a[i+1]/a[i]-r)
		}
		memo[x] = res
		return memo[x]
	}
	ans := Dfs(x, 0)
	//	fmt.Println(memo)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := Solve(n, x, a)
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

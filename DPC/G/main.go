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

	n, m := nextInt(), nextInt()
	dp := make([]int, n+1)
	memo := make([]bool, n+1)
	route := make(map[int][]int)
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		route[x] = append(route[x], y)
	}
	var F func(x int) int
	F = func(x int) int {
		if memo[x] {
			return dp[x]
		}
		memo[x] = true
		ret := 0
		for _, y := range route[x] {
			ret = Max(ret, F(y)+1)
		}
		dp[x] = ret
		return dp[x]
	}
	//fmt.Println(dp)
	ans := 0
	for i := 1; i <= n; i++ {
		ans = Max(ans, F(i))
	}
	fmt.Println(ans)
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

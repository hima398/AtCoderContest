package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n int, x, y []int) int {
	const p = int(1e9) + 7
	//0-indexed
	for i := range x {
		x[i]--
		y[i]--
	}

	e := make([][]int, n)
	for i := range x {
		e[x[i]] = append(e[x[i]], y[i])
		e[y[i]] = append(e[y[i]], x[i])
	}
	type Node struct {
		b, w int
	}

	//根から葉iを白と黒に塗り分ける組み合わせの数
	dp := make([]Node, n)
	var Dfs func(i, pre int)
	Dfs = func(i, pre int) {
		dp[i].b = 1
		dp[i].w = 1
		if len(e[i]) == 1 && e[i][0] == pre {
			return
		}

		for _, next := range e[i] {
			if next == pre {
				continue
			}
			Dfs(next, i)
			dp[i].w *= (dp[next].b + dp[next].w)
			dp[i].w %= p
			dp[i].b *= dp[next].w
			dp[i].b %= p
		}
	}
	Dfs(0, -1)
	//fmt.Println(dp)
	return (dp[0].b + dp[0].w) % p
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x, y := make([]int, n-1), make([]int, n-1)
	for i := range x {
		x[i], y[i] = nextInt(), nextInt()
	}
	ans := Solve(n, x, y)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

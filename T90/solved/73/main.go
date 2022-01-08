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

	const p = int(1e9) + 7

	n := nextInt()
	c := make([]string, n)
	for i := 0; i < n; i++ {
		c[i] = nextString()
	}
	e := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	// 01:aのみ、10:bのみ、11:aとb、00:使用しない
	dp := make([][4]int, n)
	var Dfs func(i, par int)
	Dfs = func(i, par int) {
		dpa, dpb := 1, 1
		for _, chi := range e[i] {
			if chi == par {
				continue
			}
			Dfs(chi, i)
			if c[i] == "a" {
				dpa *= dp[chi][1] + dp[chi][3]
			} else if c[i] == "b" {
				dpa *= dp[chi][2] + dp[chi][3]
			}
			dpb *= dp[chi][1] + dp[chi][2] + 2*dp[chi][3]

			dpa %= p
			dpb %= p
		}
		if c[i] == "a" {
			dp[i][1] = dpa
		} else if c[i] == "b" {
			dp[i][2] = dpa
		}
		dp[i][3] = dpb - dpa
		if dp[i][3] < 0 {
			dp[i][3] += p
		}
		dp[i][3] %= p
	}
	Dfs(0, -1)
	//fmt.Println(dp)
	fmt.Println(dp[0][3])
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

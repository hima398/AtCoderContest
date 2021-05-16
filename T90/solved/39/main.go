package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Edge struct {
	a, b int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	e := make(map[int][]int)
	edge := make([]Edge, n-1)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
		edge[i] = Edge{a, b}
	}
	dp := make([]int, n)
	visited := make([]bool, n)
	var dfs func(x int)
	dfs = func(x int) {
		visited[x] = true
		dp[x] = 1 //親からxへの経路
		for _, v := range e[x] {
			if !visited[v] {
				dfs(v)
				dp[x] += dp[v]
			}
		}
	}
	dfs(0)
	dp[0]--
	//fmt.Println(dp)
	ans := 0
	for i := 0; i < n-1; i++ {
		c := Min(dp[edge[i].a], dp[edge[i].b])
		ans += c * (n - c)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

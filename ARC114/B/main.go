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

	const p = 998244353
	n := nextInt()
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = nextInt()
	}
	visited := make([]bool, n+1)
	var dfs func(s, t int) int
	dfs = func(s, t int) int {
		visited[t] = true
		if !visited[f[t]] {
			return dfs(s, f[t])
		} else if f[t] == s {
			return 1
		} else {
			return 0
		}
	}
	c := 0
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}
		c += dfs(i, i)
	}
	//fmt.Println(len(m))
	ans := 1
	for i := 1; i <= c; i++ {
		ans *= 2
		ans %= p
	}
	fmt.Println(ans - 1)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

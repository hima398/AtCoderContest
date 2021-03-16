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
	groups := make([]int, n+1)
	var dfs func(s, t int) int
	dfs = func(s, t int) int {
		groups[t] = s
		if groups[f[t]] == 0 {
			return dfs(s, f[t])
		} else if groups[f[t]] == groups[s] {
			return 1
		} else {
			return 0
		}
	}
	c := 0
	for i := 1; i <= n; i++ {
		if groups[i] > 0 {
			continue
		}
		c += dfs(i, i)
	}
	//fmt.Println(c)
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

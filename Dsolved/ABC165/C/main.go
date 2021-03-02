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

	n, m, q := nextInt(), nextInt(), nextInt()
	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	d := make([]int, q)
	for i := 0; i < q; i++ {
		a[i], b[i], c[i], d[i] = nextInt()-1, nextInt()-1, nextInt(), nextInt()
	}

	ans := 0
	var dfs func(la []int)
	dfs = func(la []int) {
		if len(la) == n {
			score := 0
			for i := 0; i < q; i++ {
				if la[b[i]]-la[a[i]] == c[i] {
					score += d[i]
				}
			}
			ans = Max(ans, score)
			return
		}
		for i := la[len(la)-1]; i <= m; i++ {
			dfs(append(la, i))
		}
	}
	for i := 1; i <= m; i++ {
		dfs([]int{i})
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

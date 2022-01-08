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

	n, q := nextInt(), nextInt()
	a, b := make([]int, n-1), make([]int, n-1)
	c, d := make([]int, q), make([]int, q)
	for i := 0; i < n-1; i++ {
		a[i], b[i] = nextInt(), nextInt()
		a[i]--
		b[i]--
	}
	for i := 0; i < q; i++ {
		c[i], d[i] = nextInt(), nextInt()
		c[i]--
		d[i]--
	}
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	depth := make([]int, n)
	var Dfs func(i, pre, d int)
	Dfs = func(i, pre, d int) {
		//fmt.Println(i, pre)
		depth[i] = d
		for _, next := range e[i] {
			if next == pre {
				continue
			}
			Dfs(next, i, d+1)
		}
	}
	Dfs(0, -1, 0)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		if depth[c[i]]%2 == depth[d[i]]%2 {
			fmt.Fprintln(out, "Town")
		} else {
			fmt.Fprintln(out, "Road")
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

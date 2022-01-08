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
	e := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	gs := make([]int, n)
	var Dfs func(i, g int)
	Dfs = func(i, g int) {
		gs[i] = g
		for _, v := range e[i] {
			if gs[v] == 0 {
				Dfs(v, g)
			}
		}
	}
	group := 1
	for i := 0; i < n; i++ {
		if gs[i] > 0 {
			continue
		}
		Dfs(i, group)
		group++
	}
	mg := 0
	for _, g := range gs {
		mg = Max(mg, g)
	}
	ans := mg - 1
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Edge struct {
	t, d int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	e := make(map[int][]Edge)
	for i := 0; i < n-1; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], Edge{b, c})
		e[b] = append(e[b], Edge{a, c})
	}
	q, k := nextInt(), nextInt()
	k--
	ds := make([]int, n)
	var Dfs func(i, pre, d int)
	Dfs = func(i, pre, d int) {
		ds[i] = d
		for _, edge := range e[i] {
			if edge.t == pre {
				continue
			}
			Dfs(edge.t, i, d+edge.d)
		}
	}
	Dfs(k, -1, 0)
	//fmt.Println(ds)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		x, y := nextInt(), nextInt()
		x--
		y--
		fmt.Println(ds[x] + ds[y])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Edge struct {
	t, w int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	e := make(map[int][]Edge)
	for i := 0; i < n; i++ {
		u, v, w := nextInt(), nextInt(), nextInt()
		u--
		v--
		e[u] = append(e[u], Edge{v, w})
		e[v] = append(e[v], Edge{u, w})
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	ans[0] = 0
	var Dfs func(t, w, p int)
	Dfs = func(t, w, p int) {
		if w%2 == 0 {
			ans[t] = ans[p]
		} else {
			ans[t] = ans[p] ^ 1
		}
		for _, edge := range e[t] {
			if ans[edge.t] < 0 {
				Dfs(edge.t, edge.w, t)
			}
		}
	}
	for _, edge := range e[0] {
		Dfs(edge.t, edge.w, 0)
	}
	//fmt.Println(ans)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

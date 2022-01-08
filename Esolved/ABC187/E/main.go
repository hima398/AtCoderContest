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

	n := nextInt()
	a := make([]int, n-1)
	b := make([]int, n-1)
	edges := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		ai, bi := nextInt(), nextInt()
		ai--
		bi--
		a[i] = ai
		b[i] = bi
		edges[ai] = append(edges[ai], bi)
		edges[bi] = append(edges[bi], ai)
	}
	ds := make([]int, n)
	for i := range ds {
		ds[i] = -1
	}
	var Dfs func(i, d int)
	Dfs = func(i, d int) {
		ds[i] = d
		for _, v := range edges[i] {
			if ds[v] == -1 {
				Dfs(v, d+1)
			}
		}
	}
	Dfs(0, 0)
	//fmt.Println(ds)
	q := nextInt()
	ans := make([]int, n)
	for k := 0; k < q; k++ {
		t, e, x := nextInt(), nextInt(), nextInt()
		e--
		if t == 1 {
			if ds[a[e]] > ds[b[e]] {
				ans[a[e]] += x
			} else {
				ans[0] += x
				ans[b[e]] -= x
			}
		} else {
			// t == 2
			if ds[a[e]] > ds[b[e]] {
				ans[0] += x
				ans[a[e]] -= x
			} else {
				ans[b[e]] += x
			}
		}
	}
	var Dfs2 func(i, pre int)
	Dfs2 = func(i, pre int) {
		ans[i] += ans[pre]
		for _, v := range edges[i] {
			if v != pre {
				Dfs2(v, i)
			}
		}
	}
	//fmt.Println(ans)
	for _, v := range edges[0] {
		Dfs2(v, 0)
	}
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

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pair struct {
	a, b int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	e := make(map[int][]int)
	er := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		er[b] = append(er[b], a)
	}
	/*
		for i := 0; i < n; i++ {
			fmt.Println(e[i])
		}
	*/
	d := 1
	ds := make([]Pair, n)
	visited := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		for _, v := range e[i] {
			if !visited[v] {
				dfs(v)
			}
		}
		ds[i] = Pair{i, d}
		d++
	}
	for i := 0; i < n; i++ {
		if ds[i].b == 0 {
			dfs(i)
		}
	}

	sort.Slice(ds, func(i, j int) bool { return ds[i].b > ds[j].b })
	//fmt.Println(ds)
	groups := make([]int, n)
	for i := range visited {
		visited[i] = false
	}
	var dfs2 func(i, g int)
	dfs2 = func(i, g int) {
		visited[i] = true
		for _, v := range er[i] {
			if !visited[v] {
				dfs2(v, g)
			}
		}
		groups[i] = g
	}
	for _, v := range ds {
		if groups[v.a] == 0 {
			dfs2(v.a, v.a+1)
		}
	}
	mg := make(map[int]int)
	for _, v := range groups {
		mg[v]++
	}
	ans := 0
	for _, v := range mg {
		ans += v * (v - 1) / 2
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

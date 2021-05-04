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
	e := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	dist := make([]int, n)
	v1 := make([]bool, n)
	var dfs1 func(i, d int)
	dfs1 = func(i, d int) {
		v1[i] = true
		dist[i] = d
		for _, node := range e[i] {
			if !v1[node] {
				dfs1(node, d+1)
			}
		}
	}
	dfs1(0, 0)
	s := 0
	deepest := 0
	for i, v := range dist {
		if deepest < v {
			s = i
			deepest = v
		}
	}

	//fmt.Println(e)
	v2 := make([]bool, n)
	var dfs2 func(i, d int) int
	dfs2 = func(i, d int) int {
		v2[i] = true
		ret := 0
		for _, node := range e[i] {
			if !v2[node] {
				ret = Max(ret, dfs2(node, d+1))
			}
		}
		if ret == 0 {
			return d
		} else {
			return ret
		}
	}

	ans := dfs2(s, 0) + 1
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

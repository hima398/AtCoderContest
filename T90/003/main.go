package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		//e[b] = append(e[b], a)
	}
	//fmt.Println(e)
	visited := make([]bool, n)
	var dfs func(i, d int) int
	dfs = func(i, d int) int {
		visited[i] = true
		/*
			allVisted := true
			for _, node := range e[i] {
				allVisted = allVisted && visited[node]
			}
			if allVisted {
				return d
			}
		*/
		if len(e[i]) == 0 {
			return d
		}
		ret := 0
		for _, node := range e[i] {
			if !visited[node] {
				ret = Max(ret, dfs(node, d+1))
			}
		}
		return ret
	}
	var ans int
	if len(e[0]) == 1 {
		ans = dfs(0, 0) + 1
	} else {
		var d []int
		visited[0] = true
		for _, node := range e[0] {
			d = append(d, dfs(node, 1))
		}
		sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })
		ans = d[0] + d[1] + 1
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

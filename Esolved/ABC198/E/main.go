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
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = nextInt()
	}
	e := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	var cm [100001]int
	//m[0][c[0]] = true
	visited := make([]bool, n)
	var ans []int
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		cm[c[i]]++
		for _, v := range e[i] {
			if !visited[v] {
				dfs(v)
			}
		}
		if cm[c[i]] <= 1 {
			ans = append(ans, i)
		}
		cm[c[i]]--
	}
	dfs(0)

	sort.Ints(ans)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v+1)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

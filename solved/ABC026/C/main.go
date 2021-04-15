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
	m := make([]int, n+1)
	for i := 2; i <= n; i++ {
		b := nextInt()
		e[b] = append(e[b], i)
	}
	//fmt.Println(e)
	var dfs func(i int) int
	dfs = func(i int) int {
		if len(e[i]) == 0 {
			m[i] = 1
			return m[i]
		}
		if m[i] > 0 {
			return m[i]
		}
		max := 0
		min := 1 << 60
		for _, v := range e[i] {
			max = Max(max, dfs(v))
			min = Min(min, dfs(v))
		}
		m[i] = max + min + 1
		return m[i]
	}
	ans := dfs(1)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

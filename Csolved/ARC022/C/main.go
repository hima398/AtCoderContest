package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n int, a, b []int) (int, int) {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		// 0-indexed
		a[i]--
		b[i]--
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	//一番深い深さとその番号
	md, di := 0, -1
	var Dfs func(i, pre, d int)
	Dfs = func(i, pre, d int) {
		if md < d {
			md = d
			di = i
		}
		for _, next := range e[i] {
			if next == pre {
				continue
			}
			Dfs(next, i, d+1)
		}
	}
	Dfs(0, -1, 0)
	ii := di
	md, di = 0, -1
	Dfs(ii, -1, 0)
	//解は1-indexedなので1を足して返す
	return ii + 1, di + 1
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a, b := make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	ansS, ansT := Solve(n, a, b)
	fmt.Println(ansS, ansT)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

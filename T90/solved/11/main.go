package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Task struct {
	d, c, s int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ts := make([]Task, n)
	m := make(map[int][]int)
	dm := 0
	for i := 0; i < n; i++ {
		d, c, s := nextInt(), nextInt(), nextInt()
		t := Task{d, c, s}
		ts[i] = t
		m[c] = append(m[d], s)
		dm = Max(dm, d)
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].d < ts[j].d
	})
	//fmt.Println(ts)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, dm+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= dm; j++ {
			if j >= ts[i].c && j <= ts[i].d {
				//fmt.Println(i, j, ts[i].c)
				dp[i+1][j] = Max(dp[i][j], dp[i][j-ts[i].c]+ts[i].s)
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	ans := 0
	for _, v := range dp[n] {
		ans = Max(ans, v)
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

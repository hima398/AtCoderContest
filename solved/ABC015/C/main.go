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
	n, k := nextInt(), nextInt()
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, k)
		for j := 0; j < k; j++ {
			t[i][j] = nextInt()
		}
	}
	var ans []int
	var dfs func(d, j, s int)
	dfs = func(d, j, s int) {
		ss := s ^ t[d][j]
		if d == n-1 {
			ans = append(ans, ss)
			return
		} else {
			for jj := 0; jj < k; jj++ {
				dfs(d+1, jj, ss)
			}
		}
		return
	}
	for j := 0; j < k; j++ {
		dfs(0, j, 0)
	}
	sort.Ints(ans)
	if ans[0] == 0 {
		fmt.Println("Found")
	} else {
		fmt.Println("Nothing")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

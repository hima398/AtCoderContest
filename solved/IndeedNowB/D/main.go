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

	n, c := nextInt(), nextInt()
	a := nextIntSlice(n)
	mc := make([][]int, c+1)
	for i := range mc {
		mc[i] = append(mc[i], -1)
	}
	for i := range a {
		mc[a[i]] = append(mc[a[i]], i)
	}
	for i := range mc {
		mc[i] = append(mc[i], n)
	}
	//fmt.Println(mc)
	var ans []int
	s := n * (n + 1) / 2
	for i := 1; i <= c; i++ {
		ss := 0
		for j := 1; j < len(mc[i]); j++ {
			l := mc[i][j] - mc[i][j-1] - 1
			ss += l * (l + 1) / 2
		}
		ans = append(ans, s-ss)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < c; i++ {
		fmt.Fprintln(out, ans[i])
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

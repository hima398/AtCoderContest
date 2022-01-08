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
	n, m := nextInt(), nextInt()
	a := nextIntSlice(m)
	pushed := make(map[int]bool)
	var ans []int
	for i := m - 1; i >= 0; i-- {
		if !pushed[a[i]] {
			ans = append(ans, a[i])
			pushed[a[i]] = true
		}
	}
	for i := 1; i <= n; i++ {
		if !pushed[i] {
			ans = append(ans, i)
			pushed[i] = true
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
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

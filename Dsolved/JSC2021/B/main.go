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
	a := nextIntSlice(n)
	b := nextIntSlice(m)
	h := make([]int, 1001)
	for i := 0; i < n; i++ {
		h[a[i]]++
	}
	for i := 0; i < m; i++ {
		h[b[i]]++
	}
	var ans []int
	for i := 1; i < len(h); i++ {
		if h[i] == 1 {
			ans = append(ans, i)
		}
	}
	if len(ans) == 0 {
		return
	}
	fmt.Printf("%d", ans[0])
	for i := 1; i < len(ans); i++ {
		fmt.Printf(" %d", ans[i])
	}
	fmt.Println("")
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

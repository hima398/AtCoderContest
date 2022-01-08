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

	n, m, c := nextInt(), nextInt(), nextInt()
	b := nextIntSlice(m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(m)
	}
	ans := 0
	for i := 0; i < n; i++ {
		s := 0
		for j := 0; j < m; j++ {
			s += a[i][j] * b[j]
		}
		if s+c > 0 {
			ans++
		}
	}
	fmt.Println(ans)
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

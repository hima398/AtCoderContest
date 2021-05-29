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

	const p = int(1e9) + 7

	n := nextInt()
	a := make([][]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(6)
		for j := 0; j < 6; j++ {
			s[i] += a[i][j]
		}
	}
	ans := 1
	for i := 0; i < n; i++ {
		ans *= s[i]
		ans %= p
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

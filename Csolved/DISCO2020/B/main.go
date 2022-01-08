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

	const INF = 1 << 62

	n := nextInt()
	a := nextIntSlice(n)
	s := make([]int, n)
	s[0] = a[0]
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + a[i]
	}

	ans := INF
	for i := 0; i < n-1; i++ {
		l, r := s[i], s[n-1]-s[i]
		ans = Min(ans, Abs(r-l))
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

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
	var a [2][]int
	var dp [2][]int
	for i := 0; i < 2; i++ {
		a[i] = nextIntSlice(n)
		dp[i] = make([]int, n)
	}
	dp[0][0] = a[0][0]
	dp[1][0] = dp[0][0] + a[1][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + a[0][i]
	}
	for i := 1; i < n; i++ {
		dp[1][i] = Max(dp[0][i], dp[1][i-1]) + a[1][i]
	}
	fmt.Println(dp[1][n-1])
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

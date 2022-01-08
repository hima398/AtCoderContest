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
	r := nextIntSlice(n)

	dp := make([][2]int, n)

	dp[0][0] = 1
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if r[j] < r[i] {
				dp[i][0] = Max(dp[i][0], dp[j][1]+1)
			}
			if r[j] > r[i] {
				dp[i][1] = Max(dp[i][1], dp[j][0]+1)
			}
		}
	}
	ans := Max(dp[n-1][0], dp[n-1][1])
	if ans < 3 {
		fmt.Println(0)
		return
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

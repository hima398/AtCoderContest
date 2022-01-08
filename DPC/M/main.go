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
	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	//i人の子供にj個飴を分け合う方法の数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		s := make([]int, k+2)
		for j := 0; j <= k; j++ {
			s[j+1] = s[j] + dp[i][j]
			s[j+1] %= p
		}
		for j := 0; j <= k; j++ {
			dp[i+1][j] += s[j+1] - s[Max(0, j-a[i])] + p
			dp[i+1][j] %= p
		}
	}
	fmt.Println(dp[n][k])
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

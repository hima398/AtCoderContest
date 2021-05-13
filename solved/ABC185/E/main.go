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

	const INF = 3*int(1e3) + 1
	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(m)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	//fmt.Println(dp)
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			//xに関する計算
			if i > 0 {
				// a[i]を削除するスコアか、すでに求まったスコアのうち小さい方
				dp[i][j] = Min(dp[i][j], dp[i-1][j]+1)
			}
			if j > 0 {
				// b[j]を削除するスコアか、すでに求まったスコアのうち小さい方
				dp[i][j] = Min(dp[i][j], dp[i][j-1]+1)
			}

			//yに関する計算
			if i > 0 && j > 0 {
				if a[i-1] == b[j-1] {
					dp[i][j] = Min(dp[i][j], dp[i-1][j-1])
				} else {
					dp[i][j] = Min(dp[i][j], dp[i-1][j-1]+1)
				}
			}
		}
	}

	fmt.Println(dp[n][m])
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

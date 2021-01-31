package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintDP(dp [][]int) {
	for _, dpi := range dp {
		for _, v := range dpi {
			fmt.Printf("%d ", v)
		}
		fmt.Println("")
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, w := nextInt(), nextInt()
	wi, vi := make([]int, n+1), make([]int, n+1)
	sv := 0 // sum value
	for i := 1; i <= n; i++ {
		wi[i], vi[i] = nextInt(), nextInt()
		sv += vi[i]
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sv+1)
		for j := 0; j <= sv; j++ {
			dp[i][j] = 9000000000000000000
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= sv; j++ {
			if j-vi[i] >= 0 {
				//fmt.Println(i, j, wi[i])
				dp[i][j] = Min(dp[i-1][j], dp[i-1][j-vi[i]]+wi[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	//ans := 0
	for i := sv; i >= 0; i-- {
		if dp[n][i] <= w {
			fmt.Println(i)
			return
			//ans = Max(ans, i)
		}
	}
	//fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

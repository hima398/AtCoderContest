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
	for i := 1; i <= n; i++ {
		wi[i], vi[i] = nextInt(), nextInt()
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			if j-wi[i] >= 0 {
				dp[i][j] = Max(dp[i-1][j], dp[i-1][j-wi[i]]+vi[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	//PrintDP(dp)
	fmt.Println(dp[n][w])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

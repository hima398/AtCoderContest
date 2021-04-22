package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Screen struct {
	w, v int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	w := nextInt()
	n, k := nextInt(), nextInt()
	s := make([]Screen, n)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		s[i] = Screen{a, b}
	}

	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, w+1)
		}
	}
	dp[0][0][0] = 0

	for i := 0; i < n; i++ {
		for kk := 0; kk < k; kk++ {
			for j := 0; j <= w; j++ {
				if j+s[i].w <= w && kk < k {
					dp[i+1][kk+1][j+s[i].w] = Max(dp[i+1][kk+1][j+s[i].w], dp[i][kk][j]+s[i].v)
				}
				dp[i+1][kk][j] = Max(dp[i+1][kk][j], dp[i][kk][j])
			}
		}
	}
	ans := 0
	for i := 0; i <= n; i++ {
		for kk := 0; kk <= k; kk++ {
			for j := 0; j <= w; j++ {
				ans = Max(ans, dp[i][kk][j])
			}
		}
	}
	//fmt.Println(dp[5][4])
	fmt.Println(ans)
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

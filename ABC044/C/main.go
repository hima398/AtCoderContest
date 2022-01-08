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

	n, a := nextInt(), nextInt()
	x := nextIntSlice(n)
	s := 0
	for _, xi := range x {
		s += xi
	}
	//1〜i番目のカードまでの中からj枚選んで合計がkになるカードの選び方
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, s+1)
		}
	}
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k <= s; k++ {
				dp[i+1][j][k] += dp[i][j][k]
				if k+x[i] > s {
					continue
				}
				dp[i+1][j+1][k+x[i]] += dp[i][j][k]
			}
		}
	}
	ans := 0
	for j := 1; j <= n; j++ {
		if a*j > s {
			continue
		}
		ans += dp[n][j][a*j]
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

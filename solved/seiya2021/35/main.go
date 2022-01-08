package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7

	s := nextString()
	n := len(s)
	t := "chokudai"
	nt := len(t)
	dp := make([][9]int, n+1)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= nt; j++ {
			if j < nt && s[i] == t[j] {
				dp[i+1][j+1] += dp[i][j]
			}
			dp[i+1][j] += dp[i][j]
		}
		for j := 0; j <= nt; j++ {
			dp[i+1][j] %= p
		}
	}
	fmt.Println(dp[n][nt])
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

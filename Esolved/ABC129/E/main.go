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

	l := nextString()
	n := len(l)
	//MSBからi桁目まで確定して、i+1桁以降どうなってもa+b<lになれば1の組み合わせ
	dp := make([][2]int, n+1)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			if l[i] == '1' {
				dp[i+1][1] += dp[i][j]
			} else {
				dp[i+1][j] += dp[i][j]
			}
			dp[i+1][j] %= p
			if l[i] == '1' || j == 1 {
				dp[i+1][j] += dp[i][j] * 2
				dp[i+1][j] %= p
			}
		}
	}
	ans := (dp[n][0] + dp[n][1]) % p
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

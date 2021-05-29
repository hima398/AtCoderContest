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

	const mod = 13
	const p = int(1e9) + 7
	s := nextString()
	n := len(s)
	dp := make([][mod]int, n+1)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < mod; j++ {
			if s[i] == '?' {
				for k := 0; k < 10; k++ {
					nj := (10*j + k) % mod
					dp[i+1][nj] += dp[i][j]
					dp[i+1][nj] %= p
				}
			} else {
				si := int(s[i] - '0')
				nj := (10*j + si) % mod
				dp[i+1][nj] += dp[i][j]
				dp[i+1][nj] %= p
			}
		}
	}
	ans := dp[n][5]
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

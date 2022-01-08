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
	t := nextString()
	n := len(t)

	//j文字目までみた時に、ABCのうちi文字作れる組み合わせの数
	var dp [4][]int
	for i := 0; i <= 3; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for j := 1; j <= n; j++ {
		for i := 0; i <= 3; i++ {
			dp[i][j] = dp[i][j-1]
		}
		switch t[j-1] {
		case 'A':
			dp[1][j] += dp[0][j-1]
			dp[1][j] %= p
		case 'B':
			dp[2][j] += dp[1][j-1]
			dp[2][j] %= p
		case 'C':
			dp[3][j] += dp[2][j-1]
			dp[3][j] %= p
		case '?':
			for i := 0; i <= 3; i++ {
				dp[i][j] *= 3
				dp[i][j] %= p
			}
			for i := 1; i <= 3; i++ {
				dp[i][j] += dp[i-1][j-1]
				dp[i][j] %= p
			}
		}
	}
	//デバッグ出力
	//for i := 0; i <= 3; i++ {
	//	fmt.Println(dp[i])
	//}
	fmt.Println(dp[3][n])
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

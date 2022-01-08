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

	n := nextInt()
	s := nextString()
	ans := 101
	for k := 0; k < n; k++ {
		f, r := s[:k], s[k:]
		dp := make([][]int, len(f)+1)
		for i := 0; i <= len(f); i++ {
			dp[i] = make([]int, len(r)+1)
		}
		//LCS
		for i := 1; i <= len(f); i++ {
			for j := 1; j <= len(r); j++ {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
				if f[i-1] == r[j-1] {
					dp[i][j] = Max(dp[i][j], dp[i-1][j-1]+1)
				}
			}
		}
		// 文字列全体から、前後に区切った文字列の共通文字列
		// それぞれ引いたものが消す文字
		ans = Min(ans, n-2*dp[len(f)][len(r)])
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

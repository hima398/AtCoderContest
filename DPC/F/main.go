package main

import (
	"bufio"
	"fmt"
	"os"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	s, t := nextString(), nextString()
	// sのi文字目、tのj文字目まで見たとき最長共通部分の長さ
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	//fmt.Println(dp[len(s)][len(t)])
	n := dp[len(s)][len(t)]
	i, j := len(s), len(t)
	var ans string
	for n > 0 {
		if s[i-1] == t[j-1] {
			ans = string(s[i-1]) + ans
			i--
			j--
			n--
		} else if dp[i][j] == dp[i-1][j] {
			i--
		} else {
			j--
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

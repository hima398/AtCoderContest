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

	const p = int(1e9) + 7
	s := nextString()
	n := len(s)

	target := "chokudai"
	nt := len(target)
	dp := make([][]int, nt+1)
	for i := 0; i <= nt; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= nt; j++ {
			if j < nt && s[i] == target[j] {
				dp[j+1][i+1] += dp[j][i]
				//dp[j+1][i+1] %= p
			}
			dp[j][i+1] += dp[j][i]
		}
		for j := 0; j <= nt; j++ {
			dp[j][i+1] %= p
		}
	}
	/*
		for i := 0; i <= nt; i++ {
			fmt.Println(dp[i])
		}
	*/
	fmt.Println(dp[nt][n])
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

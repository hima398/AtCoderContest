package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(n int, s, x string) bool {
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, 7)
	}
	dp[n][0] = true
	for i := n; i > 0; i-- {
		for j := 0; j < 7; j++ {
			if x[i-1] == 'T' {
				dp[i-1][j] = dp[i][10*j%7] || dp[i][(10*j+int(s[i-1]-'0'))%7]
			} else {
				dp[i-1][j] = dp[i][10*j%7] && dp[i][(10*j+int(s[i-1]-'0'))%7]
			}
		}
		//fmt.Println(dp[i-1])
	}
	return dp[0][0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	x := nextString()
	if SolveCommentary(n, s, x) {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
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

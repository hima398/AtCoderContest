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

	//入力
	n := nextInt()
	s1 := nextString()
	s2 := nextString()

	dp := make([]int, n)
	if s1[0] == s2[0] {
		dp[0] = 3
	} else {
		dp[0] = 6
	}
	//
	for i := 1; i < n; i++ {

		if s1[i] == s2[i] { //i番目(0-indexed)のドミノが縦向き
			if s1[i-1] == s2[i-1] {
				dp[i] = dp[i-1] * 2
				dp[i] %= p
			} else {
				dp[i] = dp[i-1]
			}
		} else { // i番目のドミノが横向き
			// i-1番目のドミノが縦向き
			if s1[i-1] == s2[i-1] {
				dp[i] = dp[i-1] * 2
				dp[i] %= p
			} else if s1[i] == s1[i-1] && s2[i] == s2[i-1] {
				dp[i] = dp[i-1]
			} else {
				dp[i] = dp[i-1] * 3
				dp[i] %= p
			}
		}
	}
	fmt.Println(dp[n-1])
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

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

	k := nextInt()
	if k%9 != 0 {
		fmt.Println(0)
		return
	}
	const p = int(1e9) + 7
	dp := make([]int, k+1)
	dp[0] = 1
	for i := 1; i <= k; i++ {
		min := Min(i, 9)
		for j := 1; j <= min; j++ {
			dp[i] += dp[i-j]
		}
		dp[i] %= p
	}
	fmt.Println(dp[k])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

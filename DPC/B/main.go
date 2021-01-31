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

	n, k := nextInt(), nextInt()
	h := make([]int, n+1)
	for i := 1; i <= n; i++ {
		h[i] = nextInt()
	}
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 10001
	}
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + Abs(h[i-1]-h[i])
		for j := 2; j <= k; j++ {
			if i-j >= 1 {
				dp[i] = Min(dp[i], dp[i-j]+Abs(h[i]-h[i-j]))
			}
		}
	}
	//fmt.Println(dp)
	fmt.Println(dp[n])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

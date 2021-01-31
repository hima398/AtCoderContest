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
	h := make([]int, n+1)
	for i := 1; i <= n; i++ {
		h[i] = nextInt()
	}
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = Abs(h[2] - h[1])
	for i := 3; i <= n; i++ {
		dp[i] = Min(dp[i-1]+Abs(h[i]-h[i-1]), dp[i-2]+Abs(h[i]-h[i-2]))
	}
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

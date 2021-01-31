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
	a := make([]int, n+1)
	b := make([]int, n+1)
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i], b[i], c[i] = nextInt(), nextInt(), nextInt()
	}
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][1], dp[1][1], dp[2][1] = a[1], b[1], c[1]
	for i := 2; i <= n; i++ {
		dp[0][i] = Max(dp[1][i-1]+a[i], dp[2][i-1]+a[i])
		dp[1][i] = Max(dp[0][i-1]+b[i], dp[2][i-1]+b[i])
		dp[2][i] = Max(dp[0][i-1]+c[i], dp[1][i-1]+c[i])
	}
	ans := Max(dp[0][n], Max(dp[1][n], dp[2][n]))
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

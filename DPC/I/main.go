package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	p := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		p[i] = nextFloat64()
	}
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
	}
	dp[1][0] = 1.0 - p[1]
	dp[1][1] = p[1]
	for i := 2; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][0] = dp[i-1][0] * (1.0 - p[i])
			} else if j == i {
				dp[i][i] = dp[i-1][i-1] * p[i]
			} else {
				dp[i][j] = dp[i-1][j]*(1.0-p[i]) + dp[i-1][j-1]*p[i]
			}
		}
	}
	ans := 0.0
	for j := Ceil(n, 2); j <= n; j++ {
		ans += dp[n][j]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f

}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

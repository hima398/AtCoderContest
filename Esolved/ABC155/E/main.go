package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(s string) int {
	n := len(s)
	//sの上位からi桁目見た額のお金を支払う最小の枚数
	//j=0、ちょうど払える、j=1、sより多く払っている
	dp := make([][2]int, n+1)
	dp[0][1] = 1
	for i := 0; i < n; i++ {
		v := int(s[i] - '0')
		dp[i+1][0] = Min(dp[i][0]+v, dp[i][1]+10-v)
		dp[i+1][1] = Min(dp[i][0]+v+1, dp[i][1]+10-v-1)
		//fmt.Println(dp[i+1])
	}
	return dp[n][0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextString()

	ans := Solve(n)
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(k string, d int) (ans int) {
	const p = int(1e9) + 7
	n := len(k)
	//k[1]~k[i](1<=i<=n)で表される数字k'の各桁の数字ので
	//i-1番目の数値がj、総和をdで割った余りがkである個数
	dp := make([][2][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i][0] = make([]int, d)
		dp[i][1] = make([]int, d)
	}
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < d; j++ {
			for k := 0; k < 10; k++ {
				dp[i+1][1][(j+k)%d] += dp[i][1][j]
				dp[i+1][1][(j+k)%d] %= p
			}
			m := int(k[i] - '0')
			for k := 0; k < m; k++ {
				dp[i+1][1][(j+k)%d] += dp[i][0][j]
				dp[i+1][0][(j+k)%d] %= p
			}
			dp[i+1][0][(j+m)%d] += dp[i][0][j]
			dp[i+1][0][(j+m)%d] %= p
		}
		//fmt.Println(dp[i+1])
	}
	ans = dp[n][0][0] + dp[n][1][0] - 1
	if ans < 0 {
		ans += p
	}
	ans %= p
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k, d := nextString(), nextInt()
	ans := SolveHonestly(k, d)
	fmt.Println(ans)
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

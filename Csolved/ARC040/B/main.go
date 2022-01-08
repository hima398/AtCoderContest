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

	n, r := nextInt(), nextInt()
	s := nextString()

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
	}
	idx := 0
	for i := 1; i <= n; i++ {
		if s[i-1] == '.' {
			idx = i
		}
		//i番目のマスまでの歩数
		dp[0][i] = i - 1
		if s[i-1] == 'o' {
			dp[1][i] = dp[1][i-1]
		} else {
			dp[1][i] = dp[1][Max(i-r, 0)] + 1
		}
	}
	//fmt.Println(dp[0])
	//fmt.Println(dp[1])
	ans := dp[0][Max(idx-r+1, 0)] + dp[1][idx]

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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

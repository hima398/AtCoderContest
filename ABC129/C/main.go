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
	n, m := nextInt(), nextInt()
	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = nextInt()
	}
	dp := make([]int, n+1)
	dp[0] = 1
	i := 0
	if m > 0 && a[i] == 1 {
		dp[1] = 0
		i++
	} else {
		dp[1] = 1
	}
	for j := 2; j <= n; j++ {
		if i < m && a[i] == j {
			dp[j] = 0
			i++
		} else {
			dp[j] = dp[j-1] + dp[j-2]
			dp[j] %= p
		}
	}
	fmt.Println(dp[n])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

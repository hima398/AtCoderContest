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
	h, w := nextInt(), nextInt()
	a := make([]string, h)
	for i := 0; i < h; i++ {
		a[i] = nextString()
	}
	dp := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		dp[i] = make([]int, w+1)
	}
	//dp[1][1] = 1
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if i == 1 && j == 1 {
				dp[i][j] = 1
				continue
			}
			if a[i-1][j-1] == '.' {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % p
			}
		}
	}
	fmt.Println(dp[h][w])
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

var dp [2000][2000]int
var x, y, z [2000][2000]int

const M = 1000000007

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()
	s := make([]string, 2000)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	dp[0][0] = 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if s[i][j] == '#' {
				continue
			}
			if j > 0 {
				x[i][j] = (x[i][j-1] + dp[i][j-1]) % M
			}
			if i > 0 {
				y[i][j] = (y[i-1][j] + dp[i-1][j]) % M
			}
			if i > 0 && j > 0 {
				z[i][j] = (z[i-1][j-1] + dp[i-1][j-1]) % M
			}
			dp[i][j] = (x[i][j] + y[i][j] + z[i][j]) % M
		}
	}
	fmt.Println(dp[h-1][w-1])
}

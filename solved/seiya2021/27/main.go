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

	h, w := nextInt(), nextInt()
	c := make([]string, h)
	for i := 0; i < h; i++ {
		c[i] = nextString()
	}
	dp := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			if c[i][j] == '.' {
				dp[i][j] = Max(dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	fmt.Println(dp[0][0])
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

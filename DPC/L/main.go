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
	a := nextIntSlice(n)
	//区間i, jが残っているときのX-Y
	dp := make([][]int, n)
	IsCached := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		IsCached[i] = make([]bool, n)
	}

	var F func(i, j int) int
	F = func(i, j int) int {
		if IsCached[i][j] {
			return dp[i][j]
		}
		if i == j {
			dp[i][j] = a[i]
			IsCached[i][j] = true
			return dp[i][j]
		}
		dp[i][j] = Max(a[i]-F(i+1, j), a[j]-F(i, j-1))
		IsCached[i][j] = true
		return dp[i][j]
	}
	ans := F(0, n-1)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

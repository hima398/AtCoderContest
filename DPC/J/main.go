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

	const mn = 301
	var dp [mn][mn][mn]float64
	n := nextInt()
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		a := nextInt()
		m[a]++
	}
	var F func(i, j, k int) float64
	F = func(i, j, k int) float64 {
		if dp[i][j][k] > 0.0 {
			return dp[i][j][k]
		}
		if i == 0 && j == 0 && k == 0 {
			return 0.0
		}
		p0 := float64(n-i-j-k) / float64(n)
		p1 := float64(i) / float64(n)
		p2 := float64(j) / float64(n)
		p3 := float64(k) / float64(n)
		ret := 1.0
		if i > 0 {
			ret += F(i-1, j, k) * p1
		}
		if j > 0 {
			ret += F(i+1, j-1, k) * p2
		}
		if k > 0 {
			ret += F(i, j+1, k-1) * p3
		}
		ret /= (1.0 - p0)
		dp[i][j][k] = ret
		return dp[i][j][k]
	}
	fmt.Println(F(m[1], m[2], m[3]))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

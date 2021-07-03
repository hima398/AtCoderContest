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

	const INF = int(1e17) + 1
	const ms = 60

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	var dp [ms + 1][2]int
	for i := 0; i <= ms; i++ {
		for j := 0; j < 2; j++ {
			dp[i][j] = -INF
		}
	}
	dp[0][0] = 0
	for i := 0; i < 60; i++ {
		for j := 0; j < 2; j++ {
			if dp[i][j] >= 0 {
				d := ms - i - 1
				m := 1 << d
				nz, no := 0, 0
				for i := range a {
					if a[i]&m > 0 {
						no++
					} else {
						nz++
					}
				}
				j2 := j
				if k&m > 0 {
					j2 = 1
				}
				dp[i+1][j2] = Max(dp[i+1][j2], dp[i][j]+no*m)
				if k&m > 0 || j == 1 {
					dp[i+1][j] = Max(dp[i+1][j], dp[i][j]+nz*m)
				}
			}
		}
	}
	ans := Max(dp[60][0], dp[60][1])
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

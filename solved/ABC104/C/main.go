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

	d, g := nextInt(), nextInt()
	g /= 100
	p, c := make([]int, d), make([]int, d)
	// 問題数 最大 d(p0+p1...pn-1) -> 1000くらい
	nq := 0
	// 総得点数
	//得点数だけループ(最大1*100+2*100+...5500 + 10**4*10 -> 105500くらい)
	np := 0
	for i := 0; i < d; i++ {
		p[i], c[i] = nextInt(), nextInt()
		c[i] /= 100
		nq += p[i]
		np += (i+1)*p[i] + c[i]
	}
	//100*i点つけられた問題を解いた時j点取れる最小の問題数
	dp := make([][]int, d+1)
	for i := 0; i <= d; i++ {
		dp[i] = make([]int, np+1)
		for j := 0; j <= np; j++ {
			//そう問題数より明らかに大きい値を入れておく
			dp[i][j] = nq + 1
		}
	}
	dp[0][0] = 0
	for i := 0; i < d; i++ {
		for j := 0; j < np; j++ {
			for k := 0; k < p[i]; k++ {
				point := j + (i+1)*k
				if point > np {
					break
				}
				dp[i+1][point] = Min(dp[i+1][point], dp[i][j]+k)
			}
			if j+(i+1)*p[i]+c[i] <= np {
				dp[i+1][j+(i+1)*p[i]+c[i]] = Min(dp[i+1][j+(i+1)*p[i]+c[i]], dp[i][j]+p[i])
			}
		}
	}
	ans := nq + 1
	for i := g; i <= np; i++ {
		ans = Min(ans, dp[d][i])
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

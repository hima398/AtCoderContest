package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, k int, x, y []int) int {
	const INF = 1 << 60
	dm := make([][]int, n)
	for i := 0; i < n; i++ {
		dm[i] = make([]int, n)
		for j := 0; j < n; j++ {
			xd, yd := x[j]-x[i], y[j]-y[i]
			dm[i][j] = xd*xd + yd*yd
		}
	}

	m := (1 << n) - 1
	d := make([]int, m+1)
	for pat := 0; pat <= m; pat++ {
		//bd := 0
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				if (pat>>i)&1 == 1 && (pat>>j)&1 == 1 {
					d[pat] = Max(d[pat], dm[i][j])
				}
			}
		}
		//d[pat] = bd
	}
	//fmt.Println(d)

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= k; i++ {
		for pat := 1; pat <= m; pat++ {
			for j := pat; j > 0; j = (j - 1) & pat {
				dp[i][pat] = Min(dp[i][pat], Max(dp[i-1][pat-j], d[j]))
			}
		}
		//fmt.Println(dp[i])
	}
	return dp[k][m]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	x, y := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}
	ans := Solve(n, k, x, y)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

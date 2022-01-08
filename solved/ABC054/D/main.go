package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, ma, mb int, a, b, c []int) int {
	const max = 400
	const INF = 1 << 60
	dp := make([][max + 1][max + 1]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j <= max; j++ {
			for k := 0; k <= max; k++ {
				dp[i][j][k] = INF
			}
		}
	}

	dp[0][0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= max; j++ {
			for k := 0; k <= max; k++ {
				dp[i][j][k] = Min(dp[i][j][k], dp[i-1][j][k])
				nj, nk := j+a[i-1], k+b[i-1]
				if nj <= max && nk <= max {
					dp[i][nj][nk] = Min(dp[i][nj][nk], dp[i-1][j][k]+c[i-1])
				}
			}
		}
	}
	ans := INF
	for j := 1; j <= max; j++ {
		for k := 1; k <= max; k++ {
			if mb*j == ma*k {
				ans = Min(ans, dp[n][j][k])
			}
		}
	}
	if ans == INF {
		ans = -1
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, ma, mb := nextInt(), nextInt(), nextInt()
	a, b, c := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i] = nextInt(), nextInt(), nextInt()
	}
	ans := Solve(n, ma, mb, a, b, c)
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

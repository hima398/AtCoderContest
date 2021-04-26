package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INF = 1 << 60

var sc = bufio.NewScanner(os.Stdin)

func ComputeCost(a, b, c, p, q, r int) int {
	return Abs(p-a) + Abs(q-b) + Max(0, r-c)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x, y, z := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i], z[i] = nextInt(), nextInt(), nextInt()
	}
	dp := make([][]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[1][0] = 0
	for i := 1; i < 1<<n; i++ {
		if i&1 == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			//未到達場所が出発点になっていたら何も処理しない
			if (i>>j)&1 != 1 {
				continue
			}
			for k := 0; k < n; k++ {
				nextBits := i | 1<<k
				cost := ComputeCost(x[j], y[j], z[j], x[k], y[k], z[k])
				dp[nextBits][k] = Min(dp[nextBits][k], dp[i][j]+cost)
			}
		}
	}
	ans := INF
	for i := 0; i < n; i++ {
		cost := ComputeCost(x[i], y[i], z[i], x[0], y[0], z[0])
		ans = Min(ans, dp[(1<<n)-1][i]+cost)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

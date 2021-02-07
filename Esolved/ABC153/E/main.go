package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveDp(h, n int) int {

	a := make([]int, n+1)
	b := make([]int, n+1)
	const ansmax = 1000000000000000001 // 解答と比較して十分大きい値

	max := 0
	for i := 1; i <= n; i++ {
		a[i], b[i] = nextInt(), nextInt()
		max = Max(max, a[i])
	}
	dp := make([]int, h+max+1) // 0 ... n+max
	for i := 0; i <= h+max; i++ {
		dp[i] = ansmax
	}
	dp[0] = 0
	for i := 1; i <= h+max; i++ {
		for j := 1; j <= n; j++ {
			if i-a[j] < 0 {
				continue
			}
			dp[i] = Min(dp[i], dp[i-a[j]]+b[j])
		}
	}
	// h以上の最小値選ぶ
	ans := ansmax
	//fmt.Println(dp)
	for i := h; i <= h+max; i++ {
		ans = Min(ans, dp[i])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, n := nextInt(), nextInt()
	fmt.Println(SolveDp(h, n))
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

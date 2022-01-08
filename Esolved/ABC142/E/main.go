package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const INF = 1 << 60 //十分に大きな値にしておく

func PrintDp(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		var v string
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] == INF {
				v = "INF"
			} else {
				v = strconv.Itoa(dp[i][j])
			}
			fmt.Printf(" %s", v)
		}
		fmt.Println()
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := make([]int, m)
	b := make([]int, m)
	c := make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = nextInt(), nextInt()
		cs := nextIntSlice(b[i])
		ci := 0
		for _, v := range cs {
			ci |= 1 << (v - 1)
		}
		c[i] = ci
	}
	//最大でn**12 -1 = 4095
	pattern := (1 << n) - 1
	//0〜i番目の鍵を使ってjのbitパターンで開けられる宝箱の費用の最小値
	// dpは 1e3 * 4095で1e7程度
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, pattern+1)
		for j := 0; j <= pattern; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 0; i < m; i++ {
		//i+1番目の鍵を全く使わない場合の最小値をi番目からあらかじめ持ってきておく
		for j := 0; j <= pattern; j++ {
			dp[i+1][j] = dp[i][j]
		}
		for j := 0; j <= pattern; j++ {
			dp[i+1][j|c[i]] = Min(dp[i+1][j|c[i]], dp[i][j]+a[i])
		}
	}

	//デバッグ
	//PrintDp(dp)

	//全ての宝箱を開けることができない
	if dp[m][pattern] == INF {
		fmt.Println(-1)
		return
	}
	fmt.Println(dp[m][pattern])
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

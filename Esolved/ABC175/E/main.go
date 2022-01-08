package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintDp(dp [][][4]int, k int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Printf("%d ", dp[i][j][k])
		}
		fmt.Println("")
	}
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const nk = 3
	r, c, k := nextInt(), nextInt(), nextInt()
	f := make([][]int, r)
	for i := 0; i < r; i++ {
		f[i] = make([]int, c)
	}
	for i := 0; i < k; i++ {
		ri, ci, v := nextInt(), nextInt(), nextInt()
		ri--
		ci--
		f[ri][ci] = v
	}
	// i行j列で、i行目のうちアイテムをk個取った場合のスコアの最大値
	dp := make([][][nk + 1]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([][nk + 1]int, c)
	}
	if f[0][0] > 0 {
		dp[0][0][1] = f[0][0]
	}
	for i := 1; i < r; i++ {
		// 0行目でいくつアイテムを取ったかは1行目に影響しないので
		// 0行目のj列めの最大を取ってくる
		preMax := 0
		for kk := nk; kk >= 0; kk-- {
			preMax = Max(preMax, dp[i-1][0][kk])
		}
		for kk := nk; kk >= 0; kk-- {
			dp[i][0][kk] = preMax
		}
		if f[i][0] > 0 {
			for kk := nk; kk >= 1; kk-- {
				dp[i][0][kk] = Max(dp[i][0][kk], dp[i][0][kk-1]+f[i][0])
			}
		}
	}
	for j := 1; j < c; j++ {
		for kk := nk; kk >= 0; kk-- {
			dp[0][j][kk] = dp[0][j-1][kk]
		}
		if f[0][j] > 0 {
			for kk := nk; kk >= 1; kk-- {
				dp[0][j][kk] = Max(dp[0][j][kk], dp[0][j][kk-1]+f[0][j])
			}
		}
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			// i-1行目でいくつアイテムを取ったかはi行目に影響しないので
			// i-1行目のj列めの最大を取ってくる
			preMax := 0
			for kk := nk; kk >= 0; kk-- {
				preMax = Max(preMax, dp[i-1][j][kk])
			}
			for kk := nk; kk >= 0; kk-- {
				dp[i][j][kk] = Max(preMax, dp[i][j-1][kk])
			}
			if f[i][j] > 0 {
				for kk := nk; kk >= 1; kk-- {
					dp[i][j][kk] = Max(dp[i][j][kk], dp[i][j][kk-1]+f[i][j])
				}
			}
		}
	}
	ans := 0
	for kk := nk; kk >= 0; kk-- {
		ans = Max(ans, dp[r-1][c-1][kk])
	}
	//PrintDp(dp, 3)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

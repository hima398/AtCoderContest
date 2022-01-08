package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(c string) int {
	n := len(c)

	var pattern []int
	for i := 0; i < n; i++ {
		p := 0
		for j := 0; j < len(c); j++ {
			if c[j] == 'o' {
				p |= 1 << j
			}
		}
		pattern = append(pattern, p)
		c = c[1:] + string(c[0])
	}
	//ビットパターンのデバッグ出力
	for _, p := range pattern {
		fmt.Printf("%010b ", p)
	}
	fmt.Println()
	mask := (1 << n) - 1
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, mask+1)
		for j := 0; j <= mask; j++ {
			if i != 0 {
				dp[i][j] = n
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= mask; j++ {
			if i == 0 {
				break
			}
			dp[i+1][j] = Min(dp[i+1][j], dp[i][j])
		}
		smask := (1<<(i+1) - 1)
		m := mask & smask
		for j := 0; j <= m; j++ {
			for k := 0; k < len(pattern); k++ {
				nj := j | (pattern[k] & smask)
				dp[i+1][nj] = Min(dp[i+1][nj], dp[i][j]+1)
			}
		}
	}
	for i := 0; i <= n; i++ {
		fmt.Println(dp[i])
	}
	ans := dp[n][mask]
	return ans

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	c := nextString()
	//FirstSolve(c)
	n := len(c)

	var pattern []int
	for i := 0; i < n; i++ {
		p := 0
		for j := 0; j < len(c); j++ {
			if c[j] == 'o' {
				p |= 1 << j
			}
		}
		pattern = append(pattern, p)
		c = c[1:] + string(c[0])
	}
	//ビットパターンのデバッグ出力
	/*
		for _, p := range pattern {
			fmt.Printf("%010b ", p)
		}
		fmt.Println()
	*/

	mask := (1 << n) - 1
	dp := make([]int, mask+1)
	for i := 0; i <= mask; i++ {
		dp[i] = n
	}
	dp[0] = 0

	for i := 0; i <= mask; i++ {
		for j := 0; j < len(pattern); j++ {
			ni := i | pattern[j]
			dp[ni] = Min(dp[ni], dp[i]+1)
		}
	}

	ans := dp[mask]
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

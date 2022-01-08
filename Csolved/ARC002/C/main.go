package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, c string) int {
	p := []string{"AA", "AB", "AX", "AY", "BA", "BB", "BX", "BY", "XA", "XB", "XX", "XY", "YA", "YB", "YX", "YY"}

	ans := n
	for i := 0; i < len(p)-1; i++ {
		for j := i + 1; j < len(p); j++ {
			cn := n
			for k := 0; k < n-1; k++ {
				ss := string(c[k]) + string(c[k+1])
				if ss == p[i] || ss == p[j] {
					cn--
					k++
				}
			}
			ans = Min(ans, cn)
			/*
				cn = n
				for k := 1; k < n-1; k++ {
					ss := string(c[k]) + string(c[k+1])
					if ss == p[i] || ss == p[j] {
						cn--
						k++
					}
				}
				ans = Min(ans, cn)
			*/
		}
	}
	return ans
}

func Solve(n int, c string) int {
	p := []string{"AA", "AB", "AX", "AY", "BA", "BB", "BX", "BY", "XA", "XB", "XX", "XY", "YA", "YB", "YX", "YY"}
	ans := n
	for l := 0; l < len(p)-1; l++ {
		for r := l + 1; r < len(p); r++ {
			//i番目(1-indexed)までの文字列を短縮できる長さ
			dp := make([]int, n+1)
			for i := 0; i <= n; i++ {
				dp[i] = i
			}
			for i := 2; i <= n; i++ {
				dp[i] = Min(dp[i], dp[i-1]+1)
				ss := c[i-2 : i]
				if ss == p[l] || ss == p[r] {
					dp[i] = Min(dp[i], dp[i-2]+1)
				}
			}
			ans = Min(ans, dp[n])
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	c := nextString()
	fmt.Println(SolveHonestly(n, c))
	//fmt.Println(Solve(n, c))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

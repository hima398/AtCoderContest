package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, s string) int {
	ans := 0
	for i := 1; i <= n/2; i++ {
	label:
		for j := 0; j <= Ceil(n, 2)-i; j++ {
			sj := s[j : j+i]
			for k := j + i; k <= n-i; k++ {
				sk := s[k : k+i]
				//fmt.Println(i, j, k, sj, sk)
				if sj == sk {
					ans = i
					break label
				}
			}
		}
	}
	return ans
}

func SolveCommentary(n int, s string) int {
	dp := make([][]int, n+1)
	for i:=0; i<=n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i:=n-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}
		}
	}
	ans := 0
	for i:=0; i<n; i++ {
		for j:=i;j<n; j++ {
			ans = Max(ans, Min(dp[i][j], j-i))
		} 
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	//ans := SolveHonestly(n, s)
	ans := SolveCommentary(n, s)
	fmt.Println(ans)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

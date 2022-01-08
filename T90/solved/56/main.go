package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Reverse(s string) string {
	ret := ""
	for i := len(s) - 1; i >= 0; i-- {
		ret += string(s[i])
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, s := nextInt(), nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	dp := make([]map[int]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make(map[int]bool)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		for k := range dp[i] {
			if k+a[i] <= s {
				dp[i+1][k+a[i]] = true
			}
			if k+b[i] <= s {
				dp[i+1][k+b[i]] = true
			}
		}
	}
	if !dp[n][s] {
		fmt.Println("Impossible")
		return
	}
	route := ""
	j := s
	for i := n; i > 0; i-- {
		if dp[i-1][j-a[i-1]] {
			route += "A"
			j -= a[i-1]
		} else if dp[i-1][j-b[i-1]] {
			route += "B"
			j -= b[i-1]
		} else {
			fmt.Println("Not Found")
		}
	}
	ans := Reverse(route)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

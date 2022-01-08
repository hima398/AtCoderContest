package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(s string) (ans int) {
	const p = int(1e9) + 7
	n := len(s)
	dp := make([]int, n+2)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := i - 1; ; j-- {
			dp[i+2] = (dp[i+2] + dp[j+1]) % p
			if j == -1 || s[j] == s[i] {
				break
			}
		}
	}
	for i := 2; i < n+2; i++ {
		ans += dp[i]
		ans %= p
	}
	return ans
}

func FirstSolve(s string) (ans int) {
	const p = int(1e9) + 7
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 26)
	}
	var m [26][]int
	for i := 0; i < n; i++ {
		idx := s[i] - 'a'
		m[idx] = append(m[idx], i)
	}
	mi := make([]int, 26)
	for i := 1; i <= n; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				idx := mi[k+'a']
				mi[k+'a']++
				if idx >= len(m[k+'a']) {
					continue
				}

			}
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()

	ans := SolveCommentary(s)
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7

	n := nextInt()
	s := nextString()
	t := make([]string, n)
	for i := 0; i < n; i++ {
		t[i] = nextString()
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		for j := 0; j < n; j++ {
			if len(s[:i]) < len(t[j]) {
				continue
			}
			pi := i - len(t[j])
			ss := s[pi:i]
			//fmt.Println(i, pi, s[:i], ss, t[j])
			if ss == t[j] {
				dp[i] += dp[pi]
				dp[i] %= p
			}
		}
	}
	fmt.Println(dp[len(s)])
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

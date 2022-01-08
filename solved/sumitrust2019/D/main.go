package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var dp [30001][4][1000]bool

func Solve(n int, s string) int {
	dp[0][0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 1000; k++ {
				if !dp[i][j][k] {
					continue
				}
				dp[i+1][j][k] = true
				if j <= 2 {
					dp[i+1][j+1][k*10+int(s[i]-'0')] = true
				}
			}
		}
	}
	ans := 0
	for i := 0; i < 1000; i++ {
		if dp[n][3][i] {
			ans++
		}
	}
	return ans
}

func SolveHonestly(n int, s string) int {
	ans := 0
	for i := 0; i < 1000; i++ {
		v := fmt.Sprintf("%03d", i)
		j := 0
		for k := 0; k < len(s); k++ {
			if j >= 3 {
				break
			}
			if s[k] != v[j] {
				continue
			}
			j++
		}
		if j >= 3 {
			ans++
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
	//fmt.Println(SolveHonestly(n, s))
	fmt.Println(Solve(n, s))
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

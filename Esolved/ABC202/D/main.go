package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, k := nextInt(), nextInt(), nextInt()
	/*
		dp := make([][]int, a+1)
		for i := 0; i <= a; i++ {
			dp[i] = make([]int, b+1)
			for j := 0; j <= b; j++ {
				dp[i][j] = Combination(i+j, j)
			}
		}
	*/

	//データ数見積もりの名残
	//s := Combination(a+b, b)

	var F func(a, b, k int) string
	// "a"がa個、"b"がb個あるときk番目の文字
	F = func(a, b, k int) string {
		//fmt.Println(a, b, k)
		if a == 0 {
			return strings.Repeat("b", b)
		}
		if b == 0 {
			return strings.Repeat("a", a)
		}
		c := Combination(a+b-1, a-1)
		/*
			if k > dp[a-1][b] {
				return "b" + F(a, b-1, k-dp[a-1][b])
			}
		*/
		if k > c {
			return "b" + F(a, b-1, k-c)
		} else {
			return "a" + F(a-1, b, k)
		}

	}
	ans := F(a, b, k)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

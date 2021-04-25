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

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	// 石がi個残っている時、先手が勝てるかどうか
	dp := make([]bool, k+1)
	for i := 1; i <= k; i++ {
		winSecond := true
		for _, v := range a {
			if i-v >= 0 {
				winSecond = winSecond && dp[i-v]
			} else {
				break
			}
		}
		// 先手が勝てるか=相手に負けを渡せるかの否定
		dp[i] = !winSecond
	}
	//fmt.Println(dp)
	if dp[k] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
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

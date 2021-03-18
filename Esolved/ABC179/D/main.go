package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Section struct {
	l, r int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = 998244353
	n, k := nextInt(), nextInt()
	section := make([]Section, k+1)
	for i := 1; i <= k; i++ {
		section[i] = Section{l: nextInt(), r: nextInt()}
	}
	dp := make([]int, n+1)
	sdp := make([]int, n+1)
	dp[1] = 1
	sdp[1] = dp[1]
	for i := 2; i <= n; i++ {
		tdp := 0
		for j := 1; j <= k; j++ {
			if i-section[j].l >= 1 {
				tdp += sdp[i-section[j].l] - sdp[Max(i-section[j].r-1, 0)]
				if tdp < 0 {
					tdp += p
				}
				tdp %= p
				//fmt.Println(i, tdp)
			}
		}
		dp[i] += tdp
		dp[i] %= p
		sdp[i] = (dp[i] + sdp[i-1]) % p

	}
	//fmt.Println(dp)
	//fmt.Println(sdp)
	fmt.Println(dp[n])
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

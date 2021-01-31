package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const p = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextInt()

	dp := make([]int, s+1)
	dp[0] = 1
	sdp := 0
	for i := 1; i <= s; i++ {
		if i-3 >= 0 {
			sdp += dp[i-3]
			sdp %= p
			dp[i] = sdp
		}
	}
	//fmt.Println(dp)
	fmt.Println(dp[s])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

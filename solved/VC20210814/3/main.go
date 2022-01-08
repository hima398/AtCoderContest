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

	x := nextInt()

	dp := make([]bool, x+1)
	dp[0] = true
	for i := 0; i < x; i++ {
		if dp[i] {
			for j := 0; j < 6; j++ {
				nx := i + j + 100
				if nx <= x {
					dp[nx] = true
				}
			}
		}
	}
	if dp[x] {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

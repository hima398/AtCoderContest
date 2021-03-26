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

	n, t := nextInt(), nextInt()
	ans := 1001
	for i := 0; i < n; i++ {
		ci, ti := nextInt(), nextInt()
		if ti <= t {
			ans = Min(ans, ci)
		}
	}
	if ans < 1001 {
		fmt.Println(ans)
	} else {
		fmt.Println("TLE")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

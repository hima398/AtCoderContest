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

	n, k, m := nextInt(), nextInt(), nextInt()
	s := 0
	for i := 0; i < n-1; i++ {
		a := nextInt()
		s += a
	}
	ans := m*n - s
	if ans > k {
		fmt.Println(-1)
	} else {
		fmt.Println(Max(ans, 0))
	}
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	b := make([]int, n+1)
	ma := make([]int, n+1)
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a := nextInt()
		ma[i] = Max(ma[i-1], a)
	}
	for i := 1; i <= n; i++ {
		b[i] = nextInt()
	}
	for i := 1; i <= n; i++ {
		c[i] = Max(c[i-1], ma[i]*b[i])
		fmt.Println(c[i])
	}
}

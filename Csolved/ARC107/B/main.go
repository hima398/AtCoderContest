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

	F := func(n, k int) int {
		return Min(k-1, 2*n+1-k)
	}
	ans := 0
	for x := 2; x <= 2*n; x++ {
		y := x - k
		if y >= 2 && y <= 2*n {
			ans += F(n, x) * F(n, y)
		}
	}
	fmt.Println(ans)
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

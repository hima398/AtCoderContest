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

	n := nextInt()
	a := make([]int, n+1)
	sa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		sa[i] = sa[i-1] + a[i]
	}
	ans := 1 << 60
	for i := 1; i <= n-1; i++ {
		x := sa[i]
		y := sa[n] - sa[i]
		ans = Min(ans, Abs(x-y))
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

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
	w := make([]int, n)
	for i := 0; i < n; i++ {
		w[i] = nextInt()
	}
	ans := int(1e6)
	for t := 0; t < n-1; t++ {
		s1, s2 := 0, 0
		for i := 0; i <= t; i++ {
			s1 += w[i]
		}
		for i := t + 1; i < n; i++ {
			s2 += w[i]
		}
		ans = Min(ans, Abs(s1-s2))
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

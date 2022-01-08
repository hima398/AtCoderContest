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
	a, b := nextInt(), nextInt()
	p := nextIntSlice(n)
	n1, n2, n3 := 0, 0, 0
	for i := 0; i < n; i++ {
		if p[i] <= a {
			n1++
		} else if p[i] > a && p[i] <= b {
			n2++
		} else { // p[i] > b
			n3++
		}
	}
	ans := Min(n1, Min(n2, n3))
	fmt.Println(ans)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

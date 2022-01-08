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

	n, m, a, b := nextInt(), nextInt(), nextInt(), nextInt()
	c := nextIntSlice(m)
	for i := range c {
		if n <= a {
			n += b
		}

		if n < c[i] {
			fmt.Println(i + 1)
			return
		}
		n -= c[i]
	}
	fmt.Println("complete")
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

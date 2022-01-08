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
	a := nextIntSlice(n)
	b := make([]int, n)
	for i := 0; i < n-k+1; i++ {
		l, r := i, i+k
		b[l]++
		if r < n {
			b[r]--
		}
	}
	for i := 1; i < n; i++ {
		b[i] += b[i-1]
	}
	//fmt.Println(a, b)
	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i] * b[i]
	}
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

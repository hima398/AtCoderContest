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

	const p = int(1e9) + 7
	n, m := nextInt(), nextInt()
	x := nextIntSlice(n)
	y := nextIntSlice(m)

	sx := 0
	for i := 0; i < n; i++ {
		sx += (2*i - n + 1) * x[i]
		sx %= p
	}
	sy := 0
	for i := 0; i < m; i++ {
		sy += (2*i - m + 1) * y[i]
		sy %= p
	}
	ans := sx * sy % p

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

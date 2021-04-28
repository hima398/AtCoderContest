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

	const nmax = int(1e5) + 1
	n := nextInt()
	c := make([][]int, 2)
	for i := 0; i < 2; i++ {
		c[i] = make([]int, nmax)
	}
	for i := 0; i < n; i++ {
		ci, pi := nextInt(), nextInt()
		ci--
		c[ci][i] = pi
		c[ci][i] = pi
	}
	s := make([][]int, 2)
	for i := 0; i < 2; i++ {
		s[i] = make([]int, nmax)
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < nmax-1; j++ {
			s[i][j+1] = s[i][j] + c[i][j]
		}
	}
	q := nextInt()
	a, b := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		l, r := nextInt(), nextInt()
		a[i] = s[0][r] - s[0][l-1]
		b[i] = s[1][r] - s[1][l-1]
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, a[i], b[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

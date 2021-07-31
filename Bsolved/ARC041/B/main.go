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

	n, m := nextInt(), nextInt()
	b := make([][]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		s := nextString()
		b[i] = make([]int, m)
		a[i] = make([]int, m)
		for j := range s {
			b[i][j] = int(s[j] - '0')
		}
	}
	di := []int{0, -1, 0, 1}
	dj := []int{-1, 0, 1, 0}
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			v := int(1e6) + 1
			for k := 0; k < 4; k++ {
				ni, nj := i+di[k], j+dj[k]
				v = Min(v, b[ni][nj])
			}
			for k := 0; k < 4; k++ {
				ni, nj := i+di[k], j+dj[k]
				b[ni][nj] -= v
			}
			a[i][j] = v
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(out, "%d", a[i][j])
		}
		fmt.Fprintln(out, "")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

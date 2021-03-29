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

	const n = 10

	h, w := nextInt(), nextInt()
	var c [n][n]int
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = nextInt()
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				c[i][j] = Min(c[i][j], c[i][k]+c[k][j])
			}
		}
	}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] != -1 {
				ans += c[a[i][j]][1]
			}
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

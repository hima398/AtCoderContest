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
	x := make([]int, n+1)
	y := make([]int, n+1)
	h := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x[i], y[i], h[i] = nextInt(), nextInt(), nextInt()
		if h[i] > 0 {
			x[0] = x[i]
			y[0] = y[i]
			h[0] = h[i]
		}
	}
	for cx := 0; cx <= 100; cx++ {
		for cy := 0; cy <= 100; cy++ {

			ok := true
			lh := h[0] + Abs(x[0]-cx) + Abs(y[0]-cy)
			for i := 0; i < n; i++ {
				ok = ok && Max(lh-Abs(x[i]-cx)-Abs(y[i]-cy), 0) == h[i]
				if !ok {
					break
				}
			}
			if ok {
				fmt.Println(cx, cy, lh)
				return
			}
		}
	}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

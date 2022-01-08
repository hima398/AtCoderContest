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

	w, h, n := nextInt(), nextInt(), nextInt()
	minx := 0
	maxx := w
	miny := 0
	maxy := h
	for i := 0; i < n; i++ {
		x, y, a := nextInt(), nextInt(), nextInt()
		if a == 1 {
			minx = Max(minx, x)
		} else if a == 2 {
			maxx = Min(maxx, x)
		} else if a == 3 {
			miny = Max(miny, y)
		} else {
			maxy = Min(maxy, y)
		}
	}
	ans := Max(maxx-minx, 0) * Max(maxy-miny, 0)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

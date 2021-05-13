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

	n, q := nextInt(), nextInt()
	x := make([]int, n)
	y := make([]int, n)
	xd := make([]int, n)
	yd := make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
		xd[i] = x[i] - y[i] // x' = x - y
		yd[i] = x[i] + y[i] // y' = x + y
	}

	minXd, maxXd := xd[0], xd[0]
	minYd, maxYd := yd[0], yd[0]
	for i := 1; i < n; i++ {
		minXd = Min(minXd, xd[i])
		maxXd = Max(maxXd, xd[i])
		minYd = Min(minYd, yd[i])
		maxYd = Max(maxYd, yd[i])
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		qi := nextInt()
		qi--
		d1 := Abs(xd[qi] - minXd)
		d2 := Abs(xd[qi] - maxXd)
		d3 := Abs(yd[qi] - minYd)
		d4 := Abs(yd[qi] - maxYd)
		ans := Max(d1, Max(d2, Max(d3, d4)))
		fmt.Fprintln(out, ans)
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

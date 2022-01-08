package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	xa, ya, xb, yb, xc, yc := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	a, b := float64(xb-xa), float64(yb-ya)
	c, d := float64(xc-xa), float64(yc-ya)
	ans := math.Abs(a*d-b*c) / 2.0
	fmt.Println(ans)
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

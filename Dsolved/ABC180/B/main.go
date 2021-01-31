package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x := make([]int, n)
	m := 0
	y := 0.0
	c := 0
	for i := 0; i < n; i++ {
		x[i] = nextInt()
		m += Abs(x[i])
		y += float64(x[i]) * float64(x[i])
		c = Max(c, Abs(x[i]))
	}
	fmt.Println(m)
	fmt.Println(math.Sqrt(y))
	fmt.Println(c)
}

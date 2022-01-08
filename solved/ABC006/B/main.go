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

	const p = int(1e4) + 7
	n := nextInt()
	na := Max(3, n)
	a := make([]int, na)
	a[0], a[1], a[2] = 0, 0, 1
	for i := 3; i < na; i++ {
		a[i] = a[i-1] + a[i-2] + a[i-3]
		a[i] %= p
	}
	fmt.Println(a[n-1])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	// 1 <= n <= 200000
	var n int
	fmt.Scan(&n)
	// -10 ^ 8 <= a[i] <= 10 ^8
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	lenA := len(a)
	p := make([]int, n)
	q := make([]int, n)
	p[0] = a[0]
	q[0] = Max(0, a[0])
	robot := p[0]
	max := q[0]
	for i := 1; i < lenA; i++ {
		p[i] = p[i-1] + a[i]
		q[i] = Max(q[i-1], p[i])
		max = Max(max, robot+q[i])
		robot += p[i]
	}
	fmt.Println(max)
}

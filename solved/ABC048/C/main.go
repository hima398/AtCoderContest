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

	n, x := nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	q := []int{a[0]}
	ans := 0
	for i := 1; i < n; i++ {
		q = append(q, a[i])
		if q[0]+q[1] > x {
			e := q[0] + q[1] - x
			q[1] = Max(q[1]-e, 0)
			ans += e
		}
		q = q[1:]
	}
	fmt.Println(ans)
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

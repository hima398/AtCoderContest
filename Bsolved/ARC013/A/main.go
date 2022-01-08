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

	n, m, l := nextInt(), nextInt(), nextInt()
	p, q, r := nextInt(), nextInt(), nextInt()
	ans := (n / p) * (m / q) * (l / r)
	ans = Max(ans, (n/p)*(m/r)*(l/q))
	ans = Max(ans, (n/q)*(m/p)*(l/r))
	ans = Max(ans, (n/q)*(m/r)*(l/p))
	ans = Max(ans, (n/r)*(m/p)*(l/q))
	ans = Max(ans, (n/r)*(m/q)*(l/p))
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

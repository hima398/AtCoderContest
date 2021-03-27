package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, lx := nextInt(), nextInt()
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}
	x = append(x, lx)
	sort.Ints(x)
	ans := x[1] - x[0]
	for i := 1; i < len(x); i++ {
		ans = Gcd(ans, x[i]-x[i-1])
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x < y {
		x, y = y, x
	}
	return Gcd(y, x%y)
}

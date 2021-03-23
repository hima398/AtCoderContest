package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(n int) int {
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	l := make([]int, n+2)
	r := make([]int, n+2)
	for i := 0; i < n; i++ {
		l[i+1] = Gcd(a[i+1], l[i])
	}
	for i := n + 1; i > 1; i-- {
		r[i-1] = Gcd(a[i-1], r[i])
	}
	//fmt.Println(l)
	//fmt.Println(r)
	ans := 0
	m := make([]int, n+2)
	for i := 1; i <= n; i++ {
		m[i] = Gcd(l[i-1], r[i+1])
		ans = Max(ans, m[i])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	fmt.Println(SolveCommentary(n))
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

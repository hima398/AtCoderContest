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

	n := nextInt()
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}

	ps := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	ans := 1
	for _, p := range ps {
		ans *= p
	}
	for bits := 1; bits < 1<<len(ps); bits++ {
		y := 1
		for j := 0; j < len(ps); j++ {
			if bits&(1<<j) > 0 {
				y *= ps[j]
			}
		}
		ok := true
		for i := 0; i < n; i++ {
			ok = ok && (Gcd(x[i], y) > 1)
		}
		if ok {
			ans = Min(ans, y)
		}
	}
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(l, r int) int {
	const p = 2019
	ans := p
	for i := l; i <= r-1; i++ {
		for j := i + 1; j <= r; j++ {
			ans = Min(ans, (i%p)*(j%p)%p)
			if ans == 0 {
				return ans
			}
		}
	}
	return ans
}

func Solve(l, r int) int {
	const p = 2019

	r = Min(r, l+p)
	ans := p
	for i := l; i <= r-1; i++ {
		for j := i + 1; j <= r; j++ {
			ans = Min(ans, (i*j)%p)
			if ans == 0 {
				return ans
			}
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = 2019
	l, r := nextInt(), nextInt()
	//fmt.Println(SolveHonestly(l, r))
	fmt.Println(Solve(l, r))
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

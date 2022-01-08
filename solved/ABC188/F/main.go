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

	x, y := nextInt(), nextInt()
	memo := make(map[int]int)
	var F func(n int) int
	F = func(n int) int {
		if _, ok := memo[n]; ok {
			return memo[n]
		}

		if n == x {
			memo[n] = 0
			return memo[n]
		}

		if n == 1 {
			memo[n] = Abs(x - n)
			return memo[n]
		}
		if n%2 == 0 {
			memo[n] = Min(Abs(x-n), F(n/2)+1)
		} else {
			memo[n] = Min(F(n-1), F(n+1)) + 1
		}
		return memo[n]
	}

	ans := F(y)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

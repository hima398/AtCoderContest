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

	const INF = 3*int(1e8) + 1
	n, k := nextInt(), nextInt()
	x := nextIntSlice(n)
	ans := INF
	for i := 0; i < n-k+1; i++ {
		l, r := x[i], x[i+k-1]
		time := INF
		if l < 0 && r > 0 {
			//k個選んだ左端と右端で符号が異なる場合は
			//左端にろうそくをつけて行って右に行くか、右端にろうそくをつけて左に行くかのどちらか
			time = Min(2*Abs(l)+Abs(r), Abs(l)+2*Abs(r))
		} else {
			time = Max(Abs(l), Abs(r))
		}
		ans = Min(ans, time)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

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

	// k<=int(1e9)なのでシミュレートは困難
	n, k := nextInt(), nextInt()
	l := nextIntSlice(n)
	max := 0
	for _, li := range l {
		max = Max(max, li)
	}

	check := func(x int) bool {
		//全ての丸太をx以下に切り分ける回数
		c := 0
		for _, li := range l {
			c += Ceil(li, x) - 1
		}
		return c <= k
	}
	ng, ok := 0, max
	for ng+1 < ok {
		c := (ok + ng) / 2
		//fmt.Println(ng, ok, check(c))
		if check(c) {
			ok = c
		} else {
			ng = c
		}
	}
	fmt.Println(ok)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

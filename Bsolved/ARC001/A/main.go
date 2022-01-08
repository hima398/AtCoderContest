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
	m := make(map[rune]int)
	cs := nextString()
	for _, r := range cs {
		m[r]++
	}
	mn, mx := n+1, 0
	for _, v := range m {
		mn = Min(mn, v)
		mx = Max(mx, v)
	}
	if len(m) < 4 {
		mn = 0
	}
	fmt.Println(mx, mn)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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

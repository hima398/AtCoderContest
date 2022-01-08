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

	n := nextInt()
	a := nextIntSlice(n)
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	ans := 0
	var bars []int
	for k, v := range m {
		if v >= 4 {
			ans = Max(ans, k*k)
		}
		if v >= 2 {
			bars = append(bars, k)
		}
	}
	sort.Ints(bars)
	if len(bars) >= 2 {
		ans = Max(ans, bars[len(bars)-1]*bars[len(bars)-2])
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

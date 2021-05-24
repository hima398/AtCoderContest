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
	for i := 0; i < n; i++ {
		m[a[i]]++
	}
	var b []int
	c := 0
	for k, v := range m {
		if v >= 2 {
			b = append(b, k)
		}
		if v >= 4 {
			c = Max(c, k)
		}
	}
	if len(b) < 2 {
		fmt.Println(0)
		return
	}
	sort.Ints(b)
	fmt.Println(Max(b[len(b)-1]*b[len(b)-2], c*c))
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

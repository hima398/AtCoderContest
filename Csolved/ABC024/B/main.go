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
	n, t := nextInt(), nextInt()
	a := nextIntSlice(n)
	max := 0
	for _, ai := range a {
		max = Max(max, ai)
	}
	imos := make([]int, max+t+1)
	for _, ai := range a {
		imos[ai]++
		imos[ai+t]--
	}
	for i := 0; i < max+t; i++ {
		imos[i+1] += imos[i]
	}
	ans := 0
	for _, v := range imos {
		if v > 0 {
			ans++
		}
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

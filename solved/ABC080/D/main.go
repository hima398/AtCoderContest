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

	n, lc := nextInt(), nextInt()
	s := make([]int, n)
	t := make([]int, n)
	c := make([]int, n)
	tmax := 0
	for i := 0; i < n; i++ {
		s[i], t[i], c[i] = nextInt(), nextInt(), nextInt()
		c[i]--
		tmax = Max(tmax, t[i])
	}
	imos := make([][]int, lc)
	for i := 0; i < lc; i++ {
		imos[i] = make([]int, tmax+1)
	}
	for i := 0; i < n; i++ {
		imos[c[i]][s[i]-1]++
		imos[c[i]][t[i]]--
	}

	for k := 0; k < lc; k++ {
		for i := 0; i < tmax; i++ {
			imos[k][i+1] += imos[k][i]
		}
	}
	/*
		for _, v := range ss {
			fmt.Println(v)
		}
	*/
	ans := 0

	for i := 0; i < tmax+1; i++ {
		nr := 0 // 録画機の数
		for k := 0; k < lc; k++ {
			if imos[k][i] > 0 {
				nr++
			}
		}
		ans = Max(ans, nr)
	}

	fmt.Println(ans)
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

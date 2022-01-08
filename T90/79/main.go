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

	h, w := nextInt(), nextInt()
	a := make([][]int, h)
	b := make([][]int, h)
	d := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = nextIntSlice(w)
	}
	for i := 0; i < h; i++ {
		b[i] = nextIntSlice(w)
	}
	for i := 0; i < h; i++ {
		d[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			d[i][j] = b[i][j] - a[i][j]
		}
	}
	/*
		for i := 0; i < h; i++ {
			fmt.Println(d[i])
		}
	*/
	ans := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			if d[i][j] != 0 {
				d[i+1][j] -= d[i][j]
				d[i][j+1] -= d[i][j]
				d[i+1][j+1] -= d[i][j]
				ans += Abs(d[i][j])
				d[i][j] = 0
			}
		}
	}
	ok := true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if d[i][j] != 0 {
				ok = false
				break
			}
		}
	}
	if ok {
		fmt.Println("Yes")
		fmt.Println(ans)
	} else {
		fmt.Println("No")
	}
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

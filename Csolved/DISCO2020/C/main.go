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
	h, w, _ := nextInt(), nextInt(), nextInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}
	group := 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				a[i][j] = group
				group++
			}
		}
	}
	// 左から右
	for i := 0; i < h; i++ {
		for j := 1; j < w; j++ {
			if a[i][j] == 0 {
				a[i][j] = a[i][j-1]
			}
		}
	}
	// 右から左
	for i := 0; i < h; i++ {
		for j := w - 2; j >= 0; j-- {
			if a[i][j] == 0 {
				a[i][j] = a[i][j+1]
			}
		}
	}
	// 上から下
	for i := 1; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == 0 {
				a[i][j] = a[i-1][j]
			}
		}
	}
	// 下から上
	for i := h - 2; i >= 0; i-- {
		for j := 0; j < w; j++ {
			if a[i][j] == 0 {
				a[i][j] = a[i+1][j]
			}
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprintf(out, "%d ", a[i][j])
		}
		fmt.Fprintln(out)
	}
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

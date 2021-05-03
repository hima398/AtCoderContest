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
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
		}
	}
	sr := make([]int, h)
	sc := make([]int, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			sr[i] += a[i][j]
		}
	}
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			sc[j] += a[i][j]
		}
	}
	ans := make([][]int, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[i][j] = sr[i] + sc[j] - a[i][j]
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprintf(out, "%d ", ans[i][j])
		}
		fmt.Fprintln(out, "")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

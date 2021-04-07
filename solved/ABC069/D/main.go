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
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	ans := make([][]int, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]int, w)
	}
	up := 1
	idx := 0
	for i := 0; i < h; i++ {
		if up == 1 {
			for j := 0; j < w; j++ {
				ans[i][j] = idx + 1
				a[idx]--
				if a[idx] <= 0 {
					idx++
				}
			}
		} else {
			for j := w - 1; j >= 0; j-- {
				ans[i][j] = idx + 1
				a[idx]--
				if a[idx] <= 0 {
					idx++
				}
			}
		}
		up ^= 1
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < h; i++ {
		fmt.Fprintf(out, "%d", ans[i][0])
		for j := 1; j < w; j++ {
			fmt.Fprintf(out, " %d", ans[i][j])
		}
		fmt.Fprintln(out, "")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

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

	h, w, lc, q := nextInt(), nextInt(), nextInt(), nextInt()
	t, n, c := make([]int, q), make([]int, q), make([]int, q)
	mr, mc := make(map[int]bool), make(map[int]bool)
	for i := 0; i < q; i++ {
		t[i], n[i], c[i] = nextInt(), nextInt(), nextInt()
		n[i]--
		c[i]--
	}
	ans := make([]int, lc)
	for i := q - 1; i >= 0; i-- {
		switch t[i] {
		case 1:
			if mr[n[i]] {
				continue
			}
			ans[c[i]] += w - len(mc)
			mr[n[i]] = true
		case 2:
			if mc[n[i]] {
				continue
			}
			ans[c[i]] += h - len(mr)
			mc[n[i]] = true
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < lc; i++ {
		fmt.Fprintf(out, "%d ", ans[i])
	}
	fmt.Fprintln(out)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

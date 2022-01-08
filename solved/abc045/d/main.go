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

	h, w, n := nextInt(), nextInt(), nextInt()
	a, b := make([]int, n), make([]int, n)
	for k := 0; k < n; k++ {
		a[k], b[k] = nextInt(), nextInt()
	}
	m := make(map[int]map[int]int)
	for k := 0; k < n; k++ {
		for i := a[k] - 1; i <= a[k]+1; i++ {
			for j := b[k] - 1; j <= b[k]+1; j++ {
				if m[i] == nil {
					m[i] = make(map[int]int)
				}
				m[i][j]++
			}
		}
	}
	s := 0
	ans := make([]int, 10)
	for k1, _ := range m {
		for k2, v := range m[k1] {
			if k1 >= 2 && k1 <= h-1 && k2 >= 2 && k2 <= w-1 {
				ans[v]++
				s++
			}
		}
	}
	ans[0] = (h - 2) * (w - 2)
	ans[0] -= s

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

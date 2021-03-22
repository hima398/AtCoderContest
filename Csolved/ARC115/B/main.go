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

	const INF = int(1e9) + 1
	n := nextInt()
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
	}
	mi, mj := -1, -1
	minc := INF
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i][j] = nextInt()
			if c[i][j] < minc {
				mi = i
				mj = j
				minc = c[i][j]
			}
		}
	}
	//fmt.Println(mi, mj, minc)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = -INF
		b[i] = -INF
	}
	a[mi] = 0
	b[mj] = 0
	rem := 2 * (n - 1)
	for rem > 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if !(a[i] == -INF) && b[j] == -INF {
					b[j] = c[i][j] - minc - a[i]
					rem--
				}
				if !(b[j] == -INF) && a[i] == -INF {
					a[i] = c[i][j] - minc - b[j]
					rem--
				}
			}
		}
	}
	ok := true
	for i := 0; i < n; i++ {
		b[i] += minc
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ok = ok && (c[i][j] == a[i]+b[j])
		}
	}
	//fmt.Println(a, b)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	if ok {
		fmt.Fprintln(out, "Yes")
		fmt.Fprintf(out, "%d", a[0])
		for i := 1; i < n; i++ {
			fmt.Fprintf(out, " %d", a[i])
		}
		fmt.Fprintln(out, "")
		fmt.Fprintf(out, "%d", b[0])
		for i := 1; i < n; i++ {
			fmt.Fprintf(out, " %d", b[i])
		}
		fmt.Fprintln(out, "")
	} else {
		fmt.Fprintln(out, "No")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

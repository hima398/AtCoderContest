package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type C struct {
	i, v int
}

func SolveCommentary(k, n, m int, a []int) (ans []int) {
	ans = make([]int, k)
	//bd := make([]int, k)
	c := make([]C, k)
	s := 0
	for i := 0; i < k; i++ {
		ans[i] = Floor(m*a[i], n)
		c[i] = C{i, n*ans[i] - m*a[i]}
		s += ans[i]
	}
	diff := m - s
	sort.Slice(c, func(i, j int) bool { return c[i].v < c[j].v })
	for i := 0; i < diff; i++ {
		ans[c[i].i]++
	}
	return ans
}

func SolveBinarySearch(k, n, m int, a []int) (ans []int) {
	check := func(x int) bool {
		l := make([]int, k)
		r := make([]int, k)
		sl, sr := 0, 0
		for i := 0; i < k; i++ {
			l[i] = Max(0, m*a[i]-x+n-1) / n
			r[i] = (m*a[i] + x) / n
			sl += l[i]
			sr += r[i]
		}
		return sl <= m && m <= sr
	}
	ng, ok := -1, n*m
	for ng+1 < ok {
		x := (ng + ok) / 2
		if check(x) {
			ok = x
		} else {
			ng = x
		}
	}
	ans = make([]int, k)
	r := make([]int, k)
	s := 0
	for i := 0; i < k; i++ {
		ans[i] = Max(0, m*a[i]-ok+n-1) / n
		s += ans[i]
		r[i] = (m*a[i] + ok) / n
	}
	for i := 0; i < k; i++ {
		x := Min(m-s, r[i]-ans[i])
		ans[i] += x
		s += x
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k, n, m := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(k)

	ans := SolveCommentary(k, n, m, a)
	//ans := SolveBinarySearch(k, n, m, a)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out, "")
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	return x / y
}

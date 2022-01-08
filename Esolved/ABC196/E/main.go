package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func F(x, a, t int) (ret int) {
	switch t {
	case 1:
		ret = x + a
	case 2:
		ret = Max(x, a)
	case 3:
		ret = Min(x, a)
	}
	return ret
}

func SolveHonestly(n, q int, a, t, x []int) (ans []int) {

	for _, xi := range x {
		fi := xi
		for i := range a {
			fi = F(fi, a[i], t[i])
		}
		ans = append(ans, fi)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = int(1e14) + 1
	n := nextInt()
	a := make([]int, n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], t[i] = nextInt(), nextInt()
	}
	q := nextInt()
	x := nextIntSlice(q)

	l, h := -INF, INF
	add := 0
	for i := 0; i < n; i++ {
		switch t[i] {
		case 1:
			l += a[i]
			h += a[i]
			add += a[i]
		case 2:
			l = Max(l, a[i])
			h = Max(h, a[i])
		case 3:
			l = Min(l, a[i])
			h = Min(h, a[i])
		}
	}

	var ans []int
	for _, xi := range x {
		v := Min(h, Max(l, xi+add))
		ans = append(ans, v)
	}

	//ans := SolveHonestly(n, q, a, t, x)

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

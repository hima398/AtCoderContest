package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = int(1e12)
	a, b, q := nextInt(), nextInt(), nextInt()
	s := make([]int, a+2)
	t := make([]int, b+2)
	s[0] = -INF
	for i := 1; i <= a; i++ {
		s[i] = nextInt()
	}
	s[a+1] = INF
	t[0] = -INF
	for i := 1; i <= b; i++ {
		t[i] = nextInt()
	}
	t[b+1] = INF
	//制約の中にs, tは同じ値がなく昇順とあるので、この問題ではソート不要
	//sort.Ints(s)
	//sort.Ints(t)
	for i := 0; i < q; i++ {
		x := nextInt()
		si := sort.Search(len(s), func(idx int) bool {
			return x <= s[idx]
		})
		ti := sort.Search(len(t), func(idx int) bool {
			return x <= t[idx]
		})
		// 左右に動く場合
		ds := Min(x-s[si-1], s[si]-x)
		dt := Min(x-t[ti-1], t[ti]-x)
		lr := Min(2*ds+dt, ds+2*dt)
		// 左にだけ進む場合
		l := Max(x-s[si-1], x-t[ti-1])
		// 右にだけ進む場合
		r := Max(s[si]-x, t[ti]-x)

		ans := Min(Min(l, r), lr)
		fmt.Println(ans)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

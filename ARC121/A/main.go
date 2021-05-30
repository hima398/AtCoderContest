package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type House struct {
	i, x, y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	xs := make([]House, n)
	ys := make([]House, n)
	for i := range xs {
		xs[i].i = i
		xs[i].x, xs[i].y = nextInt(), nextInt()
	}
	copy(ys, xs)
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].x < xs[j].x
	})
	sort.Slice(ys, func(i, j int) bool {
		return ys[i].y < ys[j].y
	})
	//fmt.Println(xs)
	//fmt.Println(ys)
	var s []int
	if xs[0].i == ys[0].i && xs[n-1].i == ys[n-1].i {
		s = append(s, Max(xs[n-1].x-xs[0].x, ys[n-1].y-ys[0].y))
		s = append(s, Max(xs[n-1].x-xs[1].x, ys[n-1].y-ys[1].y))
		s = append(s, Max(xs[n-2].x-xs[0].x, ys[n-2].y-ys[0].y))
	} else {
		s = append(s, xs[n-1].x-xs[0].x)
		s = append(s, ys[n-1].y-ys[0].y)
		s = append(s, Max(xs[n-1].x-xs[1].x, ys[n-1].y-ys[1].y))
		s = append(s, Max(xs[n-2].x-xs[0].x, ys[n-2].y-ys[0].y))
	}
	sort.Ints(s)
	//fmt.Println(s)
	ans := s[len(s)-2]

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

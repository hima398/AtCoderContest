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
	n, a, b := nextInt(), nextInt(), nextInt()
	s := nextIntSlice(n)
	min, max := INF, 0

	for _, si := range s {
		min, max = Min(min, si), Max(max, si)
	}
	// sの最大値と最小値の差
	diff := max - min
	// 後からdiffの値で割るため、0であれば計算が不能
	// つまり、maxとminの差が0である同じ値ばかりの配列の場合は計算不可能
	if diff <= 0 {
		fmt.Println(-1)
		return
	}
	p := float64(b) / float64(diff)
	ss := 0.0
	for _, si := range s {
		ss += p * float64(si)
	}
	// aとsの平均の差
	q := float64(a) - ss/float64(n)
	fmt.Println(p, q)
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

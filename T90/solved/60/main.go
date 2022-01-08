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

	n := nextInt()
	a := nextIntSlice(n)
	f := make([]int, n) //前方からi番目の要素を使う最長増加部分列の長さ
	var fv []int
	for i, v := range a {
		idx := sort.Search(len(fv), func(j int) bool {
			return v <= fv[j]
		})
		if idx >= len(fv) {
			fv = append(fv, v)
		} else {
			fv[idx] = v
		}
		f[i] = idx + 1
	}
	//fmt.Println(fv)
	//fmt.Println(f)

	ra := make([]int, n)
	for i, v := range a {
		ra[n-i-1] = v
	}
	r := make([]int, n) //後方からi番目の要素を使う最長部分列の長さ
	var rv []int
	for i, v := range ra {
		idx := sort.Search(len(rv), func(j int) bool {
			return v <= rv[j]
		})
		if idx >= len(rv) {
			rv = append(rv, v)
		} else {
			rv[idx] = v
		}
		r[i] = idx + 1
	}
	//fmt.Println(r)
	ans := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, f[i]+r[n-i-1]-1)
	}
	fmt.Println(ans)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

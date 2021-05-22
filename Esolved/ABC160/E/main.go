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

	x, y, a, b, c := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	p := nextIntSlice(a)
	sort.Slice(p, func(i, j int) bool { return p[i] > p[j] })
	p = p[:x]
	q := nextIntSlice(b)
	sort.Slice(q, func(i, j int) bool { return q[i] > q[j] })
	q = q[:y]
	r := nextIntSlice(c)

	apples := append(p, q...)
	apples = append(apples, r...)
	sort.Slice(apples, func(i, j int) bool {
		return apples[i] > apples[j]
	})
	//fmt.Println(apples)
	ans := 0
	for i := 0; i < x+y; i++ {
		ans += apples[i]
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

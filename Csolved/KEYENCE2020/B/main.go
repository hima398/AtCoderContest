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

	type Robot struct {
		l, r int
	}
	n := nextInt()
	rs := make([]Robot, n)
	for i := 0; i < n; i++ {
		x, l := nextInt(), nextInt()
		rs[i].l, rs[i].r = x-l, x+l
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].r < rs[j].r
	})
	cur := -(int(1e9) + 1)
	ans := 0
	for i := 0; i < n; i++ {
		if cur <= rs[i].l {
			ans++
			cur = rs[i].r
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

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

	n, l := nextInt(), nextInt()
	k := nextInt()
	a := nextIntSlice(n)
	ok, ng := 0, l
	check := func(x int) bool {
		pre := 0
		idx := 0
		for i := 0; i < n; i++ {
			if a[i]-pre >= x && l-a[i] >= x {
				idx++
				pre = a[i]
			}
		}
		return idx >= k
	}
	for ng-ok > 1 {
		c := (ng + ok) / 2
		//fmt.Printf("ok = %d, ng = %d, check() = %b\n", ok, ng, check(c))
		if check(c) {
			ok = c
		} else {
			ng = c
		}
	}
	fmt.Println(ok)
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

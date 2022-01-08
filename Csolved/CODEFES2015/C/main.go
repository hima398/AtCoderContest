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

	n := nextInt()
	a := nextIntSlice(n)
	ans := 0
	for len(a) > 1 {
		na := len(a)
		l := 2*a[0] + a[1] + 1
		r := 2*a[na-1] + a[na-2] + 1
		if l <= r {
			ans += l
			v := a[0] + a[1] + 2
			a = a[2:]
			a[0] += v
		} else {
			ans += r
			v := a[na-1] + a[na-2] + 2
			a = a[:na-2]
			a[len(a)-1] += v
		}
		//fmt.Println(a)
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

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
	bs := make([]int, n)

	for i := 0; i < n; i++ {
		bs[i] = a[i] - (i + 1)
	}
	sort.Ints(bs)

	var b int
	if n%2 == 0 {
		b = (bs[n/2] + bs[n/2-1]) / 2
	} else {
		b = bs[n/2]
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += Abs(a[i] - (b + i + 1))
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

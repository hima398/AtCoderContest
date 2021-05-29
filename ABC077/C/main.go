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
	b := nextIntSlice(n)
	c := nextIntSlice(n)
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)
	ans := 0
	for _, v := range b {
		i := sort.Search(n, func(i int) bool {
			return v <= a[i]
		})

		j := sort.Search(n, func(j int) bool {
			return v < c[j]
		})

		//fmt.Println(i, j)
		ans += i * (n - j)
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

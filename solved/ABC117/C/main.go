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

	n, m := nextInt(), nextInt()
	if m == 1 || n >= m {
		fmt.Println(0)
		return
	}
	x := nextIntSlice(m)
	sort.Ints(x)
	var d []int
	for i := 0; i < m-1; i++ {
		d = append(d, x[i+1]-x[i])
	}
	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})
	//fmt.Println(d)
	ans := 0
	for i := n - 1; i < len(d); i++ {
		ans += d[i]
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

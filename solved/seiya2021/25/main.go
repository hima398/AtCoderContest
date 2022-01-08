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
	m := make(map[int][]int)
	for i, v := range a {
		m[v] = append(m[v], i)
	}
	ans := 0
	//fmt.Println(m)
	for i := 0; i < n-1; i++ {
		nn := len(m[a[i]])
		idx := sort.Search(nn, func(j int) bool {
			return i <= m[a[i]][j]
		})
		v := n - i - 1 - (nn - idx - 1)
		//fmt.Println(idx, v)
		ans += v
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

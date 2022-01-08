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
	x, y := nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(m)
	i := 0
	j := 0
	ans := 0
	for {
		j = sort.Search(m, func(idx int) bool {
			return a[i]+x <= b[idx]
		})
		//fmt.Println(j)
		if j >= m {
			break
		}
		i = sort.Search(n, func(idx int) bool {
			return b[j]+y <= a[idx]
		})
		ans++
		if i >= n {
			break
		}
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

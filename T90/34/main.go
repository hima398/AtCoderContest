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

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	m := make(map[int]int)
	ans := 1
	m[a[0]]++
	r := 1
	for l := 0; l < n; l++ {
		//fmt.Println(m)
		for r < n && len(m) <= k {
			if _, ok := m[a[r]]; !ok && len(m) == k {
				break
			}
			m[a[r]]++
			r++
		}
		ans = Max(ans, r-l)
		if r == n {
			break
		}
		m[a[l]]--
		if m[a[l]] == 0 {
			delete(m, a[l])
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

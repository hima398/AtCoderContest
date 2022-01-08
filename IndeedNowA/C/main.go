package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, q int, s, k []int) []int {
	max := 0
	for i := 0; i < n; i++ {
		max = Max(max, s[i])
	}
	if max == 0 {
		return []int{0}
	}
	ms := make([]int, max+1)
	for i := 0; i < n; i++ {
		ms[s[i]]++
	}
	ss := make([]int, max)
	ss[0] = ms[max]
	for i := 1; i < max; i++ {
		ss[i] = ss[i-1] + ms[max-i]
	}

	var ans []int
	for j := 0; j < q; j++ {
		idx := sort.Search(len(ss), func(i int) bool {
			return k[j] < ss[i]
		})
		if idx != max {
			ans = append(ans, max-idx+1)
		} else {
			ans = append(ans, 0)
		}
	}

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextIntSlice(n)
	q := nextInt()
	k := nextIntSlice(q)

	ans := Solve(n, q, s, k)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ans[i])
	}
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

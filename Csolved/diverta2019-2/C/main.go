package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Answer struct {
	x, y int
}

func Solve(n int, a []int) (int, []Answer) {
	var p, m []int
	for i := 0; i < n; i++ {
		if a[i] >= 0 {
			p = append(p, a[i])
		} else {
			m = append(m, a[i])
		}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})
	sort.Ints(m)
	sort.Ints(a)
	if len(m) == 0 {
		np := len(p)
		m = append(m, p[np-1])
		p = p[:np-1]
	}
	if len(p) == 0 {
		nm := len(m)
		p = append(p, m[nm-1])
		m = m[:nm-1]
	}
	var ans []Answer
	cur := m[0]
	for i := 0; i < len(p)-1; i++ {
		ans = append(ans, Answer{cur, p[i]})
		cur -= p[i]
	}
	ans = append(ans, Answer{p[len(p)-1], cur})
	cur = p[len(p)-1] - cur
	for i := 1; i < len(m); i++ {
		ans = append(ans, Answer{cur, m[i]})
		cur -= m[i]
	}
	return cur, ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	m, ans := Solve(n, a)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintln(out, m)
	for _, v := range ans {
		fmt.Fprintln(out, v.x, v.y)
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

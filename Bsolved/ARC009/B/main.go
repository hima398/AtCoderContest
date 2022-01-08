package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pair struct {
	x, y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	m := make(map[int]int)
	b := nextIntSlice(10)
	for i, bi := range b {
		m[bi] = i
	}
	Convert := func(x int) (y int) {
		w := 1
		for x > 0 {
			y += m[(x%10)] * w
			x /= 10
			w *= 10
		}
		return y
	}
	n := nextInt()
	a := make([]Pair, n)
	for i := 0; i < n; i++ {
		a[i].x = nextInt()
		a[i].y = Convert(a[i].x)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].y < a[j].y
	})
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, a[i].x)
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

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
	a, b int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	q := nextInt()
	bi := nextIntSlice(q)
	var b []Pair
	for i := 0; i < q; i++ {
		b = append(b, Pair{i, bi[i]})
	}
	sort.Ints(a)
	sort.Slice(b, func(i, j int) bool { return b[i].b < b[j].b })
	var ans []Pair
	j := 0
	k := 0
	for i := 0; i < q; i++ {
		if j == n-1 {
			ans = append(ans, Pair{b[k].a, Abs(a[j] - b[k].b)})
			k++
			continue
		}
		for j < n-1 && Abs(a[j]-b[k].b) >= Abs(a[j+1]-b[k].b) {
			j++
		}
		ans = append(ans, Pair{b[k].a, Abs(a[j] - b[k].b)})
		k++
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i].a < ans[j].a })
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := range ans {
		fmt.Fprintln(out, ans[i].b)
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

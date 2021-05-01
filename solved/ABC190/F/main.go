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

	n := nextInt()
	a := nextIntSlice(n)
	f := NewFenwick(n)
	ans := make([]int, n)
	for i, v := range a {
		f.Add(v, 1)
		ans[0] += i + 1 - f.Sum(v)
	}
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + (n - 1 - 2*a[i-1])
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
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

type Fenwick struct {
	data []int
}

func NewFenwick(size int) *Fenwick {
	var f Fenwick
	f.data = make([]int, size)
	return &f
}

func (this *Fenwick) Add(idx, v int) {
	for i := idx; i < len(this.data); i = i | (i + 1) {
		this.data[i] += v
	}
}

func (this *Fenwick) Sum(idx int) int {
	ret := 0
	for i := idx; i >= 0; i = (i & (i + 1)) - 1 {
		ret += this.data[i]
	}
	return ret
}

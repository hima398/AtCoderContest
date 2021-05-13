package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Brick struct {
	l, r int
}

type SegmentTree struct {
	size  int
	nodes []int
	lazy  []int
}

func NewSegmentTree(n int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([]int, 2*st.size)
	st.lazy = make([]int, 2*st.size)
	return st
}

// k番目のノードの遅延評価を行う
func (this *SegmentTree) Eval(k int) {
	if k < this.size {
		this.lazy[k*2] = Max(this.lazy[k*2], this.lazy[k])
		this.lazy[k*2+1] = Max(this.lazy[k*2+1], this.lazy[k])
	}
	this.nodes[k] = Max(this.nodes[k], this.lazy[k])
	this.lazy[k] = 0
}

func (this *SegmentTree) UpdateRecursively(a, b, x, k, l, r int) {
	this.Eval(k)
	if a >= r || b <= l {
		return
	}
	if a <= l && b >= r {
		this.lazy[k] = x
		this.Eval(k)
		return
	}
	this.UpdateRecursively(a, b, x, 2*k, l, (l+r)/2)
	this.UpdateRecursively(a, b, x, 2*k+1, (l+r)/2, r)
	this.nodes[k] = Max(this.nodes[2*k], this.nodes[2*k+1])
}

func (this *SegmentTree) Update(l, r, x int) {
	this.UpdateRecursively(l, r, x, 1, 0, this.size)
}

func (this *SegmentTree) QueryRecursively(a, b, k, l, r int) int {
	this.Eval(k)
	if a >= r || b <= l {
		return -1
	}
	if a <= l && b >= r {
		return this.nodes[k]
	}
	lv := this.QueryRecursively(a, b, 2*k, l, (l+r)/2)
	rv := this.QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return Max(lv, rv)
}

func (this *SegmentTree) Query(l, r int) int {
	return this.QueryRecursively(l, r, 1, 0, this.size)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	w, n := nextInt(), nextInt()
	st := NewSegmentTree(w)
	bricks := make([]Brick, n)
	for i := 0; i < n; i++ {
		l, r := nextInt(), nextInt()
		bricks[i] = Brick{l, r}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, b := range bricks {
		h := st.Query(b.l-1, b.r) + 1
		st.Update(b.l-1, b.r, h)
		fmt.Fprintln(out, h)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

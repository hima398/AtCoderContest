package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type SegmentTree struct {
	size  int
	nodes []int
}

func NewSegmentTree(n int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([]int, 2*st.size)
	return st
}

func (this *SegmentTree) QueryRecursively(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return 0
	}

	// [a, b)が[l, r)を完全に含んでいる
	if a <= l && b >= r {
		return this.nodes[k]
	}

	vl := this.QueryRecursively(a, b, 2*k, l, (l+r)/2)
	vr := this.QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return vl ^ vr
}

// [l, r)の区間の値をxorした結果を返す
func (this *SegmentTree) Query(l, r int) int {
	return this.QueryRecursively(l, r, 1, 0, this.size)
}

func (this *SegmentTree) Update(k, x int) {
	k += this.size
	this.nodes[k] = this.nodes[k] ^ x
	for k > 1 {
		k /= 2
		this.nodes[k] = this.nodes[k*2] ^ this.nodes[k*2+1]
	}
	//fmt.Println(this.nodes)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	st := NewSegmentTree(n)
	//fmt.Println(st.size)
	for i, v := range a {
		st.Update(i, v)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		t, x, y := nextInt(), nextInt(), nextInt()
		if t == 1 {
			x--
			st.Update(x, y)
		} else {
			// t == 2
			x-- // 0-indexedのためデクリメント
			fmt.Fprintln(out, st.Query(x, y))
			//fmt.Println(st.Query(x, y))
		}
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

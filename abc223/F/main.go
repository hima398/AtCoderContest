package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

//正しい括弧にするために左に必要な"("と右に必要な")"の数
type Node struct {
	l, r int
}

type SegmentTree struct {
	size  int
	nodes []Node
}

func NewSegmentTree(n int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([]Node, 2*st.size)
	for i := range st.nodes {
		st.nodes[i] = Node{0, 0}
	}
	return st
}

// [l, r)の区間の値をxorした結果を返す
func (st *SegmentTree) Query(l, r int) Node {
	var QueryRecursively func(a, b, k, l, r int) Node
	QueryRecursively = func(a, b, k, l, r int) Node {
		// [a, b)と[l, r)が交差しない
		if a >= r || b <= l {
			return Node{0, 0}
		}

		// [a, b)が[l, r)を完全に含んでいる
		if a <= l && b >= r {
			return st.nodes[k]
		}

		vl := QueryRecursively(a, b, 2*k, l, (l+r)/2)
		vr := QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
		return Node{vl.l + Max(vr.l-vl.r, 0), vr.r + Max(vl.r-vr.l, 0)}
	}

	return QueryRecursively(l, r, 1, 0, st.size)
}

func (st *SegmentTree) Update(k int, node Node) {
	k += st.size
	st.nodes[k] = node
	for k > 1 {
		k /= 2
		vl := st.nodes[2*k]
		vr := st.nodes[2*k+1]
		st.nodes[k] = Node{vl.l + Max(vr.l-vl.r, 0), vr.r + Max(vl.r-vr.l, 0)}
	}
	//fmt.Println(this.nodes)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	s := nextString()
	ns := make([]Node, n)
	st := NewSegmentTree(n)
	for i := range s {
		if s[i] == '(' {
			ns[i].r = 1
		} else {
			ns[i].l = 1
		}
		st.Update(i, ns[i])
	}
	//fmt.Println(st.nodes)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < q; i++ {
		t, l, r := nextInt(), nextInt(), nextInt()
		l--
		r--
		switch t {
		case 1:
			ns[l], ns[r] = ns[r], ns[l]
			st.Update(l, ns[l])
			st.Update(r, ns[r])
		case 2:
			//Queryメソッドは区間[l, r)を想定しているのでrをインクリメントし直す
			node := st.Query(l, r+1)
			//fmt.Println(node)
			if node.l == 0 && node.r == 0 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

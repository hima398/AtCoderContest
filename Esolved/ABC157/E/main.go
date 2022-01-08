package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type SegmentTree struct {
	n    int
	size int
	s    []string
	//nodes []map[string]int
	nodes []int
}

func NewSegmentTree(n int, s string) *SegmentTree {
	st := new(SegmentTree)
	st.n = n
	st.size = 1
	for st.size < st.n {
		st.size *= 2
	}
	st.s = strings.Split(s, "")
	st.nodes = make([]int, 2*st.size)
	//for i := 0; i < 2*st.size; i++ {
	//	st.nodes[i] = 0
	//}
	for i := 0; i < st.n; i++ {
		idx := i + st.size
		st.nodes[idx] |= 1 << (st.s[i][0] - 'a')
	}
	for i := st.size - 1; i >= 1; i-- {
		st.nodes[i] |= st.nodes[2*i]
		st.nodes[i] |= st.nodes[2*i+1]
	}
	return st
}

func (st *SegmentTree) QueryRecursively(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return 0
	}

	// [a, b)が[l, r)を完全に含んでいる
	if a <= l && b >= r {
		return st.nodes[k]
	}

	ml := st.QueryRecursively(a, b, 2*k, l, (l+r)/2)
	mr := st.QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return ml | mr
}

// [l, r)の区間の文字の種類を返す
func (st *SegmentTree) Query(l, r int) int {
	m := st.QueryRecursively(l, r, 1, 0, st.size)
	return bits.OnesCount(uint(m))
}

//i番目(0-indexed)の文字列をxに更新する
func (st *SegmentTree) Update(i int, x string) {
	//ox := st.s[i]
	//st.s[i] = x
	idx := i + st.size
	st.nodes[idx] = 1 << (x[0] - 'a')
	for idx > 1 {
		idx /= 2
		st.nodes[idx] = st.nodes[2*idx] | st.nodes[2*idx+1]
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	q := nextInt()
	st := NewSegmentTree(n, s)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		t := nextInt()
		switch t {
		case 1:
			i, c := nextInt(), nextString()
			i--
			st.Update(i, c)
			//fmt.Println(st.nodes)
		case 2:
			l, r := nextInt(), nextInt()
			l--
			ans := st.Query(l, r)
			//fmt.Println(ans)
			fmt.Fprintln(out, ans)
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

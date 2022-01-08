package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type People struct {
	i, v int
}
type Edge struct {
	s, t int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	// yの最大値
	const mxy = 200000

	es := make([][]Edge, mxy+1)
	ps := make([][]People, mxy+1)

	n, m := nextInt(), nextInt()
	for i := 0; i < m; i++ {
		a, b, y := nextInt(), nextInt(), nextInt()
		es[y] = append(es[y], Edge{a, b})
	}
	q := nextInt()
	for i := 0; i < q; i++ {
		v, w := nextInt(), nextInt()
		ps[w] = append(ps[w], People{i, v})
	}
	uf := NewUnionFind(n)
	ans := make([]int, q)
	for y := mxy; y >= 0; y-- {
		//y年以前の道を使わない人が行ける都市の数
		for _, p := range ps[y] {
			ans[p.i] = uf.Size(p.v)
		}
		//y年にできた道をUnionFindで繋ぐ
		for _, e := range es[y] {
			uf.Unite(e.s, e.t)
		}
	}
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

type UnionFind struct {
	par  []int // parent numbers
	rank []int // height of tree
	size []int
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.par = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.par[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (this *UnionFind) Find(x int) int {
	if this.par[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		this.par[x] = this.Find(this.par[x])
		return this.par[x]
	}
}

func (this *UnionFind) Size(x int) int {
	return this.size[this.Find(x)]
}

func (this *UnionFind) ExistSameUnion(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Unite(x, y int) {
	x = this.Find(x)
	y = this.Find(y)
	if x == y {
		return
	}
	// rank
	if this.rank[x] < this.rank[y] {
		//yがrootの木にxがrootの木を結合する
		this.par[x] = y
		this.size[y] += this.size[x]
	} else {
		// this.rank[x] >= this.rank[y]
		//xがrootの木にyがrootの木を結合する
		this.par[y] = x
		this.size[x] += this.size[y]
		if this.rank[x] == this.rank[y] {
			this.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.par)
	fmt.Println(u.rank)
	fmt.Println(u.size)
}

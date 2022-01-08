package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(n int, x, y []int) int {
	type Pos struct {
		i, x int
	}
	var xs, ys []Pos
	for i := 0; i < n; i++ {
		xs = append(xs, Pos{i, x[i]})
		ys = append(ys, Pos{i, y[i]})
	}
	Sort := func(s []Pos) {
		sort.Slice(s, func(i, j int) bool {
			return s[i].x < s[j].x
		})
	}
	Sort(xs)
	Sort(ys)
	type Node struct {
		s, t, d int
	}
	var ns []Node
	for i := 0; i < n-1; i++ {
		ns = append(ns, Node{xs[i].i, xs[i+1].i, xs[i+1].x - xs[i].x})
		ns = append(ns, Node{ys[i].i, ys[i+1].i, ys[i+1].x - ys[i].x})
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].d < ns[j].d
	})
	ans := 0
	uf := NewUnionFind(n)
	root := ns[0].s
	for _, node := range ns {
		if !uf.ExistSameUnion(node.s, node.t) {
			ans += node.d
			uf.Unite(node.s, node.t)
		}
		if uf.Size(root) >= n {
			break
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := FirstSolve(n, x, y)
	fmt.Println(ans)
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

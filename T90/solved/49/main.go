package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, m int, c, l, r []int) (ans int) {
	type Node struct {
		c, l, r int
	}
	nodes := make([]Node, m)
	for i := 0; i < m; i++ {
		nodes[i] = Node{c[i], l[i], r[i]}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].c < nodes[j].c
	})
	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		l, r := nodes[i].l-1, nodes[i].r
		if !uf.ExistSameUnion(l, r) {
			uf.Unite(l, r)
			ans += nodes[i].c
		}
	}
	for i := 0; i < n; i++ {
		// 0〜nまでのうち連結でないノードが存在する場合
		if !uf.ExistSameUnion(i, i+1) {
			return -1
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	c, l, r := make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		c[i], l[i], r[i] = nextInt(), nextInt(), nextInt()
	}
	ans := Solve(n, m, c, l, r)
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

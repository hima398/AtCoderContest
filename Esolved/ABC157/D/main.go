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

	n, m, k := nextInt(), nextInt(), nextInt()
	uf := NewUnionFind(n)
	friends := make(map[int][]int)
	blocks := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		friends[a] = append(friends[a], b)
		friends[b] = append(friends[b], a)
		uf.Unite(a, b)
	}
	for i := 0; i < k; i++ {
		c, d := nextInt(), nextInt()
		c--
		d--
		blocks[c] = append(blocks[c], d)
		blocks[d] = append(blocks[d], c)
	}
	//PrintUnionFind(uf)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = uf.GetUnionSize(i) - 1 //同じグループの中から自身を引く
		ans[i] -= len(friends[i])       //直近の友人の数を引く
		for _, v := range blocks[i] {
			if uf.ExistSameUnion(i, v) {
				ans[i]-- //同じグループにいるブロック関係を引く
			}
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out, "")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type UnionFind struct {
	par  []int // parent numbers
	rank []int // height of tree
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.par = make([]int, n+1)
	u.rank = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.par[i] = -1
		u.rank[i] = 0
	}
	return u
}

func (this *UnionFind) Find(x int) int {
	if this.par[x] < 0 {
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

func (this *UnionFind) ExistSameUnion(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Unite(x, y int) {
	x = this.Find(x)
	y = this.Find(y)
	if x == y {
		return
	}
	// raink
	if this.rank[x] < this.rank[y] {
		this.par[y] += this.par[x]
		this.par[x] = y
	} else {
		// this.rank[x] >= this.rank[y]
		this.par[x] += this.par[y]
		this.par[y] = x
		if this.rank[x] == this.rank[y] {
			this.rank[x]++
		}
	}
}

func (this *UnionFind) GetUnionSize(x int) int {
	root := this.Find(x)
	return -this.par[root]
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.par)
	fmt.Println(u.rank)
}

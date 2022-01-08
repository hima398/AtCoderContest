package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveUnionFind() {
	l, q := nextInt(), nextInt()
	c, x := make([]int, q), make([]int, q)

	sl := []int{0, l}
	for i := 0; i < q; i++ {
		c[i], x[i] = nextInt(), nextInt()
		if c[i] == 1 {
			sl = append(sl, x[i])
		}
	}
	n := len(sl)
	sort.Ints(sl)
	//fmt.Println(sl)

	var ans []int
	uf := NewUnionFind(sl)
	//PrintUnionFind(uf)
	for i := q - 1; i >= 0; i-- {
		idx := sort.Search(n, func(j int) bool {
			return x[i] < sl[j]
		})
		switch c[i] {
		case 1:
			uf.Unite(idx-1, idx)
		case 2:
			ans = append(ans, uf.Size(idx))
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintln(out, ans[i])
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	SolveUnionFind()
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

// この問題用にスライスで初期化をする
func NewUnionFind(s []int) *UnionFind {
	if len(s) <= 0 {
		return nil
	}
	u := new(UnionFind)
	n := len(s)
	// for accessing index without minus 1
	u.par = make([]int, n)
	u.rank = make([]int, n)
	u.size = make([]int, n)
	for i := 0; i < n; i++ {
		u.par[i] = i
		u.rank[i] = 0
		//u.size[i] = 1
	}
	for i := 0; i < n-1; i++ {
		u.size[i+1] = s[i+1] - s[i]
	}
	return u
}

func (uf *UnionFind) InitSize(s []int) {
	for i := 1; i < len(uf.size); i++ {
		uf.size[i] = s[i] - s[i-1]
	}
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

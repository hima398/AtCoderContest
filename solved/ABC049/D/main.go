package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pair struct {
	a, b int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k, l := nextInt(), nextInt(), nextInt()
	road := NewUnionFind(n)
	for i := 0; i < k; i++ {
		p, q := nextInt(), nextInt()
		road.Unite(p, q)
	}
	line := NewUnionFind(n)
	for i := 0; i < l; i++ {
		r, s := nextInt(), nextInt()
		line.Unite(r, s)
	}
	//PrintUnionFind(road)
	//PrintUnionFind(line)
	m := make(map[Pair]int)
	for i := 1; i <= n; i++ {
		m[Pair{road.Find(i), line.Find(i)}]++
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "%d", m[Pair{road.Find(1), line.Find(1)}])
	for i := 2; i <= n; i++ {
		fmt.Fprintf(out, " %d", m[Pair{road.Find(i), line.Find(i)}])
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
		u.par[i] = i
		u.rank[i] = 0
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
		this.par[x] = y
	} else {
		// this.rank[x] >= this.rank[y]
		this.par[y] = x
		if this.rank[x] == this.rank[y] {
			this.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.par)
	fmt.Println(u.rank)
}

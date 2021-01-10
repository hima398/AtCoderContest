package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	u := NewUnionFind(n)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		u.Unite(a, b)
	}
	groups := make(map[int]int)
	for i := 1; i <= n; i++ {
		groups[u.Find(i)]++
	}
	ans := 0
	for _, v := range groups {
		ans = Max(ans, v)
	}
	//PrintUnionFind(u)
	fmt.Println(ans)
}

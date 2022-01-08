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

	n, m := nextInt(), nextInt()
	ans := []int{(n - 1) * n / 2}
	a, b := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = nextInt(), nextInt()
		a[i]--
		b[i]--
	}
	uf := NewUnionFind(n)

	for i := m - 1; i > 0; i-- {
		if uf.Find(a[i]) == uf.Find(b[i]) {
			ans = append(ans, ans[len(ans)-1])
		} else {
			w := uf.Num(a[i]) * uf.Num(b[i])
			v := ans[len(ans)-1] - w
			ans = append(ans, v)
			uf.Unite(a[i], b[i])
		}
	}
	//fmt.Println(ans)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := m - 1; i >= 0; i-- {
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
	num  []int
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.par = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.num = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.par[i] = i
		u.rank[i] = 0
		u.num[i] = 1
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

//xを含む連携成分の個数を返す
func (this *UnionFind) Num(x int) int {
	return this.num[this.Find(x)]
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
		this.num[y] += this.num[x]
	} else {
		// this.rank[x] >= this.rank[y]
		this.par[y] = x
		this.num[x] += this.num[y]
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n int, a []int) int {
	m := make(map[int]int)
	/*
		for i := range a {
			m[a[i]] = a[i]
		}
	*/
	ret := 0
	for i := 0; i < n/2; i++ {
		x, y := a[i], a[n-i-1]
		//fmt.Println(x, y)
		bufx := []int{x}
		for {
			if _, ok := m[x]; !ok {
				break
			}
			x = m[x]
			bufx = append(bufx, x)
		}
		bufy := []int{y}
		for {
			if _, ok := m[y]; !ok {
				break
			}
			y = m[y]
			bufy = append(bufy, y)
		}
		//fmt.Println(buf)
		if x != y {
			if len(bufx) < len(bufy) {
				for _, v := range bufx {
					m[v] = y
				}
			} else {
				//len(bufx) > len(bufy)
				for _, v := range bufy {
					m[v] = x
				}
			}
			ret++
		}
	}
	//fmt.Println(m)
	return ret
}

func FirstSolve(n int, a []int) (ans int) {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[n-i-1]
	}
	//ans1 := Solve(n, a)
	//fmt.Println(ans1)
	ans2 := Solve(n, b)
	fmt.Println(ans2)
	//fmt.Println(Min(ans1, ans2))
	return ans2
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	max := 0
	for _, v := range a {
		max = Max(max, v)
	}
	uf := NewUnionFind(max)
	ans := 0
	for i := 0; i < n/2; i++ {
		x := a[i] - 1
		y := a[n-i-1] - 1
		if uf.Find(x) != uf.Find(y) {
			uf.Unite(x, y)
			ans++
		}
	}
	//ans = FirstSolve(n, a)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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

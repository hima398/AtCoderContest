package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n, k int, a [][]int) int {
	const p = 998244353

	//mc := make(map[int]bool)
	ufc := NewUnionFind(n)
	for x := 0; x < n-1; x++ {
		for y := x + 1; y < n; y++ {
			ok := true
			for i := 0; i < n; i++ {
				ok = ok && a[i][x]+a[i][y] <= k
			}
			if ok {
				//mc[x] = true
				//mc[y] = true
				ufc.Unite(x, y)
			}
		}
	}
	//PrintUnionFind(ufc)

	//mr := make(map[int]bool)
	ufr := NewUnionFind(n)
	for x := 0; x < n-1; x++ {
		for y := x + 1; y < n; y++ {
			ok := true
			for j := 0; j < n; j++ {
				ok = ok && a[x][j]+a[y][j] <= k
			}
			if ok {
				//mr[x] = true
				//mr[y] = true
				ufr.Unite(x, y)
			}
		}
	}
	//PrintUnionFind(ufr)

	fac := make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % p
	}
	//fmt.Println(fac)
	mr := make(map[int]int)
	mc := make(map[int]int)
	ans := 1
	for i := 0; i < n; i++ {
		mr[ufr.Find(i)]++
	}
	for i := 0; i < n; i++ {
		mc[ufc.Find(i)]++
	}
	for _, v := range mr {
		ans *= fac[v]
		ans %= p
	}
	for _, v := range mc {
		ans *= fac[v]
		ans %= p
	}
	//for i := 1; i <= nr; i++ {
	//	ans *= i
	//	ans %= p
	//}
	//for j := 1; j <= nc; j++ {
	//	ans *= j
	//	ans %= p
	//}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(n)
	}
	ans := Solve(n, k, a)
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

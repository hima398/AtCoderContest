package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(n, q int, a, b, c, d []int, e map[int][]int) {
	depth := make([]int, n)
	parents := make([][21]int, n)
	var Dfs func(i, p, v int)
	Dfs = func(i, p, v int) {
		depth[i] = v
		parents[i][0] = p
		for _, chi := range e[i] {
			if chi == p {
				continue
			}
			Dfs(chi, i, v+1)
		}
	}
	Dfs(0, -1, 0)
	//fmt.Println(parents)
	for i := 0; i < 20; i++ {
		for j := 0; j < n; j++ {
			if parents[j][i] < 0 {
				parents[j][i+1] = -1
			} else {
				parents[j][i+1] = parents[parents[j][i]][i]
			}
		}
	}
	//fmt.Println(parents)

	Lca := func(u, v int) int {
		if depth[u] > depth[v] {
			u, v = v, u
		}
		for i := 0; i < 21; i++ {
			if ((depth[v]-depth[u])>>i)&1 == 1 {
				v = parents[v][i]
			}
		}
		if u == v {
			return u
		}
		for i := 20; i >= 0; i-- {
			if parents[u][i] != parents[v][i] {
				u = parents[u][i]
				v = parents[v][i]
			}
		}
		return parents[u][0]
	}

	for i := 0; i < q; i++ {
		//fmt.Println(c[i], d[i])
		idx := Lca(c[i], d[i])
		var dis int
		//fmt.Println(idx)
		dis = 2*depth[c[i]] + 2*depth[d[i]] - 4*depth[idx]

		out := bufio.NewWriter(os.Stdout)
		defer out.Flush()
		if (dis/2)%2 == 1 {
			fmt.Fprintln(out, "Road")
		} else {
			fmt.Fprintln(out, "Town")
		}
	}

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	a, b := make([]int, n-1), make([]int, n-1)
	e := make(map[int][]int)
	c, d := make([]int, q), make([]int, q)

	for i := 0; i < n-1; i++ {
		a[i], b[i] = nextInt(), nextInt()
		a[i]--
		b[i]--
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	for i := 0; i < q; i++ {
		c[i], d[i] = nextInt(), nextInt()
		c[i]--
		d[i]--
	}

	depth := make([]int, n)
	var Dfs func(i, p, v int)
	Dfs = func(i, p, v int) {
		depth[i] = v % 2
		for _, chi := range e[i] {
			if chi == p {
				continue
			}
			Dfs(chi, i, v+1)
		}
	}
	Dfs(0, -1, 0)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		if depth[c[i]] != depth[d[i]] {
			fmt.Fprintln(out, "Road")
		} else {
			fmt.Fprintln(out, "Town")
		}
	}
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

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}

type IntQueue struct {
	q []int
}

func NewIntQueue() *IntQueue {

	return new(IntQueue)
}
func (this *IntQueue) Push(v int) {
	this.q = append(this.q, v)
}

func (this *IntQueue) Pop() (int, error) {
	if this.Size() == 0 {
		return -1, errors.New("")
	}
	ret := this.q[0]
	this.q = this.q[1:]
	return ret, nil
}

func (this *IntQueue) Size() int {
	return len(this.q)
}

func (this *IntQueue) PrintQueue() {
	fmt.Println(this.q)
}

type Pos struct {
	X int
	Y int
	D int
}

type Queue struct {
	ps []Pos
}

func NewQueue() *Queue {
	return new(Queue)
}

func (this *Queue) Push(p Pos) {
	this.ps = append(this.ps, p)
}

func (this *Queue) Pop() *Pos {
	if len(this.ps) == 0 {
		return nil
	}
	p := this.ps[0]
	this.ps = this.ps[1:]
	return &p
}

func (this *Queue) Find(x, y int) bool {
	for _, v := range this.ps {
		if x == v.X && y == v.Y {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
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

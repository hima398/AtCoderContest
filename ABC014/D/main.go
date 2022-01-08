package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const MaxNode = int(1e5)
const MaxLogNode = 17 //16.61
var parent [MaxLogNode + 1][MaxNode + 1]int
var depth [MaxNode + 1]int

var e map[int][]int

type Node struct {
	i, d int
}

func bfs(f, t, n int) int {
	q := []Node{Node{f, 0}}
	v := make([]bool, n)
	v[f] = true
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p.i == t {
			return p.d
		}
		for _, n := range e[p.i] {
			if !v[n] {
				q = append(q, Node{n, p.d + 1})
				v[n] = true
			}
		}
	}
	return -1
}

func Dfs(i, par, d int) {
	parent[0][i] = par
	depth[i] = d

	for _, next := range e[i] {
		if next == par {
			continue
		}
		Dfs(next, i, d+1)
	}
}

func BinarySearchLCAInit(v int) {
	Dfs(0, -1, 0)
	for i := 0; i < MaxLogNode; i++ {
		for j := 0; j < v; j++ {
			if parent[i][j] < 0 {
				parent[i+1][j] = -1
			} else {
				parent[i+1][j] = parent[i][parent[i][j]]
			}
		}
	}
}

func BinarySearchLCA(u, v int) int {

	// depth(u) <= depth(v)になるように調整する
	if depth[u] > depth[v] {
		u, v = v, u
	}
	for i := 0; i < MaxLogNode; i++ {
		if ((depth[v]-depth[u])>>i)&1 == 1 {
			v = parent[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := MaxLogNode - 1; i >= 0; i-- {
		if parent[i][u] != parent[i][v] {
			u = parent[i][u]
			v = parent[i][v]
		}
	}
	return parent[0][u]

}

func Solve(n int, x, y []int, q int, a, b []int) []int {
	e = make(map[int][]int)
	for i := 0; i < n-1; i++ {
		// 0-indexed
		x[i]--
		y[i]--
		e[x[i]] = append(e[x[i]], y[i])
		e[y[i]] = append(e[y[i]], x[i])
	}
	for i := 0; i < q; i++ {
		a[i]--
		b[i]--
	}

	BinarySearchLCAInit(n)
	var ans []int
	//for i := 0; i < n; i++ {
	//	fmt.Printf("%d ", depth[i])
	//}
	//fmt.Println()
	for i := 0; i < q; i++ {
		node := BinarySearchLCA(a[i], b[i])
		//fmt.Println(node)
		dist := depth[a[i]] + depth[b[i]] - 2*depth[node]
		ans = append(ans, dist+1)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x, y := make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}
	q := nextInt()
	a, b := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	ans := Solve(n, x, y, q, a, b)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
	if x < y {
		x, y = y, x
	}
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

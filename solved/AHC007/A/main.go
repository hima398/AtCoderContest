package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const N = 400
const M = 1995

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	i, x, y int
}

type Edge struct {
	i, s, t, d int
}

func SoveHonestly(p []Pos, e []Edge) {
	// 結果の出力
	//out := bufio.NewWriter(os.Stdout)
	//defer out.Flush()

	for i := 0; i < M; i++ {
		//l := nextInt()
		nextInt()
		fmt.Println(1)
		//fmt.Fprintln(out, l, 1)

	}
}

func FirstSolve(p []Pos, e []Edge) {
	uf := NewUnionFind(N)
	used := make([]bool, M)
	for i := 0; i < M; i++ {
		if !uf.ExistSameUnion(e[i].s, e[i].t) {
			uf.Unite(e[i].s, e[i].t)
			used[i] = true
		}
	}
	for i := 0; i < M; i++ {
		nextInt()
		if used[i] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].d < pq[j].d
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func ComputeDistance(p1, p2 Pos) int {
	d := (p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y)
	return int(math.Sqrt(float64(d)))
}

func Solve2(p []Pos, e []Edge) {
	ee := make([][]int, N)
	for i := 0; i < M; i++ {
		u, v := e[i].s, e[i].t
		ee[u] = append(ee[u], v)
		ee[v] = append(ee[v], u)
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := 0; i < M; i++ {
		u, v := e[i].s, e[i].t
		d := ComputeDistance(p[u], p[v])
		e[i].i = i
		e[i].d = d
		heap.Push(pq, e[i])
	}
	uf := NewUnionFind(N)
	used := make([]bool, M)
	for pq.Len() > 0 {
		edge := heap.Pop(pq).(Edge)
		if !uf.ExistSameUnion(edge.s, edge.t) {
			uf.Unite(edge.s, edge.t)
			used[edge.i] = true
		}
	}
	for i := 0; i < M; i++ {
		nextInt()
		if used[i] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

func Solve3(p []Pos, e []Edge) {
	const INF = 1 << 60
	dist := make([][]int, N)
	via := make([][]int, N)
	emap := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		via[i] = make([]int, N)
		emap[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dist[i][j] = INF
			emap[i][j] = -1
		}
	}
	for i := 0; i < M; i++ {
		s, t := e[i].s, e[i].t
		dist[s][t] = ComputeDistance(p[s], p[t])
		via[s][t] = s
		emap[s][t] = i
	}
	num := 0
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if emap[i][j] >= 0 {
				num++
			}
		}
	}
	//fmt.Println("Num = ", num)
	// 経由点
	for k := 0; k < N; k++ {
		// 始点
		for i := 0; i < N-1; i++ {
			// 終点
			for j := i + 1; j < N; j++ {
				if i == k || j == k {
					continue
				}
				newDist := dist[i][k] + dist[k][j]
				if dist[i][j] > newDist {
					dist[i][j] = newDist
					via[i][j] = k
					emap[i][j] = -1
				}
			}
		}
	}
	num = 0
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if emap[i][j] >= 0 {
				num++
			}
		}
	}
	//fmt.Println("Num = ", num)

	used := make([]bool, M)
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if emap[i][j] >= 0 {
				used[emap[i][j]] = true
			}
		}
	}
	for i := 0; i < M; i++ {
		nextInt()
		if used[i] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

func Solve4(p []Pos, e []Edge) {
	for i := 0; i < N; i++ {
		p[i].i = i
	}
	for i := 0; i < M; i++ {
		e[i].i = i
	}
	const INF = 1 << 60
	v := make(map[int]Pos)
	//ee := make([][]int, N)
	//for i := 0; i < M; i++ {
	//	s, t := e[i].s, e[i].t
	//	ee[s] = append(ee[s], t)
	//}
	ee := make(map[int]Edge)

	dx := []int{0, 800, 0, 800, 400}
	dy := []int{0, 0, 800, 800, 400}
	var totalDist [5]int
	var used [5][]bool
	for ii := 0; ii < 5; ii++ {
		used[ii] = make([]bool, M)
	}
	for ii := 0; ii < 5; ii++ {
		center := Pos{-1, dx[ii], dy[ii]}
		start := p[0]
		dist := ComputeDistance(center, p[0])
		for i := 1; i < N; i++ {
			newDist := ComputeDistance(center, p[i])
			if dist > newDist {
				dist = newDist
				start = p[i]
			}
		}
		v[start.i] = start

		for len(v) < len(p) {
			dist := INF
			addPoint := -1
			addEdge := -1
			for _, pos := range v {
				for i := 0; i < M; i++ {
					// 始点がpos
					if pos.i != e[i].s && pos.i != e[i].t {
						continue
					}
					// 終点がvにいたら無視
					if _, ok := v[e[i].t]; pos.i == e[i].s && ok {
						continue
					}
					if _, ok := v[e[i].s]; pos.i == e[i].t && ok {
						continue
					}

					//fmt.Println(len(v), i)
					var newDist int
					if pos.i == e[i].s {
						newDist = ComputeDistance(pos, p[e[i].t])
					} else {
						newDist = ComputeDistance(pos, p[e[i].s])
					}
					if dist > newDist {
						dist = newDist
						if pos.i == e[i].s {
							addPoint = e[i].t
						} else {
							addPoint = e[i].s
						}
						addEdge = i
					}
				}
			}
			totalDist[ii] += dist
			v[addPoint] = p[addPoint]
			ee[addEdge] = e[addEdge]
		}
		//used := make([]bool, M)
		//fmt.Println("Len V = ", len(v), "Len EE = ", len(ee))
		for k := range ee {
			used[ii][k] = true
		}
	}
	totalDistMin := INF
	idx := -1
	for ii := 0; ii < 5; ii++ {
		if totalDistMin > totalDist[ii] {
			totalDistMin = totalDist[ii]
			idx = ii
		}
	}
	for i := 0; i < M; i++ {
		nextInt()
		if used[idx][i] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}

}

type Node struct {
	root int
}

func Solve5(p []Pos, e []Edge) {
	const SkipDist = 50
	ee := make([][]int, N)
	for i := 0; i < M; i++ {
		u, v := e[i].s, e[i].t
		ee[u] = append(ee[u], v)
		ee[v] = append(ee[v], u)
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := 0; i < M; i++ {
		u, v := e[i].s, e[i].t
		d := ComputeDistance(p[u], p[v])
		e[i].i = i
		e[i].d = d
		heap.Push(pq, e[i])
	}
	uf := NewUnionFind(N)
	used := make([]bool, M)

	for pq.Len() > 0 {
		edge := heap.Pop(pq).(Edge)
		if edge.d > SkipDist {
			continue
		}
		if !uf.ExistSameUnion(edge.s, edge.t) {
			uf.Unite(edge.s, edge.t)
			used[edge.i] = true
		}
	}
	nodes := make(map[int]Node)
	for i := 0; i < N; i++ {
		root := uf.Find(i)
		if _, ok := nodes[root]; !ok {
			nodes[root] = Node{root}
		}
	}
	fmt.Println("Len Nodes = ", len(nodes))
	/*
		for i := 0; i < M; i++ {
			nextInt()
			if used[i] {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	*/
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	p := make([]Pos, N)
	e := make([]Edge, M)
	//e := make([][]int, N)

	for i := 0; i < N; i++ {
		p[i].x, p[i].y = nextInt(), nextInt()
	}
	for i := 0; i < M; i++ {
		e[i].s, e[i].t = nextInt(), nextInt()
	}

	//SoveHonestly(p, e)
	//FirstSolve(p, e)
	//Solve2(p, e)
	//Solve3(p, e)
	//Solve4(p, e)
	Solve5(p, e)
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
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
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

type Comb struct {
	n, p int
	fac  []int // Factional(i) mod p
	finv []int // 1/Factional(i) mod p
	inv  []int // 1/i mod p
}

func NewCombination(n, p int) *Comb {
	c := new(Comb)
	c.n = n
	c.p = p
	c.fac = make([]int, n+1)
	c.finv = make([]int, n+1)
	c.inv = make([]int, n+1)

	c.fac[0] = 1
	c.fac[1] = 1
	c.finv[0] = 1
	c.finv[1] = 1
	c.inv[1] = 1
	for i := 2; i <= n; i++ {
		c.fac[i] = c.fac[i-1] * i % p
		c.inv[i] = p - c.inv[p%i]*(p/i)%p
		c.finv[i] = c.finv[i-1] * c.inv[i] % p
	}
	return c
}

func (c *Comb) Factional(x int) int {
	return c.fac[x]
}

func (c *Comb) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	ret := c.fac[n] * c.finv[k]
	ret %= c.p
	ret *= c.finv[n-k]
	ret %= c.p
	return ret
}

//重複組み合わせ H
func (c *Comb) DuplicateCombination(n, k int) int {
	return c.Combination(n+k-1, k)
}
func (c *Comb) Inv(x int) int {
	return c.inv[x]
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

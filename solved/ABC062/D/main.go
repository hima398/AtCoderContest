package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type PriorityQueue struct {
	greater bool
	q       []int
}

func NewPriorityQueue(greater bool) *PriorityQueue {
	pq := &PriorityQueue{greater: greater}
	return pq
}
func (this PriorityQueue) Len() int {
	return len(this.q)
}

func (this PriorityQueue) Less(i, j int) bool {
	if this.greater {
		return this.q[i] > this.q[j]
	} else {
		return this.q[i] < this.q[j]
	}
}

func (this PriorityQueue) Swap(i, j int) {
	this.q[i], this.q[j] = this.q[j], this.q[i]
}

func (this PriorityQueue) Top() int {
	return this.q[0]
	//return this.q[this.Len()-1]
}

func (this *PriorityQueue) Push(item interface{}) {
	this.q = append(this.q, item.(int))
}

func (this *PriorityQueue) Pop() interface{} {
	pq := *this
	n := pq.Len()
	item := pq.q[n-1]
	this.q = pq.q[0 : n-1]
	return item
}

func (this PriorityQueue) PrintQueue() {
	fmt.Println(this.q)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 3*int(1e14) + 1
	n := nextInt()
	a := nextIntSlice(3 * n)

	q1 := NewPriorityQueue(false)
	heap.Init(q1)
	sa := make([]int, 2*n+1)
	for i := 0; i < n; i++ {
		sa[i+1] = sa[i] + a[i]
		heap.Push(q1, a[i])
	}
	for i := n; i < n*2; i++ {
		t := q1.Top()
		if a[i] > t {
			sa[i+1] = sa[i] - t + a[i]
			heap.Pop(q1)
			heap.Push(q1, a[i])
		} else {
			sa[i+1] = sa[i]
		}
		//q1.PrintQueue()
	}

	q2 := NewPriorityQueue(true)
	heap.Init(q2)
	sa2 := make([]int, 2*n+1)
	for i := 0; i < n; i++ {
		sa2[i+1] = sa2[i] + a[3*n-i-1]
		heap.Push(q2, a[3*n-i-1])
		//q2.PrintQueue()
	}
	for i := n; i < 2*n; i++ {
		t := q2.Top()
		if a[3*n-i-1] < t {
			sa2[i+1] = sa2[i] - t + a[3*n-i-1]
			heap.Pop(q2)
			heap.Push(q2, a[3*n-i-1])
		} else {
			sa2[i+1] = sa2[i]
		}
		//q2.PrintQueue()
	}
	ans := -INF
	//fmt.Println(sa, sa2)
	for i := n; i <= 2*n; i++ {
		ans = Max(ans, sa[i]-sa2[3*n-i])
	}
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

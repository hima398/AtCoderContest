package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	f int
	t int
	i int
}

type Queue struct {
	q []Item
}

func NewQueue() *Queue {

	return new(Queue)
}
func (this *Queue) Push(v Item) {
	this.q = append(this.q, v)
}

func (this *Queue) Pop() Item {
	if this.Size() == 0 {
		return Item{0, 0, 0}
	}
	ret := this.q[0]
	this.q = this.q[1:]
	return ret
}

func (this *Queue) Size() int {
	return len(this.q)
}

func (this *Queue) PrintQueue() {
	for _, v := range this.q {
		fmt.Println(v)
	}
}

func DeleteElement(n []int, v int) []int {
	var ret []int
	idx := 0
	for idx < len(n) {
		if n[idx] == v {
			break
		}
		idx++
	}
	if idx > len(n) {
		return n
	}
	ret = append(ret, n[:idx]...)
	ret = append(ret, n[idx+1:]...)
	return ret
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	nodes := make(map[int][]Item)
	//e := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		nodes[a] = append(nodes[a], Item{a, b, i})
		nodes[b] = append(nodes[b], Item{b, a, i})
	}
	//fmt.Println(nodes)
	k := 0
	for _, v := range nodes {
		k = Max(k, len(v))
	}
	//fmt.Println(k)
	colors := make([]int, k)
	for i := 0; i < k; i++ {
		colors[i] = i + 1
	}
	ans := make([]int, n-1)
	v := make([]bool, n) // visited
	q := NewQueue()
	q.Push(Item{1, 0, 0})
	v[0] = true
	for {
		// parent
		p := q.Pop()
		cc := make([]int, k)
		copy(cc, colors)
		// nodeに繋がる道に色を塗る
		for _, e := range nodes[p.f] {
			//fmt.Println(p.f, e, cc)
			if ans[e.i] != 0 {
				cc = DeleteElement(cc, ans[e.i])
			} else {
				ans[e.i] = cc[0]
				cc = DeleteElement(cc, cc[0])
			}
		}
		// 次のnodeをキューに積む
		for _, e := range nodes[p.f] {
			if !v[e.t-1] {
				q.Push(Item{e.t, 0, 0})
			}
			v[e.t-1] = true
		}

		if q.Size() == 0 {
			break
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintln(out, k)
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

type Pos struct {
	X int
	Y int
	D int
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

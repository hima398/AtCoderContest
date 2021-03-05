package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func IntPow(x, y int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x
		}
		y >>= 1
		x = x * x
	}
	return ret
}

func CountBits(n int) int {
	ret := 0
	for i := 1; i <= 64; i <<= 1 {
		if n&i == i {
			ret++
		}
	}
	return ret
}

var x [16][16]int
var y [16][16]int

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	ln := nextInt()
	a := make([]int, ln+1)
	for i := 1; i < ln; i++ {
		a[i] = nextInt()
		for j := 1; j <= a[i]; j++ {
			x[i][j], y[i][j] = nextInt(), nextInt()
		}
	}

	n := IntPow(2, ln) - 1
	ans := 0
	for bits := n; bits >= 1; bits-- {
		// 証言に基づく正直者リスト
		ok := true
		for i := 1; i <= ln; i++ {
			for j := 1; j <= a[i]; j++ {
				if bits&(1<<(i-1)) == 0 {
					continue
				}
				if (bits>>(x[i][j]-1)&1)^y[i][j] == 1 {
					ok = false
				}
			}
			if ok {
				ans = Max(ans, CountBits(bits))
			}
		}
	}
	fmt.Println(ans)
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

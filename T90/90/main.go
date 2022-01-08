package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const p = 998244353

var sc = bufio.NewScanner(os.Stdin)

// 小課題1 n<=1e5までフィボナッチ数列を計算
func SubTask1(n, k int) int {
	m := make(map[int]bool)
	f := []int{1, 1}
	for i := 0; i < n; i++ {
		fn := len(f)
		nf := (f[fn-1] + f[fn-2]) % p

		m[nf] = true
		f = append(f, nf)
	}

	return f[n+1]
}

// 小課題2 行列を使った高速フィボナッチ数列
func SubTask2(n, k int) int {
	type Matrix [2][2]int

	Mul := func(a, b *Matrix, p int) *Matrix {
		ret := new(Matrix)
		ret[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
		ret[0][0] %= p
		ret[0][1] = a[0][0]*b[1][0] + a[0][1]*b[1][1]
		ret[0][1] %= p
		ret[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
		ret[1][0] %= p
		ret[1][1] = a[1][0]*b[1][0] + a[1][1]*b[1][1]
		ret[1][1] %= p
		return ret
	}
	ModPow := func(a *Matrix, n, p int) *Matrix {
		ret := new(Matrix)
		ret[0][0] = a[0][0]
		ret[0][1] = a[0][1]
		ret[1][0] = a[1][0]
		ret[1][1] = a[1][1]
		for n > 0 {
			if n%2 == 1 {
				ret = Mul(ret, a, p)
			}
			n >>= 1
			a = Mul(a, a, p)
		}
		return ret
	}

	mtx := new(Matrix)
	mtx[0][0] = 0
	mtx[0][1] = 1
	mtx[1][0] = 1
	mtx[1][1] = 1

	ret := ModPow(mtx, n, p)
	return ret[1][1]
}

// 全網羅
func SubTask3(n, k int) int {
	// 小課題3 全網羅
	const INF = 1 << 60

	var s [][]int
	var Dfs func(i int, ps []int)
	Dfs = func(i int, ps []int) {
		ps = append(ps, i)
		if len(ps) == n {
			s = append(s, ps)
			return
		}

		for i := 0; i <= k; i++ {
			nps := make([]int, len(ps))
			copy(nps, ps)
			Dfs(i, nps)
		}
	}
	for i := 0; i <= k; i++ {
		Dfs(i, make([]int, 0))
	}
	//fmt.Println(s)
	ans := 0
	for _, ps := range s {
		ok := true
		for l := 0; l < n; l++ {
			for r := l; r < n; r++ {
				min := INF
				for i := l; i <= r; i++ {
					min = Min(min, ps[i])
				}
				ok = ok && min*(r-l+1) <= k
			}
		}
		if ok {
			ans++
			ans %= p
		}
	}
	return ans

}

func SubTask4(n, k int) int {
	PrintDp := func(dp [][]int) {
		for i := 0; i < n; i++ {
			fmt.Println(dp[i])
		}
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[0][j] = 1
		}
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			dp[i][j] = dp[i-1][j] * ((j - 1) / i)
		}
	}
	ans := 0
	for i := 0; i <= k; i++ {
		ans += dp[n][i]
	}
	return ans
}

func SolveHonestly(n, k int) int {
	if k == 1 {
		if n <= int(1e5) {
			return SubTask1(n, k)
		} else {
			return SubTask2(n, k)
		}
	}
	if n <= 6 && k <= 6 {
		return SubTask3(n, k)
	} else if n <= 100 && k <= 100 {
		return SubTask4(n, k)
	} else {
		return 0
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	fmt.Println(SolveHonestly(n, k))
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

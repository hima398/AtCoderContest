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

type Point struct {
	t, a int
}

func FirstSolve() {
	h, w := nextInt(), nextInt()
	if h == 1 && w == 1 {
		fmt.Println("Draw")
		return
	}
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		s := nextString()
		for j := range s {
			if s[j] == '+' {
				a[i][j] = 1
			} else {
				a[i][j] = -1
			}
		}
	}
	//fmt.Println(a)
	dp := make([][]Point, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]Point, w)
	}
	// j == 0
	for i := 1; i < h; i++ {
		if i%2 == 1 {
			dp[i][0].t = dp[i-1][0].t + a[i][0]
			dp[i][0].a = dp[i-1][0].a
		} else {
			dp[i][0].t = dp[i-1][0].t
			dp[i][0].a = dp[i-1][0].a + a[i][0]
		}
	}
	// i == 0
	for j := 1; j < w; j++ {
		if j%2 == 1 {
			dp[0][j].t = dp[0][j-1].t + a[0][j]
			dp[0][j].a = dp[0][j-1].a
		} else {
			dp[0][j].t = dp[0][j-1].t
			dp[0][j].a = dp[0][j-1].a + a[0][j]
		}
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if (i+j)%2 == 1 {
				dp[i][j].t = Max(dp[i][j-1].t, dp[i-1][j].t) + a[i][j]
				dp[i][j].a = Max(dp[i][j-1].a, dp[i-1][j].a)
			} else {
				dp[i][j].t = Max(dp[i][j-1].t, dp[i-1][j].t)
				dp[i][j].a = Max(dp[i][j-1].a, dp[i-1][j].a) + a[i][j]
			}
		}
	}
	/*
		for i := 0; i < h; i++ {
			fmt.Println(dp[i])
		}
	*/
	//fmt.Println(dp[h][w])
	/*
		turn := ((h - 1) + (w - 1)) % 2
		takahashi, aoki := -1, -1
		if turn == 0 {
			aoki = dp[h-1][w-1]
			if h == 1 {
				takahashi = dp[h-1][w-2]
			} else if w == 1 {
				takahashi = dp[h-2][w-1]
			} else {
				takahashi = Max(dp[h-2][w-1], dp[h-1][w-2])
			}
		} else {
			takahashi = dp[h-1][w-1]
			if h == 1 {
				aoki = dp[h-1][w-2]
			} else if w == 1 {
				aoki = dp[h-2][w-1]
			} else {
				aoki = Max(dp[h-2][w-1], dp[h-1][w-2])
			}
		}
	*/
	//fmt.Println(turn, takahashi, aoki)
	if dp[h-1][w-1].a > dp[h-1][w-1].t {
		fmt.Println("Aoki")
	} else if dp[h-1][w-1].a < dp[h-1][w-1].t {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Draw")
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	if h == 1 && w == 1 {
		fmt.Println("Draw")
		return
	}
	INF := h*w + 1
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		s := nextString()
		a[i] = make([]int, w)
		for j := range s {
			if s[j] == '+' {
				a[i][j] = 1
			} else {
				a[i][j] = -1
			}
		}
	}
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
		for j := 0; j < w; j++ {
			dp[i][j] = -INF
		}
	}
	//dp[0][0] = 0
	dp[h-1][w-1] = 0
	// j==w-1
	for i := h - 2; i >= 0; i-- {
		if (w-1+i)%2 == 0 {
			dp[i][w-1] = dp[i+1][w-1] + a[i+1][w-1]
		} else {
			dp[i][w-1] = dp[i+1][w-1] - a[i+1][w-1]
		}
	}
	// i==h-1
	for j := w - 2; j >= 0; j-- {
		if (h-1+j)%2 == 0 {
			dp[h-1][j] = dp[h-1][j+1] + a[h-1][j+1]
		} else {
			dp[h-1][j] = dp[h-1][j+1] - a[h-1][j+1]
		}
	}
	//fmt.Println(dp)
	for i := h - 2; i >= 0; i-- {
		for j := w - 2; j >= 0; j-- {
			if (i+j)%2 == 0 {
				dp[i][j] = Max(dp[i+1][j]+a[i+1][j], dp[i][j+1]+a[i][j+1])
			} else {
				dp[i][j] = Min(dp[i+1][j]-a[i+1][j], dp[i][j+1]-a[i][j+1])
			}
		}
	}
	/*
		for i := 0; i < h; i++ {
			fmt.Println(dp[i])
		}
	*/
	//fmt.Println(dp[h-1][w-1])
	if dp[0][0] > 0 {
		fmt.Println("Takahashi")
	} else if dp[0][0] < 0 {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Draw")
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

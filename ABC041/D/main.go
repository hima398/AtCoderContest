package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestry(n int, e [][]int) int {
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, i+1)
	}

	ans := 0
	for {
		indexOf := make(map[int]int)
		for i, v := range s {
			indexOf[v] = i
		}
		ok := true
		for x := 1; x <= n; x++ {
			for _, y := range e[x] {
				ok = ok && indexOf[x] < indexOf[y]
			}
		}
		if ok {
			ans++
		}
		//fmt.Println(ok, s)
		if !NextPermutation(sort.IntSlice(s)) {
			break
		}
	}
	return ans
}
func Solve(n int, e [][]int) int {
	set := make(map[string]bool)
	for y := 1; y <= n; y++ {
		for _, x := range e[y] {
			var s []string
			for i := 1; i <= n; i++ {
				if i == x || i == y {
					continue
				}
				s = append(s, strconv.Itoa(i))
			}
			fmt.Println(s)
			for i := 0; i < n-1; i++ {
				for j := i + 1; j < n; j++ {
					//fmt.Println(i, j)
					t := make([]string, len(s))
					copy(t, s)
					/*
						if i == 0 {
							t = append([]string{strconv.Itoa(y)}, t...)
						} else {
							t = append(append(t[:i], strconv.Itoa(y)), t[i+1:]...)
						}
					*/
					if i < n-2 {
						t = append(t[:i+1], t[i:]...)
						t[i] = strconv.Itoa(y)

						t = append(t[:j+1], t[j:]...)
						t[j] = strconv.Itoa(x)
					} else {
						t = append(t, strconv.Itoa(y))
						t = append(t, strconv.Itoa(x))
					}
					fmt.Println(strings.Join(t, ""))
					set[strings.Join(t, "")] = true
				}
			}
		}
	}
	ans := Factional(n) - len(set)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	//e := make([][]int, n+1)
	me := make([][]bool, n)
	for i := 0; i < n; i++ {
		me[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		//e[x] = append(e[x], y)
		//e[y] = append(e[y], x)
		x--
		y--
		//e[x] = append(e[y], x)
		me[y][x] = true
	}
	max := 1 << n
	dp := make([]int, max+1)
	dp[0] = 1
	fmt.Println(n, max)
	for i := 0; i < max; i++ {
		for j := 0; j < n; i++ {
			if i&(1<<j) == 0 {
				ok := true
				for k := 0; k < n; k++ {
					if i&(1<<k) > 0 {
						if me[k][j] {
							ok = false
						}
					}
				}
				fmt.Println(i, i|(1<<j))
				if ok {
					dp[i|(1<<j)] += dp[i]
				}
			}
		}
	}
	fmt.Println(dp[(1<<n)-1])
	//fmt.Println(SolveHonestry(n, e))
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

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(s1, s2, s3 string) {
	m := make(map[byte]int)
	for i := range s1 {
		m[s1[i]]++
	}
	for i := range s2 {
		m[s2[i]]++
	}
	for i := range s3 {
		m[s3[i]]++
	}
	// 文字の種類が10より大きい場合は0〜9を割り当てられない
	if len(m) > 10 {
		fmt.Println("UNSOLVABLE")
		return
	}
	var cs []byte
	for k := range m {
		cs = append(cs, k)
	}

	//m2 := make(map[string]int)
	mapping := make([]byte, 10)
	Check := func() bool {
		m := make(map[byte]byte)
		for i := 0; i < 10; i++ {
			if mapping[i] != 0 {
				m[mapping[i]] = byte(i + '0')
			}
		}
		//var nss1, nss2, nss3 []string
		var nss1, nss2, nss3 []byte
		for i := range s1 {
			nss1 = append(nss1, m[s1[i]])
		}
		for i := range s2 {
			nss2 = append(nss2, m[s2[i]])
		}
		for i := range s3 {
			nss3 = append(nss3, m[s3[i]])
		}
		//ns1, ns2, ns3 := strings.Join(nss1, ""), strings.Join(nss2, ""), strings.Join(nss3, "")
		ns1, ns2, ns3 := string(nss1), string(nss2), string(nss3)
		n1, _ := strconv.Atoi(ns1)
		n2, _ := strconv.Atoi(ns2)
		n3, _ := strconv.Atoi(ns3)
		if n1 == 0 || n2 == 0 || n3 == 0 {
			return false
		}
		if strconv.Itoa(n1) != ns1 || strconv.Itoa(n2) != ns2 || strconv.Itoa(n3) != ns3 {
			return false
		}

		if n1+n2 == n3 {
			fmt.Println(n1)
			fmt.Println(n2)
			fmt.Println(n3)
			return true
		} else {
			return false
		}
	}
	var Dfs func(n int) bool
	Dfs = func(n int) bool {
		if n == len(cs) {
			//fmt.Println(mapping)
			return Check()
		}
		for i := 0; i < 10; i++ {
			//m2[string(cs[n])] = i
			if mapping[i] == 0 {
				mapping[i] = cs[n]
				if Dfs(n + 1) {
					return true
				}
				mapping[i] = 0
			}
		}
		return false
	}
	canSolve := Dfs(0)
	if !canSolve {
		fmt.Println("UNSOLVABLE")
	}
}

//文字の種類が8〜10くらいになると手元でも5s〜10s以上かかる
func SecondSolve(s1, s2, s3 string) {

	pattern := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := make(map[byte]int)
	for i := range s1 {
		m[s1[i]]++
	}
	for i := range s2 {
		m[s2[i]]++
	}
	for i := range s3 {
		m[s3[i]]++
	}
	// 文字の種類が10より大きい場合は0〜9を割り当てられない
	if len(m) > 10 {
		fmt.Println("UNSOLVABLE")
		return
	}
	var cs []byte
	for k := range m {
		cs = append(cs, k)
	}

	for {
		mapping := make(map[byte]byte)
		for i, v := range pattern {
			if i < len(cs) {
				mapping[cs[i]] = byte(v + '0')
			}
		}
		if mapping[s1[0]] == '0' || mapping[s2[0]] == '0' || mapping[s3[0]] == '0' {
		} else {
			var n1, n2, n3 int
			for i := 0; i < len(s1); i++ {
				n1 = n1*10 + int(mapping[s1[i]]-'0')
			}
			for i := 0; i < len(s2); i++ {
				n2 = n2*10 + int(mapping[s2[i]]-'0')
			}
			for i := 0; i < len(s3); i++ {
				n3 = n3*10 + int(mapping[s3[i]]-'0')
			}

			if n1+n2 == n3 {
				fmt.Println(n1)
				fmt.Println(n2)
				fmt.Println(n3)
				return
			}
		}
		if !NextPermutation(sort.IntSlice(pattern)) {
			break
		}
	}
	fmt.Println("UNSOLVABLE")
	return
}

func ThirdSolve(s1, s2, s3 string) {
	m := make(map[byte]bool)
	for i := range s1 {
		m[s1[i]] = true
	}
	for i := range s2 {
		m[s2[i]] = true
	}
	for i := range s3 {
		m[s3[i]] = true
	}
	// 文字の種類が10より大きい場合は0〜9を割り当てられない
	if len(m) > 10 {
		fmt.Println("UNSOLVABLE")
		return
	}
	var cs []byte
	for k := range m {
		cs = append(cs, k)
	}

	Check := func(used []byte) bool {
		m := make(map[byte]byte)
		//for i := 0; i < 10; i++ {
		for i := 0; i < len(used); i++ {
			m[used[i]] = byte(i + '0')
		}
		//var nss1, nss2, nss3 []string
		var nss1, nss2, nss3 []byte
		for i := range s1 {
			nss1 = append(nss1, m[s1[i]])
		}
		for i := range s2 {
			nss2 = append(nss2, m[s2[i]])
		}
		for i := range s3 {
			nss3 = append(nss3, m[s3[i]])
		}
		//ns1, ns2, ns3 := strings.Join(nss1, ""), strings.Join(nss2, ""), strings.Join(nss3, "")
		ns1, ns2, ns3 := string(nss1), string(nss2), string(nss3)
		n1, _ := strconv.Atoi(ns1)
		n2, _ := strconv.Atoi(ns2)
		n3, _ := strconv.Atoi(ns3)
		if n1 == 0 || n2 == 0 || n3 == 0 {
			return false
		}
		if strconv.Itoa(n1) != ns1 || strconv.Itoa(n2) != ns2 || strconv.Itoa(n3) != ns3 {
			return false
		}

		if n1+n2 == n3 {
			fmt.Println(n1)
			fmt.Println(n2)
			fmt.Println(n3)
			return true
		} else {
			return false
		}
	}

	//var Dfs func(used []byte, unused map[byte]bool) bool
	//Dfs = func(used []byte, unused map[byte]bool) bool {
	var patterns [][]byte
	var Dfs func(used []byte) bool
	Dfs = func(used []byte) bool {
		if len(used) == len(cs) {
			//return Check(used)
			patterns = append(patterns, used)
		}
		for k := range m {
			nextUsed := make([]byte, len(used))
			copy(nextUsed, used)
			nextUsed = append(nextUsed, k)
			//nextUnused := make(map[byte]bool)
			//for k, v := range unused {
			//	nextUnused[k] = v
			//}
			//delete(nextUnused, k)
			delete(m, k)
			Dfs(nextUsed)
			m[k] = true
		}
		return false
	}
	used := make([]byte, 0)
	//Dfs(used, m)
	Dfs(used)
	//fmt.Println(patterns)
	for _, pattern := range patterns {
		if Check(pattern) {
			return
		}
	}
	fmt.Println("UNSOLVABLE")
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//fmt.Println(Factional(10))
	s1 := nextString()
	s2 := nextString()
	s3 := nextString()

	//FirstSolve(s1, s2, s3)
	SecondSolve(s1, s2, s3)
	//ThirdSolve(s1, s2, s3)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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

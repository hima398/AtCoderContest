package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, a []int) int {
	const p = 998244353

	ps := make([]int, n+1)
	ps[0] = 1
	for i := 1; i < n; i++ {
		ps[i] = ps[i-1] * 2 % p
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] <= a[j] {
				ans += ps[j-i-1] //2のj-i-1乗
				ans %= p
			}
		}
	}
	return ans
}

type BinaryIndexTree struct {
	n     int
	nodes []int
}

func NewBinaryIndexTree(n int) BinaryIndexTree {
	var ret BinaryIndexTree
	ret.n = n
	ret.nodes = make([]int, n+1)
	return ret
}

func (b BinaryIndexTree) add(x, a, p int) {
	x++
	for x <= b.n {
		b.nodes[x] = (b.nodes[x] + a) % p
		x += x & -x
	}
}

func (b BinaryIndexTree) get(x, p int) (ret int) {
	x++
	for x > 0 {
		ret = (ret + b.nodes[x]) % p
		x -= x & -x
	}
	return ret
}

func CompressSlice(s []int) (int, []int) {
	m := make(map[int]bool)
	for _, si := range s {
		m[si] = true
	}
	var p []int
	for k := range m {
		p = append(p, k)
	}
	sort.Ints(p)
	cs := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		idx := sort.SearchInts(p, s[i])
		cs[i] = idx
	}
	return len(m), cs
}

func SolveCommentary(n int, a []int) int {
	const p = 998244353
	div := (p + 1) / 2 // 2**(-1)
	nn, ca := CompressSlice(a)
	//fmt.Println(a)
	//fmt.Println(nn, ca)

	biTree := NewBinaryIndexTree(nn)
	pows, divs := make([]int, n+1), make([]int, n+1)
	pows[0] = 1
	divs[0] = 1
	for i := 1; i <= n; i++ {
		pows[i] = 2 * pows[i-1] % p
		divs[i] = divs[i-1] * div % p
	}
	var ans int
	for i := 0; i < n; i++ {
		//ans += Pow(2, i, p) * biTree.get(ca[i], p)
		ans += pows[i] * biTree.get(ca[i], p)
		ans %= p
		biTree.add(ca[i], divs[i+1], p)
		//biTree.add(ca[i], Pow(div, i+1, p), p)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)

	//ans := SolveHonestly(n, a)
	ans := SolveCommentary(n, a)
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

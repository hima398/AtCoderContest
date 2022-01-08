package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, h, a []int) int {
	maxH := 0
	for i := 0; i < n; i++ {
		maxH = Max(maxH, h[i])
	}
	// i番目の花までの中で、単調増加の末尾がjになるキレイさの最大
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxH+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= maxH; j++ {
			dp[i][j] = dp[i-1][j]
		}
		for j := 1; j <= h[i-1]; j++ {
			dp[i][h[i-1]] = Max(dp[i][h[i-1]], dp[i-1][j]+a[i-1])
		}
		//fmt.Println(dp[i])
	}
	ans := 0
	for j := 0; j <= maxH; j++ {
		ans = Max(ans, dp[n][j])
	}
	return ans
}

const INF = 1 << 60

type SegmentTree struct {
	size  int
	nodes []int
}

func NewSegmentTree(n int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([]int, 2*st.size)
	for i := range st.nodes {
		st.nodes[i] = 0
	}
	return st
}

func (this *SegmentTree) QueryRecursively(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return 0
	}

	// [a, b)が[l, r)を完全に含んでいる
	if a <= l && b >= r {
		return this.nodes[k]
	}

	vl := this.QueryRecursively(a, b, 2*k, l, (l+r)/2)
	vr := this.QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return Max(vl, vr)
}

// [l, r)の区間の値をxorした結果を返す
func (this *SegmentTree) Query(l, r int) int {
	return this.QueryRecursively(l, r, 1, 0, this.size)
}

func (this *SegmentTree) Update(k, x int) {
	k += this.size
	this.nodes[k] = Max(this.nodes[k], x)
	for k > 1 {
		k /= 2
		this.nodes[k] = Max(this.nodes[k*2], this.nodes[k*2+1])
	}
	//fmt.Println(this.nodes)
}

func Solve(n int, h, a []int) int {
	maxH := 0
	for i := 0; i < n; i++ {
		maxH = Max(maxH, h[i])
	}

	//単調増加の末尾がiになるキレイさの最大値
	//を、セグ木を使って高速化
	st := NewSegmentTree(maxH + 1)
	for i := 0; i < n; i++ {
		max := st.Query(0, h[i])
		st.Update(h[i], max+a[i])
		//fmt.Println(st.nodes)
	}
	return st.Query(0, maxH+1)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	h := nextIntSlice(n)
	a := nextIntSlice(n)

	//ans := SolveHonestly(n, h, a)
	ans := Solve(n, h, a)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

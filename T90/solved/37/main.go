package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INF = int(1e9) + 1

var sc = bufio.NewScanner(os.Stdin)

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
		st.nodes[i] = -INF
	}
	return st
}

func (this *SegmentTree) QueryRecursively(a, b, k, l, r int) int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return -INF
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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	w, n := nextInt(), nextInt()
	//N種類の料理の中から何種類か選んで1つずつ作り、香辛料をW消費したときの料理の価値
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			dp[i][j] = -INF
		}
	}
	dp[0][0] = 0
	l, r, v := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		l[i], r[i], v[i] = nextInt(), nextInt(), nextInt()
	}

	for i := 0; i < n; i++ {
		st := NewSegmentTree(w + 2)
		for j := 0; j <= w; j++ {
			st.Update(j, dp[i][j])
		}
		for j := 0; j <= w; j++ {
			// i番目の料理を作らない
			dp[i+1][j] = dp[i][j]

			// i番目の料理を作る
			li, ri := Max(j-r[i], 0), Max(j-l[i]+1, 0)
			if li == ri {
				continue
			}
			max := st.Query(li, ri)
			if max != -INF {
				dp[i+1][j] = Max(dp[i+1][j], max+v[i])
			}
		}
	}
	var ans int
	if dp[n][w] == -INF {
		ans = -1
	} else {
		ans = dp[n][w]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

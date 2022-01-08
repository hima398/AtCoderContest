package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type BinaryIndexTree struct {
	n     int
	nodes []int
}

func New(n int) *BinaryIndexTree {
	ret := new(BinaryIndexTree)
	ret.n = n
	ret.nodes = make([]int, n+1)
	return ret
}

func (b *BinaryIndexTree) Update(i, v int) {
	i++
	for i <= b.n {
		//b.nodes[i] = b.eval(b.nodes[i], v)
		b.nodes[i] = Max(b.nodes[i], v)
		i += i & -i
	}
}

func (b *BinaryIndexTree) Get(i int) int {
	i++
	ret := 0
	for i > 0 {
		//ret = b.eval(b.nodes[i], ret)
		ret = Max(b.nodes[i], ret)
		i -= i & -i
	}
	return ret
}

func SolveHonestly(n int, w, h []int) int {

	memo := make([]int, n)
	var F func(i int) int
	F = func(i int) int {
		if memo[i] > 0 {
			return memo[i]
		}
		//fmt.Println(i)
		ret := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if w[i] > w[j] && h[i] > h[j] {
				ret = Max(ret, F(j))
			}
		}
		ret++
		memo[i] = ret
		return memo[i]
	}

	return memo[n-1]
}

func Solve(n int, w, h []int) int {
	type Box struct {
		h, w int
	}
	var bs []Box
	for i := 0; i < n; i++ {
		bs = append(bs, Box{h[i], w[i]})
	}
	sort.Slice(bs, func(i, j int) bool {
		// hが同じ箱はwが大きい方が最大値を取りやすいので
		// wが先に評価されるようにwを降順にしておく
		if bs[i].h == bs[j].h {
			return bs[i].w > bs[j].w
		}
		return bs[i].h < bs[j].h
	})
	wx := 0
	for _, wi := range w {
		wx = Max(wx, wi)
	}
	bitree := New(wx)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		// dp[i] = Max(dp[0], dp[1], .... dp[i-1])+1
		dp[i] = bitree.Get(bs[i].w-1) + 1
		bitree.Update(bs[i].w, dp[i])
	}
	//fmt.Println(bitree.nodes)
	//fmt.Println(dp)
	ans := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, dp[i])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	w, h := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		w[i], h[i] = nextInt(), nextInt()
	}

	//ans := SolveHonestly(n, w, h)
	ans := Solve(n, w, h)
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

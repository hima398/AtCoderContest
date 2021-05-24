package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INF = 1 << 60

var sc = bufio.NewScanner(os.Stdin)

type Ball struct {
	l, r, c int
}

func Dist(f, t, m int) int {
	if (f < m && m < t) || (f > m && m > t) {
		return Abs(t - f)
	} else {
		return Abs(m-f) + Abs(t-m)
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	cmax := 0
	x, c := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], c[i] = nextInt(), nextInt()
		cmax = Max(cmax, c[i])
	}
	b := make([]Ball, cmax+1)
	for i := 0; i <= cmax; i++ {
		b[i].l = INF
		b[i].r = -INF
	}
	for i := 0; i < n; i++ {
		b[c[i]].c = c[i]
		b[c[i]].r = Max(b[c[i]].r, x[i])
		b[c[i]].l = Min(b[c[i]].l, x[i])
	}
	b[0].l = 0
	b[0].r = 0
	//fmt.Println(b)
	dp := make([][]int, cmax+1)
	for i := 0; i <= cmax; i++ {
		dp[i] = make([]int, 2)
	}
	//fmt.Println(b)
	for i := 0; i < cmax; i++ {
		//if b[i+1].r == -INF && b[i+1].l == INF {
		// 存在しない色の処理
		if b[i+1].c == 0 {
			dp[i+1][0] = dp[i][0]
			dp[i+1][1] = dp[i][1]
			// ボールは存在しないけども最左と最右の値を参照するのでコピーしておく
			b[i+1].l = b[i].l
			b[i+1].r = b[i].r
			continue
		}
		// c[i]のボールのうち最左からc[i+1]のボールの最左に向かう距離
		dl1 := Dist(b[i].l, b[i+1].l, b[i+1].r)
		// c[i]のボールのうち最右からc[i+1]のボールの最左に向かう距離
		dl2 := Dist(b[i].r, b[i+1].l, b[i+1].r)

		dp[i+1][0] = Min(dp[i][0]+dl1, dp[i][1]+dl2)

		// rignt
		// c[i]のボールのうち最左からc[i+1]のボールの最右に向かう距離
		dr1 := Dist(b[i].l, b[i+1].r, b[i+1].l)
		// c[i]のボールのうち最右からc[i+1]のボールの最右に向かう距離
		dr2 := Dist(b[i].r, b[i+1].r, b[i+1].l)

		dp[i+1][1] = Min(dp[i][0]+dr1, dp[i][1]+dr2)
	}
	//fmt.Println(dp[cmax][0], dp[cmax][1])
	ans := Min(dp[cmax][0]+Abs(b[cmax].l), dp[cmax][1]+Abs(b[cmax].r))
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

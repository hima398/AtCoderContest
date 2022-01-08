package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7
	const INF = int(1e5) + 1
	n, m := nextInt(), nextInt()
	e := make(map[int][]int)
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		x--
		y--
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	var q []int
	q = append(q, 0)
	ds := make([]int, n)
	dp := make([]int, n)
	for i := range ds {
		ds[i] = INF
	}
	ds[0] = 0
	dp[0] = 1
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		//ds[node.i] = node.d
		// 親のノードに辿り着くまでの経路の個数を注目しているノードに辿り着くまでの経路の個数に足す
		for _, par := range e[i] {
			if ds[par] == ds[i]-1 {
				dp[i] += dp[par]
				dp[i] %= p
			}
		}
		// 子ノードを探索するためキューに積む
		for _, chi := range e[i] {
			if ds[chi] == INF {
				q = append(q, chi)
				ds[chi] = ds[i] + 1
			}
		}
	}
	//fmt.Println(dp)
	fmt.Println(dp[n-1])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

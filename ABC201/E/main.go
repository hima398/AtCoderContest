package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Edge struct {
	u, v, w int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7
	n := nextInt()
	//u, v, w := make([]int, n-1), make([]int, n-1), make([]int, n-1)
	e := make(map[int][]Edge)
	edge := make([]Edge, n-1)
	for i := 0; i < n-1; i++ {
		u, v, w := nextInt(), nextInt(), nextInt()
		u--
		v--
		e[u] = append(e[u], Edge{u, v, w})
		e[v] = append(e[v], Edge{v, u, w})
		edge[i] = Edge{u, v, w}
	}
	//fmt.Println(e)
	// 頂点0から各頂点までの距離
	dp := make([]int, n)
	var dfs func(x, p, w int)
	dfs = func(x, p, w int) {
		//fmt.Println(x, p, w)
		dp[x] = w
		//fmt.Println(e[x])
		for _, v := range e[x] {
			//木のDFSなので行き先が親であるものだけ除外する
			if v.v == p {
				continue
			}
			dfs(v.v, x, w^v.w)
		}
	}
	dfs(0, -1, 0)
	//fmt.Println(dp)
	ans := 0
	for k := 0; k < 60; k++ {
		nz, no := 0, 0
		for i := 0; i < n; i++ {
			if dp[i]&(1<<k) == 0 {
				nz++
			} else {
				no++
			}
		}
		s := nz * no
		for kk := 0; kk < k; kk++ {
			s *= 2
			s %= p
		}
		ans += s
		ans %= p
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

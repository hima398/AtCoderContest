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

	const INF = 3 * int(1e9)

	n, m := nextInt(), nextInt()
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		e[i] = make([]int, n)
		for j := 0; j < n; j++ {
			e[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		u, v, l := nextInt(), nextInt(), nextInt()
		u--
		v--
		e[u][v] = l
		e[v][u] = l
	}

	for k := 1; k < n; k++ {
		for i := 1; i < n; i++ {
			for j := 1; j < n; j++ {
				if e[i][k] != INF && e[k][j] != INF {
					e[i][j] = Min(e[i][j], e[i][k]+e[k][j])
				}
			}
		}
	}
	ans := INF
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ans = Min(ans, e[0][i]+e[i][j]+e[j][0])
		}
	}
	if ans == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve() {
	h, w, n, m := nextInt(), nextInt(), nextInt(), nextInt()
	// 0:dark, 1:light, 2:bulb, 3:block
	g := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		g[i] = make([]int, w+1)
	}
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
		g[a[i]][b[i]] = 2
	}
	for i := 0; i < m; i++ {
		c, d := nextInt(), nextInt()
		g[c][d] = 3
	}
	ans := 0
	//0:dark, 1:light, 2:bulb, 3:block
	for k := 0; k < n; k++ {
		for i := a[k] - 1; i > 0; i-- {
			if g[i][b[k]] > 1 {
				break
			}
			if g[i][b[k]] == 0 {
				g[i][b[k]] = 1
				ans++
			}
		}
		for j := b[k] - 1; j > 0; j-- {
			if g[a[k]][j] > 1 {
				break
			}
			if g[a[k]][j] == 0 {
				g[a[k]][j] = 1
				ans++
			}
		}
		for i := a[k] + 1; i <= h; i++ {
			if g[i][b[k]] > 1 {
				break
			}
			if g[i][b[k]] == 0 {
				g[i][b[k]] = 1
				ans++

			}
		}
		for j := b[k] + 1; j <= w; j++ {
			if g[a[k]][j] > 1 {
				break
			}
			if g[a[k]][j] == 0 {
				g[a[k]][j] = 1
				ans++
			}
		}
	}
	ans += n
	fmt.Println(ans)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	FirstSolve()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

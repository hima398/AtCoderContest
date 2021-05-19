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

	const INF = int(1e6) + 1
	h, w := nextInt(), nextInt()
	rs, cs := nextInt(), nextInt()
	rs--
	cs--
	rt, ct := nextInt(), nextInt()
	rt--
	ct--
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	v := make([][]bool, h)
	ans := make([][]int, h)
	for i := 0; i < h; i++ {
		v[i] = make([]bool, w)
		ans[i] = make([]int, w)
		for j := 0; j < w; j++ {
			ans[i][j] = INF
		}
	}
	dr := []int{0, -1, 0, 1}
	dc := []int{-1, 0, 1, 0}
	var dfs func(pr, pc, r, c, t int)
	dfs = func(pr, pc, r, c, t int) {
		v[r][c] = true
		ans[r][c] = Min(ans[r][c], t)
		for i := 0; i < 4; i++ {
			nr, nc := r+dr[i], c+dc[i]
			if nr >= 0 && nr < h && nc >= 0 && nc < w && s[nr][nc] == '.' {
				//上下方向の移動
				if i%2 == 0 {
					if pc-c == 0 {
						if t+1 < ans[nr][nc] {
							dfs(r, c, nr, nc, t+1)
						}
					} else {
						if t < ans[nr][nc] {
							dfs(r, c, nr, nc, t)
						}
					}
				} else {
					//左右方向の移動
					if pr-r == 0 {
						if t+1 < ans[nr][nc] {
							dfs(r, c, nr, nc, t+1)
						}
					} else {
						if t < ans[nr][nc] {
							dfs(r, c, nr, nc, t)
						}
					}
				}
			}
		}
	}
	v[rs][cs] = true
	ans[rs][cs] = 0
	for i := 0; i < 4; i++ {
		nr, nc := rs+dr[i], cs+dc[i]
		if nr >= 0 && nr < h && nc >= 0 && nc < w && s[nr][nc] == '.' {
			dfs(rs, cs, nr, nc, 0)
		}
	}

	/*
		for i := 0; i < h; i++ {
			fmt.Printf("%7d\n", ans[i])
		}
	*/

	fmt.Println(ans[rt][ct])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

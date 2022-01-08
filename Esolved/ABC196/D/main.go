package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(h, w, a, b int) int {
	used := make([][]bool, h)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
	}
	var dfs func(i, j, a, b int) int
	dfs = func(i, j, a, b int) int {
		if a < 0 || b < 0 {
			return 0
		}
		if j == w {
			j = 0
			i++
		}
		if i == h {
			return 1
		}

		if used[i][j] {
			return dfs(i, j+1, a, b)
		}

		used[i][j] = true
		ret := dfs(i, j+1, a, b-1)

		if j+1 < w && !used[i][j+1] {
			used[i][j+1] = true
			ret += dfs(i, j+1, a-1, b)
			used[i][j+1] = false
		}

		if i+1 < h && !used[i+1][j] {
			used[i+1][j] = true
			ret += dfs(i, j+1, a-1, b)
			used[i+1][j] = false
		}

		used[i][j] = false
		return ret
	}
	ans := dfs(0, 0, a, b)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, a, b := nextInt(), nextInt(), nextInt(), nextInt()
	fmt.Println(SolveCommentary(h, w, a, b))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

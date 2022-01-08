package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Cell struct {
	i, j, w int
}

func PrintField(f1, f2 [][]int) {
	for i := 0; i < len(f2); i++ {
		fmt.Println(f1[i], f2[i])
	}
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7

	h, w := nextInt(), nextInt()
	a := make([][]int, h)
	var cs []Cell
	for i := 0; i < h; i++ {
		a[i] = nextIntSlice(w)
		for j := 0; j < w; j++ {
			cs = append(cs, Cell{i, j, a[i][j]})
		}
	}
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].w > cs[j].w
	})
	s := make([][]int, h)
	for i := 0; i < h; i++ {
		s[i] = make([]int, w)
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	// O(h*w)程度のループ
	for i := 0; i < len(cs); i++ {
		c := cs[i]
		// i, jに留まるパターン
		s[c.i][c.j] = 1
		for k := 0; k < 4; k++ {
			ni, nj := c.i+di[k], c.j+dj[k]
			if ni >= 0 && ni < h && nj >= 0 && nj < w {
				if c.w < a[ni][nj] {
					s[c.i][c.j] += s[ni][nj]
					s[c.i][c.j] %= p
				}
			}
		}
	}
	//PrintField(a, s)
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans += s[i][j]
			ans %= p
		}
	}
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

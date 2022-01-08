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

	h, w := nextInt(), nextInt()
	f := make([]string, h)
	for i := range f {
		f[i] = nextString()
	}
	// 左、上、右、下の順
	diri := []int{0, -1, 0, 1}
	dirj := []int{-1, 0, 1, 0}
	v := make([][]bool, h)
	for i := range v {
		v[i] = make([]bool, w)
	}
	var Dfs func(i, j, si, sj int) int
	Dfs = func(i, j, si, sj int) int {
		//fmt.Println("  ", i, j, v[i][j])
		if i == si && j == sj && v[i][j] {
			return 0
		}
		v[i][j] = true
		ret := -(h*w + 1)
		for k := 0; k < 4; k++ {
			ni, nj := i, j
			ni += diri[k]
			nj += dirj[k]
			if ni >= 0 && ni < h && nj >= 0 && nj < w && f[ni][nj] == '.' {
				if !v[ni][nj] || (ni == si && nj == sj) {
					dist := Dfs(ni, nj, si, sj)
					ret = Max(ret, dist+1)
				}
			}
		}
		v[i][j] = false
		return ret
	}
	ans := -1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			//fmt.Println(i, j)
			ans = Max(ans, Dfs(i, j, i, j))
		}
	}
	if ans <= 2 {
		ans = -1
	}
	fmt.Println(ans)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

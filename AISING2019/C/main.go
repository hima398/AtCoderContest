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
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	v := make([][]bool, h)
	for i := 0; i < h; i++ {
		v[i] = make([]bool, w)
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	var Dfs func(i, j int) (a, b int)
	Dfs = func(i, j int) (a, b int) {
		v[i][j] = true
		a = 1
		for k := 0; k < 4; k++ {
			ni := i + di[k]
			nj := j + dj[k]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			//fmt.Println(v[ni][nj])
			if v[ni][nj] {
				continue
			}
			//fmt.Println(s[i][j], s[ni][nj])
			if s[i][j] != s[ni][nj] {
				r1, r2 := Dfs(ni, nj)
				//fmt.Println(r1, r2)
				a += r2
				b += r1
			}
		}
		return a, b
	}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !v[i][j] {
				b, w := Dfs(i, j)
				ans += b * w
			}
		}
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

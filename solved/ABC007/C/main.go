package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	i, j, d int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	si, sj := nextInt(), nextInt()
	si--
	sj--
	gi, gj := nextInt(), nextInt()
	gi--
	gj--
	c := make([]string, h)
	for i := 0; i < h; i++ {
		c[i] = nextString()
	}
	var q []Pos
	ans := make([][]int, h)
	v := make([][]bool, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]int, w)
		v[i] = make([]bool, w)
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	q = append(q, Pos{si, sj, 0})
	v[si][sj] = true
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		ans[p.i][p.j] = p.d
		for k := 0; k < 4; k++ {
			ni, nj := p.i+di[k], p.j+dj[k]
			if ni >= 0 && ni < h && nj >= 0 && nj < w {
				if c[ni][nj] == '.' && !v[ni][nj] {
					q = append(q, Pos{ni, nj, p.d + 1})
					v[ni][nj] = true
				}
			}
		}
	}
	fmt.Println(ans[gi][gj])
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

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

func PrintField(d [][]int) {
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 25*int(1e4) + 1
	h, w := nextInt(), nextInt()
	c := make([]string, h)
	d := make([][]int, h)
	v := make([][]bool, h)
	si, sj, gi, gj := -1, -1, -1, -1
	for i := 0; i < h; i++ {
		c[i] = nextString()
		d[i] = make([]int, w)
		v[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			if c[i][j] == 's' {
				si, sj = i, j
			}
			if c[i][j] == 'g' {
				gi, gj = i, j
			}
			d[i][j] = INF
		}
	}
	var q0, q1 []Pos
	q0 = append(q0, Pos{si, sj, 0})
	v[si][sj] = true
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for len(q0) > 0 || len(q1) > 0 {
		var p Pos
		if len(q0) > 0 {
			p = q0[0]
			q0 = q0[1:]
		} else {
			p = q1[0]
			q1 = q1[1:]
		}
		d[p.i][p.j] = p.d
		for k := 0; k < 4; k++ {
			ni, nj := p.i+di[k], p.j+dj[k]
			if ni >= 0 && ni < h && nj >= 0 && nj < w {
				if !v[ni][nj] {
					if c[ni][nj] == '#' {
						q1 = append(q1, Pos{ni, nj, p.d + 1})
					} else {
						q0 = append(q0, Pos{ni, nj, p.d})
					}
					v[ni][nj] = true
				}
			}
		}
	}
	//PrintField(d)
	if d[gi][gj] <= 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintAns(ans [][]int) {
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}

type Node struct {
	i, j, w int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 1 << 60
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	ans := make([][]int, h)
	v1, v2 := make([][]bool, h), make([][]bool, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]int, w)
		v1[i] = make([]bool, w)
		v2[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			ans[i][j] = INF
		}
	}
	var q1, q2 []Node
	q1 = append(q1, Node{0, 0, 0})
	v1[0][0] = true
	for len(q1) > 0 || len(q2) > 0 {
		var p Node
		if len(q1) > 0 {
			p = q1[0]
			q1 = q1[1:]
		} else {
			p = q2[0]
			q2 = q2[1:]
		}
		ans[p.i][p.j] = Min(ans[p.i][p.j], p.w)

		//進める道があれば進む
		for k := 0; k < 4; k++ {
			ni, nj := p.i+di[k], p.j+dj[k]
			if ni >= 0 && ni < h && nj >= 0 && nj < w {
				if s[ni][nj] == '.' && !v1[ni][nj] {
					q1 = append(q1, Node{ni, nj, p.w})
					v1[ni][nj] = true
				}
			}
		}

		//周辺の壁壊しを行う
		for ii := -2; ii <= 2; ii++ {
			for jj := -2; jj <= 2; jj++ {
				if (ii == 0 && jj == 0) || (Abs(ii) == 2 && Abs(jj) == 2) {
					continue
				}
				ni, nj := p.i+ii, p.j+jj
				if ni >= 0 && ni < h && nj >= 0 && nj < w {
					if !v2[ni][nj] {
						q2 = append(q2, Node{ni, nj, p.w + 1})
						v2[ni][nj] = true
					}
				}
			}
		}
	}
	//PrintAns(ans)
	fmt.Println(ans[h-1][w-1])
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

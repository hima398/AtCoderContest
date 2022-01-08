package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	x, y int
}

func Dist(x1, y1, x2, y2 int) int64 {
	return int64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func SolveHonestly(n int, p []Pos) int {
	var ans int
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				for l := k + 1; l < n; l++ {
					if Dist(p[i].x, p[i].y, p[j].x, p[j].y) == Dist(p[k].x, p[k].y, p[l].x, p[l].y) && Dist(p[i].x, p[i].y, p[k].x, p[k].y) == Dist(p[j].x, p[j].y, p[l].x, p[l].y) && Dist(p[i].x, p[i].y, p[l].x, p[l].y) == Dist(p[j].x, p[j].y, p[k].x, p[k].y) {
						ans++
					}
				}
			}
		}
	}
	return ans
}

type Line struct {
	x1, y1, x2, y2 int
}

func SolveDiagonal() int {
	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
	}
	m := make(map[int64][]Line)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			m[Dist(x[i], y[i], x[j], y[j])] = append(m[Dist(x[i], y[i], x[j], y[j])], Line{x[i], y[i], x[j], y[j]})
		}
	}

	var ans int
	for k, v := range m {
		if k == 1 || len(v) < 2 {
			continue
		}

		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				ok := true
				//m2 := make(map[int]int)
				ok = ok && (k == Dist(v[i].x1, v[i].y1, v[j].x1, v[j].y1)+Dist(v[i].x1, v[i].y1, v[j].x2, v[j].y2))
				ok = ok && (k == Dist(v[i].x2, v[i].y2, v[j].x1, v[j].y1)+Dist(v[i].x2, v[i].y2, v[j].x2, v[j].y2))
				if ok {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	mx := make(map[int][]int)
	for i := 0; i < n; i++ {
		x[i], y[i] = nextInt(), nextInt()
		mx[x[i]] = append(mx[x[i]], y[i])
	}
	for _, v := range mx {
		sort.Ints(v)
	}

	//fmt.Println(mx)
	my := make(map[int]map[int]int)
	// x = kの数だけループ
	for _, ys := range mx {
		// x = kに点が2つに満たないものはスキップする
		if len(ys) < 2 {
			continue
		}
		// x = kにある2点を選択
		for i := range ys {
			for j := i + 1; j < len(ys); j++ {
				y1, y2 := ys[i], ys[j]
				if _, ok := my[y1]; !ok {
					my[y1] = make(map[int]int)
				}
				my[y1][y2]++
			}
		}
	}
	var ans int
	for y1 := range my {
		// v = y1, y2の点を持つx = kの数
		for _, v := range my[y1] {
			ans += v * (v - 1) / 2
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

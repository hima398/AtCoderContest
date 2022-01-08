package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INF = 1 << 60

var sc = bufio.NewScanner(os.Stdin)

var f []string
var d [20][20][20][20]int

func SolveWarshallFloyd(h, w int) int {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for ii := 0; ii < h; ii++ {
				for jj := 0; jj < w; jj++ {
					d[i][j][ii][jj] = INF
				}
			}
		}
	}
	// 経由点
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// 開始点
			for ii := 0; ii < h; ii++ {
				for jj := 0; jj < w; jj++ {
					// 終点
					for iii := 0; iii < h; iii++ {
						for jjj := 0; jjj < w; jjj++ {
							d[ii][jj][iii][jjj] = Min(d[ii][jj][iii][jjj], d[ii][jj][i][j]+d[i][j][iii][jjj])
						}
					}
				}
			}
		}
	}
	return 0
}

func Solve(h, w int) int {
	ans := 0
	// 開始点を全網羅
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if f[i][j] == '#' {
				continue
			}
			d := make([][]int, h)
			for ii := 0; ii < h; ii++ {
				d[ii] = make([]int, w)
			}
			for ii := 0; ii < h; ii++ {
				for jj := 0; jj < w; jj++ {
					d[ii][jj] = -1
				}
			}
			maxd := 0
			q := NewQueue()
			fp := Pos{i, j, 0} // First Pos
			q.Push(fp)
			d[i][j] = 0
			for {
				pos := q.Pop()
				d[pos.I][pos.J] = pos.D
				maxd = Max(maxd, d[pos.I][pos.J])
				if pos.I-1 >= 0 && d[pos.I-1][pos.J] < 0 && f[pos.I-1][pos.J] == '.' && !q.Find(pos.I-1, pos.J) {
					np := Pos{pos.I - 1, pos.J, pos.D + 1}
					q.Push(np)
				}
				if pos.J-1 >= 0 && d[pos.I][pos.J-1] < 0 && f[pos.I][pos.J-1] == '.' && !q.Find(pos.I, pos.J-1) {
					np := Pos{pos.I, pos.J - 1, pos.D + 1}
					q.Push(np)
				}
				if pos.J+1 < w && d[pos.I][pos.J+1] < 0 && f[pos.I][pos.J+1] == '.' && !q.Find(pos.I, pos.J+1) {
					np := Pos{pos.I, pos.J + 1, pos.D + 1}
					q.Push(np)
				}
				if pos.I+1 < h && d[pos.I+1][pos.J] < 0 && f[pos.I+1][pos.J] == '.' && !q.Find(pos.I+1, pos.J) {
					np := Pos{pos.I + 1, pos.J, pos.D + 1}
					q.Push(np)
				}
				if q.Size() == 0 {
					break
				}
			}
			ans = Max(ans, maxd)
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	f = make([]string, h)
	for i := 0; i < h; i++ {
		f[i] = nextString()
	}
	//fmt.Println(SolveWarshallFloyd(h, w))
	fmt.Println(Solve(h, w))
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Pos struct {
	I int
	J int
	D int
}

type Queue struct {
	ps []Pos
}

func NewQueue() *Queue {
	return new(Queue)
}

func (this *Queue) Push(p Pos) {
	this.ps = append(this.ps, p)
}

func (this *Queue) Pop() *Pos {
	if len(this.ps) == 0 {
		return nil
	}
	p := this.ps[0]
	this.ps = this.ps[1:]
	return &p
}

func (this *Queue) Find(x, y int) bool {
	for _, v := range this.ps {
		if x == v.I && y == v.J {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
}

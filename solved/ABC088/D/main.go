package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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
		if x == v.J && y == v.I {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = 50*50 + 1

	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	f := make([][]int, h)
	v := make([][]bool, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	for i := 0; i < h; i++ {
		f[i] = make([]int, w)
		v[i] = make([]bool, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			f[i][j] = INF
		}
	}
	// 経路を探索
	q := NewQueue()
	q.Push(Pos{0, 0, 1})
	v[0][0] = true
	for q.Size() > 0 {
		p := q.Pop()
		f[p.I][p.J] = p.D
		for i := 0; i < 4; i++ {
			npi := p.I + di[i]
			npj := p.J + dj[i]
			if npi >= 0 && npi < h && npj >= 0 && npj < w {
				if !v[npi][npj] && s[npi][npj] == '.' && f[npi][npj] > p.D+1 {
					q.Push(Pos{npi, npj, p.D + 1})
					v[npi][npj] = true
				}
			}
		}
	}
	// たどり着けない
	if f[h-1][w-1] == INF {
		fmt.Println(-1)
		return
	}
	// スコアを計算
	black := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				black++
			}
		}
	}
	ans := h*w - f[h-1][w-1] - black
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

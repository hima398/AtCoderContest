package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	dirh := []int{-1, 0, 1, 0}
	dirw := []int{0, -1, 0, 1}
	warph := []int{-2, -2, -2, -2, -2, -1, -1, -1, -1, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	warpw := []int{-2, -1, 0, 1, 2, -2, -1, 1, 2, -2, 2, -2, -1, 1, 2, -2, -1, 0, 1, 2}

	h, w := nextInt(), nextInt()
	ch, cw := nextInt(), nextInt()
	ch--
	cw--
	dh, dw := nextInt(), nextInt()
	dh--
	dw--

	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	d := make([][]int, h)
	v0 := make([][]bool, h)
	v1 := make([][]bool, h)
	fixed := make([][]bool, h)
	for i := 0; i < h; i++ {
		d[i] = make([]int, w)
		v0[i] = make([]bool, w)
		v1[i] = make([]bool, w)
		fixed[i] = make([]bool, w)
	}
	q0 := NewQueue()
	q1 := NewQueue()
	q0.Push(Pos{ch, cw, 0})
	v0[ch][cw] = true
	for q0.Size() > 0 || q1.Size() > 0 {
		var pos *Pos
		if q0.Size() > 0 {
			pos = q0.Pop()
		} else {
			pos = q1.Pop()
		}

		if !fixed[pos.I][pos.J] {
			d[pos.I][pos.J] = pos.D
			fixed[pos.I][pos.J] = true
			if pos.I == dh && pos.J == dw {
				fmt.Println(d[dh][dw])
				return
			}
		}

		for dir := 0; dir < 4; dir++ {
			ni := pos.I + dirh[dir]
			nj := pos.J + dirw[dir]
			if ni >= 0 && ni < h && nj >= 0 && nj < w && s[ni][nj] == '.' && !v0[ni][nj] {
				q0.Push(Pos{ni, nj, pos.D})
				v0[ni][nj] = true
			}
		}
		for dir := 0; dir < len(warph); dir++ {
			ni := pos.I + warph[dir]
			nj := pos.J + warpw[dir]
			if ni >= 0 && ni < h && nj >= 0 && nj < w && s[ni][nj] == '.' && !v0[ni][nj] && !v1[ni][nj] {
				q1.Push(Pos{ni, nj, pos.D + 1})
				v1[ni][nj] = true
			}
		}
	}
	fmt.Println(-1)

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

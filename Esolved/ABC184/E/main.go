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

	const INF = 2000*2000 + 1

	h, w := nextInt(), nextInt()
	a := make([]string, h)
	d := make([][]int, h)
	var warp [26][]Pos
	var si, sj, gi, gj int
	for i := 0; i < h; i++ {
		a[i] = nextString()
		d[i] = make([]int, w)
		for j := range d[i] {
			if a[i][j] == 'S' {
				si, sj = i, j
			}
			if a[i][j] == 'G' {
				gi, gj = i, j
			}
			idx := int(a[i][j]) - int('a')
			if idx >= 0 && idx < 26 {
				warp[idx] = append(warp[idx], Pos{i, j, 0})
			}
			d[i][j] = INF
		}
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	q := NewQueue()
	d[si][sj] = 0
	q.Push(Pos{si, sj, 0})
	for q.Size() > 0 {
		p := q.Pop()
		//d[p.I][p.J] = Min(d[p.I][p.J], p.D)

		for k := 0; k < len(di); k++ {
			ni, nj := p.I+di[k], p.J+dj[k]
			if ni >= 0 && ni < h && nj >= 0 && nj < w {
				if a[ni][nj] != '#' && d[ni][nj] == INF {
					d[ni][nj] = p.D + 1
					q.Push(Pos{ni, nj, d[ni][nj]})
				}
			}
		}
		idx := int(a[p.I][p.J]) - int('a')
		if idx >= 0 && idx < 26 {
			for _, wp := range warp[idx] {
				if p.I == wp.I && p.J == wp.J {
					continue
				}
				d[wp.I][wp.J] = p.D + 1
				wp.D = d[wp.I][wp.J]
				//fmt.Println(idx, wp)
				q.Push(wp)
			}
			warp[idx] = []Pos{}
		}
	}
	/*
		for i := 0; i < h; i++ {
			fmt.Println(d[i])
		}
	*/
	ans := d[gi][gj]
	if ans == INF {
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

func (this *Queue) Find(i, j int) bool {
	for _, v := range this.ps {
		if j == v.J && i == v.I {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	r, c int
	dir  int
}

type DoubleEndedQueue struct {
	front []Pos
	rear  []Pos
}

func NewDoubleEndedQueue() *DoubleEndedQueue {
	return &DoubleEndedQueue{}
}

func (deq *DoubleEndedQueue) PushFront(p Pos) {
	deq.front = append(deq.front, p)
}

func (deq *DoubleEndedQueue) PushBack(p Pos) {
	deq.rear = append(deq.rear, p)
}

func (deq *DoubleEndedQueue) Pop() (ret Pos) {
	if len(deq.front) > 0 {
		nf := len(deq.front)
		ret = deq.front[nf-1]
		deq.front = deq.front[:nf-1]
	} else if len(deq.rear) > 0 {
		ret = deq.rear[0]
		deq.rear = deq.rear[1:]
	} else {
		//キューが空の時にPopされるエラー処理を書くのだけど省略
	}
	return ret
}

func (deq *DoubleEndedQueue) Len() int {
	return len(deq.front) + len(deq.rear)
}

func (deq *DoubleEndedQueue) Print() {
	fmt.Println(deq.front)
	fmt.Println(deq.rear)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = int(1e6) + 1
	h, w := nextInt(), nextInt()
	rs, cs := nextInt(), nextInt()
	rs--
	cs--
	rt, ct := nextInt(), nextInt()
	rt--
	ct--
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	c := make([][][4]int, h)
	for i := 0; i < h; i++ {
		c[i] = make([][4]int, w)
		for j := 0; j < w; j++ {
			for k := 0; k < 4; k++ {
				c[i][j][k] = INF
			}
		}
	}
	dr := []int{0, -1, 0, 1}
	dc := []int{-1, 0, 1, 0}
	//var q []Pos
	q := NewDoubleEndedQueue()
	for k := 0; k < 4; k++ {
		c[rs][cs][k] = 0
	}
	for k := 0; k < 4; k++ {
		q.PushBack(Pos{rs, cs, k})
	}
	for q.Len() > 0 {
		//q.Print()
		p := q.Pop()
		/*
			if p.cost >= 0 {
				c[p.r][p.c][p.dir] = p.cost
			}
		*/
		for k := 0; k < 4; k++ {
			cost := c[p.r][p.c][p.dir]
			if p.dir != k {
				cost++
			}
			nr, nc := p.r+dr[k], p.c+dc[k]

			if nr >= 0 && nr < h && nc >= 0 && nc < w {
				if s[nr][nc] == '.' && c[nr][nc][k] > cost {
					c[nr][nc][k] = cost
					if p.dir != k {
						q.PushBack(Pos{nr, nc, k})
					} else {
						q.PushFront(Pos{nr, nc, k})
					}
				}
			}
		}
	}

	//fmt.Println(c)
	ans := INF
	for k := 0; k < 4; k++ {
		ans = Min(ans, c[rt][ct][k])
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

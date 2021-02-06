package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	X int
	Y int
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
		if x == v.X && y == v.Y {
			return true
		}
	}
	return false
}

func (this *Queue) Size() int {
	return len(this.ps)
}

func SolveBFS(n, x, y int) []int {
	x--
	y--
	dist := make([]int, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dist[j] = 2001
		}
		q := NewQueue()
		p := Pos{i, 0, 0}
		q.Push(p)
		for {
			v := q.Pop()
			if dist[v.X] >= 2001 {
				dist[v.X] = v.D
			}
			if v.X-1 >= 0 && dist[v.X-1] >= 2001 {
				np := Pos{v.X - 1, 0, v.D + 1}
				q.Push(np)
			}
			if v.X+1 < n && dist[v.X+1] >= 2001 {
				np := Pos{v.X + 1, 0, v.D + 1}
				q.Push(np)
			}
			if v.X == x && dist[y] >= 2001 {
				np := Pos{y, 0, v.D + 1}
				q.Push(np)
			}
			if v.X == y && dist[x] >= 2001 {
				np := Pos{x, 0, v.D + 1}
				q.Push(np)
			}
			if q.Size() == 0 {
				break
			}
		}
		for j := 0; j < n; j++ {
			ans[dist[j]]++
		}
	}
	for i := 0; i < n; i++ {
		ans[i] /= 2
	}
	return ans
}

func Solve(n, x, y int) []int {

	ans := make([]int, n)
	for i := 1; i < n; i++ {
		for j := i; j <= n; j++ {
			min := Min(2001, Min(Abs(j-i), Min(Abs(x-i)+Abs(j-y)+1, Abs(y-i)+Abs(j-x)+1)))
			ans[min]++
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x, y := nextInt(), nextInt(), nextInt()
	//ans := SolveBFS(n, x, y)
	ans := Solve(n, x, y)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 1; i < n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

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

	n := nextInt()
	e := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	//fmt.Println(e)
	start := -1
	for i := 0; i < n; i++ {
		if len(e[i]) == 1 {
			start = i
			break
		}
	}
	//fmt.Println(start)

	q := NewQueue()
	visited := make([]bool, n)
	q.Push(Pos{start, 0})
	visited[start] = true
	ans := make([][]int, 2)
	for q.Size() > 0 {
		p := q.Pop()
		if p == nil {
			break
		}
		ans[p.D%2] = append(ans[p.D%2], p.X)
		for _, v := range e[p.X] {
			if !visited[v] {
				q.Push(Pos{v, p.D + 1})
				visited[v] = true
			}
		}
	}
	//fmt.Println(ans[0])
	//fmt.Println(ans[1])
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	if len(ans[0]) >= n/2 {
		for i := 0; i < n/2; i++ {
			fmt.Fprintf(out, "%d ", ans[0][i]+1)
		}
	} else {
		for i := 0; i < n/2; i++ {
			fmt.Fprintf(out, "%d ", ans[1][i]+1)
		}
	}
	fmt.Fprintln(out, "")

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type Pos struct {
	X int
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

func (this *Queue) Size() int {
	return len(this.ps)
}

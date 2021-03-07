package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	t int
	i int
}

type Item struct {
	i int // index of node
	c int // color
}

type Queue struct {
	q []Item
}

func NewQueue() *Queue {

	return new(Queue)
}
func (this *Queue) Push(v Item) {
	this.q = append(this.q, v)
}

func (this *Queue) Pop() Item {
	if this.Size() == 0 {
		return Item{0, 0}
	}
	ret := this.q[0]
	this.q = this.q[1:]
	return ret
}

func (this *Queue) Size() int {
	return len(this.q)
}

func (this *Queue) PrintQueue() {
	for _, v := range this.q {
		fmt.Println(v)
	}
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	edge := make(map[int][]Edge)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		edge[a] = append(edge[a], Edge{b, i})
		edge[b] = append(edge[b], Edge{a, i})
	}
	//fmt.Println(edge)
	k := 0
	for _, v := range edge {
		k = Max(k, len(v))
	}
	//fmt.Println(k)

	ans := make([]int, n-1)
	v := make([]bool, n) // visited
	q := NewQueue()
	q.Push(Item{0, 0})
	v[0] = true
	for q.Size() > 0 {
		// parent
		p := q.Pop()
		color := 1
		for _, e := range edge[p.i] {
			if v[e.t] || ans[e.i] > 0 {
				continue
			}
			//fmt.Println(p.f, e, cc)
			if p.c == color {
				color++
			}
			ans[e.i] = color
			q.Push(Item{e.t, color})
			color++
			v[e.t] = true
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintln(out, k)
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

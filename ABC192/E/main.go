package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007
const INF = 1 << 60

var sc = bufio.NewScanner(os.Stdin)

type Edge struct {
	B int
	C int
	T int
	K int
}

type Graph struct {
	E [][]Edge
}

func NewGraph(n int) *Graph {
	e := make([][]Edge, n)
	return &Graph{e}
}

func (this *Graph) AddEdge(a, b, t, k int) {
	a--
	b--
	this.E[a] = append(this.E[a], Edge{b, 0, t, k})
}

func (this *Graph) Len() int {
	return len(this.E)
}

func (this *Graph) Dijkstra(x, y int) int {
	n := this.Len()
	q := &PriorityQueue{}
	heap.Init(q)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[x] = 0
	//TODO
	init := func(b, c int) {
		heap.Push(q, Edge{b, c, 0, 0})
	}
	push := func(b, c, t, k int) {
		var nc int
		if c%k == 0 {
			nc = c + t
		} else {
			// c + ((c/k + 1) * k) - c + t
			nc = ((c/k + 1) * k) + t
		}
		if dist[b] > nc {
			dist[b] = nc
			heap.Push(q, Edge{b, nc, t, k})
		}
	}
	init(x, 0)
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		//fmt.Println("cur", cur)
		if dist[cur.B] < cur.C {
			continue
		}
		//fmt.Println("Edge", this.E[cur.B])
		for _, e := range this.E[cur.B] {
			push(e.B, e.C+cur.C, e.T, e.K)
		}
	}
	//fmt.Println(dist)
	return dist[y]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, x, y := nextInt(), nextInt(), nextInt(), nextInt()
	x--
	y--
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, t, k := nextInt(), nextInt(), nextInt(), nextInt()
		g.AddEdge(a, b, t, k)
		g.AddEdge(b, a, t, k)
	}
	d := g.Dijkstra(x, y)
	if d == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(d)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].C < pq[j].C
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

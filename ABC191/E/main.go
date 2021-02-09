package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

const INF = 1 << 60

var sc = bufio.NewScanner(os.Stdin)

type Edge struct {
	T int
	C int
}

type Graph struct {
	E [][]Edge
}

func NewGraph(n int) *Graph {
	e := make([][]Edge, n)
	return &Graph{e}
}

func (this *Graph) AddEdge(s, t, c int) {
	s--
	t--
	this.E[s] = append(this.E[s], Edge{t, c})
}

func (this *Graph) Len() int {
	return len(this.E)
}

func (this *Graph) Dijkstra(s int) int {
	n := this.Len()
	q := &PriorityQueue{}
	heap.Init(q)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	push := func(t, c int) {
		if dist[t] > c {
			dist[t] = c
			heap.Push(q, Edge{t, c})
		}
	}
	push(s, 0)
	dist[s] = INF
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if dist[cur.T] < cur.C {
			continue
		}
		for _, e := range this.E[cur.T] {
			push(e.T, cur.C+e.C)
		}
	}
	return dist[s]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		g.AddEdge(a, b, c)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		d := g.Dijkstra(i)
		if d == INF {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, d)
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

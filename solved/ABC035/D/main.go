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
	t, c int
}

type PriorityQueue []Edge

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].c < pq[j].c
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

type Graph struct {
	e    [][]Edge
	dist []int
}

func NewGraph(n int) *Graph {
	e := make([][]Edge, n)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	return &Graph{e, dist}
}

func (g *Graph) AddEdge(s, t, c int) {
	g.e[s] = append(g.e[s], Edge{t, c})
}
func (g *Graph) GetDist(i int) int {
	return g.dist[i]
}
func (g *Graph) Len() int {
	return len(g.e)
}

func (g *Graph) Dijkstra(s int) {
	q := NewPriorityQueue()
	heap.Init(q)
	push := func(t, c int) {
		if g.dist[t] > c {
			g.dist[t] = c
			heap.Push(q, Edge{t, c})
		}
	}
	push(s, 0)
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if g.dist[cur.t] < cur.c {
			continue
		}
		for _, e := range g.e[cur.t] {
			push(e.t, cur.c+e.c)
		}
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, t := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	g1 := NewGraph(n)
	g2 := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		a--
		b--
		g1.AddEdge(a, b, c)
		g2.AddEdge(b, a, c)
	}
	g1.Dijkstra(0)
	g2.Dijkstra(0)
	ans := 0
	for i := 0; i < n; i++ {
		money := Max(t-g1.GetDist(i)-g2.GetDist(i), 0) * a[i]
		ans = Max(ans, money)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

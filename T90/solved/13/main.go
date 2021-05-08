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

type Graph struct {
	e [][]Edge
}

func NewGraph(n int) *Graph {
	e := make([][]Edge, n)
	return &Graph{e}
}

func (this *Graph) AddEdge(s, t, c int) {
	s--
	t--
	this.e[s] = append(this.e[s], Edge{t, c})
}

func (this *Graph) Len() int {
	return len(this.e)
}

func (this *Graph) Dijkstra(s int) []int {
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
		if dist[cur.t] < cur.c {
			continue
		}
		for _, e := range this.e[cur.t] {
			push(e.t, cur.c+e.c)
		}
	}
	return dist
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	//e := make(map[int][]Edge)
	g := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		g.AddEdge(a, b, c)
		g.AddEdge(b, a, c)
	}
	d1 := g.Dijkstra(0)
	dn := g.Dijkstra(n - 1)
	//fmt.Println(d1)
	//fmt.Println(dn)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintln(out, d1[n-1])
	for i := 1; i < n-1; i++ {
		ans := d1[i] + dn[i]
		fmt.Fprintln(out, ans)
	}
	fmt.Fprintln(out, d1[n-1])
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

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

const INF = 1 << 60
const size = 301

var sc = bufio.NewScanner(os.Stdin)

var memo [size][size]int

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

func (this *Graph) Dijkstra(s, t int) int {
	if s > t {
		s, t = t, s
	}
	if memo[s][t] != INF {
		return memo[s][t]
	}
	n := this.Len()
	q := &PriorityQueue{}
	heap.Init(q)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[s] = 0
	push := func(t, c int) {
		if dist[t] > c {
			dist[t] = c
			heap.Push(q, Edge{t, c})
		}
	}
	heap.Push(q, Edge{s, 0})
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		//fmt.Println("cur", cur)
		if dist[cur.T] < cur.C {
			continue
		}
		//fmt.Println("Edge", this.E[cur.B])
		for _, e := range this.E[cur.T] {
			push(e.T, e.C+cur.C)
		}
	}
	memo[s][t] = dist[t]
	return memo[s][t]
}

func SolveDijkstra() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			memo[i][j] = INF
		}
	}
	n, m := nextInt(), nextInt()

	g := NewGraph(n)
	for i := 0; i < m; i++ {
		a, b, t := nextInt(), nextInt(), nextInt()
		g.AddEdge(a, b, t)
		g.AddEdge(b, a, t)
	}
	ans := INF
	for i := 0; i < n; i++ {
		max := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			max = Max(max, g.Dijkstra(i, j))
		}
		ans = Min(ans, max)
	}
	fmt.Println(ans)

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}
	for i := 0; i < m; i++ {
		a, b, t := nextInt(), nextInt(), nextInt()
		a--
		b--
		dist[a][b] = t
		dist[b][a] = t
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	ans := INF
	for i := 0; i < n; i++ {
		max := 0
		for j := 0; j < n; j++ {
			max = Max(max, dist[i][j])
		}
		ans = Min(ans, max)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

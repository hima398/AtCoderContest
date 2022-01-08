package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Edge struct {
	r, f, t, c int
}

type Graph struct {
	e [][]Edge
}

//type Graph [][]Edge

func NewGraph(n int) *Graph {
	g := new(Graph)
	g.e = make([][]Edge, n)
	return g
}

func (g Graph) Len() int {
	return len(g.e)
}

func (g Graph) Redge(e *Edge) *Edge {
	return &g.e[e.t][e.r]
}

func (g Graph) Flow(e *Edge, f int) {
	e.c -= f
	g.Redge(e).c += f
}

func (g Graph) Add(f, t, c int) {
	fr := len(g.e[f])
	tr := len(g.e[t])
	g.e[f] = append(g.e[f], Edge{tr, f, t, c})
	g.e[t] = append(g.e[t], Edge{fr, t, f, c})
}

type FordFlukerson struct {
	inf     int
	visited []bool
}

func NewFordFlukerson(inf, nodes int) *FordFlukerson {
	ff := new(FordFlukerson)
	ff.inf = inf
	ff.visited = make([]bool, nodes)
	return ff
}

func (ff *FordFlukerson) Solve(g *Graph, s, t int) (ans int) {

	var Dfs func(g *Graph, i, t, v int) int
	Dfs = func(g *Graph, i, t, v int) int {
		if i == t {
			return v
		}
		ff.visited[i] = true
		for _, e := range g.e[i] {
			if ff.visited[e.t] {
				continue
			}
			if e.c == 0 {
				continue
			}

			flow := Dfs(g, e.t, t, Min(v, e.c))
			if flow == 0 {
				continue
			}
			g.Flow(&e, flow)

			return flow
		}
		return 0
	}

	for {
		for i := range ff.visited {
			ff.visited[i] = false
		}
		flow := Dfs(g, s, t, ff.inf)
		fmt.Println(flow)
		if flow == 0 {
			return ans
		}
		ans += flow
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, g, e := nextInt(), nextInt(), nextInt()
	graph := NewGraph(n + 1)
	p := nextIntSlice(g)
	//edges := make([][]int, n)
	for i := 0; i < e; i++ {
		a, b := nextInt(), nextInt()
		//a--
		//b--
		graph.Add(a, b, 1)
		graph.Add(b, a, 1)
		//edges[a] = append(edges[a], b)
		//edges[b] = append(edges[b], a)
	}
	for _, pi := range p {
		graph.Add(pi-1, n, 1)
		graph.Add(n, pi-1, 1)
	}
	ff := NewFordFlukerson(1<<60, graph.Len())
	ans := ff.Solve(graph, 0, n)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

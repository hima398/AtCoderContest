package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Item struct {
	s, i, j, k int
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].s > pq[j].s
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

/*
func SolveHonestly(a, b, c []int, k int) (ans []int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	q := new(PriorityQueue)
	heap.Init(q)

	for _, ai := range a {
		for _, bi := range b {
			for _, ci := range c {
				heap.Push(q, ai+bi+ci)
			}
		}
	}
	for i := 0; i < k; i++ {
		v := heap.Pop(q)
		ans = append(ans, v.(int))
	}
	return ans
}
*/

type Pos struct {
	i, j, k int
}

func Solve(a, b, c []int, k int) (ans []int) {
	const INF = int(1e10) + 1
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	m := make(map[Pos]bool)
	q := new(PriorityQueue)
	heap.Init(q)
	heap.Push(q, Item{a[0] + b[0] + c[0], 0, 0, 0})
	m[Pos{0, 0, 0}] = true
	//da, db, dc := INF, INF, INF
	for len(ans) < k {
		item := heap.Pop(q).(Item)
		ans = append(ans, item.s)
		ni, nj, nk := item.i+1, item.j+1, item.k+1
		if ni < len(a) && !m[Pos{ni, item.j, item.k}] {
			heap.Push(q, Item{a[ni] + b[item.j] + c[item.k], ni, item.j, item.k})
			m[Pos{ni, item.j, item.k}] = true
		}
		if nj < len(b) && !m[Pos{item.i, nj, item.k}] {
			heap.Push(q, Item{a[item.i] + b[nj] + c[item.k], item.i, nj, item.k})
			m[Pos{item.i, nj, item.k}] = true
		}
		if nk < len(c) && !m[Pos{item.i, item.j, nk}] {
			heap.Push(q, Item{a[item.i] + b[item.j] + c[nk], item.i, item.j, nk})
			m[Pos{item.i, item.j, nk}] = true
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x, y, z, k := nextInt(), nextInt(), nextInt(), nextInt()
	a := nextIntSlice(x)
	b := nextIntSlice(y)
	c := nextIntSlice(z)

	//ans := SolveHonestly(a, b, c, k)
	ans := Solve(a, b, c, k)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
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

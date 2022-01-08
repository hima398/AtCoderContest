package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	s := *pq
	n := len(s)
	item := s[n-1]
	*pq = s[0 : n-1]
	return item
}

func Solve(n, m int, a, b []int) (ans []int) {
	e := make([][]int, n)
	v := make([]int, n)
	for i := 0; i < m; i++ {
		a[i]--
		b[i]--
		e[a[i]] = append(e[a[i]], b[i])
		v[b[i]]++
	}
	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < n; i++ {
		if v[i] == 0 {
			heap.Push(pq, i)
		}
	}
	for pq.Len() > 0 {
		//fmt.Println("q = ", q, "pq = ", pq)
		//fmt.Println("v = ", v)
		//node := pq.Pop().(int)
		node := heap.Pop(pq).(int)
		ans = append(ans, node+1)
		for _, next := range e[node] {
			v[next]--
			if v[next] == 0 {
				heap.Push(pq, next)
			}
		}
	}

	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a, b := make([]int, m), make([]int, m)

	for i := 0; i < m; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	ans := Solve(n, m, a, b)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	if len(ans) < n {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintf(out, "%d", ans[0])
	for i := 1; i < n; i++ {
		fmt.Fprintf(out, " %d", ans[i])
	}
	fmt.Fprintln(out)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

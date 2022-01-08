package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Sushi struct {
	t, d int
}

type PriorityQueue struct {
	q     []Sushi
	isAsc bool
}

func (pq PriorityQueue) Len() int {
	return len(pq.q)
}
func (pq *PriorityQueue) Less(i, j int) bool {
	if pq.isAsc {
		return pq.q[i].d < pq.q[j].d
	} else {
		return pq.q[i].d > pq.q[j].d
	}
}
func (pq *PriorityQueue) Swap(i, j int) {
	pq.q[i], pq.q[j] = pq.q[j], pq.q[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	pq.q = append(pq.q, item.(Sushi))
}

func (pq *PriorityQueue) Pop() interface{} {
	s := pq.q
	n := len(s)
	item := s[n-1]
	pq.q = s[0 : n-1]
	return item
}

func Solve(n, k int, t, d []int) int {

	q1 := &PriorityQueue{}
	heap.Init(q1)
	for i := 0; i < n; i++ {
		heap.Push(q1, Sushi{t[i], d[i]})
	}
	q2 := &PriorityQueue{}
	heap.Init(q2)
	q2.isAsc = true
	m := make(map[int]int)
	ans := 0
	for i := 0; i < k; i++ {
		s := heap.Pop(q1).(Sushi)
		if m[s.t] >= 1 {
			heap.Push(q2, s)
		}
		m[s.t]++
		ans += s.d
	}

	ans += len(m) * len(m)

	sum := ans
	for len(m) < k && q1.Len() > 0 && q2.Len() > 0 {
		p := len(m)
		s1 := heap.Pop(q1).(Sushi)
		s2 := heap.Pop(q2).(Sushi)
		m[s2.t]--
		if m[s1.t] >= 1 {
			heap.Push(q2, s1)
		}
		m[s1.t]++
		sum = sum - s2.d + s1.d - p*p + len(m)*len(m)
		//fmt.Println(ans, p, m)
		ans = Max(ans, sum)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	t, d := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		t[i], d[i] = nextInt(), nextInt()
	}
	ans := Solve(n, k, t, d)
	fmt.Println(ans)
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

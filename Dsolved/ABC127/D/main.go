package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Cards struct {
	v, n int
}

type PriorityQueue []Cards

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].v > pq[j].v
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Cards))
}
func (pq *PriorityQueue) Pop() interface{} {
	s := *pq
	n := len(s)
	item := s[n-1]
	*pq = s[0 : n-1]
	return item
}

func FirstSolve() {
	n, m := nextInt(), nextInt()
	q := &PriorityQueue{}
	heap.Init(q)
	for i := 0; i < n; i++ {
		heap.Push(q, Cards{v: nextInt(), n: 1})
	}
	//fmt.Println(q)
	for i := 0; i < m; i++ {
		b, c := nextInt(), nextInt()
		heap.Push(q, Cards{v: c, n: b})
	}
	//fmt.Println(q)
	ans := 0
	i := 0
	for {
		item := heap.Pop(q).(Cards)
		if i+item.n >= n {
			ans += item.v * (n - i)
			break
		}
		ans += item.v * item.n
		i += item.n
	}
	fmt.Println(ans)

}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	FirstSolve()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Item struct {
	V int
}

// Priority Queue
type PQ []Item

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].V > pq[j].V }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(i interface{}) {
	*pq = append(*pq, i.(Item))
}

func (pq *PQ) Pop() interface{} {
	s := *pq //スライスとしてアクセスできるようにしておく
	n := len(s)
	i := s[n-1]
	*pq = s[0 : n-1]
	return i
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	//a := make([]int, n)
	a := &PQ{}
	heap.Init(a)
	for i := 0; i < n; i++ {
		ai := nextInt()
		heap.Push(a, Item{ai})
	}
	for i := m; i > 0; i-- {
		iv := heap.Pop(a).(Item).V
		//fmt.Println(iv)
		iv /= 2
		heap.Push(a, Item{iv})
	}
	ans := 0
	for a.Len() > 0 {
		ans += heap.Pop(a).(Item).V
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

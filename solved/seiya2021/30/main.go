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
	es := *pq // EdgeのSlice
	n := pq.Len()
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func (pq *PriorityQueue) PrintQueue() {
	fmt.Println(pq)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	var ans []int
	var offsets []int
	pq := &PriorityQueue{}
	heap.Init(pq)
	offset := 0
	for i := 0; i < q; i++ {
		t := nextInt()
		switch t {
		case 1:
			x := nextInt()
			heap.Push(pq, x-offset)
		case 2:
			x := nextInt()
			offset += x
		case 3:
			ball := heap.Pop(pq).(int)
			ans = append(ans, ball)
			offsets = append(offsets, offset)
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := range ans {
		fmt.Fprintln(out, ans[i]+offsets[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

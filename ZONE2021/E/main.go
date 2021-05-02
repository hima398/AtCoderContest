package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const INF = 1 << 60

type Cell struct {
	i, j, k int
	c       int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	r, c := nextInt(), nextInt()
	a := make([][]int, r)
	for i := 0; i < r; i++ {
		a[i] = make([]int, c-1)
		for j := 0; j < c-1; j++ {
			a[i][j] = nextInt()
		}
	}
	/*
		for i := range a {
			fmt.Println(a[i])
		}
	*/
	b := make([][]int, r-1)
	for i := 0; i < r-1; i++ {
		b[i] = make([]int, c)
		for j := 0; j < c; j++ {
			b[i][j] = nextInt()
		}
	}
	/*
		for i := range b {
			fmt.Println(b[i])
		}
	*/
	cost := make([][][]int, 2)
	for k := 0; k < 2; k++ {
		cost[k] = make([][]int, r)
		for i := 0; i < r; i++ {
			cost[k][i] = make([]int, c)
			for j := 0; j < c; j++ {
				cost[k][i][j] = INF
			}
		}
	}
	//cost[0][0][0] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	push := func(i, j, k, c int) {
		if cost[k][i][j] > c {
			cost[k][i][j] = c
			heap.Push(pq, Cell{i, j, k, c})
		}
	}
	push(0, 0, 0, 0)
	//Dikstraで最短経路を探索
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Cell)
		//重複の確認
		if cost[cur.k][cur.i][cur.j] < cur.c {
			continue
		}
		//次のステップをキューに詰む
		if cur.k == 0 {
			if cur.j < c-1 {
				push(cur.i, cur.j+1, cur.k, cur.c+a[cur.i][cur.j])
			}
			if cur.j > 0 {
				push(cur.i, cur.j-1, cur.k, cur.c+a[cur.i][cur.j-1])
			}
			if cur.i < r-1 {
				push(cur.i+1, cur.j, cur.k, cur.c+b[cur.i][cur.j])
			}
			push(cur.i, cur.j, cur.k+1, cur.c+1)
		} else {
			// cur.k == 1
			if cur.i > 0 {
				push(cur.i-1, cur.j, cur.k, cur.c+1)
			}
			push(cur.i, cur.j, cur.k-1, cur.c)
		}
	}

	/*
		for i := range cost[0] {
			fmt.Println(cost[0][i])
		}
	*/

	fmt.Println(cost[0][r-1][c-1])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type PriorityQueue []Cell

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
	*pq = append(*pq, item.(Cell))
}

func (pq *PriorityQueue) Pop() interface{} {
	s := *pq
	n := len(s)
	item := s[n-1]
	*pq = s[0 : n-1]
	return item
}

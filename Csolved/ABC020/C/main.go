package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	i, j int
}

type Edge struct {
	i, j, c int
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
	es := *pq
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const INF = int(1e12) + 1
	h, w, t := nextInt(), nextInt(), nextInt()
	f := make([]string, h)
	for i := 0; i < h; i++ {
		f[i] = nextString()
	}
	// スタートとゴールの座標を探す
	var s, g Pos
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if f[i][j] == 'S' {
				s = Pos{i, j}
			}
			if f[i][j] == 'G' {
				g = Pos{i, j}
			}
		}
	}
	//ダイクストラ(Dijkstra)法で黒マスのコストxの場合t以内にゴールに着けるか探索する
	Check := func(x int) bool {
		pq := &PriorityQueue{}
		heap.Init(pq)
		costs := make([][]int, h)
		for i := 0; i < h; i++ {
			costs[i] = make([]int, w)
			for j := 0; j < w; j++ {
				costs[i][j] = INF
			}
		}
		costs[s.i][s.j] = 0
		heap.Push(pq, Edge{s.i, s.j, 0})

		di := []int{-1, 0, 1, 0}
		dj := []int{0, -1, 0, 1}
		for pq.Len() > 0 {
			cur := heap.Pop(pq).(Edge)
			// 行先にすでに小さいコストがあればループの最初に戻る
			if costs[cur.i][cur.j] < cur.c {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if ni >= 0 && ni < h && nj >= 0 && nj < w {
					var c int
					if f[ni][nj] == '#' {
						c = x
					} else {
						// f[ni][nj] == '.' or 'S' or 'G'
						c = 1
					}
					//移動さきのコストより少ないコストの場合だけキューに積む
					if costs[ni][nj] > cur.c+c {
						costs[ni][nj] = cur.c + c
						heap.Push(pq, Edge{ni, nj, cur.c + c})
					}
				}
			}
		}
		/*
			fmt.Printf("x = %d\n", x)
			for i := 0; i < h; i++ {
				fmt.Println(costs[i])
			}
		*/
		return costs[g.i][g.j] <= t
	}
	ok, ng := 1, t+1
	for ng-ok > 1 {
		c := (ng + ok) / 2
		if Check(c) {
			ok = c
		} else {
			ng = c
		}
	}
	fmt.Println(ok)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

const INF = 1 << 60

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	e := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}
	k := nextInt()
	c := nextIntSlice(k)
	for i := range c {
		c[i]--
	}
	//fmt.Println(c)
	mb := (1 << k) - 1
	var cs []int
	for i := 1; i <= mb; i++ {
		cs = append(cs, i)
	}
	sort.Slice(cs, func(i, j int) bool {
		if bits.OnesCount(uint(cs[i])) < bits.OnesCount(uint(cs[j])) {
			return true
		}
		if bits.OnesCount(uint(cs[i])) == bits.OnesCount(uint(cs[j])) && cs[i] < cs[j] {
			return true
		}
		return false
	})
	//fmt.Println(cs)
	type DistKey struct {
		k1, k2 int
	}
	dist := make(map[DistKey]int)
	for _, ci1 := range c {
		for _, ci2 := range c {
			if ci1 == ci2 {
				dist[DistKey{ci1, ci2}] = 0
			} else {
				dist[DistKey{ci1, ci2}] = INF
			}
		}
	}

	//fmt.Println(dist)

	for _, ci := range c {
		q := NewQueue()
		visited := make([]bool, n)
		q.Push(Pos{ci, 0})
		visited[ci] = true
		for q.Size() > 0 {
			p := q.Pop()
			if v, ok := dist[DistKey{ci, p.i}]; ok {
				dist[DistKey{ci, p.i}] = Min(v, p.d)
			}
			for _, v := range e[p.i] {
				if !visited[v] {
					q.Push(Pos{v, p.d + 1})
					visited[v] = true
				}
			}
		}
	}
	//fmt.Println(dist)

	dp := make([][]int, mb+1)
	for i := 0; i <= mb; i++ {
		dp[i] = make([]int, k)
		for j := 0; j < k; j++ {
			dp[i][j] = INF
		}
	}
	for i := range c {
		dp[1<<(i)][i] = 1
	}
	/*
		for i := 0; i <= len(cs); i++ {
			fmt.Println(dp[i])
		}
	*/

	for _, b := range cs {
		for i1 := range c {
			for i2 := range c {
				if b&(1<<i1) == 0 {
					if i1 != i2 {
						nb := b | (1 << i1)
						if d, ok := dist[DistKey{c[i1], c[i2]}]; ok {
							dp[nb][i1] = Min(dp[nb][i1], dp[b][i2]+d)
						}
					}
				}
			}
		}
	}
	/*
		for i := 0; i <= mb; i++ {
			fmt.Println(dp[i])
		}
	*/
	ans := INF
	for i := range c {
		ans = Min(ans, dp[mb][i])
	}
	if ans == INF {
		ans = -1
	}
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

type Pos struct {
	i, d int
}
type Queue struct {
	q []Pos
}

func NewQueue() *Queue {

	return new(Queue)
}
func (this *Queue) Push(v Pos) {
	this.q = append(this.q, v)
}

func (this *Queue) Pop() *Pos {
	if this.Size() == 0 {
		return nil
	}
	ret := this.q[0]
	this.q = this.q[1:]
	return &ret
}

func (this *Queue) Size() int {
	return len(this.q)
}

func (this *Queue) PrintQueue() {
	fmt.Println(this.q)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

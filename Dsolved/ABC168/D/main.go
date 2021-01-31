package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Queue struct {
	q []int
}

func NewQueue() *Queue {
	return new(Queue)
}

func (this *Queue) Push(i int) {
	this.q = append(this.q, i)
}

func (this *Queue) Pop() int {
	if this.Size() == 0 {
		return -1
	}
	i := this.q[0]
	this.q = this.q[1:]
	return i
}

func (this *Queue) Size() int {
	return len(this.q)
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	r := make(map[int][]int) // route
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		r[a] = append(r[a], b)
		r[b] = append(r[b], a)
	}
	//fmt.Println(r)

	v := make([]bool, n+1) // visited
	sp := make([]int, n+1) // sign point

	q := NewQueue()
	q.Push(1)
	v[1] = true
	cnt := 1
	for {
		par := q.Pop()
		for _, chi := range r[par] {
			if !v[chi] {
				sp[chi] = par
				q.Push(chi)
				cnt++
			}
			v[chi] = true
		}
		if q.Size() == 0 {
			break
		}
	}
	ok := cnt == n
	/*
		for i := 2; i <= n; i++ {
			ok = ok && sp[i] != 0
		}
	*/
	if ok {
		fmt.Println("Yes")
		for i := 2; i <= n; i++ {
			fmt.Fprintln(out, sp[i])
		}
	} else {
		fmt.Fprintln(out, "No")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type IntQueue struct {
	q []int
}

func NewIntQueue() *IntQueue {
	return new(IntQueue)
}
func (this *IntQueue) Push(v int) {
	this.q = append(this.q, v)
}

func (this *IntQueue) Pop() int {
	ret := this.q[0]
	this.q = this.q[1:]
	return ret
}

func (this *IntQueue) Size() int {
	return len(this.q)
}

func (this *IntQueue) PrintQueue() {
	fmt.Println(this.q)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k := nextInt()
	queue := NewIntQueue()
	for i := 1; i <= 9; i++ {
		queue.Push(i)
	}
	//queue.PrintQueue()
	for i := 1; i <= 100000; i++ {
		v := queue.Pop()
		if i == k {
			// Print Answer
			fmt.Println(v)
			return
		}
		v1 := v % 10
		for j := v1 - 1; j <= v1+1; j++ {
			if j == -1 || j == 10 {
				continue
			}
			nv := (10*v + j)
			queue.Push(nv)
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

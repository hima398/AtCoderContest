package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Worker struct {
	t, i int
}

func FirstSolve(n int) int {
	mins := int(1e6)
	a := make([]Worker, n)
	b := make([]Worker, n)

	for i := 0; i < n; i++ {
		a[i] = Worker{nextInt(), i}
		b[i] = Worker{nextInt(), i}
		mins = Min(mins, a[i].t+b[i].t)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].t < a[j].t })
	sort.Slice(b, func(i, j int) bool { return b[i].t < b[j].t })
	if a[0].i != b[0].i {
		return Min(mins, Max(a[0].t, b[0].t))
	} else {
		return Min(mins, Min(Max(a[0].t, b[1].t), Max(a[1].t, b[0].t)))
	}
}

func Solve(n int) int {
	ans := 2*int(1e5) + 1
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				ans = Min(ans, a[i]+b[i])
			} else {
				ans = Min(ans, Max(a[i], b[j]))
			}
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()

	//fmt.Println(FirstSolve(n))
	fmt.Println(Solve(n))

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

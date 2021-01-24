package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	var route []int
	var loop []int
	visited := make([]int, n+1)
	for current := 1; visited[current] < 2; current = a[current] {
		if visited[current] == 1 {
			loop = append(loop, current)
		} else {
			route = append(route, current)
		}
		visited[current]++
	}
	b := len(route) - len(loop)
	if k < b {
		fmt.Println(route[k])
	} else {
		k -= b
		fmt.Println(loop[k%len(loop)])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

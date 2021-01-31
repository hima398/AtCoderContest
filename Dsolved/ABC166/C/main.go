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

	n, m := nextInt(), nextInt()
	h := make([]int, n+1)
	hst := make([]int, n+1) // heighest
	for i := 1; i <= n; i++ {
		h[i] = nextInt()
	}
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		hst[a] = Max(hst[a], h[b])
		hst[b] = Max(hst[b], h[a])
	}
	//fmt.Println(hst)
	ans := 0
	for i := 1; i <= n; i++ {
		if h[i] > hst[i] {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

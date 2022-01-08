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
	n, m, lx, ly := nextInt(), nextInt(), nextInt(), nextInt()
	x := make([]int, n)
	y := make([]int, m)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}
	for i := 0; i < m; i++ {
		y[i] = nextInt()
	}
	for z := lx + 1; z <= ly; z++ {
		ok := true
		for i := 0; i < n; i++ {
			ok = ok && x[i] < z
		}
		for i := 0; i < m; i++ {
			ok = ok && y[i] >= z
		}
		if ok {
			fmt.Println("No War")
			return
		}
	}
	fmt.Println("War")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	z := make([]int, n)
	w := make([]int, n)
	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		z[i] = x + y
		w[i] = x - y
	}
	sort.Ints(z)
	sort.Ints(w)
	//fmt.Println(d)
	ans := Max(z[n-1]-z[0], w[n-1]-w[0])
	fmt.Println(ans)
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

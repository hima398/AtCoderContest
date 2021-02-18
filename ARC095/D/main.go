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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	sort.Ints(a)
	ai := a[n-1]
	aj := a[n-2]
	pre := Abs((ai+1)/2 - aj)
	for j := n - 3; j >= 0; j-- {
		c := Abs((ai+1)/2 - a[j])
		if c > pre {
			break
		}
		aj = a[j]
		pre = c
	}
	fmt.Println(ai, aj)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Combination(N, K int) int {
	if K <= 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

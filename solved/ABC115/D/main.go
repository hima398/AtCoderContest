package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var size [51]int
var patties [51]int

var sc = bufio.NewScanner(os.Stdin)

func Patties(n int) int {
	if patties[n] > 0 {
		return patties[n]
	}
	if n == 0 {
		patties[0] = 1
		return patties[0]
	}
	patties[n] = 2*Patties(n-1) + 1
	return patties[n]
}

func Solve(n, x int) int {
	if n == 0 {
		return 1
	}
	if x == 1 {
		return 0
	} else if x > 1 && x <= 1+size[n-1] {
		return Solve(n-1, x-1)
	} else if x > 1+size[n-1] && x <= 2+size[n-1] {
		return Patties(n-1) + 1
	} else if x > 2+size[n-1] && x <= 2+2*size[n-1] {
		return Patties(n-1) + 1 + Solve(n-1, x-2-size[n-1])
	} else {
		return 2*Patties(n-1) + 1
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x := nextInt(), nextInt()
	size[0] = 1
	for i := 1; i <= n; i++ {
		size[i] = 2*size[i-1] + 3
	}
	//fmt.Println(size)
	fmt.Println(Solve(n, x))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

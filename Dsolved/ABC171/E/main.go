package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func SolveHonestly(n int, a []int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		t := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			t ^= a[j]
		}
		ret[i] = t
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		s ^= a[i]
	}
	//ans := SolveHonestly(n, a)
	fmt.Printf("%d", s^a[0])
	for i := 1; i < n; i++ {
		fmt.Printf(" %d", s^a[i])
	}
	fmt.Println("")
}

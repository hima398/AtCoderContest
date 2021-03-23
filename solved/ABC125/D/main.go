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

	n := nextInt()
	a := make([]int, n)
	sum := 0
	nm := 0
	hasZero := false
	min := int(1e9) + 1
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		sum += Abs(a[i])
		if a[i] < 0 {
			nm++
		} else if a[i] == 0 {
			hasZero = true
		}
		min = Min(min, Abs(a[i]))
	}
	if hasZero || nm%2 == 0 {
		fmt.Println(sum)
	} else {
		fmt.Println(sum - 2*min)
	}
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

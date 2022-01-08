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

	n, _ := nextInt(), nextInt()
	a := nextIntSlice(n)
	mx := make([]int, n)
	mx[n-1] = a[n-1]
	for i := n - 1; i > 0; i-- {
		mx[i-1] = Max(mx[i], a[i-1])
	}
	//fmt.Println(mx)
	var max int
	for i := 0; i < n; i++ {
		mx[i] = mx[i] - a[i]
		max = Max(max, mx[i])
	}
	var ans int
	for i := 0; i < n; i++ {
		if mx[i] == max {
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

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

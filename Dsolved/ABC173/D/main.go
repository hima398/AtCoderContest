package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	ans := 0
	for i := 0; i < n-1; i++ {
		ans += a[(i+1)/2]
	}
	fmt.Println(ans)
}

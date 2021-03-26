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
	n, k := nextInt(), nextInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}
	sort.Ints(h)
	ans := int(1e9) + 1
	for i := 0; i <= n-k; i++ {
		//fmt.Println(i, i+k-1)
		ans = Min(ans, h[i+k-1]-h[i])
	}
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

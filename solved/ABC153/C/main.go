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
	// 全員必殺技で勝てる
	if k >= n {
		fmt.Println(0)
		return
	}
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}
	sort.Ints(h)
	m := n - k
	ans := 0
	for i := 0; i < m; i++ {
		ans += h[i]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

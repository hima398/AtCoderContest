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

	var p []int
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		p = append(p, b)
		p = append(p, a-b)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(p)))
	ans := 0
	for i := 0; i < k; i++ {
		ans += p[i]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	// 1 <= n <= 2 * 10**5
	// 1 <= w
	n, w := nextInt(), nextInt()
	// 0 < si < ti < 2 * 10**5
	a := make([]int, 200001)
	// pi <= 10**9 int64?
	for i := 0; i < n; i++ {
		s, t, p := nextInt(), nextInt(), nextInt()
		a[s] += p
		a[t] -= p
	}
	for i := 0; i < 200000; i++ {
		a[i+1] += a[i]
		if a[i] > w {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

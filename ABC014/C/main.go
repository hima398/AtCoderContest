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

	const max = int(1e6) + 2
	n := nextInt()
	m := make([]int, max)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		m[a]++
		m[b+1]--
	}
	for i := 1; i <= max-1; i++ {
		m[i] += m[i-1]
	}
	ans := -1
	for i := 0; i < max-1; i++ {
		if ans < m[i] {
			ans = m[i]
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

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
	b := make([]int, n)
	for i := 0; i < n-1; i++ {
		b[i] = nextInt()
	}
	ans := b[0]
	for i := 0; i < n-1; i++ {
		ans += Min(b[i], b[i+1])
	}
	ans += b[n-2]
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

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

	k, n := nextInt(), nextInt()
	max := 0
	ao := nextInt()
	pa := ao
	for i := 1; i < n; i++ {
		a := nextInt()
		max = Max(max, a-pa)
		pa = a
	}
	// この時点でpaはAn
	max = Max(max, k+ao-pa)
	ans := k - max
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

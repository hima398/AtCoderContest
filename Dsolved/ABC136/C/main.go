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
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}

	for i := n - 1; i > 0; i-- {
		if h[i] >= h[i-1] {
			continue
		}
		// h[i] < h[i-1]
		if h[i-1]-h[i] > 1 {
			fmt.Println("No")
			return
		}
		h[i-1]--
	}
	fmt.Println("Yes")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

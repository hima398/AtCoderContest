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

	n, k := nextInt(), nextInt()
	ba := make([]int, int(1e5)+1)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		ba[a] += b
	}
	idx := 0
	for i, v := range ba {
		idx += v
		if k <= idx {
			fmt.Println(i)
			return
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

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

	const INF = int(1e9) + 1

	n := nextInt()
	min, imin := INF, -1
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		if min > b {
			imin = a
			min = b
		}
	}
	ans := imin + min
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

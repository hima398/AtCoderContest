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

	n, a, b := nextInt(), nextInt(), nextInt()
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		d := x[i+1] - x[i]
		if d*a <= b {
			ans += d * a
		} else {
			ans += b
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

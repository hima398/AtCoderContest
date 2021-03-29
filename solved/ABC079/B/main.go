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
	l := make([]int, n+1)
	l[0] = 2
	l[1] = 1
	for i := 2; i <= n; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	fmt.Println(l[n])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

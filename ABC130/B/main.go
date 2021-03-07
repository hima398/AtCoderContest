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

	n, x := nextInt(), nextInt()
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = nextInt()
	}
	d := 0
	ans := 1
	for i := 0; i < n; i++ {
		if d+l[i] > x {
			break
		}
		d += l[i]
		ans++
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

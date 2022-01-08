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

	n, q := nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < q; i++ {
		l, r, t := nextInt(), nextInt(), nextInt()
		l--
		for j := l; j < r; j++ {
			a[j] = t
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, ans := range a {
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

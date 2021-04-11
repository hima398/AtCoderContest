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
	t := make([]int, n)
	st := 0
	for i := 0; i < n; i++ {
		t[i] = nextInt()
		st += t[i]
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	m := nextInt()
	for i := 0; i < m; i++ {
		p, x := nextInt(), nextInt()
		ans := st - t[p-1] + x
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

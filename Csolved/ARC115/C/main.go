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
	e := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := i; j <= n; j += i {
			if e[j] == 0 {
				e[j] = i
			}
		}
	}
	ans := make([]int, n+1)
	ans[1] = 1
	for i := 2; i <= n; i++ {
		ans[i] = ans[i/e[i]] + 1
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "%d", ans[1])
	for i := 2; i <= n; i++ {
		fmt.Fprintf(out, " %d", ans[i])
	}
	fmt.Fprintln(out, "")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

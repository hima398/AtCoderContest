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
	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	for i := n; i > 0; i-- {
		s := 0
		for j := 2 * i; j <= n; j += i {
			s ^= b[j]
		}
		b[i] = s ^ a[i]
	}
	var ans []int
	for i := 1; i <= n; i++ {
		if b[i] == 1 {
			ans = append(ans, i)
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintln(out, len(ans))
	if len(ans) > 0 {
		fmt.Fprintf(out, "%d", ans[0])
		for i := 1; i < len(ans); i++ {
			fmt.Fprintf(out, "% d", ans[i])
		}
		fmt.Fprintln(out, "")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

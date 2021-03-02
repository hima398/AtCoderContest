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
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}

	m := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			m[1] -= 2 * a[i]
		} else {
			m[1] += 2 * a[i]
		}
	}
	m[1] /= 2
	for i := 2; i <= n; i++ {
		m[i] = 2*a[i-1] - m[i-1]
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "%d", m[1])
	for i := 2; i <= n; i++ {
		fmt.Fprintf(out, " %d", m[i])
	}
	fmt.Fprintln(out, "")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

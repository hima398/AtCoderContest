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
	ans := make([]int, n+1)
	for x := 1; x <= 100; x++ {
		for y := x; y <= 100; y++ {
			for z := y; z <= 100; z++ {
				nn := x*x + y*y + z*z + x*y + y*z + z*x
				if nn <= n {
					if x == y && y == z && z == x {
						ans[nn]++
					} else if x != y && y != z && z != x {
						ans[nn] += 6
					} else {
						ans[nn] += 3
					}
				}
			}
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

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

	f := make([]int, int(1e4)+1)
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 100; y++ {
			for z := 1; z <= 100; z++ {
				n := x*x + y*y + z*z + x*y + y*z + z*x
				if n <= int(1e4) {
					f[n]++
				}
			}
		}
	}

	ln := nextInt()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 1; i <= ln; i++ {
		fmt.Fprintln(out, f[i])
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

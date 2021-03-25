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

	n, ly := nextInt(), nextInt()
	for x := 0; x <= n; x++ {
		for y := 0; y <= n-x; y++ {
			z := n - x - y
			if 10000*x+5000*y+1000*z == ly {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

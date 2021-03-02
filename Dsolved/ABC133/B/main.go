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

	n, d := nextInt(), nextInt()
	var x [10][10]int
	for i := 0; i < n; i++ {
		for j := 0; j < d; j++ {
			x[i][j] = nextInt()
		}
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			dist := 0
			for k := 0; k < d; k++ {
				dd := x[i][k] - x[j][k]
				dist += dd * dd
			}
			for k := 1; k*k <= dist; k++ {
				if dist%k == 0 && dist/k == k {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

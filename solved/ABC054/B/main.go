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

	n, m := nextInt(), nextInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = nextString()
	}
	b := make([]string, m)
	for i := 0; i < m; i++ {
		b[i] = nextString()
	}
	for i := 0; i < n-m+1; i++ {
		for j := 0; j < n-m+1; j++ {
			ok := true
			for ii := 0; ii < m; ii++ {
				for jj := 0; jj < m; jj++ {
					//fmt.Println(i+ii, j+jj, a[i+ii][j+jj], i, j, b[ii][jj])
					ok = ok && a[i+ii][j+jj] == b[ii][jj]
				}
			}
			if ok {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

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
	d1 := m - 2*n
	if d1 == 0 {
		fmt.Println(n, 0, 0)
		return
	} else if d1 < 0 {
		fmt.Println(-1, -1, -1)
		return
	}
	d2 := m - 4*n
	if d2 > 0 {
		fmt.Println(-1, -1, -1)
		return
	}
	if d1-n < 0 {
		fmt.Println(n-d1, d1, 0)
		return
	} else {
		// d1 - n >= 0
		//fmt.Println(2*n-d1, d1-n)
		for i := 0; i <= 2*n-d1; i++ {
			for j := 0; j <= d1-n; j++ {
				if (n-i-j)*2+i*3+j*4 == m {
					fmt.Println(n-i-j, i, j)
					return
				}
			}
		}
		fmt.Println(-1, -1, -1)
		return
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

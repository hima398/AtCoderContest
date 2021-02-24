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

	// r < g < b
	// a:r, b:g, c:b
	a, b, c := nextInt(), nextInt(), nextInt()
	k := nextInt()
	for i := 0; i < k; i++ {
		if a < b && b < c {
			fmt.Println("Yes")
			return
		}
		if b >= c {
			c *= 2
		} else if a >= b {
			b *= 2
		}
	}
	if a < b && b < c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

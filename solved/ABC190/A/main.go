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

	a, b, c := nextInt(), nextInt(), nextInt()
	if c == 0 {
		if a <= b {
			fmt.Println("Aoki")
		} else {
			fmt.Println("Takahashi")
		}
	} else {
		if a >= b {
			fmt.Println("Takahashi")
		} else {
			fmt.Println("Aoki")
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

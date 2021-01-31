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

	x := nextInt()
	for a := -118; a <= 119; a++ {
		for b := -118; b <= 119; b++ {
			if a*a*a*a*a-b*b*b*b*b == x {
				fmt.Println(a, b)
				return
			}
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

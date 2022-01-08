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
	x1 := x / 100
	x2 := x % 100

	if x1 >= 20 {
		fmt.Println(1)
	} else {
		if x2 > 5*x1 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

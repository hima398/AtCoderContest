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

	a, b := nextInt(), nextInt()
	ab := a + b
	if ab >= 15 && b >= 8 {
		fmt.Println(1)
	} else if ab >= 10 && b >= 3 {
		fmt.Println(2)
	} else if ab >= 3 {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

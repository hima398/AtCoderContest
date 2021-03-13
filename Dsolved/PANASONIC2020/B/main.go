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

	h, w := nextInt(), nextInt()
	if h == 1 || w == 1 {
		fmt.Println(1)
	} else if h%2 == 1 && w%2 == 1 {
		fmt.Println(h*w/2 + 1)
	} else {
		fmt.Println(h * w / 2)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

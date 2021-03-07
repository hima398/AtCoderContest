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

	w, h, x, y := nextInt(), nextInt(), nextInt(), nextInt()
	ans := float64(w*h) / 2.0
	valid := 0
	if 2*x == w && 2*y == h {
		valid = 1
	}
	fmt.Println(ans, valid)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

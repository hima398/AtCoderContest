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

	w, h := nextInt(), nextInt()
	if w%16 == 0 && h%9 == 0 {
		fmt.Println("16:9")
	} else if w%4 == 0 && h%3 == 0 {
		fmt.Println("4:3")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

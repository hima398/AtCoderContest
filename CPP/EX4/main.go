package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

const v = 60 * 60 * 24 * 365

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	fmt.Println(v)
	fmt.Println(2 * v)
	fmt.Println(5 * v)
	fmt.Println(10 * v)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextString(), nextString()
	if len(a) >= len(b) {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

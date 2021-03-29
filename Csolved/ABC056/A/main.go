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
	if (a == "H" && b == "H") || (a == "D" && b == "D") {
		fmt.Println("H")
	} else {
		fmt.Println("D")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

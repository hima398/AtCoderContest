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

	c := nextString()
	if c[0] == c[1] && c[1] == c[2] {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

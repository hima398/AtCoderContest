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
	fmt.Println(string(c[0] + 1))
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

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

	x := nextString()
	fmt.Println(x[0] - 'A' + 1)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

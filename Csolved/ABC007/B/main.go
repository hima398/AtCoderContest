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

	a := nextString()
	if a == "a" {
		fmt.Println(-1)
	} else if len(a) == 1 {
		fmt.Println(string(a[0] - 1))
	} else {
		fmt.Println(a[:len(a)-1])
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

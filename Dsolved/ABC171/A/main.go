package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	if s[0] >= 'A' && s[0] <= 'Z' {
		fmt.Println("A")
	} else if s[0] >= 'a' && s[0] <= 'z' {
		fmt.Println("a")
	}
}

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

	s := nextString()
	if s == "AAA" || s == "BBB" {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

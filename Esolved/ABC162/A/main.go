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

	n := nextString()
	if n[0] == '7' || n[1] == '7' || n[2] == '7' {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

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
	if len(a) > len(b) {
		fmt.Println("GREATER")
		return
	} else if len(a) < len(b) {
		fmt.Println("LESS")
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			fmt.Println("GREATER")
			return
		} else if a[i] < b[i] {
			fmt.Println("LESS")
			return
		}
	}
	fmt.Println("EQUAL")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

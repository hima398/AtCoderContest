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
	m := make(map[rune]bool)
	for _, v := range s {
		if m[v] {
			fmt.Println("no")
			return
		}
		m[v] = true
	}
	fmt.Println("yes")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

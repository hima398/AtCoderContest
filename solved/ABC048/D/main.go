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
	if s[0] == s[len(s)-1] {
		// axax...xa
		if len(s)%2 == 1 {
			fmt.Println("Second")
		} else {
			fmt.Println("First")
		}
	} else {
		// axax...ax
		if len(s)%2 == 1 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

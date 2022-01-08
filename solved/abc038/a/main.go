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
	if s[len(s)-1] == 'T' {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

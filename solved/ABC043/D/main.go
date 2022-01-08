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
	n := len(s)
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			fmt.Println(i+1, i+2)
			return
		}
		if i < n-2 && s[i] == s[i+2] {
			fmt.Println(i+1, i+3)
			return
		}
	}
	fmt.Println(-1, -1)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

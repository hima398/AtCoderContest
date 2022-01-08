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
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ss := s[:i] + s[j:]
			if ss == "keyence" {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

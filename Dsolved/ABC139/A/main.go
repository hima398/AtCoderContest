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

	s, t := nextString(), nextString()
	ans := 0
	for i := range s {
		if s[i] == t[i] {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

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
	nw := 0
	for i := 0; i < n; i++ {
		if s[i] == 'W' {
			nw++
		}
	}
	ans := 0
	idx := 0
	for i := 0; i < n; i++ {
		if s[i] == 'W' {
			ans += i - idx
			idx++
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

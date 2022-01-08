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
	m := make(map[byte]bool)
	for i := range s {
		m[s[i]] = true
	}
	for i := 0; i < 26; i++ {
		c := byte(i) + 'a'
		if !m[c] {
			fmt.Println(string(c))
			return
		}
	}
	fmt.Println("None")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

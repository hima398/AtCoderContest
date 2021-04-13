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

	c := nextString()
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
		fmt.Println("vowel")
	} else {
		fmt.Println("consonant")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

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
	i := 0
	for {
		if s[i] == 'A' {
			break
		}
		i++
	}
	j := len(s) - 1
	for {
		if s[j] == 'Z' {
			break
		}
		j--
	}
	fmt.Println(j - i + 1)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

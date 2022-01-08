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
	ok := true
	for i := 1; i <= len(s); i++ {
		if i%2 == 1 {
			// Odd
			if s[i-1] == 'L' {
				ok = false
				break
			}
		} else {
			// Even
			if s[i-1] == 'R' {
				ok = false
				break
			}
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

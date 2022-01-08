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
	if s[0] != 'A' {
		fmt.Println("WA")
		return
	}
	nc := 0
	nu := 0 // 大文字の数
	for i := 2; i < len(s)-1; i++ {
		if s[i] == 'C' {
			nc++
		} else if s[i]-'A' >= 0 && s[i]-'A' < 26 {
			nu++
		}
	}

	if s[1]-'A' >= 0 && s[1]-'A' < 26 {
		nu++
	}

	if s[len(s)-1]-'A' >= 0 && s[len(s)-1]-'A' < 26 {
		nu++
	}
	if nc == 1 && nu == 0 {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

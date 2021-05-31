package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	t := nextString()
	canWin := true
	for i := range s {
		if s[i] != t[i] {
			if s[i] == '@' {
				canWin = canWin && strings.Contains("atcoder", string(t[i]))
			} else if t[i] == '@' {
				canWin = canWin && strings.Contains("atcoder", string(s[i]))
			} else {
				canWin = false
			}
		}
	}
	if canWin {
		fmt.Println("You can win")
	} else {
		fmt.Println("You will lose")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

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
	ng, np := 0, 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if np == ng {
			if s[i] == 'p' {
				ans--
			}
			// s[i] == 'g'ならあいこ
			ng++
		} else {
			if s[i] == 'g' {
				ans++
			}
			// s[i] == 'p'ならあいこ
			np++
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

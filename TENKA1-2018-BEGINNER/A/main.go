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
	if len(s) == 2 {
		fmt.Println(s)
		return
	} else if len(s) == 3 {
		ans := string(s[2]) + string(s[1]) + string(s[0])
		fmt.Println(ans)
		return
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
